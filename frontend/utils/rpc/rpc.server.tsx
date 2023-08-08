import {messages as wrapper} from '@/messages/wrapper'
import { cookies } from 'next/headers';

async function rpcRequest(url: string, request: wrapper.Wrapper, privateRequest: boolean) {
    let headers: HeadersInit = {}
    if (privateRequest) {
        const cookieStore = cookies();
        const sessionCookie = cookieStore.get(process.env.SESSION_COOKIE_NAME!);
        if (sessionCookie === undefined) {
            throw Error("Trying to make private request without session cookie");
        }
        const cookieString = `${process.env.SESSION_COOKIE_NAME!}=${sessionCookie.value}`; 
        headers = {
            cookie: cookieString,
        }
    }
    const requestData = request.serialize();
    const response = await fetch(url,
        {
            method: "POST",
            body: requestData,
            headers: headers,
            cache: "force-cache",
        }
    );
    if (response.status === 500) {
        // Most likely no other data was attached to this response...
        throw new Error("Internal server error")
    }
    const resonseBlob = await response.blob();
    const responseData = await resonseBlob.arrayBuffer();
    return wrapper.Wrapper.deserialize(new Uint8Array(responseData));
}

export async function ServerPublicRPCRequest(request: wrapper.Wrapper, identifier: string) {
    request.api_token = process.env.SERVER_RPC_API_TOKEN!;
    return rpcRequest(`${process.env.API_URL!}/rpc/server?id=${identifier}`, request, false);
}

export async function ServerPrivateRPCRequest(request: wrapper.Wrapper, identifier: string) {
    request.api_token = process.env.SERVER_RPC_API_TOKEN!;
    return rpcRequest(`${process.env.API_URL!}/rpc/private/server?id=${identifier}`, request, true);
}

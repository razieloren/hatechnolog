import {messages as wrapper} from '@/messages/wrapper'

async function rpcRequest(url: string, request: wrapper.Wrapper) {
    const requestData = request.serialize();
    const response = await fetch(url,
        {
            method: "POST",
            body: requestData,
            mode: "cors",
            credentials: "include",
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

export async function ClientPublicRPCRequest(request: wrapper.Wrapper) {
    request.api_token = process.env.CLIENT_RPC_API_TOKEN!;
    return rpcRequest(`${process.env.API_URL!}/rpc/client`, request);
}

export async function ClientPrivateRPCRequest(request: wrapper.Wrapper) {
    request.api_token = process.env.CLIENT_RPC_API_TOKEN!;
    return rpcRequest(`${process.env.API_URL!}/rpc/private/client`, request);
}

import {messages as wrapper} from '@/messages/wrapper'

async function rpcRequest(url: string, request: wrapper.Wrapper) {
    const requestData = request.serialize();
    try {
        const response = await fetch(url,
            {
                method: "POST",
                body: requestData,
            }
        );
        if (!response.ok) {
            throw new Error(`response is not ok: ${response.status}`)
        }
        const resonseBlob = await response.blob();
        const responseData = await resonseBlob.arrayBuffer();
        return wrapper.Wrapper.deserialize(new Uint8Array(responseData));
    } catch (e) {
        console.error(`Error while making RPC request: ${e}`)
        throw(e);
    }
}

export async function ServerRPCRequest(request: wrapper.Wrapper) {
    request.api_token = process.env.SERVER_RPC_API_TOKEN!;
    return rpcRequest(process.env.SERVER_RPC_API_URL!, request);
}

export async function ClientRPCRequest(request: wrapper.Wrapper) {
    request.api_token = process.env.CLIENT_RPC_API_TOKEN!;
    return rpcRequest(process.env.CLIENT_RPC_API_URL!, request);
}

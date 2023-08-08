import { ServerPublicRPCRequest, ServerPrivateRPCRequest } from "./rpc.server";
import {messages as wrapper} from '@/messages/wrapper'
import {messages as user_api} from '@/messages/user/api'

export async function GetMyUser(): Promise<user_api.GetUserResponse> {
    const wrappedResponse = await ServerPrivateRPCRequest(new wrapper.Wrapper({
        get_user_request: new user_api.GetUserRequest({})
    }), "myuser");
    if (!wrappedResponse.has_get_user_response) {
        throw new Error("no get_user_response");
    }
    return wrappedResponse.get_user_response;
}

export async function PublicGetUser(handle: string): Promise<user_api.GetUserResponse> {
    const wrappedResponse = await ServerPublicRPCRequest(new wrapper.Wrapper({
        get_user_request: new user_api.GetUserRequest({
            handle: handle,
        })
    }), `getuser-${handle}`);
    if (!wrappedResponse.has_get_user_response) {
        throw new Error("no get_user_response");
    }
    return wrappedResponse.get_user_response;
}
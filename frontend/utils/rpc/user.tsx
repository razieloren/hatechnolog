import { ClientPrivateRPCRequest } from "./rpc";
import {messages as wrapper} from '@/messages/wrapper'
import {messages as user_api} from '@/messages/user/api'

export async function FetchMyUser(): Promise<user_api.GetUserResponse> {
    const wrappedResponse = await ClientPrivateRPCRequest(new wrapper.Wrapper({
        get_user_request: new user_api.GetUserRequest({})
    }));
    if (!wrappedResponse.has_get_user_response) {
        throw new Error("bad message type in response");
    }
    return wrappedResponse.get_user_response;
}
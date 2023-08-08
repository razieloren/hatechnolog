import { ClientPrivateRPCRequest } from "./rpc.client";
import {messages as wrapper} from '@/messages/wrapper'
import {messages as user_api} from '@/messages/user/api'

export async function RefreshDiscordUser(): Promise<user_api.UpdateDiscordUserResponse> {
    const wrappedResponse = await ClientPrivateRPCRequest(new wrapper.Wrapper({
        update_discord_user_request: new user_api.UpdateDiscordUserRequest({})
    }), "refreshdiscorduser");
    if (!wrappedResponse.has_update_discord_user_response) {
        throw new Error("no update_discrd_user_response");
    }
    return wrappedResponse.update_discord_user_response;
}
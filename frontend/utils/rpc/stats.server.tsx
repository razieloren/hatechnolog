import { ServerPublicRPCRequest } from "./rpc.server";
import {messages as wrapper} from '@/messages/wrapper'
import {messages as stats_api} from '@/messages/stats/api'
import {messages as stats_helpers} from '@/messages/stats/helpers'

export type StatsProps = {
    valid: boolean;
    discord: stats_helpers.LatestDiscordStats;
    youtube: stats_helpers.LatestYoutubeStats;
    github: stats_helpers.LatestGithubStats;
}

export async function GetLatestStats(): Promise<StatsProps> {
    const wrappedResponse = await ServerPublicRPCRequest(new wrapper.Wrapper({
        get_latest_stats_request: new stats_api.GetLatestStatsRequest({
            discord_guild: process.env.DISCORD_GUILD_NAME!,
            youtube_channel: process.env.YOUTUBE_CHANNEL_NAME!,
            github_repo: process.env.GITHUB_CHANNEL_NAME!,
        })
    }), "lateststats");
    if (!wrappedResponse.has_get_latest_stats_response) {
        throw new Error("no latest_stats_response");
    }
    const stats = wrappedResponse.get_latest_stats_response;
    return {
        valid: stats.discord_stats.valid && stats.youtube_stats.valid && stats.github_stats.valid,
        discord: stats.discord_stats,
        youtube: stats.youtube_stats,
        github: stats.github_stats,
    }
}
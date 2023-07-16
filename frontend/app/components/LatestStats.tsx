'use client'

import { useEffect, useState } from "react"
import { Websocket } from "websocket-ts"
import {messages as wrapper} from '../messages/wrapper'
import {messages as stats_api} from '../messages/stats/api'
import {messages as stats_helpers} from '../messages/stats/helpers'
import { InitAPIWS } from "../wsapi"
import { setInterval } from "timers/promises"
import { wrap } from "module"

type LatestStatsProps = {
    discord_latest_stats: stats_helpers.LatestDiscordStats;
    youtube_latest_stats: stats_helpers.LatestYoutubeStats;
    github_latest_stats: stats_helpers.LatestGithubStats;
}

export default function LatestStats() {
    const [latestStats, setLatestStats]  = useState<LatestStatsProps>({
        discord_latest_stats: new stats_helpers.LatestDiscordStats(),
        youtube_latest_stats: new stats_helpers.LatestYoutubeStats(),
        github_latest_stats: new stats_helpers.LatestGithubStats(),
    });

    function onReady(ws: Websocket) {
        const wrapped = new wrapper.Wrapper({
            latest_stats_push_request: new stats_api.LatestStatsPushRequest({
                discord_guild: "הטכנולוג",
                youtube_channel: "itsnapgaming",
                github_repo: "hatechnolog"
            })
        });
        ws.send(wrapped.serialize());
    }

    function onMessage(ws: Websocket, data: any) {
        const wrapped = wrapper.Wrapper.deserialize(data);
        if (!wrapped.has_latest_stats_push_response) {
            return
        }
        const response = wrapped.latest_stats_push_response
        setLatestStats({
            discord_latest_stats: response.discord_stats,
            youtube_latest_stats: response.youtube_stats,
            github_latest_stats: response.github_stats,
        });
    }

    useEffect(() => {
        const ws = InitAPIWS(wrapper.Endpoint.latest_stats_push, onReady, onMessage);
        return () => ws.close();
    }, []);

    return (
        <div>
            <p>
                {latestStats.discord_latest_stats.total_members}
            </p>
            <p>
                {latestStats.youtube_latest_stats.subscribers}
            </p>
            <p>
                {latestStats.github_latest_stats.contributors}
            </p>
        </div>
    )
}
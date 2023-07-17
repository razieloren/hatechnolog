'use client'

import Link from 'next/link'
import Image from 'next/image'
import { useEffect, useState } from "react"
import { Websocket } from "websocket-ts"
import {messages as wrapper} from '../messages/wrapper'
import {messages as stats_api} from '../messages/stats/api'
import {messages as stats_helpers} from '../messages/stats/helpers'
import { InitAPIWS } from "../wsapi"
import discordLogo from '../../public/social/discord_logo.png'
import youtubeLogo from '../../public/social/youtube_logo.png'
import GithubLatestStats from './GithubStats'

type LatestStatsProps = {
    discord: stats_helpers.LatestDiscordStats;
    youtube: stats_helpers.LatestYoutubeStats;
    github: stats_helpers.LatestGithubStats;
}

export default function LatestStats() {
    const [latestStats, setLatestStats]  = useState<LatestStatsProps>({
        discord: new stats_helpers.LatestDiscordStats(),
        youtube: new stats_helpers.LatestYoutubeStats(),
        github: new stats_helpers.LatestGithubStats(),
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
            discord: response.discord_stats,
            youtube: response.youtube_stats,
            github: response.github_stats,
        });
    }

    useEffect(() => {
        const ws = InitAPIWS(wrapper.Endpoint.latest_stats_push, onReady, onMessage);
        return () => ws.close();
    }, []);

    const avgJoinHours = (latestStats.discord.join_avg_sec / 60 / 60).toLocaleString()
    return (
        <div className="stats">
            <GithubLatestStats
                contributors={latestStats.github.contributors}
                commits={latestStats.github.commits}
            />
           <ul className="stats-social-list">
                <li className="stats-social-item">
                    <Link 
                        className="stats-social-link"
                        href={process.env.DISCORD_INVITE_URL!}
                        rel="noopener noreferrer nofollow" 
                        target="_blank"
                    >
                        <Image className="stats-social-logo" src={discordLogo} alt="Hatechnolog Discord Server"/>
                    </Link>
                    <ul className="stats-social-item-list">
                        <li><span className="stats-number">{latestStats.discord.total_members.toLocaleString()}</span> משתמשות ומשתמשים</li>
                        <li><span className="stats-number">{latestStats.discord.new_members.toLocaleString()}</span> הצטרפו ב-<span className="stats-number">{latestStats.discord.new_members_period_days.toLocaleString()}</span> ימים האחרונים</li>
                        <li>כל <span className="stats-number">{avgJoinHours}</span> שעות מצטרפת עוד משתמשת</li>
                    </ul>
                </li>
                <li className="stats-social-item">
                    <Link 
                            className="stats-social-link"
                            href={process.env.YOUTUBE_CHANNEL_URL!}
                            rel="noopener noreferrer nofollow" 
                            target="_blank"
                        >
                        <Image className="stats-social-logo" src={youtubeLogo} alt="Hatechnolog Youtube Channel"/>
                    </Link>
                    <ul className="stats-social-item-list">
                        <li><span className="stats-number">{latestStats.youtube.subscribers.toLocaleString()}</span> רשומות ורשומים</li>
                        <li><span className="stats-number">{latestStats.youtube.views.toLocaleString()}</span> צפיות בכל הקורסים</li>
                    </ul>
                </li>
            </ul>
        </div>
    )
}
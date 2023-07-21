import Link from 'next/link'
import Image from 'next/image'
import discordLogo from '@/public/social/discord_logo_purple.svg'
import youtubeLogo from '@/public/social/youtube_logo_red.svg'
import blackStar from '@/public/misc/black_star.svg'
import twitterLogo from '@/public/social/twitter_logo_blue.svg'
import { StatsProps } from '@/utils/rpc/stats'


export default async function Social(props: StatsProps) {
    const avgJoinHours = (props.discord.join_avg_sec / 60 / 60).toLocaleString();
    return (
        <div className="flex flex-col justify-center items-center">
            <span className="text-xl font-bold">כל הפלטפורמות של קהילת הטכנולוג</span>
            <div className="flex xl:flex-row flex-col gap-8 my-2">
                <ul className="list-none my-3 flex flex-col items-center">
                    <li className="my-1">
                        <Link 
                            href={process.env.DISCORD_INVITE_URL!}
                            target="_blank"
                        >
                            <button className="social-button">
                                <Image className="w-8 h-8" src={discordLogo} alt="Hatechnolog Discord"/>
                                {props.valid ?
                                <span>הצטרפו ל-<span className="text-lime-300 font-bold">{props.discord.total_members.toLocaleString()}</span> משתמשים</span>
                                :
                                <span>הצטרפו לשרת הדיסקורד הקהילתי</span>
                                }
                            </button>
                        </Link>
                    </li>
                    {props.valid &&
                    <>
                        <li className="my-1">
                            <div className="flex gap-2">
                                <Image src={blackStar} alt="Discord Fact"/>
                                <span>ב-<span className="text-lime-300 font-bold">{props.discord.new_members_period_days.toLocaleString()}</span> ימים האחרונים הצטרפו <span className="text-lime-300 font-bold">{props.discord.new_members.toLocaleString()}</span> משתמשים חדשים</span>
                            </div>
                        </li>
                        <li className="my-1">
                            <div className="flex gap-2">
                                <Image src={blackStar} alt="Discord Fact"/>
                                <span>מצטרף משתמש חדש לקהילה כל <span className="text-lime-300 font-bold">{avgJoinHours}</span> שעות</span>
                            </div>
                        </li>
                    </>
                    }
                </ul>
                <ul className="list-none my-3 flex flex-col items-center">
                    <li className="my-1">
                        <Link 
                            href={process.env.YOUTUBE_CHANNEL_URL!}
                            target="_blank"
                        >
                            <button className="social-button">
                                <Image className="w-8 h-8" src={youtubeLogo} alt="Hatechnolog Youtube"/>
                                {props.valid ?
                                <span>הצטרפו ל-<span className="text-lime-300 font-bold">{props.youtube.subscribers.toLocaleString()}</span> רשומים</span>
                                :
                                <span>הירשמו לערוץ היוטיוב הקהילתי</span>
                                }
                            </button>
                        </Link>
                    </li>
                    {props.valid &&
                    <>
                        <li className="my-1">
                            <div className="flex gap-2">
                                <Image src={blackStar} alt="Youtube Fact"/>
                                <span><span className="text-lime-300 font-bold">{props.youtube.views.toLocaleString()}</span> צפיות בכל הקורסים</span>
                            </div>
                        </li>
                    </>
                    }
                </ul>
                <ul className="list-none my-3 flex flex-col items-center">
                    <li className="my-1">
                        <Link 
                        href={process.env.TWITTER_PROFILE_URL!}
                        target="_blank"
                        >
                            <button className="social-button">
                                <Image className="w-8 h-8" src={twitterLogo} alt="Hatechnolog Twitter"/>
                                <span>עדכונים שוטפים וחדשות</span>
                            </button>
                        </Link>
                    </li>
                </ul>
            </div>
        </div>
    )
}

/*
<div className="stats">
            <LoadableComponent
                isLoading={!latestStats.github_stats.valid}
                caption="טוען את הנתונים מגיטהאב..."
                component={<GithubLatestStats
                    contributors={latestStats.github_stats.contributors}
                    commits={latestStats.github_stats.commits}
                />}
            />
           <ul className="stats-social-list">
                <li className="stats-social-item">
                    <Link 
                        className="stats-social-link"
                        href={process.env.DISCORD_INVITE_URL!}

                        target="_blank"
                    >
                        <Image className="stats-social-logo" src={discordLogo} alt="Hatechnolog Discord Server"/>
                    </Link>
                    <LoadableComponent
                        isLoading={!props.discord.valid}
                        caption="טוען את הנתונים מדיסקורד..."
                        component={(
                            <ul className="stats-social-item-list">
                                <li><span className="stats-number">{props.discord.total_members.toLocaleString()}</span> משתמשות ומשתמשים</li>
                                <li><span className="stats-number">{props.discord.new_members.toLocaleString()}</span> הצטרפו ב-<span className="stats-number">{props.discord.new_members_period_days.toLocaleString()}</span> ימים האחרונים</li>
                                <li>כל <span className="stats-number">{avgJoinHours}</span> שעות מצטרפת עוד משתמשת</li>
                            </ul>
                        )}
                    />
                </li>
                <li className="stats-social-item">

                        <Image className="stats-social-logo" src={youtubeLogo} alt="Hatechnolog Youtube Channel"/>
                    </Link>
                    <LoadableComponent
                        isLoading={!latestStats.discord_stats.valid}
                        caption="טוען את הנתונים מיוטיוב..."
                        component={(
                            <ul className="stats-social-item-list">
                                <li><span className="stats-number">{props.youtube.subscribers.toLocaleString()}</span> רשומות ורשומים</li>
                                <li><span className="stats-number">{props.youtube.views.toLocaleString()}</span> צפיות בכל הקורסים</li>
                            </ul>
                        )}
                    />
                </li>
            </ul>
        </div>*/
import Link from 'next/link'
import Image from 'next/image'
import discordLogo from '@/public/social/discord_logo.svg'
import youtubeLogo from '@/public/social/youtube_logo_red.svg'
import goldenStar from '@/public/misc/golden_star.svg'
import githubLogo from '@/public/social/github_logo_blue.svg'
import { GetLatestStats } from '@/utils/rpc/stats.server'


export default async function Social() {
    const stats = await GetLatestStats();
    const avgJoinHours = (stats.discord.join_avg_sec / 60 / 60).toLocaleString();
    return (
        <div className="flex flex-col gap-10 items-center">
            <div className="flex flex-col gap-4 items-center">
                <Link href={process.env.YOUTUBE_CHANNEL_URL!} target="_blank"
                    className="flex flex-col w-fit items-center gap-2 bg-cta px-4 py-2 rounded-md">
                    <button className="flex gap-4 items-center">
                        <Image className="w-8 h-8" src={youtubeLogo} alt="Hatechnolog Discord"/>
                        <div className="flex flex-col gap-1 items-start">
                            <span className="text-lg font-bold">קורסי תכנות בחינם</span>
                            {stats.valid ?
                                <span>הצטרפו ל-<span className="text-lime-300 font-bold">{stats.youtube.subscribers.toLocaleString()}</span> רשומים</span>
                                :
                                <span>הירשמו לערוץ היוטיוב הקהילתי</span>
                            }
                        </div>
                    </button>
                </Link>
                {stats.valid &&
                <div className="flex flex-col gap-2">
                    <div className="flex gap-2">
                        <Image src={goldenStar} alt="Youtube Fact"/>
                        <span><span className="text-lime-300 font-bold">{stats.youtube.views.toLocaleString()}</span> צפיות בכל הקורסים</span>
                    </div>
                </div>
                }
            </div>
            <div className="flex flex-col gap-4 items-center">
                <Link href={process.env.DISCORD_INVITE_URL!} target="_blank"
                    className="flex flex-col w-fit items-center gap-2 bg-cta px-4 py-2 rounded-md">
                    <button className="flex gap-4 items-center">
                        <Image className="w-8 h-8" src={discordLogo} alt="Hatechnolog Discord"/>
                        <div className="flex flex-col gap-1 items-start">
                            <span className="text-lg font-bold">השרת הקהילתי</span>
                            {stats.valid ?
                                <span>הצטרפו ל-<span className="text-lime-300 font-bold">{stats.discord.total_members.toLocaleString()}</span> משתמשים</span>
                                :
                                <span>הצטרפו לשרת הדיסקורד הקהילתי</span>
                            }
                        </div>
                    </button>
                </Link>
                {stats.valid &&
                    <div className="flex flex-col gap-2">
                        <div className="flex gap-2">
                            <Image src={goldenStar} alt="Discord Fact"/>
                            <span>ב-<span className="text-lime-300 font-bold">{stats.discord.new_members_period_days.toLocaleString()}</span> ימים האחרונים הצטרפו <span className="text-lime-300 font-bold">{stats.discord.new_members.toLocaleString()}</span> משתמשים חדשים</span>
                        </div>
                        <div className="flex gap-2">
                            <Image src={goldenStar} alt="Discord Fact"/>
                            <span>מצטרף משתמש חדש לקהילה כל <span className="text-lime-300 font-bold">{avgJoinHours}</span> שעות</span>
                        </div>
                    </div>
                }
            </div>
            <div className="flex items-center gap-4">
                <Link 
                    href={process.env.MAIN_REPO_URL!}
                    target="_blank"
                >
                    <Image className="h-8 w-8" src={githubLogo} alt="Hatechnolog Github"/>
                </Link>
                <div>
                    <span>האתר הזה נבנה באהבה ע״י </span>
                    <Link
                        href={process.env.MAIN_REPO_CONTRIBUTORS_URL!}
                        target="_blank"
                        className="in-link"
                    >
                        {stats.github.contributors === 1 ?
                            <span>מפתח אחד</span>
                            :
                            <span>{`${stats.github.contributors} מפתחות ומפתחים`}</span>
                        }
                    </Link>
                    <span>, </span>
                    <Link
                        href={process.env.MAIN_REPO_COMMITS_URL!}
                        target="_blank"
                        className="in-link"
                    >{`ב-${stats.github.commits} קומיטים`}</Link>
                    <span>, תרמו קוד עוד היום!</span>
                </div>
            </div>
        </div>
    )
}
import Link from 'next/link'
import Image from 'next/image'
import { StatsProps } from '@/utils/rpc/stats'
import githubLogo from '@/public/social/github_logo_blue.svg'

export default function Credits(props: StatsProps) {
    return (
        <div className="flex items-center justify-center gap-4">
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
                    className="text-purple-300 font-bold"
                >
                    {props.github.contributors === 1 ?
                        <span>מפתח אחד</span>
                        :
                        <span>{`${props.github.contributors} מפתחות ומפתחים`}</span>
                    }
                </Link>
                <span>, </span>
                <Link
                    href={process.env.MAIN_REPO_COMMITS_URL!}
                    target="_blank"
                    className="text-purple-300 font-bold"
                >{`ב-${props.github.commits} קומיטים`}</Link>
                <span>, תרמו קוד עוד היום!</span>
            </div>
        </div>
    )
}
import React from "react";
import Image from "next/image";
import Link from "next/link";
import githubLogo from '../../public/social/github_logo.png'

interface GithubLatestStatsProps {
    contributors: number;
    commits: number;
}

const GithubLatestStats: React.FC<GithubLatestStatsProps> = ({
    contributors, commits
}: GithubLatestStatsProps) => {
    return (
        <div className="stats-github">
            <Link 
                className="stats-link"
                href={process.env.MAIN_REPO_URL!}
                rel="noopener noreferrer nofollow" 
                target="_blank"
            >
                <Image className="stats-github-logo" src={githubLogo} alt="Hatechnolog Github"/>
            </Link>
            <div>
                <span>האתר הזה נבנה ע״י </span>
                <Link
                    className="stats-link"
                    href={process.env.MAIN_REPO_CONTRIBUTORS_URL!}
                    rel="noopener noreferrer nofollow"
                    target="_blank"
                >
                    {contributors === 1 ?
                        <span>מפתח אחד</span>
                        :
                        <span>{`${contributors} מפתחות ומפתחים`}</span>
                    }
                </Link>
                <span>, </span>
                <Link
                    className="stats-link"
                    href={process.env.MAIN_REPO_COMMITS_URL!}
                    rel="noopener noreferrer nofollow"
                    target="_blank"
                >{`ב-${commits} קומיטים`}</Link>
            </div>
        </div>
    )
}

export default GithubLatestStats;
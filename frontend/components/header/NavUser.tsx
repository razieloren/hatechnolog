'use client'

import Link from "next/link"
import Image from "next/image"
import { IsLoggedIn } from "@/utils/cookies"
import { FetchMyUser } from "@/utils/rpc/user"
import discordLogo from "@/public/social/discord_logo_purple.svg"
import { useEffect, useState } from "react"

type LoggedUserProps = {
    loggedIn: boolean;
    handle: string;
    avater_url: string;
}

export default function NavUser() {
    const [loggedUser, setLoggedUser] = useState<LoggedUserProps>({
        loggedIn: false,
        handle: "",
        avater_url: "",
    })
    useEffect(() => {
        if (IsLoggedIn()) {
            FetchMyUser().then(user => {
                setLoggedUser({
                    loggedIn: true,
                    handle: user.handle,
                    avater_url: user.avatar_url,
                });
            }).catch(() => {
                // Maybe we just logged out.
            })
        }
    }, []);
    return (
        <>
        {loggedUser.loggedIn ?
        <>
        <span>Hey, {loggedUser.handle}</span>
        <Image src={loggedUser.avater_url} width={60} height={60} className="h-10 w-10" alt="User avatar"/>
        <Link
        href={`${process.env.API_URL!}/auth/logout?return=${encodeURIComponent(window.location.pathname)}`}
        rel="nofollow"
        >
            <button className="font-inherit p-1 text-gray-900 rounded-lg group bg-gradient-to-br from-teal-300 to-lime-300 group-hover:from-teal-300 group-hover:to-lime-300 focus:ring-4 focus:outline-none focus:ring-lime-200">
                <div className="flex items-center gap-2 py-2 px-4 text-white hover:text-black justify-center transition-all ease-in duration-75 bg-black rounded-md group-hover:bg-opacity-0">
                    <span className="font-bold sm:text-xl text-md">
                        התנתקו
                    </span>
                </div>
            </button>
        </Link>
        </>
        :
        <Link
            href={`${process.env.API_URL!}/auth/login/discord?return=${encodeURIComponent(window.location.pathname)}`}
            rel="nofollow"
        >
            <button className="font-inherit p-1 text-gray-900 rounded-lg group bg-gradient-to-br from-teal-300 to-lime-300 group-hover:from-teal-300 group-hover:to-lime-300 focus:ring-4 focus:outline-none focus:ring-lime-200">
                <div className="flex items-center gap-2 py-2 px-4 text-white hover:text-black justify-center transition-all ease-in duration-75 bg-black rounded-md group-hover:bg-opacity-0">
                    <span className="font-bold sm:text-xl text-md">
                        התחברו
                    </span>
                    <Image className="sm:w-6 sm:h-6 w-4 h-4" src={discordLogo} alt="Discord Logo"/>
                </div>
            </button>
        </Link>
        }
        </>
    )
}
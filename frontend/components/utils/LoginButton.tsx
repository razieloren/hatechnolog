'use client'

import Link from "next/link"
import Image from "next/image"
import { usePathname } from "next/navigation"
import discordLogo from "@/public/social/discord_logo.svg"

export default function LoginButton() {
    return (
        <Link
            href={`${process.env.API_URL!}/auth/login/discord?redirect=${encodeURIComponent(usePathname())}`}
            rel="nofollow"
        >
            <button className="font-inherit bg-cta rounded-lg">
                <div className="flex items-center gap-2 py-2 px-2 justify-center">
                    <Image className="w-4 h-4" src={discordLogo} alt="Discord Logo"/>
                    <span className="sm:block hidden font-bold text-md">
                        התחברו
                    </span>
                </div>
            </button>
        </Link>
    )
}
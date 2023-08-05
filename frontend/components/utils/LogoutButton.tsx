'use client'

import Link from "next/link"
import Image from "next/image"
import logoutIcon from "@/public/misc/logout.svg"
import { usePathname } from "next/navigation"

export default function LogoutButton() {
    return (
        <Link
        href={`${process.env.API_URL!}/auth/logout?redirect=${encodeURIComponent(usePathname())}`}
        rel="nofollow"
        >
            <button className="bg-red-600 rounded-full pr-2 pl-1 py-1 flex itens-center justify-center">
                <Image className="h-4 w-4" src={logoutIcon} alt="Hatechnolog Logout"/>
            </button>
        </Link>
    )
}
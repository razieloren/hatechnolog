import Link from "next/link"
import Image from "next/image"
import discordLogo from "@/public/social/discord_logo_purple.svg"

export default function NavUser() {
    return (
        <Link
            href={`${process.env.AUTH_API_URL!}/login/discord`}
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
    )
}
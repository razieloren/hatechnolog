import Link from "next/link"
import Image from "next/image"
import { GetMyUser } from "@/utils/rpc/user.server"
import {messages as user_api} from "@/messages/user/api"
import LogoutButton from "@/components/utils/LogoutButton"
import LoginButton from "@/components/utils/LoginButton"
import externalLinkLogo from "@/public/misc/external_link.svg"
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import {DropdownMenu, DropdownMenuTrigger, DropdownMenuContent, DropdownMenuLabel, DropdownMenuSeparator, DropdownMenuItem, DropdownMenuGroup} from '@/components/ui/dropdown-menu'

export default async function NavUser() {
    let isLoggedIn = false;
    let user = new user_api.GetUserResponse();
    try {
        user = await GetMyUser();
        isLoggedIn = true;
    } catch (e) {
        // Maybe we must logged out...
    }

    return (
        <>
        {isLoggedIn ?
        <div className="flex gap-4 items-center">
            <DropdownMenu>
                <DropdownMenuTrigger asChild>
                    <Avatar>
                        <AvatarImage src={user.avatar_url} />
                        <AvatarFallback>{user.handle[0].toUpperCase()}</AvatarFallback>
                    </Avatar>
                </DropdownMenuTrigger>
                <DropdownMenuContent className="w-56 flex flex-col items-center">
                    <DropdownMenuLabel>
                        <Link className="text-words flex gap-1 items-center" href={`/users/${user.handle}`} target='_blank' >
                            {user.handle.toUpperCase()}
                            <Image src={externalLinkLogo} className="h-6 w-6" alt="Profile"/>
                        </Link>
                    </DropdownMenuLabel>
                    <DropdownMenuSeparator />
                    <DropdownMenuGroup>
                        <DropdownMenuItem>
                            <LogoutButton/>
                        </DropdownMenuItem>
                    </DropdownMenuGroup>
                </DropdownMenuContent>
            </DropdownMenu>
        </div>
        :
        <LoginButton/>
        }
        </>
    )
}
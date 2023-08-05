import Image from 'next/image';
import NavUser from "./NavUser";
import menuIcon from "@/public/misc/menu_icon.svg";
import { NavigationMenu, NavigationMenuList, NavigationMenuItem, NavigationMenuTrigger, NavigationMenuContent, NavigationMenuLink } from '@/components/ui/navigation-menu';
import Link from 'next/link';

type NavbarItem = {
    title: string
    href: string;
}

const NarBarItems: NavbarItem[] = [
    {
        title: "בלוג",
        href: "/blog"
    },
    {
        title: "קורסים",
        href: "/courses"
    },
    {
        title: "צרו קשר",
        href: "/contact"
    },
]

export default function Navbar() {
    return (
        <nav className="flex justify-between items-center w-full"> 
            <div className="sm:flex sm:gap-10 gap-8 text-l hidden items-center">
                {NarBarItems.map(item => {
                    return (
                        <Link href={item.href}>
                            <button className="hover:text-white font-bold">
                            {item.title}
                            </button>
                        </Link>
                    )
                })}
            </div>
            <div className="sm:invisible visible">
                <NavigationMenu>
                    <NavigationMenuList>
                        <NavigationMenuItem>
                        <NavigationMenuTrigger>
                            <Image src={menuIcon} className="h-6 w-6" alt="Hatechnolog Menu"/>
                        </NavigationMenuTrigger>
                        <NavigationMenuContent>
                            <ul className="flex flex-col gap-4 text-center px-4 py-2">
                                {NarBarItems.map(item => {
                                    return (
                                        <Link href={item.href}>
                                            <li key={item.title} className="hover:text-white text-l font-bold text-words w-16">
                                            {item.title}
                                            </li>
                                        </Link>
                                    )
                                })}
                            </ul>
                        </NavigationMenuContent>
                        </NavigationMenuItem>
                    </NavigationMenuList>
                </NavigationMenu>
            </div>
            <NavUser/>
        </nav>
    )
}
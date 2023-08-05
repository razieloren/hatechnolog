import Link from 'next/link'
import Image from 'next/image'
import mainLogo from '@/public/hatechnolog_logo.svg'
import Navbar from './Navbar'
import { Separator } from '@/components/ui/separator'

export default function Header() {
    return (
        <div className="flex flex-col items-right gap-4 py-4">
            <div className="flex gap-8">
                <Link href="/">
                    <Image className="h-16 w-16" src={mainLogo} alt="Hatechnolog Logo"/>
                </Link>
                <Navbar/>
            </div>
            <Separator/>
        </div>

    )
}

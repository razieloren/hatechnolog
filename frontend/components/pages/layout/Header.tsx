import Link from 'next/link'
import Image from 'next/image'
import mainLogo from '@/public/hatechnolog_logo.svg'
import Navbar from './Navbar'
import { Separator } from '@/components/ui/separator'

export default function Header() {
    return (
        <div className="flex flex-col items-right gap-4 py-4">
            <div itemScope itemType='https://schema.org/Organization' className="flex gap-8">
                <Link itemProp='url' href="/">
                    <Image itemProp='logo' className="h-16 w-16" src={mainLogo} alt="Hatechnolog Logo"/>
                </Link>
                <span className='hidden' itemProp='legalName' content="הטכנולוג"/>
                <span className='hidden' itemProp='name' content="הטכנולוג"/>
                <Navbar/>
            </div>
            <Separator/>
        </div>

    )
}

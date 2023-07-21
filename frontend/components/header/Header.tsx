import Link from 'next/link'
import Image from 'next/image'
import mainLogo from '@/public/hatechnolog_logo.svg'
import Navbar from './Navbar'

export default function Header() {
    return (
        <div className="w-full responsive-flex">
            <div className="basis-1/3">
                <Link className="flex sm:gap-4 gap-2 items-center" href="/">
                    <Image className="sm:h-20 sm:w-20 h-12 w-12" src={mainLogo} alt="Hatechnolog Logo"/>
                    <span className="sm:text-6xl text-4xl font-bold">הטכנולוג</span>
                </Link>
            </div>
            <div className="basis-2/3">
                <Navbar/>
            </div>
        </div>
    )
}
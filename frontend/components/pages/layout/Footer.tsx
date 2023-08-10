import Link from 'next/link'
import Image from 'next/image'
import twitterLogo from '@/public/social/twitter_logo_blue.svg'
import patreonLogo from '@/public/social/patreon_logo.svg'
import sitemapLogo from '@/public/misc/sitemap.svg'

export default function Footer() {
    return (
        <div className="flex gap-2 items-center justify-center">
            <Link
                className="in-link"
                href={process.env.TWITTER_PROFILE_URL!}
                target="_blank"
                >
                <Image src={twitterLogo} className='w-6 h-6' alt="Hatechnolog Twitter"/>
            </Link>
            <span>|</span>
            <Link
                className="in-link"
                href={process.env.PATREON_URL!}
                target="_blank"
                >
                <Image src={patreonLogo} className='w-6 h-6' alt="Hatechnolog Patreon"/>
            </Link>
            <span>|</span>
            <Link
                className="in-link"
                href={`${process.env.API_URL!}/sitemap.xml`}
                target="_blank"
                >
                <Image src={sitemapLogo} className='w-6 h-6' alt="Hatechnolog Sitemap"/>
            </Link>
            <span>|</span>
            <Link className="in-link" href="/terms">
                תנאים
            </Link>
            <span>|</span>
            <Link className="in-link" href="/privacy">
                פרטיות
            </Link>
        </div>
    )
}
import { GetCoursesTeasers } from '@/utils/rpc/courses.server'
import { Metadata } from 'next';
import Image from 'next/image';
import Link from 'next/link';

export const metadata: Metadata = {
    metadataBase: new URL("https://hatechnolog.com"),
    title: `הטכנולוג - קורסים`,
    description: `הקורסים המצולמים של קהילת הטכנולוג - ללמוד לתכנת לבד, עם קהילה שלמה, וכמובן - בחינם.`,
    publisher: "הטכנולוג",
    alternates: {
        canonical: `/courses`
    },
    openGraph: {
        title: `הטכנולוג - קורסים`,
        description: `הקורסים המצולמים של קהילת הטכנולוג - ללמוד לתכנת לבד, עם קהילה שלמה, וכמובן - בחינם.`,
        url: "https://hatechnolog.com/courses",
        siteName: "הטכנולוג",
        images: {
            url: "https://hatechnolog.fra1.cdn.digitaloceanspaces.com/technolog_logo.png",
            secureUrl: "https://hatechnolog.fra1.cdn.digitaloceanspaces.com/technolog_logo.png",
            alt: "Hatechnolog Logo",
            width: 256,
            height: 256,
        }
    },
    twitter: {
        site: "https://hatechnolog.com",
        creator: "@hatechnolog",
        description: `הקורסים המצולמים של קהילת הטכנולוג - ללמוד לתכנת לבד, עם קהילה שלמה, וכמובן - בחינם.`,
        title: `הטכנולוג - קורסים`,
        images: {
            url: "https://hatechnolog.fra1.cdn.digitaloceanspaces.com/technolog_logo.png",
            secureUrl: "https://hatechnolog.fra1.cdn.digitaloceanspaces.com/technolog_logo.png",
            alt: "Hatechnolog Logo",
            width: 256,
            height: 256,
        }
    }
}

export default async function CoursesPage() {
    const teasers = await GetCoursesTeasers();
    return (
        <div className="flex flex-col gap-8 items-center text-center justify-center">
            <h1 className="text-xl font-bold">הקורסים של הטכנולוג</h1>
            <div className="grid grid-cols-1 xl:grid-cols-2 gap-8">
                {teasers.map(teaser => {
                    return (
                        <div key={teaser.slug} itemScope itemType='https://schema.org/Course'>
                            <Link itemProp='url' className="relative hover:drop-shadow-xl" href={`/courses/${teaser.slug}`} target="_blank">
                                <img itemProp='image' src={`https://i.ytimg.com/vi/${teaser.main_video_id}/maxresdefault.jpg`} width={480} height={270} alt={teaser.title}/>
                                <div className="absolute top-0 left-auto w-full px-3 py-5 bg-opacity-70 bg-black">
                                    <h2 itemProp='name' className="font-bold text-lg">{teaser.title}</h2>
                                </div>
                            </Link>
                            <span hidden itemProp='description' content={teaser.description}/>
                            <span hidden itemProp='provider' itemScope itemType='https://schema.org/Person'>
                                <span hidden itemProp='name' content="raznick"/>
                                <span hidden itemProp='url' content="https://hatechnolog.com/users/raznick"/>
                            </span>
                        </div>
                    )
                })}
            </div>
        </div>
    )
}
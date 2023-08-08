import { GetCoursesTeasers } from '@/utils/rpc/courses.server'
import { Metadata } from 'next';
import Image from 'next/image';
import Link from 'next/link';

export const metadata: Metadata = {
    metadataBase: new URL("https://hatechnolog.com"),
    title: `הטכנולוג - קורסים`,
    description: `בלוג הטכנולוג - קורסי תכנות בחינם`,
    publisher: "הטכנולוג",
    alternates: {
        canonical: `/courses`
    },
    openGraph: {
        title: `הטכנולוג - קורסים`,
        url: "https://hatechnolog.com/courses",
        siteName: "הטכנולוג",
    },
}

export default async function CoursesPage() {
    const teasers = await GetCoursesTeasers();
    return (
        <div className="flex flex-col gap-8 items-center text-center justify-center">
            <h1 className="text-xl font-bold">הקורסים של הטכנולוג</h1>
            <div className="grid grid-cols-1 xl:grid-cols-2 gap-8">
                {teasers.map(teaser => {
                    return (
                        <Link key={teaser.slug} className="relative hover:drop-shadow-xl" href={`/courses/${teaser.slug}`} target="_blank">
                            <img src={`https://i.ytimg.com/vi/${teaser.main_video_id}/maxresdefault.jpg`} width={480} height={270} alt={teaser.title}/>
                            <div className="absolute top-0 left-auto w-full px-3 py-5 bg-opacity-70 bg-black">
                                <h2 className="font-bold text-lg">{teaser.title}</h2>
                            </div>
                        </Link>
                    )
                })}
            </div>
        </div>
    )
}
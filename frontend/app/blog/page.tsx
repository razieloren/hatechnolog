import { GetPostsTeasers } from '@/utils/rpc/content.server';
import { Metadata } from 'next';
import Link from 'next/link';

export const metadata: Metadata = {
    metadataBase: new URL("https://hatechnolog.com"),
    title: `הטכנולוג - פוסטים`,
    description: `בלוג הטכנולוג - פוסטים טכנולוגים`,
    publisher: "הטכנולוג",
    alternates: {
        canonical: `/blog`
    },
    openGraph: {
        title: `הטכנולוג - פוסטים`,
        url: "https://hatechnolog.com/blog",
        siteName: "הטכנולוג",
    },
}

export default async function PostsPage() {
    const teasers = await GetPostsTeasers();
    return (
        <div className="flex flex-col gap-4">
            <div>
                <h1 className="font-bold text-xl">כל הפוסטים</h1>
                <Link className="in-link" href="category" target="_blank">
                    <span>(לכל הקטגוריות)</span>
                </Link>
            </div>
            {teasers.map(teaser => {
                return (
                    <div key={teaser.slug} className="flex flex-col gap-1 py-2 px-4 bg-secondary rounded-lg w-fit">
                        <Link className="in-link text-lg" href={`/blog/${teaser.slug}`} target="_blank">
                            <h2>{teaser.title}</h2>
                        </Link>
                        <div className="flex gap-2 text-sm">
                            <Link href={`/users/${teaser.author}`} target="_blank">
                                <span>{teaser.author}</span>
                            </Link>
                            <span>|</span>
                            <Link href={`/category/${teaser.category}`} target="_blank">
                                <span>{teaser.category}</span>
                            </Link>
                            <span>|</span>
                            <span>{new Date(teaser.published.seconds * 1000).toLocaleDateString()}</span>
                        </div>
                    </div>
                )
            })}
        </div>
    )
}
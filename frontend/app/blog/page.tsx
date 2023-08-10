import { GetPostsTeasers } from '@/utils/rpc/content.server';
import { Metadata } from 'next';
import Link from 'next/link';

export const metadata: Metadata = {
    metadataBase: new URL("https://hatechnolog.com"),
    title: `הטכנולוג - פוסטים`,
    description: `הבלוג של קהילת הטכנולוג - כל הפוסטים מכל הקטגוריות. התחילו עוד היום ללמוד לתכנת לבד ולחקור לבד, וגם לכתוב פוסטים בעצמכם!`,
    publisher: "הטכנולוג",
    alternates: {
        canonical: `/blog`
    },
    openGraph: {
        title: `הטכנולוג - פוסטים`,
        description: `הבלוג של קהילת הטכנולוג - כל הפוסטים מכל הקטגוריות. התחילו עוד היום ללמוד לתכנת לבד ולחקור לבד, וגם לכתוב פוסטים בעצמכם!`,
        url: "https://hatechnolog.com/blog",
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
        description: `הבלוג של קהילת הטכנולוג - כל הפוסטים מכל הקטגוריות. התחילו עוד היום ללמוד לתכנת לבד ולחקור לבד, וגם לכתוב פוסטים בעצמכם!`,
        title: `הטכנולוג - פוסטים`,
        images: {
            url: "https://hatechnolog.fra1.cdn.digitaloceanspaces.com/technolog_logo.png",
            secureUrl: "https://hatechnolog.fra1.cdn.digitaloceanspaces.com/technolog_logo.png",
            alt: "Hatechnolog Logo",
            width: 256,
            height: 256,
        }
    }
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
                    <div itemScope itemType='https://schema.org/NewsArticle' key={teaser.slug} className="flex flex-col gap-1 py-2 px-4 bg-secondary rounded-lg w-fit">
                        <Link itemProp='url' className="in-link text-lg" href={`/blog/${teaser.slug}`} target="_blank">
                            <h2 itemProp='headline'>{teaser.title}</h2>
                        </Link>
                        <div className="flex gap-2 text-sm">
                            <span itemProp='author' itemScope itemType='https://schema.org/Person'>
                                <Link itemProp='url' href={`/users/${teaser.author}`} target="_blank">
                                    <span itemProp='name'>{teaser.author}</span>
                                </Link>
                            </span>
                            <span>|</span>
                            <Link href={`/category/${teaser.category}`} target="_blank">
                                <span itemProp='genre'>{teaser.category}</span>
                            </Link>
                            <span>|</span>
                            <span itemProp='datePublished' content={new Date(teaser.published.seconds * 1000).toISOString()}>{new Date(teaser.published.seconds * 1000).toLocaleDateString()}</span>
                            <span className='hidden' itemProp='dateModified' content={new Date(teaser.edited.seconds * 1000).toISOString()}/>
                        </div>
                    </div>
                )
            })}
        </div>
    )
}
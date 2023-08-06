import { GetCategoriesTeasers } from '@/utils/rpc/content.server';
import { Metadata } from 'next';
import Link from 'next/link';

export const metadata: Metadata = {
    metadataBase: new URL("https://hatechnolog.com"),
    title: `הטכנולוג - קטגוריות`,
    description: `בלוג הטכנולוג - קטגוריות`,
    publisher: "הטכנולוג",
    alternates: {
        canonical: `/category`
    },
    openGraph: {
        title: `הטכנולוג - קטגוריות`,
        url: "https://hatechnolog.com/category",
        siteName: "הטכנולוג",
    },
}

export default async function CategoriesPage() {
    const teasers = await GetCategoriesTeasers();
    return (
        <div className="flex flex-col gap-4">
            <h1 className="font-bold text-xl">כל הקטגוריות</h1>
            {teasers.map(teaser => {
                return (
                    <Link key={teaser.slug} className="flex gap-2" href={`/category/${teaser.slug}`} target="_blank">
                        <h2 className="in-link">{teaser.name}</h2>
                        <span>-</span>
                        <span>{teaser.description}</span>
                    </Link>
                )
            })}
        </div>
    )
}
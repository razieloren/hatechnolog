import { GetCategoriesTeasers } from '@/utils/rpc/content.server';
import { Metadata } from 'next';
import Link from 'next/link';

export const metadata: Metadata = {
    metadataBase: new URL("https://hatechnolog.com"),
    title: `הטכנולוג - קטגוריות`,
    description: `כל קטגוריות הפוסטים של קהילת הטכנולוג, ללמוד לתכנת לבד מעולם לא הייתה משימה פשוטה יותר, היכנסו לעולם התוכן הטכנולוגי שאתם הכי אוהבים!`,
    publisher: "הטכנולוג",
    alternates: {
        canonical: `/category`
    },
    openGraph: {
        title: `הטכנולוג - קטגוריות`,
        description: `כל קטגוריות הפוסטים של קהילת הטכנולוג, ללמוד לתכנת לבד מעולם לא הייתה משימה פשוטה יותר, היכנסו לעולם התוכן הטכנולוגי שאתם הכי אוהבים!`,
        url: "https://hatechnolog.com/category",
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
        description: `כל קטגוריות הפוסטים של קהילת הטכנולוג, ללמוד לתכנת לבד מעולם לא הייתה משימה פשוטה יותר, היכנסו לעולם התוכן הטכנולוגי שאתם הכי אוהבים!`,
        title: `הטכנולוג - קטגוריות`,
        images: {
            url: "https://hatechnolog.fra1.cdn.digitaloceanspaces.com/technolog_logo.png",
            secureUrl: "https://hatechnolog.fra1.cdn.digitaloceanspaces.com/technolog_logo.png",
            alt: "Hatechnolog Logo",
            width: 256,
            height: 256,
        }
    }
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
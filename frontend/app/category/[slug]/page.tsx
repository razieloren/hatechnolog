import { GetCategoriesTeasers, GetCategory } from "@/utils/rpc/content.server";
import { Metadata, ResolvedMetadata } from "next";
import Link from "next/link";

type CategoryPageProps = {
    params: {
        slug: string;
    };
}

export async function generateMetadata(
    { params }: CategoryPageProps,
    parent: ResolvedMetadata
  ): Promise<Metadata> {
    const content = await GetCategory(params.slug);   
    return {
        metadataBase: new URL("https://hatechnolog.com"),
        title: `הטכנולוג - ${content.teaser.name}`,
        description: `כל הפוסטים מקטגוריית ${content.teaser.name}`,
        publisher: "הטכנולוג",
        alternates: {
            canonical: `/category/${content.teaser.slug}`
        },
        openGraph: {
            title: `הטכנולוג - ${content.teaser.name}`,
            description: `כל הפוסטים מקטגוריית ${content.teaser.name}`,
            url: `https://hatechnolog.com/category/${content.teaser.slug}`,
            siteName: "הטכנולוג",
        },
    }
}

export async function generateStaticParams() {
    const teasers = await GetCategoriesTeasers()
    return teasers.map(teaser => ({
        slug: teaser.slug,
    }))
}

export default async function PostPage(props: CategoryPageProps) {
    const details = await GetCategory(props.params.slug);
    return (
        <div className="flex flex-col gap-2">
            <h1 className="font-bold text-xl">{details.teaser.name}</h1>
            <p className="text-md">{details.teaser.description}</p>
            <div className="my-4 flex flex-col gap-4">
                {details.contents.map(content => {
                    return (
                        <Link key={content.slug} className="in-link" href={`/blog/${content.slug}`} target="_blank">
                            <h1>{content.title}</h1>
                        </Link>
                    )
                })}
            </div>
        </div>
    )
}
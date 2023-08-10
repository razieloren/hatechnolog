import { GetCategory } from "@/utils/rpc/content.server";
import { Metadata } from "next";
import Link from "next/link";

type CategoryPageProps = {
    params: {
        slug: string;
    };
}

export async function generateMetadata(
    { params }: CategoryPageProps,
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
            description: content.teaser.description,
            url: `https://hatechnolog.com/category/${content.teaser.slug}`,
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
            description: content.teaser.description,
            title: `הטכנולוג - ${content.teaser.name}`,
            images: {
                url: "https://hatechnolog.fra1.cdn.digitaloceanspaces.com/technolog_logo.png",
                secureUrl: "https://hatechnolog.fra1.cdn.digitaloceanspaces.com/technolog_logo.png",
                alt: "Hatechnolog Logo",
                width: 256,
                height: 256,
            }
        }
    }
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
                        <div key={content.slug} className="flex gap-2" itemScope itemType="https://schema.org/NewsArticle">
                            <Link itemProp="url" className="in-link" href={`${content.type === "posts" ? "/blog" : ""}/${content.slug}`} target="_blank">
                                <h1 itemProp="headline">{content.title}</h1>
                            </Link>
                            <span>מאת</span>
                            <span itemProp="author" itemScope itemType="https://schema.org/Person">
                                <Link itemProp="url" className="in-link" href={`/users/${content.author}`} target="_blank">
                                    <h1 itemProp="name">{content.author}</h1>
                                </Link>
                            </span>
                        </div>
                    )
                })}
            </div>
        </div>
    )
}
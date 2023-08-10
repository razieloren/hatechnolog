import Content from "@/components/pages/index/Content";
import { GetPost } from "@/utils/rpc/content.server";
import { Metadata } from "next";

type PostPageProps = {
    params: {
        slug: string;
    };
}

export async function generateMetadata(
    { params }: PostPageProps,
  ): Promise<Metadata> {
    const content = await GetPost(params.slug);   
    return {
        metadataBase: new URL("https://hatechnolog.com"),
        title: `הטכנולוג - ${content.teaser.title}`,
        description: content.teaser.description,
        keywords: [content.teaser.category],
        authors: [{name: content.teaser.author, url: `/users/${content.teaser.author}`}],
        creator: content.teaser.author,
        publisher: "הטכנולוג",
        alternates: {
            canonical: `/blog/${content.teaser.slug}`
        },
        openGraph: {
            title: `הטכנולוג - ${content.teaser.title}`,
            description: content.teaser.description,
            url: `https://hatechnolog.com/blog/${content.teaser.slug}`,
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
            title: `הטכנולוג - ${content.teaser.title}`,
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

export default async function PostPage(props: PostPageProps) {
    const details = await GetPost(props.params.slug);
    return <Content details={details} />
}
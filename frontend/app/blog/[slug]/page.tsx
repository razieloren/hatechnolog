import Content from "@/components/pages/index/Content";
import { GetPost, GetPostsTeasers } from "@/utils/rpc/content.server";
import { Metadata, ResolvedMetadata } from "next";

type PostPageProps = {
    params: {
        slug: string;
    };
}

export async function generateMetadata(
    { params }: PostPageProps,
    parent: ResolvedMetadata
  ): Promise<Metadata> {
    const content = await GetPost(params.slug);   
    return {
        metadataBase: new URL("https://hatechnolog.com"),
        title: `הטכנולוג - ${content.teaser.title}`,
        description: `${content.teaser.title} by ${content.teaser.author}`,
        keywords: [content.teaser.category],
        authors: [{name: content.teaser.author, url: `/users/${content.teaser.author}`}],
        creator: content.teaser.author,
        publisher: "הטכנולוג",
        alternates: {
            canonical: `/blog/${content.teaser.slug}`
        },
        openGraph: {
            title: `הטכנולוג - ${content.teaser.title}`,
            description: `${content.teaser.title} by ${content.teaser.author}`,
            url: `https://hatechnolog.com/blog/${content.teaser.slug}`,
            siteName: "הטכנולוג",
        },
    }
}

export async function generateStaticParams() {
    const teasers = await GetPostsTeasers()
    return teasers.map(teaser => ({
        slug: teaser.slug,
    }))
}

export default async function PostPage(props: PostPageProps) {
    const details = await GetPost(props.params.slug);
    return <Content details={details} />
}
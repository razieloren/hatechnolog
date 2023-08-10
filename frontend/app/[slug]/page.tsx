import Content from '@/components/pages/index/Content';
import { GetPage } from '@/utils/rpc/content.server';
import { Metadata } from 'next';

type AppPageProps = {
    params: {
        slug: string;
    };
}

export async function generateMetadata(
    { params }: AppPageProps,
  ): Promise<Metadata> {
    const content = await GetPage(params.slug);   
    return {
        metadataBase: new URL("https://hatechnolog.com"),
        title: `הטכנולוג - ${content.teaser.title}`,
        description: content.teaser.description,
        keywords: [content.teaser.category],
        authors: [{name: content.teaser.author, url: `/users/${content.teaser.author}`}],
        creator: content.teaser.author,
        publisher: "הטכנולוג",
        alternates: {
            canonical: `/${content.teaser.slug}`
        },
        openGraph: {
            title: `הטכנולוג - ${content.teaser.title}`,
            description: content.teaser.description,
            url: `https://hatechnolog.com/${content.teaser.slug}`,
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

export default async function AppPage(props: AppPageProps) {
    const content = await GetPage(props.params.slug);
    return <Content details={content}/>
}
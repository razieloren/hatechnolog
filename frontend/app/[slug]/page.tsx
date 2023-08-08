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
        description: `${content.teaser.title} by ${content.teaser.author}`,
        keywords: [content.teaser.category],
        authors: [{name: content.teaser.author, url: `/users/${content.teaser.author}`}],
        creator: content.teaser.author,
        publisher: "הטכנולוג",
        alternates: {
            canonical: `/${content.teaser.slug}`
        },
        openGraph: {
            title: `הטכנולוג - ${content.teaser.title}`,
            description: `${content.teaser.title} by ${content.teaser.author}`,
            url: `https://hatechnolog.com/${content.teaser.slug}`,
            siteName: "הטכנולוג",
        },
    }
}

export default async function AppPage(props: AppPageProps) {
    const content = await GetPage(props.params.slug);
    return <Content details={content}/>
}
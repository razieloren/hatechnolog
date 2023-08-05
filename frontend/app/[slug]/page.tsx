import { GetPage } from '@/utils/rpc/content.server';
import Link from 'next/link'
import { inflate } from 'pako';

type AppPageProps = {
    params: {
        slug: string;
    };
}

export default async function AppPage(props: AppPageProps) {
    const content = await GetPage(props.params.slug);
    const direction = content.ltr ? "ltr" : "rtl";
    const publishStr = new Date(content.teaser.published.seconds * 1000).toLocaleDateString();
    const editStr = new Date(content.edited.seconds * 1000).toLocaleDateString();
    const showEdited = publishStr !== editStr;
    const pageContent = new TextDecoder().decode(inflate(content.compressed_content));
    return (
        <div dir={direction} className="flex flex-col gap-6">
            <div className="flex flex-col gap-2 items-center self-center">
                <h1 className="text-4xl">{content.teaser.title}</h1>
                <div>
                    <span>פורסם: {publishStr}</span>
                    {showEdited &&
                    <span> (נערך: {editStr})</span>}
                    <Link href={`/users/${content.teaser.author}`} target="_blank">
                        <span>, {content.teaser.author}</span>
                    </Link>
                </div>
            </div>
            <div className="md-file flex flex-col gap-4 text-md" dangerouslySetInnerHTML={{__html: pageContent}}/>
        </div>
    )
}
import {messages as helpers} from '@/messages/content/helpers'
import Link from 'next/link';
import { inflate } from 'pako';

type ContentProps = {
    details: helpers.ContentDetails
}

export default function Content(props: ContentProps) {
    const content = props.details
    const direction = content.ltr ? "ltr" : "rtl";
    const publishStr = new Date(content.teaser.published.seconds * 1000).toLocaleDateString();
    const pageContent = new TextDecoder().decode(inflate(content.compressed_content));
    return (
        <div dir={direction} className="flex flex-col gap-6">
            <div className="flex flex-col gap-2 items-center self-center">
                <h1 className="text-4xl">{content.teaser.title}</h1>
                <div dir="ltr" className="flex gap-2">
                    <span>{publishStr}</span>
                    <span>|</span>
                    <Link href={`/users/${content.teaser.author}`} target="_blank">
                        <span>{content.teaser.author}</span>
                    </Link>
                    <span>|</span>
                    <Link href={`/category/${content.teaser.category}`} target="_blank">
                        <span>{content.teaser.category}</span>
                    </Link>
                </div>
            </div>
            <div className="md-file flex flex-col gap-4 text-md" dangerouslySetInnerHTML={{__html: pageContent}}/>
        </div>
    )
}
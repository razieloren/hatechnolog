import {messages as helpers} from '@/messages/content/helpers'
import Link from 'next/link';
import { inflate } from 'pako';

type ContentProps = {
    details: helpers.ContentDetails
}

export default function Content(props: ContentProps) {
    const content = props.details
    const direction = content.ltr ? "ltr" : "rtl";
    const edited = new Date(content.teaser.edited.seconds * 1000);
    const publish = new Date(content.teaser.published.seconds * 1000);
    const pageContent = new TextDecoder().decode(inflate(content.compressed_content));
    return (
        <div dir={direction} className="flex flex-col gap-6" itemScope itemType='https://schema.org/NewsArticle'>
            <div className="flex flex-col gap-2 items-center self-center">
                <h1 className="text-4xl" itemProp='headline'>{content.teaser.title}</h1>
                <div dir="ltr" className="flex gap-2">
                    <span itemProp='datePublished' content={publish.toISOString()}>{publish.toLocaleDateString()}</span>
                    <span itemProp='dateModified' content={edited.toISOString()}/>
                    <span>|</span>
                    <span itemProp='author' itemScope itemType='https://schema.org/Person'>
                        <Link itemProp='url' href={`/users/${content.teaser.author}`} target="_blank">
                            <span itemProp='name'>{content.teaser.author}</span>
                        </Link>
                    </span>
                    <span>|</span>
                    <Link href={`/category/${content.teaser.category}`} target="_blank">
                        <span itemProp='genre'>{content.teaser.category}</span>
                    </Link>
                </div>
            </div>
            <div itemProp='articleBody' className="md-file flex flex-col gap-4 text-md" dangerouslySetInnerHTML={{__html: pageContent}}/>
        </div>
    )
}
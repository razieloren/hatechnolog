import { GetCategory } from "@/utils/rpc/content.server";
import Link from "next/link";

type CategoryPageProps = {
    params: {
        slug: string;
    };
}

export default async function PostPage(props: CategoryPageProps) {
    const details = await GetCategory(props.params.slug);
    return (
        <div>
            <h1>{details.teaser.name}</h1>
            {details.contents.map(content => {
                return (
                    <Link href={`/blog/${content.slug}`} target="_blank">
                        <h1>{content.title}</h1>
                    </Link>
                )
            })}
        </div>
    )
}
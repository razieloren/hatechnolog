import { GetPost } from "@/utils/rpc/content.server";

type PostPageProps = {
    params: {
        slug: string;
    };
}

export default async function PostPage(props: PostPageProps) {
    const details = await GetPost(props.params.slug);
    return (
        <h1>{details.teaser.title} | {details.teaser.category}</h1>
    )
}
import { GetPostsTeasers } from '@/utils/rpc/content.server';
import Link from 'next/link';

export default async function PostsPage() {
    const teasers = await GetPostsTeasers();
    return (
        <div>
            {teasers.map(teaser => {
                return (
                    <Link href={`/blog/${teaser.slug}`} target="_blank">
                        <h1>{teaser.title}</h1>
                    </Link>
                )
            })}
        </div>
    )
}
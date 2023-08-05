import { GetCategoriesTeasers } from '@/utils/rpc/content.server';
import Link from 'next/link';

export default async function CategoriesPage() {
    const teasers = await GetCategoriesTeasers();
    return (
        <div>
            {teasers.map(teaser => {
                return (
                    <Link href={`/blog/category/${teaser.slug}`} target="_blank">
                        <h1>{teaser.name}</h1>
                    </Link>
                )
            })}
        </div>
    )
}
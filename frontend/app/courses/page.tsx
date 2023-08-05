import { GetCoursesTeasers } from '@/utils/rpc/courses.server'
import Link from 'next/link';

export default async function CoursesPage() {
    const teasers = await GetCoursesTeasers();
    return (
        <div>
            {teasers.map(teaser => {
                return (
                    <Link href={`/courses/${teaser.slug}`} target="_blank">
                        <h1>{teaser.title}</h1>
                    </Link>
                )
            })}
        </div>
    )
}
import { GetCourse } from "@/utils/rpc/courses.server";

type CoursePageProps = {
    params: {
        slug: string;
    };
}

export default async function CoursePage(props: CoursePageProps) {
    const details = await GetCourse(props.params.slug);
    return (
        <h1>{details.teaser.title} | {details.playlist_id}</h1>
    )
}
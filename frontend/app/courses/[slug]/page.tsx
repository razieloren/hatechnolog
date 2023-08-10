import { GetCourse } from "@/utils/rpc/courses.server";
import { Separator } from '@/components/ui/separator'
import { Metadata } from "next";

type CoursePageProps = {
    params: {
        slug: string;
    };
}

export async function generateMetadata(
    { params }: CoursePageProps,
  ): Promise<Metadata> {
    const content = await GetCourse(params.slug);   
    return {
        metadataBase: new URL("https://hatechnolog.com"),
        title: `הטכנולוג - ${content.teaser.title}`,
        description: content.teaser.description,
        publisher: "הטכנולוג",
        alternates: {
            canonical: `/courses/${content.teaser.slug}`
        },
        openGraph: {
            title: `הטכנולוג - ${content.teaser.title}`,
            description: content.teaser.description,
            url: `https://hatechnolog.com/courses/${content.teaser.slug}`,
            siteName: "הטכנולוג",
            images: {
                url: "https://hatechnolog.fra1.cdn.digitaloceanspaces.com/technolog_logo.png",
                secureUrl: "https://hatechnolog.fra1.cdn.digitaloceanspaces.com/technolog_logo.png",
                alt: "Hatechnolog Logo",
                width: 256,
                height: 256,
            }
        },
        twitter: {
            site: "https://hatechnolog.com",
            creator: "@hatechnolog",
            description: content.teaser.description,
            title: `הטכנולוג - ${content.teaser.title}`,
            images: {
                url: "https://hatechnolog.fra1.cdn.digitaloceanspaces.com/technolog_logo.png",
                secureUrl: "https://hatechnolog.fra1.cdn.digitaloceanspaces.com/technolog_logo.png",
                alt: "Hatechnolog Logo",
                width: 256,
                height: 256,
            }
        }
    }
}

export default async function CoursePage(props: CoursePageProps) {
    const details = await GetCourse(props.params.slug);
    return (
        <div itemScope itemType='https://schema.org/Course' className="flex flex-col gap-4 items-center text-center justify-center">
            <h1 itemProp="name" className="font-bold text-xl">{details.teaser.title}</h1>
            <h2 itemProp="description" className="texl-lg">{details.teaser.description}</h2>
            <Separator className="w-1/4 bg-cta"/>
            <iframe
                className="md:w-96 w-64 aspect-video rounded-lg p-1 my-4 bg-cta"
                src={`https://www.youtube.com/embed/videoseries?list=${details.playlist_id}`}
                title={details.teaser.title}
                allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
                allowFullScreen>
            </iframe>
            <span itemProp='provider' itemScope itemType='https://schema.org/Person'>
                <span className='hidden' itemProp='name' content="raznick"/>
                <span className='hidden' itemProp='url' content="https://hatechnolog.com/users/raznick"/>
            </span>
        </div>
    )
}
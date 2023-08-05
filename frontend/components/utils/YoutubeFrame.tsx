export type YoutubeFrameProps = {
    title: string;
    description: string;
    videoId: string;
    frameWidth: number;
    frameHeight: number;
}

export default function YoutubeFrame(props: YoutubeFrameProps) {
    return (
        <div className="flex flex-col gap-3 text-right bg-secondary w-fit px-4 py-4 rounded-lg border-black border">
            <h4 className="text-md">{props.title}</h4>
            <p className="text-sm">{props.description}</p>
            <iframe
                className="rounded-md" 
                width={`${props.frameWidth}`} height={`${props.frameHeight}`}
                src={`https://www.youtube.com/embed/${props.videoId}`}
                title={props.title}
                allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
                allowFullScreen></iframe>
        </div>
    )
}
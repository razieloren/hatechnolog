import YoutubeFrame, { YoutubeFrameProps } from "../../utils/YoutubeFrame"
import Social from "./Social"

export default function Hero() {
    return (
        <div className="flex flex-col gap-1 items-center">
            <h1 className="font-bold text-xl">פורומים טכנולוגיים, בלוג מחקר וקורסי תכנות - בחינם.</h1>
            <h2 className="text-l my-4">בואו לתכנת, זה כיף! עם הקהילה הטכנולוגית הגדולה במדינה ❤️</h2>
            <Social/>
        </div>
    )
}
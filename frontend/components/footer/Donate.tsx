import Link from "next/link";

export default function Donate() {
    return (
        <div className="flex flex-col gap-1">
            <span className="font-bold">אוהבים את התוכן שלנו? רוצים שהקהילה תישאר חינמית לעד? אז קדימה - זה יותר זול מספוטיפיי :)</span>
            <span>
                <span>יש שתי אפשרויות לתמוך כספית בקהילת הטכנולוג:   </span>
                <Link
                    href={process.env.YOUTUBE_MEMBER_JOIN_URL!} target="_blank"
                    className="font-bold text-red-500"
                >
                    Youtube Membership עבור צופי הקורסים
                </Link>
                <span> ו-</span>
                <Link
                    href={process.env.DISCORD_MEMBER_JOIN_URL!} target="_blank"
                    className="font-bold text-purple-400"
                >
                    Discord Subscription עבור פעילי הפורום
                </Link>
            </span>
        </div>
    )
}
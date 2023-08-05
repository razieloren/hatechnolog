import { PublicGetUser } from "@/utils/rpc/user.server";
import { Metadata } from "next";
import Image from "next/image";

type UserPageProps = {
    params: {
        handle: string;
    };
}

export async function generateMetadata(props: UserPageProps): Promise<Metadata> {
    const user = await PublicGetUser(props.params.handle);
    const joinedAt = new Date(user.member_since.seconds * 1000);
    return {
        title: `הטכנולוג - ${props.params.handle}`,
        description: `תאריך הצטרפות:`
    }
}

export default async function UserPage(props: UserPageProps) {
    const now = new Date();
    const user = await PublicGetUser(props.params.handle);
    const joinDate = new Date(user.member_since.seconds * 1000);
    const joinedDays = Math.floor((now.getTime() - joinDate.getTime()) / (1000 * 60 * 60 * 24));
    return (
        <div className="w-full">
            <div className="flex flex-col gap-4">
                <div className="flex gap-6 items-center bg-purple-800 w-fit py-2 px-4 rounded-lg">
                    <Image className="w-20 h-20 rounded-full border-2 border-black" src={user.avatar_url} width={256} height={256} alt={`Hatechnolog ${user.handle} avatar`}/>
                    <h1 className="text-2xl font-bold">פרופיל: {user.handle}</h1>
                </div>
                <div className="flex gap-2 text-xl">
                    <span className="font-bold">ותק:</span>
                    <span>{joinedDays} ימים</span>
                </div>
            </div>
        </div>
    )
}
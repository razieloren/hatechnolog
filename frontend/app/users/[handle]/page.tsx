import { PublicGetUser } from "@/utils/rpc/user.server";
import { Metadata } from "next";
import Image from "next/image";

type UserPageProps = {
    params: {
        handle: string;
    };
}

export async function generateMetadata(
    { params }: UserPageProps,
  ): Promise<Metadata> {
    const content = await PublicGetUser(params.handle);   
    return {
        metadataBase: new URL("https://hatechnolog.com"),
        title: `הטכנולוג - משתמש - ${content.handle}`,
        description: `פרופיל המשתמש של ${content.handle}`,
        alternates: {
            canonical: `/users/${content.handle}`
        },
        openGraph: {
            title: `הטכנולוג - משתמש - ${content.handle}`,
            description: `פרופיל המשתמש של ${content.handle}`,
            url: `https://hatechnolog.com/users/${content.handle}`,
            siteName: "הטכנולוג",
        },
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
                <div className="flex gap-4 items-center">
                    <Image className="w-14 h-14 rounded-full" src={user.avatar_url} width={256} height={256} alt={`Hatechnolog ${user.handle} avatar`}/>
                    <h1 className="text-xl font-bold">{user.handle}</h1>
                </div>
                <div className="flex gap-2 text-md">
                    <span className="font-bold">ותק:</span>
                    <span>{joinedDays} ימים</span>
                </div>
            </div>
        </div>
    )
}
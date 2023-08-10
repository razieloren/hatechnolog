import { PublicGetUser } from "@/utils/rpc/user.server";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import { Metadata } from "next";
import Image from "next/image";
import Link from "next/link";

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
            images: {
                url: content.avatar_url,
                secureUrl: content.avatar_url,
                alt: `Hatechnolog User Avatar of ${content.handle}`,
                width: 256,
                height: 256,
            }
        },
        twitter: {
            site: "https://hatechnolog.com",
            creator: "@hatechnolog",
            description: `פרופיל המשתמש של ${content.handle}`,
            title: `הטכנולוג - משתמש - ${content.handle}`,
            images: {
                url: content.avatar_url,
                secureUrl: content.avatar_url,
                alt: `Hatechnolog User Avatar of ${content.handle}`,
                width: 256,
                height: 256,
            }
        }
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
                <div className="flex gap-4 items-center" itemScope itemType="https://schema.org/Person">
                    <Avatar>
                        <AvatarImage itemProp="image" content={user.avatar_url} src={user.avatar_url} />
                        <AvatarFallback>{user.handle[0].toUpperCase()}</AvatarFallback>
                    </Avatar>
                    <h1 itemProp="name" className="text-xl font-bold">{user.handle}</h1>
                    <span hidden itemProp="url" content={`/users/${user.handle}`} />
                </div>
                <div className="flex gap-2 text-md">
                    <span className="font-bold">ותק:</span>
                    <span>{joinedDays} ימים</span>
                </div>
                {user.contents.length !== 0 &&
                    <div className="flex flex-col gap-2">
                        <span className="text-md">פוסטים:</span>
                        {user.contents.map(content => {
                            return (
                                <Link key={content.slug} className="in-link" href={`/blog/${content.slug}`} target="_blank">
                                    <h1>{content.title}</h1>
                                </Link>
                            )
                        })}
                    </div>
                }
            </div>
        </div>
    )
}
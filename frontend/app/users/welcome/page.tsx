import {messages as user_api} from '@/messages/user/api'
import { GetMyUser } from '@/utils/rpc/user.server';
import Link from 'next/link';
import Image from 'next/image';
import failureIcon from '@/public/misc/failure_icon.svg';
import UserRedirect from '@/components/utils/UserRedirect';
import RefreshDiscordUserButton from '@/components/utils/RefreshDiscordUserButton';

export default async function UserWelcomePage() {
    let isLoggedIn = false;
    let user = new user_api.GetUserResponse();
    try {
        user = await GetMyUser();
        isLoggedIn = true;
    } catch (e) {
        // Maybe we must logged out...
    }
    console.log(user.state);

    return (
        <div className="text-lg font-bold text-center">
        {isLoggedIn ?
        <>
        {user.state !== user_api.UserState.CREATED ?
        <div className="flex flex-col gap-6 items-center text-center justify-center">
            <span>הגיע הזמן להשלים את ההרשמה לקהילה!</span>
            <div className="flex flex-col gap-4 bg-purple-950 rounded-lg p-6 items-start justify-start text-start">
                <span className="text-xl">הגדרות משתמש דיסקורד</span>
                <div>
                {!user.is_hatechnolog_member &&
                    <div className="flex gap-2 items-center justify-start text-start">
                        <Image className="h-10 w-10" src={failureIcon} alt="Setup Failure"/>
                        <div className="flex flex-col gap-2">
                            <span>לא מצאנו אותך בשרת הדיסוקרד של הקהילה 😢</span>
                            <Link href={process.env.DISCORD_INVITE_URL!} target="_blank">
                                <button className="bg-green-600 px-2 hover:bg-green-800 text-sm rounded-md">לחצו להצטרפות!</button>
                            </Link>
                        </div>
                    </div>
                }
                </div>
                <div>
                {!user.email_verified &&
                    <div className="flex gap-2 items-center justify-start text-start">
                        <Image className="h-10 w-10" src={failureIcon} alt="Hatechnolog Setup Missing"/>
                        <div className="flex flex-col gap-2">
                            <span>כתובת האימייל שלך אינה מאומתת</span>
                            <Link href="https://support.discord.com/hc/en-us/articles/213219267-Resending-Verification-Email" target="_blank">
                                <button className="bg-green-600 hover:bg-green-800 px-2 text-sm rounded-md">איך לשלוח מייל אימות מחדש?</button>
                            </Link>
                        </div>
                    </div>
                }
                </div>
                <RefreshDiscordUserButton/>
            </div>
        </div>
        :
        <UserRedirect/>
        }
        </>
        :
        <div>
            <span>כדאי להתחבר קודם... ¯\_(ツ)_/¯</span>
        </div>
        }
        </div>
    )
}
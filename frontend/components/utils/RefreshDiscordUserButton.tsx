'use client'

import Image from 'next/image'
import refreshIcon from '@/public/misc/refresh.svg'
import { MouseEvent, useState } from 'react'
import { RefreshDiscordUser } from '@/utils/rpc/user.client';
import { useRouter } from 'next/navigation';

type ErrorContextProps = {
    hasError: boolean;
    errorCount: number;
}

export default function RefreshDiscordUserButton() {
    const router = useRouter();
    const [errorContext, setErrorContext] = useState<ErrorContextProps>({
        hasError: false,
        errorCount: 0,
    });

    function onUpdateClick(e: MouseEvent<HTMLButtonElement>) {
        e.preventDefault();
        RefreshDiscordUser().then(() => {
            router.refresh();
        }).catch((e) => {
            console.log('Error:', e);
            setErrorContext(prev => {
                return {
                    hasError: true,
                    errorCount: prev.errorCount + 1,
                }
            });
        })
    }
    return (
        <div>
            <button onClick={onUpdateClick} className="bg-orange-600 w-fit self-center p-2 hover:bg-orange-800 rounded-md">
                <div className="flex gap-2 items-center">
                    <span>בדקו מחדש בבקשה</span>
                    <Image className="w-6 h-6" src={refreshIcon} alt="Refresh"/>
                </div>
            </button>
            {errorContext.hasError &&
                <span>נראה שזה לא עבד ({errorContext.errorCount})... האם באמת כל ההגדרות נכונות?</span>
            }
        </div>
    )
}
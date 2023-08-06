'use client'

import { useRouter, useSearchParams } from 'next/navigation';
import { useEffect } from 'react';

export default function UserRedirect() {
    const router = useRouter();
    const searchParams = useSearchParams();

    useEffect(() => {
        try {
            router.push(searchParams.get(process.env.REDIRECT_PARAM!) || "/");
        } catch(e) {
            router.push("/");
        }
    });

    return (
        <h1>איזה כיף לראות אותך 🥳</h1>
    )
}
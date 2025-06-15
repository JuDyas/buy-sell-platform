'use client';
import { useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { useUser } from '@/context/UserContext';

export function useAuthRedirect({ redirectTo = "/", fetchOnMount = true } = {}) {
    const router = useRouter();
    const { user, fetchUser } = useUser();

    useEffect(() => {
        if (fetchOnMount) fetchUser();
    }, [fetchUser, fetchOnMount]);

    useEffect(() => {
        if (user) {
            router.replace(redirectTo);
        }
    }, [user, router, redirectTo]);

    return {
        user,
        loading: user === undefined,
    };
}
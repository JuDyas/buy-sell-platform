'use client';
import { useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { useUser } from '@/context/UserContext';

export function useAuthRedirect({ redirectTo = "/", requireAuth = true, fetchOnMount = true } = {}) {
    const router = useRouter();
    const { user, fetchUser } = useUser();

    useEffect(() => {
        if (fetchOnMount) fetchUser();
    }, [fetchUser, fetchOnMount]);

    useEffect(() => {
        if (requireAuth) {
            if (user === null) {
                router.replace(redirectTo);
            }
        } else {
            if (user) {
                router.replace(redirectTo);
            }
        }
    }, [user, router, redirectTo, requireAuth]);

    return {
        user,
        loading: user === undefined,
    };
}
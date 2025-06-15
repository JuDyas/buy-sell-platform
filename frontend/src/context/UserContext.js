"use client"
import {createContext, useCallback, useContext, useState} from 'react';

export const UserContext = createContext(null)

export function UserProvider({ children }) {
    const [user, setUser] = useState(undefined)

    const fetchUser  = useCallback (async () => {
        const token = localStorage.getItem('token')
        if (!token) {
            setUser(null);
            return;
        }

        try {
            const res = await fetch(process.env.NEXT_PUBLIC_API_URL + '/users/me', {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
                credentials: 'include',
                cache: 'no-store',
            })

            if (res.ok) {
                const data = await res.json()
                setUser(data)
            } else {
                setUser(null)
            }
        } catch {
            setUser(null)
        }
    }, []);

    return (
        <UserContext.Provider value={{user, fetchUser}}>
            {children}
        </UserContext.Provider>
    )
}

export function useUser() {
    return useContext(UserContext)
}
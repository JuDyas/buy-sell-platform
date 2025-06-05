"use client";

import { useCallback, useEffect, useRef, useState } from "react";
import Link from "next/link";
import Image from "next/image";

/**
 * Header component that displays a navigation header.
 * The component includes user authentication logic, a dynamic user avatar,
 * and a dropdown menu for user-specific actions or authentication links.
 *
 * Internal logic:
 * - Fetches user data using a stored token.
 * - Toggles dropdown menu visibility and handles clicks outside the menu to close it.
 * - Provides user avatar image source URL handling.
 * - Logs the user out by clearing the token and reloading the page.
 *
 * @return {JSX.Element} Returns the header JSX structure, containing navigation and user menu.
 */

export default function Header() {
    const [user, setUser] = useState(null);
    const [menuOpen, setMenuOpen] = useState(false);
    const menuRef = useRef(null);

    const getToken = useCallback(() => {
        if (typeof window !== "undefined") {
            return localStorage.getItem("token");
        }
        return null;
    }, []);

    const getAvatarUrl = (user) => {
        if (!user?.avatar_url) return "icons/user.svg";
        // Если путь уже абсолютный (http, https) — возвращаем как есть
        if (/^https?:\/\//.test(user.avatar_url)) return user.avatar_url;
        // Иначе собираем с базовым API урлом (без завершающего /)
        const base = process.env.NEXT_PUBLIC_API_URL_IMG?.replace(/\/+$/, "");
        return base + user.avatar_url;
    };

    useEffect(() => {
        const token = getToken();
        if (!token) {
            setUser(null);
            return;
        }

        fetch(process.env.NEXT_PUBLIC_API_URL + "/users/me", {
            headers: {
                Authorization: `Bearer ${token}`,
            },
            credentials: "include",
            cache: "no-store",
        })
            .then((res) => (res.ok ? res.json() : Promise.reject()))
            .then((data) => setUser(data))
            .catch(() => setUser(null));
    }, [getToken]);

    useEffect(() => {
        function handleClickOutside(event) {
            if (menuRef.current && !menuRef.current.contains(event.target)) {
                setMenuOpen(false);
            }
        }
        if (menuOpen) {
            document.addEventListener("mousedown", handleClickOutside);
        }
        return () => {
            document.removeEventListener("mousedown", handleClickOutside);
        };
    }, [menuOpen]);

    const handleLogout = useCallback(() => {
        localStorage.removeItem("token");
        window.location.reload();
    }, []);

    return (
        <header className="bg-white border-b">
            <div className="container mx-auto px-4 py-3 flex items-center justify-between">
                <Link href="/">
                  <span className="text-2xl font-bold tracking-tight text-[#2563eb] cursor-pointer select-none">
                    GoSell
                  </span>
                </Link>
                <div className="relative" ref={menuRef}>
                    <button
                        className="flex items-center focus:outline-none transition duration-100"
                        onClick={() => setMenuOpen((open) => !open)}
                        aria-label={user ? "Меню пользователя" : "Меню входа"}
                    >
                        <div className="w-10 h-10 relative">
                            <Image
                                src={getAvatarUrl(user)}
                                alt="User"
                                fill
                                className="rounded-full border bg-gray-100 object-cover"
                                priority
                            />
                        </div>
                    </button>
                    {menuOpen && (
                        <div className="absolute right-0 mt-2 w-56 bg-white border rounded-lg shadow-lg z-50 py-2 animate-fade-in">
                            {user ? (
                                <div>
                                    <div className="px-4 pt-2 text-sm text-gray-900 font-medium truncate">
                                        {user.username || user.email || "Пользователь"}
                                    </div>
                                    <div className="px-4 pb-2 text-sm text-gray-500 truncate border-b">
                                        {user.email}
                                    </div>
                                    <Link
                                        href="/profile"
                                        className="block px-4 py-2 text-left hover:bg-gray-50 text-gray-700"
                                    >
                                        Профіль
                                    </Link>
                                    <button
                                        onClick={handleLogout}
                                        className="block w-full px-4 py-2 text-left hover:bg-gray-50 text-gray-700"
                                    >
                                        Вийти
                                    </button>
                                </div>
                            ) : (
                                <div>
                                    <Link
                                        href="/login"
                                        className="block px-4 py-2 hover:bg-gray-50 text-gray-700"
                                    >
                                        Війти
                                    </Link>
                                    <Link
                                        href="/register"
                                        className="block px-4 py-2 hover:bg-gray-50 text-gray-700"
                                    >
                                        Зареєструватися
                                    </Link>
                                </div>
                            )}
                        </div>
                    )}
                </div>
            </div>
        </header>
    );
}
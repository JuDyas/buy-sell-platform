"use client";

import { useCallback, useEffect, useRef, useState } from "react";
import Link from "next/link";
import Image from "next/image";
import { useUser } from "@/context/UserContext";
import { useRouter } from "next/navigation";

export default function Header() {
    const {user, fetchUser} = useUser();
    const [menuOpen, setMenuOpen] = useState(false);
    const menuRef = useRef(null);

    const [categories, setCategories] = useState([]);
    const [catMenuOpen, setCatMenuOpen] = useState(false);
    const catRef = useRef(null);
    const router = useRouter();

    useEffect(() => {
        async function fetchCategories() {
            try {
                const res = await fetch(process.env.NEXT_PUBLIC_API_URL + '/categories');
                const data = await res.json();
                if (data?.categories && Array.isArray(data.categories)) {
                    setCategories(data.categories);
                } else {
                    setCategories([]);
                }
            } catch {
                setCategories([]);
            }
        }
        fetchCategories();
    }, []);

    const getAvatarUrl = (user) => {
        if (!user?.avatar_url) return "/icons/user.svg";
        if (/^https?:\/\//.test(user.avatar_url)) return user.avatar_url;
        const base = process.env.NEXT_PUBLIC_API_URL_IMG?.replace(/\/+$/, "");
        return base + user.avatar_url;
    };

    useEffect(() => {
        fetchUser();
    }, [fetchUser]);

    useEffect(() => {
        function handleClickOutside(event) {
            if (menuRef.current && !menuRef.current.contains(event.target)) {
                setMenuOpen(false);
            }
            if (catRef.current && !catRef.current.contains(event.target)) {
                setCatMenuOpen(false);
            }
        }
        if (menuOpen || catMenuOpen) {
            document.addEventListener("mousedown", handleClickOutside);
        }
        return () => {
            document.removeEventListener("mousedown", handleClickOutside);
        };
    }, [menuOpen, catMenuOpen]);

    const handleLogout = useCallback(() => {
        localStorage.removeItem("token");
        window.location.reload();
    }, []);

    const handleCategoryClick = (id) => {
        setCatMenuOpen(false);
        router.push(`/categories/${id}`);
    };

    return (
        <header className="bg-white border-b">
            <div className="container mx-auto px-4 py-1.5 flex items-center justify-between">
                <Link href="/">
                  <span className="text-2xl font-bold tracking-tight text-[#2563eb] cursor-pointer select-none">
                    GoSell
                  </span>
                </Link>

                <div className="flex items-center gap-3">
                    <div className="relative" ref={catRef}>
                        <button
                            onClick={() => setCatMenuOpen((prev) => !prev)}
                            className="px-4 py-2 bg-gray-100 hover:bg-gray-200 text-gray-800 rounded-lg text-sm font-medium flex items-center gap-1 transition"
                        >
                            Категорії
                            <svg className={`w-4 h-4 ml-1 transition-transform ${catMenuOpen ? "rotate-180" : ""}`} fill="none" stroke="currentColor" strokeWidth={2} viewBox="0 0 24 24">
                                <path strokeLinecap="round" strokeLinejoin="round" d="M19 9l-7 7-7-7" />
                            </svg>
                        </button>
                        {catMenuOpen && (
                            <div className="absolute left-0 mt-2 min-w-[12rem] bg-white border rounded-lg shadow-lg z-50 py-1 animate-fade-in">
                                {categories.length === 0 &&
                                    <div className="px-4 py-2 text-sm text-gray-500">Завантаження...</div>
                                }
                                {categories.map(cat => (
                                    <button
                                        key={cat.id}
                                        onClick={() => handleCategoryClick(cat.id)}
                                        className="block w-full text-left px-4 py-2 text-gray-700 hover:bg-gray-100 transition"
                                    >
                                        {cat.name}
                                    </button>
                                ))}
                            </div>
                        )}
                    </div>

                    <Link
                        href="/add"
                        className="px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg text-sm font-semibold transition-colors"
                    >
                        Створити оголошення
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
                                            {user.username || user.email}
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
                                            className="block w-full text-left px-4 py-2 hover:bg-gray-100 text-gray-700"
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
                                            Вхід
                                        </Link>
                                        <Link
                                            href="/register"
                                            className="block px-4 py-2 hover:bg-gray-50 text-gray-700"
                                        >
                                            Реєстрація
                                        </Link>
                                    </div>
                                )}
                            </div>
                        )}
                    </div>
                </div>
            </div>
        </header>
    );
}
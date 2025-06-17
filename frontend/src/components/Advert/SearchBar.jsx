"use client"

import React, { useState } from "react";

export default function SearchBar({ onSearch, initialQuery = "" }) {
    const [query, setQuery] = useState(initialQuery);

    const handleSearch = (e) => {
        e.preventDefault();
        if (query.trim()) {
            if (onSearch) {
                onSearch(query);
            } else {
                window.open(`/search?q=${encodeURIComponent(query)}`, "_blank");
            }
        }
    };

    return (
        <form
            onSubmit={handleSearch}
            className="
                flex
                items-center
                w-full
                my-4
                bg-white
                rounded-xl
                shadow-md
                border
                border-gray-200
                px-5 py-1
                gap-3
                max-w-2xl
                mx-auto
            "
        >
            <input
                type="text"
                placeholder="Пошук по оголошеннях..."
                value={query}
                onChange={(e) => setQuery(e.target.value)}
                className="
                    flex-1
                    bg-transparent
                    outline-none
                    py-2
                    px-1
                    text-lg
                    rounded
                    placeholder-gray-400
                    text-gray-800
                "
                autoComplete="off"
            />
            <button
                type="submit"
                className="
                    flex
                    items-center
                    gap-2
                    bg-blue-600
                    text-white
                    font-semibold
                    text-base
                    px-4
                    py-1.5
                    rounded-full
                    transition
                    hover:bg-blue-700
                    shadow
                    active:scale-95
                    duration-100
                "
            >
                <svg className="w-5 h-5" fill="none" stroke="currentColor" strokeWidth="2" viewBox="0 0 24 24">
                    <circle cx="11" cy="11" r="7" stroke="currentColor" strokeWidth="2"></circle>
                    <line x1="16.65" y1="16.65" x2="21" y2="21" stroke="currentColor" strokeWidth="2" strokeLinecap="round"/>
                </svg>
                Пошук
            </button>
        </form>
    );
}
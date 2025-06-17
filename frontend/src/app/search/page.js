"use client";

import { useState, useEffect } from "react";
import { useSearchParams, useRouter } from "next/navigation";
import SearchBar from "@/components/Advert/SearchBar";
import AdvertsGrid from "@/components/Advert/AdvertGrid";

export default function SearchResultPage() {
    const API_URL = process.env.NEXT_PUBLIC_API_URL;
    const searchParams = useSearchParams();
    const query = searchParams.get("q") || "";
    const [adverts, setAdverts] = useState([]);
    const [loading, setLoading] = useState(true);
    const router = useRouter();

    const handleSearch = (newQuery) => {
        router.push(`/search?q=${encodeURIComponent(newQuery)}`);
    };

    useEffect(() => {
        if (!query) {
            setAdverts([]);
            setLoading(false);
            return;
        }
        setLoading(true);
        fetch(`${API_URL}/adds/search`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ query }),
        })
            .then((res) => res.json())
            .then((data) => setAdverts(data.adverts || []))
            .finally(() => setLoading(false));
    }, [query]);

    return (
        <div className="min-h-[85vh]">
            <SearchBar onSearch={handleSearch} initialQuery={query} />
            <AdvertsGrid adverts={adverts} loading={loading} />
        </div>
    );
}
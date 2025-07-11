'use client';

import { useEffect, useState } from 'react';
import Link from 'next/link';

const API_URL = process.env.NEXT_PUBLIC_API_URL;
const IMG_URL = process.env.NEXT_PUBLIC_API_URL_IMG?.replace(/\/+$/, '');

export default function AdvertsGrid({ adverts: externalAdverts, categoryId }) {
    const [adverts, setAdverts] = useState(externalAdverts || []);
    const [loading, setLoading] = useState(!externalAdverts);
    const [error, setError] = useState('');

    useEffect(() => {
        if (externalAdverts) {
            setAdverts(externalAdverts);
            setLoading(false);
            setError('');
            return;
        }

        const fetchCategoryAdverts = async () => {
            setLoading(true);
            setError('');
            try {
                const res = await fetch(`${API_URL}/categories/${categoryId}/adds`);
                const data = await res.json();
                if (Array.isArray(data)) {
                    setAdverts(data);
                } else if (Array.isArray(data?.adverts)) {
                    setAdverts(data.adverts);
                } else {
                    setError('Не вдалося отримати оголошення');
                }
            } catch {
                setError('Помилка з’єднання з сервером');
            } finally {
                setLoading(false);
            }
        };

        const fetchAllAdverts = async () => {
            setLoading(true);
            setError('');
            try {
                const res = await fetch(`${API_URL}/adds`);
                const data = await res.json();
                if (Array.isArray(data?.adverts)) {
                    setAdverts(data.adverts);
                } else {
                    setError('Не вдалося отримати оголошення');
                }
            } catch {
                setError('Помилка з’єднання з сервером');
            } finally {
                setLoading(false);
            }
        };

        if (categoryId) {
            fetchCategoryAdverts();
        } else {
            fetchAllAdverts();
        }
    }, [externalAdverts, categoryId]);

    if (loading) {
        return <div className="py-10 text-center text-gray-600">Завантаження…</div>;
    }
    if (error) {
        return <div className="py-10 text-center text-red-500">{error}</div>;
    }
    if (!adverts || adverts.length === 0) {
        return <div className="py-10 text-center text-gray-500">Нічого не знайдено</div>;
    }

    return (
        <div className="container mx-auto grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6 p-4">
            {adverts.map(ad => (
                <Link
                    href={`/add/${ad.ID}`}
                    key={ad.ID}
                    className="block bg-white rounded-xl shadow hover:shadow-lg transition overflow-hidden cursor-pointer border border-gray-100"
                >
                    <div className="aspect-[4/3] bg-gray-200 flex items-center justify-center overflow-hidden">
                        {ad.Images?.length > 0 ? (
                            <img
                                src={
                                    ad.Images[0].startsWith('http')
                                        ? ad.Images[0]
                                        : `${IMG_URL}/${ad.Images[0].replace(/^\.?\//, '')}`
                                }
                                alt={ad.Title}
                                className="object-cover w-full h-full"
                            />
                        ) : (
                            <div className="text-gray-400 text-4xl w-full h-full flex items-center justify-center">
                                <span>🖼️</span>
                            </div>
                        )}
                    </div>
                    <div className="p-4">
                        <div className="font-bold text-lg line-clamp-1 mb-1">{ad.Title}</div>
                        <div className="text-gray-500 text-sm mb-2 line-clamp-2">{ad.Description}</div>
                        <div className="font-semibold text-blue-600 text-right">{ad.Price} ₴</div>
                    </div>
                </Link>
            ))}
        </div>
    );
}
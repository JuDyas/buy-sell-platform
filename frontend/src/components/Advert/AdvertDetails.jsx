'use client';

import { useEffect, useState } from 'react';

const API_URL = process.env.NEXT_PUBLIC_API_URL;
const IMG_URL = process.env.NEXT_PUBLIC_API_URL_IMG?.replace(/\/+$/, '');

export default function AdvertDetails({ advertId }) {
    const [advert, setAdvert] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState('');
    const [imgIdx, setImgIdx] = useState(0);
    const [author, setAuthor] = useState(null);
    const [category, setCategory] = useState(null);
    const [showContacts, setShowContacts] = useState(false);

    useEffect(() => {
        if (!advertId) return;
        const fetchAdvert = async () => {
            setLoading(true);
            try {
                const res = await fetch(`${API_URL}/adds/${advertId}`);
                const data = await res.json();
                if (data && (data.ID || data.advert || data.adverts)) {
                    const ad = data.advert || data || (data.adverts?.[0] ?? null);
                    setAdvert(ad);
                    setImgIdx(0);

                    if (ad.AuthorID) {
                        fetch(`${API_URL}/users/${ad.AuthorID}`)
                            .then(r => r.json())
                            .then(setAuthor)
                            .catch(() => setAuthor(null));
                    } else {
                        setAuthor(null);
                    }

                    if (ad.Category) {
                        fetch(`${API_URL}/categories/${ad.Category}`)
                            .then(r => r.json())
                            .then(setCategory)
                            .catch(() => setCategory(null));
                    } else {
                        setCategory(null);
                    }
                } else {
                    setError('Оголошення не знайдено');
                }
            } catch {
                setError('Помилка отримання оголошення');
            } finally {
                setLoading(false);
            }
        };
        fetchAdvert();
    }, [advertId]);

    if (loading) return <div className="py-10 text-center">Завантаження…</div>;
    if (error) return <div className="py-10 text-center text-red-500">{error}</div>;
    if (!advert) return null;

    const images = advert.Images || [];
    const handlePrev = () => setImgIdx((old) => (old - 1 + images.length) % images.length);
    const handleNext = () => setImgIdx((old) => (old + 1) % images.length);

    const getAvatarUrl = (user) => {
        if (!user?.avatar_url) return "/icons/user.svg";
        if (/^https?:\/\//.test(user.avatar_url)) return user.avatar_url;
        return IMG_URL + user.avatar_url;
    };

    return (
        <div className="container mx-auto py-10 px-4 flex flex-col gap-6">
            <div className="flex flex-row gap-6">
                {/* Левая колонка: 70% — фото и описание */}
                <div className="w-[70%] flex flex-col gap-6">
                    <div className="rounded-xl bg-white shadow p-5 flex items-center justify-center relative">
                        {images.length > 1 && (
                            <button
                                className="
                                absolute left-0 top-1/2 -translate-y-1/2
                                bg-gray-200 hover:bg-gray-300
                                px-2 pb-1.5 ml-1.5 rounded-full shadow transition"
                                onClick={handlePrev}
                                aria-label="Попередне фото"
                            >
                                <span className="text-2xl">{'‹'}</span>
                            </button>
                        )}
                        {images.length ? (
                            <img
                                src={
                                    images[imgIdx].startsWith('http')
                                        ? images[imgIdx]
                                        : `${IMG_URL}/${images[imgIdx].replace(/^\.?\//, '')}`
                                }
                                alt={advert.Title}
                                className="w-full max-h-[38rem] object-contain rounded"
                            />
                        ) : (
                            <div className="w-full h-[200px] bg-gray-200 flex justify-center items-center rounded text-3xl text-gray-400">
                                🖼️
                            </div>
                        )}
                        {images.length > 1 && (
                            <button
                                className="absolute right-0 top-1/2 -translate-y-1/2
                                bg-gray-200 hover:bg-gray-300
                                px-2 pb-1.5 mr-1.5 rounded-full shadow transition"
                                onClick={handleNext}
                                aria-label="Наступне фото"
                            >
                                <span className="text-2xl">{'›'}</span>
                            </button>
                        )}
                    </div>
                    {/* Описание и категория под фотографиями */}
                    <div className="bg-white rounded-xl shadow p-6">
                        <div className="mb-2 text-sm text-gray-500">
                            Категорія: {category?.name || '—'}
                        </div>
                        <div className="font-semibold mb-1">Опис</div>
                        <div className="text-gray-700 whitespace-pre-line">{advert.Description}</div>
                    </div>
                </div>

                {/* Правая колонка: 30% — детали, продавец, местоположение */}
                <div className="w-[30%] flex flex-col gap-4">
                    <div className="bg-white rounded-xl shadow p-6 flex flex-col gap-4">
                        <h1 className="text-2xl font-bold">{advert.Title}</h1>
                        <div className="text-lg font-bold text-blue-700">{advert.Price} ₴</div>

                        {showContacts && (
                            <div className="mb-3 py-2 px-3 bg-blue-50 rounded">
                                <div className="text-sm text-gray-800">
                                    <span className="font-semibold">Email:</span>{' '}
                                    {author?.email || '—'}
                                </div>
                                {author?.phone && (
                                    <div className="text-sm text-gray-800">
                                        <span className="font-semibold">Телефон:</span>{' '}
                                        {author.phone}
                                    </div>
                                )}
                            </div>
                        )}

                        <button
                            className="w-full bg-blue-600 hover:bg-blue-700 text-white rounded px-4 py-2 transition"
                            onClick={() => setShowContacts(v => !v)}
                        >
                            {showContacts ? "Сховати контакти" : "Показати контакти продавця"}
                        </button>
                    </div>
                    <div className="bg-white rounded-xl shadow p-6 flex flex-col gap-2">
                        <div className="font-medium text-sm text-gray-500">Користувач</div>
                        <div className="flex gap-3 items-center">
                            <img
                                src={getAvatarUrl(author)}
                                alt="avatar"
                                className="w-10 h-10 rounded-full object-cover bg-gray-100 border"
                            />
                            <div>
                                <div className="font-semibold text-lg">
                                    {author?.username || '—'}
                                </div>
                                <div className="text-sm text-gray-400">
                                    Реєстрація:{" "}
                                    {author?.created_at
                                        ? new Date(author.created_at).toLocaleDateString()
                                        : '—'}
                                </div>
                            </div>
                        </div>
                        <button
                            className="mt-2 bg-gray-100 hover:bg-gray-200 text-blue-700 rounded px-3 py-1 text-sm transition"
                            onClick={() =>
                                author?.id &&
                                window.open(`/profile/${author.id}`, '_blank')
                            }
                        >
                            Перейти до профілю
                        </button>
                    </div>
                    <div className="bg-white h-32 rounded-xl shadow p-6 flex flex-col">
                        <div className="font-medium text-sm text-gray-500 mb-1">Місцезнаходження</div>
                        <div className="text-gray-800">
                            {advert.Location || advert.location || advert.city || '—'}
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}
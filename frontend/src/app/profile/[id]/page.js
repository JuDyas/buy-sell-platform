'use client';

import React, { useEffect, useState } from 'react';
import AdvertsGrid from '@/components/Advert/AdvertGrid'; // Смени путь при необходимости

const API_URL = process.env.NEXT_PUBLIC_API_URL;
const IMG_URL = process.env.NEXT_PUBLIC_API_URL_IMG?.replace(/\/+$/, '');

export default function ProfilePage({ params }) {
    const { id } = React.use(params);
    const [user, setUser] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState('');

    useEffect(() => {
        if (!id) return;
        const fetchUser = async () => {
            setLoading(true);
            setError('');
            try {
                const res = await fetch(`${API_URL}/users/${id}`);
                if (!res.ok) throw new Error();
                const data = await res.json();
                setUser(data);
            } catch {
                setError('Не вдалося отримати дані користувача');
            } finally {
                setLoading(false);
            }
        };
        fetchUser();
    }, [id]);


    return (
        <div className="min-h-[85vh] flex flex-col md:flex-row gap-8 container mx-auto py-8">
            {/* Левая панель — user info */}
            <div className="md:w-1/3 w-full bg-white rounded-xl shadow p-6 flex flex-col items-center gap-4">
                {loading && <div>Завантаження профілю…</div>}
                {error && <div className="text-red-500">{error}</div>}
                {user && (
                    <>
                        {/* Аватар */}
                        <img
                            src={
                                user.avatar_url
                                    ? (user.avatar_url.startsWith('http')
                                        ? user.avatar_url
                                        : `${IMG_URL}/${user.avatar_url.replace(/^\.?\//, '')}`)
                                    : '/icons/user.svg'
                            }
                            alt={user.username}
                            className="w-32 h-32 object-cover rounded-full border"
                        />
                        {/* Имя */}
                        <div className="text-xl font-bold">{user.username || 'Без імені'}</div>
                        {/* Телефон */}
                        {user.phone && (
                            <div className="text-gray-700 flex items-center gap-1">
                                <span className="material-icons text-base">Телефон:</span>
                                <span>{user.phone}</span>
                            </div>
                        )}
                        {/* Местоположение */}
                        {user.location && (
                            <div className="text-gray-700 flex items-center gap-1">
                                <span className="material-icons text-base">Місцезнаходження:</span>
                                <span>{user.location}</span>
                            </div>
                        )}
                        {/* Почта */}
                        <div className="text-gray-700 flex items-center gap-1">
                            <span className="material-icons text-base">Пошта:</span>
                            <span>{user.email}</span>
                        </div>
                        <div className="text-xs text-gray-400 mt-2">
                            Зареєстрований: {user.created_at && (new Date(user.created_at).toLocaleDateString('uk-UA'))}
                        </div>
                    </>
                )}
            </div>
            {/* Правая панель — объявления */}
            <div className="md:w-2/3 w-full">
                <h2 className="text-2xl font-semibold mb-4">Оголошення користувача</h2>
                <AdvertsGrid userId={id} />
            </div>
        </div>
    );
}
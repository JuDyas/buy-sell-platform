'use client';
import { useEffect, useState, useRef } from 'react';

export default function UserProfileForm() {
    const [form, setForm] = useState({
        username: '',
        email: '',
        password: '',
        phone: '',
        location: '',
        avatar_url: '',
    });
    const [loading, setLoading] = useState(false);
    const [avatarUploading, setAvatarUploading] = useState(false);
    const [error, setError] = useState('');
    const [message, setMessage] = useState('');
    const fileInputRef = useRef(null);

    useEffect(() => {
        const load = async () => {
            setLoading(true);
            setError('');
            try {
                const apiBase = process.env.NEXT_PUBLIC_API_URL;
                const token = localStorage.getItem('token');
                const res = await fetch(`${apiBase}/users/me`, {
                    headers: { 'Authorization': `Bearer ${token}` },
                });
                const data = await res.json();
                if (res.ok) {
                    setForm({
                        username: data.username || '',
                        email: data.email || '',
                        password: '',
                        phone: data.phone || '',
                        location: data.location || '',
                        avatar_url: data.avatar_url || '',
                    });
                } else {
                    setError(data.message || 'Не вдалося отримати дані профілю');
                }
            } catch {
                setError('Помилка з’єднання');
            }
            setLoading(false);
        };
        load();
    }, []);

    const handleChange = (e) => {
        setForm(prev => ({
            ...prev,
            [e.target.name]: e.target.value,
        }));
    };

    const handleAvatarChange = async (e) => {
        if (!e.target.files?.[0]) return;

        setAvatarUploading(true);
        setMessage('');
        setError('');
        try {
            const apiBase = process.env.NEXT_PUBLIC_API_URL;
            const token = localStorage.getItem('token');
            const formData = new FormData();
            formData.append('file', e.target.files[0]);

            const res = await fetch(`${apiBase}/users/upload-avatar`, {
                method: 'POST',
                headers: { 'Authorization': `Bearer ${token}` },
                body: formData,
            });

            const data = await res.json();
            if (res.ok && data.avatarURL) {
                setForm(prev => ({ ...prev, avatar_url: data.avatarURL }));
                setMessage('Аватар успішно оновлено!');
            } else {
                setError(data.message || 'Не вдалося завантажити аватар');
            }
        } catch {
            setError('Помилка при завантаженні аватара');
        }
        setAvatarUploading(false);
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        setLoading(true);
        setError('');
        setMessage('');

        const body = {};
        for (const key in form) {
            if (form[key]) body[key === 'avatar_url' ? 'avatar_url' : key] = form[key];
        }

        try {
            const apiBase = process.env.NEXT_PUBLIC_API_URL;
            const token = localStorage.getItem('token');
            const res = await fetch(`${apiBase}/users`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`,
                },
                body: JSON.stringify(body),
            });
            const data = await res.json();
            if (res.ok) {
                setMessage('Профіль оновлено!');
                setForm(prev => ({ ...prev, password: '' }));
            } else {
                setError(data.message || 'Помилка при оновленні профілю');
            }
        } catch {
            setError('Помилка з’єднання');
        }
        setLoading(false);
    };

    return (
        <form className="w-6/12 flex flex-col p-5 space-y-5" onSubmit={handleSubmit}>
            <div>
                <div>
                    <div className="flex flex-row ">
                        {form.avatar_url && (
                            <img src={process.env.NEXT_PUBLIC_API_URL_IMG + form.avatar_url} alt="Avatar" className="w-28 h-28 rounded-full object-cover mb-3" />
                        )}
                        <div className="flex flex-col justify-center ml-4">
                            <h2 className="text-2xl font-bold mb-4 px-4 m-2 bg-amb">Мій профіль</h2>
                            <input
                                ref={fileInputRef}
                                type="file"
                                accept="image/*"
                                style={{ display: 'none' }}
                                onChange={handleAvatarChange}
                            />
                            <button
                                type="button"
                                onClick={() => fileInputRef.current?.click()}
                                className="px-4 py-1 bg-gray-200 rounded hover:bg-gray-300 transition"
                                disabled={avatarUploading}
                            >
                                {avatarUploading ? 'Завантаження...' : form.avatar_url ? 'Змінити аватар' : 'Завантажити аватар'}
                            </button>
                        </div>
                    </div>
                </div>
            </div>

            <div className="bg-white rounded-lg p-8 space-y-6 shadow-lg">
                <div className="flex flex-col max-w-sm gap-3">
                    <div className="">
                        <label className="block mb-1 font-medium pl-1 text-[0.75em]">Ім'я користувача</label>
                        <input
                            type="text"
                            name="username"
                            value={form.username}
                            onChange={handleChange}
                            className="input input-bordered w-full p-2 bg-[#f2f4f5] rounded-md"
                            required
                            minLength={3}
                            suppressHydrationWarning={true}
                        />
                    </div>
                    <div>
                        <label className="block mb-1 font-medium pl-1 text-[0.75em]">Email</label>
                        <input
                            type="email"
                            name="email"
                            value={form.email}
                            onChange={handleChange}
                            className="input input-bordered w-full p-2 bg-[#f2f4f5] rounded-md"
                            required
                            suppressHydrationWarning={true}
                        />
                    </div>
                    <div>
                        <label className="block mb-1 pl-1 text-[0.75em]">Новий пароль (не обов’язково)</label>
                        <input
                            type="password"
                            name="password"
                            value={form.password}
                            onChange={handleChange}
                            className="input input-bordered w-full p-2 bg-[#f2f4f5] rounded-md"
                            autoComplete="new-password"
                            suppressHydrationWarning={true}
                        />
                    </div>
                    <div>
                        <label className="block mb-1 pl-1 text-[0.75em]">Телефон</label>
                        <input
                            type="tel"
                            name="phone"
                            value={form.phone}
                            onChange={handleChange}
                            className="input input-bordered w-full p-2 bg-[#f2f4f5] rounded-md"
                            suppressHydrationWarning={true}
                        />
                    </div>
                    <div>
                        <label className="block mb-1 pl-1 text-[0.75em]">Локація</label>
                        <input
                            type="text"
                            name="location"
                            value={form.location}
                            onChange={handleChange}
                            className="input input-bordered w-full p-2 bg-[#f2f4f5] rounded-md"
                            suppressHydrationWarning={true}
                        />
                    </div>
                    {error && <div className="text-red-500">{error}</div>}
                    {message && <div className="text-green-500">{message}</div>}
                    <button
                        type="submit"
                        disabled={loading}
                        className="w-full bg-blue-600 hover:bg-blue-700 mt-1 text-white py-2 rounded font-semibold transition-colors disabled:opacity-60"
                    >
                        {loading ? 'Оновлення...' : 'Зберегти зміни'}
                    </button>
                </div>
            </div>
        </form>
    );
}
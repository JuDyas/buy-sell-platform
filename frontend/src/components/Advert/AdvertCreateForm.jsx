'use client';
import { useState, useEffect, useRef } from 'react';
import { useAuthRedirect } from '@/hooks/useAuthRedirect';

export default function CreateAdvertForm() {
    useAuthRedirect({redirectTo: '/login', requireAuth: true, fetchOnMount: true });

    const [form, setForm] = useState({
        title: '',
        description: '',
        category: '',
        price: '',
        images: [],
    });

    const [categories, setCategories] = useState([]);
    const [loading, setLoading] = useState(false);
    const [imgLoading, setImgLoading] = useState(false);
    const [message, setMessage] = useState('');
    const [imgError, setImgError] = useState('');
    const fileInputRef = useRef(null);

    useEffect(() => {
        const fetchCats = async () => {
            try {
                const res = await fetch(process.env.NEXT_PUBLIC_API_URL + '/categories');
                const data = await res.json();
                if (Array.isArray(data)) {
                    setCategories(data);
                } else if (Array.isArray(data.categories)) {
                    setCategories(data.categories);
                } else {
                    setCategories([]);
                }
            } catch {
                setCategories([]);
            }
        };
        fetchCats();
    }, []);

    const handleChange = (e) => {
        const { name, value } = e.target;
        setForm(f => ({ ...f, [name]: value }));
    };

    const handleFiles = async (e) => {
        const files = Array.from(e.target.files);
        if (files.length + form.images.length > 10) {
            setImgError('Максимум 10 зображень.');
            return;
        }
        setImgLoading(true);
        setImgError('');
        const body = new FormData();
        files.forEach(f => body.append('images', f));

        try {
            const res = await fetch(process.env.NEXT_PUBLIC_API_URL + '/adds/upload-images', {
                method: 'POST',
                body,
            });
            const data = await res.json();
            if (Array.isArray(data.images)) {
                setForm(f => ({ ...f, images: [...f.images, ...data.images] }));
            } else {
                setImgError('Не вдалося завантажити зображення');
            }
        } catch {
            setImgError('Помилка при завантаженні зображень');
        } finally {
            setImgLoading(false);
        }
    };

    const handleRemoveImg = (img) => {
        setForm(f => ({
            ...f,
            images: f.images.filter(i => i !== img)
        }));
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        setLoading(true);
        setMessage('');
        const token = typeof window !== 'undefined' ? localStorage.getItem('token') : null;
        try {
            const res = await fetch(process.env.NEXT_PUBLIC_API_URL + '/adds', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    ...(token ? { Authorization: `Bearer ${token}` } : {}),
                },
                body: JSON.stringify({
                    title: form.title,
                    description: form.description,
                    category_id: form.category,
                    price: Number(form.price),
                    images: form.images,
                }),
            });
            if (res.ok) {
                setMessage('Оголошення створено успішно!');
                setForm({
                    title: '',
                    description: '',
                    category: '',
                    price: '',
                    images: [],
                });
            } else {
                const data = await res.json();
                setMessage(data.message || 'Помилка створення оголошення');
            }
        } catch {
            setMessage('Помилка з\'єднання');
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="min-h-[80vh] flex flex-col justify-center items-center bg-[#f2f4f5] py-8">
            <form
                className="w-full max-w-2xl flex flex-col p-5 space-y-6"
                onSubmit={handleSubmit}
            >
                <div className="flex flex-col items-center bg-white rounded-lg p-8 space-y-6 shadow-lg">
                    <h2 className="text-2xl font-bold mb-4 text-gray-800">
                        Нове оголошення
                    </h2>

                    {/* Фото */}
                    <div className="flex flex-col max-w-sm w-full">
                        <label className="block font-medium text-[0.75em] pl-1 mb-1 text-gray-700">
                            Фото (до 10 шт)
                        </label>
                        <div className="flex items-center gap-3">
                            <input
                                ref={fileInputRef}
                                type="file"
                                accept="image/*"
                                multiple
                                style={{ display: 'none' }}
                                disabled={imgLoading || form.images.length >= 10}
                                onChange={handleFiles}
                            />

                            <button
                                type="button"
                                onClick={() => fileInputRef.current?.click()}
                                className="px-4 py-2 bg-gray-400 hover:bg-gray-600 text-white rounded font-semibold transition-colors disabled:opacity-60"
                                disabled={imgLoading || form.images.length >= 10}
                            >
                                Вибрати зображення
                            </button>
                            {imgLoading && <span className="text-blue-500 text-sm">Завантаження...</span>}
                            {imgError && <span className="text-red-500 text-sm">{imgError}</span>}
                        </div>
                        <div className="flex flex-wrap gap-2 mt-2">
                            {form.images.map((img, idx) => (
                                <div key={img} className="relative w-20 h-20 border rounded overflow-hidden">
                                    <img
                                        src={
                                            (process.env.NEXT_PUBLIC_API_URL_IMG?.replace(/\/+$/, '') || '') +
                                            '/' +
                                            img.replace(/^\/+/, '')
                                        }
                                        alt={`img-${idx}`}
                                        className="object-cover w-full h-full"
                                    />
                                    <button
                                        type="button"
                                        onClick={() => handleRemoveImg(img)}
                                        className="absolute top-0 right-0 bg-white bg-opacity-70 rounded-bl px-2 py-0.5 text-xs text-red-600 hover:bg-opacity-100"
                                        aria-label="Видалити"
                                    >
                                        ✕
                                    </button>
                                </div>
                            ))}
                        </div>
                    </div>

                    <div className="flex flex-col gap-3 max-w-sm w-full">
                        <div>
                            <label className="block mb-1 font-medium pl-1 text-[0.75em]">Заголовок</label>
                            <input
                                name="title"
                                className="input input-bordered w-full p-2 bg-[#f2f4f5] rounded-md"
                                placeholder="Заголовок"
                                value={form.title}
                                required
                                maxLength={200}
                                onChange={handleChange}
                                autoComplete="off"
                            />
                        </div>

                        <div>
                            <label className="block mb-1 font-medium pl-1 text-[0.75em]">Опис</label>
                            <textarea
                                name="description"
                                className="input input-bordered w-full p-2 bg-[#f2f4f5] rounded-md"
                                placeholder="Опис"
                                value={form.description}
                                required
                                maxLength={1000}
                                rows={2}
                                onChange={handleChange}
                            />
                        </div>

                        <div>
                            <label className="block mb-1 font-medium pl-1 text-[0.75em]">Категорія</label>
                            <select
                                required
                                name="category"
                                value={form.category_id}
                                onChange={handleChange}
                                className="input input-bordered w-full p-2 bg-[#f2f4f5] rounded-md"
                            >
                                <option value="">Оберіть категорію...</option>
                                {(Array.isArray(categories) ? categories : []).map(cat =>
                                    <option key={cat.id} value={cat.id}>
                                        {cat.name || cat.title}
                                    </option>
                                )}
                            </select>
                        </div>

                        <div>
                            <label className="block mb-1 font-medium pl-1 text-[0.75em]">Ціна (грн)</label>
                            <input
                                name="price"
                                type="number"
                                className="input input-bordered w-full p-2 bg-[#f2f4f5] rounded-md"
                                placeholder="Ціна (грн)"
                                value={form.price}
                                required
                                min={0}
                                step={1}
                                onChange={handleChange}
                                autoComplete="off"
                            />
                        </div>

                        {message && (
                            <div className={`text-center text-sm mt-1 ${message.includes('успішно') ? 'text-green-500' : 'text-red-500'}`}>{message}</div>
                        )}

                        <button
                            type="submit"
                            disabled={loading}
                            className="w-full bg-blue-600 hover:bg-blue-700 mt-1 text-white py-2 rounded font-semibold transition-colors disabled:opacity-60"
                        >
                            {loading ? 'Публікація...' : 'Створити'}
                        </button>                    </div>
                </div>
            </form>
        </div>
    );
}
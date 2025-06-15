'use client';
import { useState } from 'react';

export default function RegisterForm({ onSuccess }) {
    const [email, setEmail] = useState('');
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [repeat, setRepeat] = useState('');
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();
        setError('');
        if (password !== repeat) {
            setError('Паролі не співпадають');
            return;
        }
        setLoading(true);
        try {
            const apiBase = process.env.NEXT_PUBLIC_API_URL;
            const res = await fetch(`${apiBase}/users/register`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ email, username, password }),
            });
            const data = await res.json();
            if (res.ok && data.token) {
                localStorage.setItem('token', data.token);
                onSuccess?.(data);
            } else {
                setError(data.message || 'Помилка, спробуйте ще раз!');
            }
        } catch {
            setError('Помилка зєднання');
        }
        setLoading(false);
    };

    return (
        <form
            onSubmit={handleSubmit}
            className="w-full max-w-sm mx-auto bg-white rounded-xl shadow-md p-8 flex flex-col gap-4"
        >
            <input
                name="email"
                type="email"
                className="input input-bordered w-full px-3 py-2 rounded-md border-gray-300"
                value={email}
                required
                onChange={e => setEmail(e.target.value)}
                placeholder="Ваш email"
            />
            <input
                name="username"
                type="text"
                className="input input-bordered w-full px-3 py-2 rounded-md border-gray-300"
                value={username}
                required
                minLength={3}
                maxLength={32}
                autoCapitalize="off"
                autoComplete="username"
                onChange={e => setUsername(e.target.value)}
                placeholder="Ім'я користувача"
            />
            <input
                name="password"
                type="password"
                className="input input-bordered w-full px-3 py-2 rounded-md border-gray-300"
                value={password}
                required
                autoComplete="new-password"
                onChange={e => setPassword(e.target.value)}
                placeholder="Пароль"
            />
            <input
                name="repeat_password"
                type="password"
                className="input input-bordered w-full px-3 py-2 rounded-md border-gray-300"
                value={repeat}
                required
                autoComplete="new-password"
                onChange={e => setRepeat(e.target.value)}
                placeholder="Повторіть пароль"
            />
            {error && <div className="text-red-500 text-sm">{error}</div>}
            <button
                type="submit"
                disabled={loading}
                className="w-full bg-blue-600 hover:bg-blue-700 text-white py-2 rounded font-semibold transition-colors disabled:opacity-60"
            >
                {loading ? 'Завантаження...' : 'Зареєструватися'}
            </button>
        </form>
    );
}
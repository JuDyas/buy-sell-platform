'use client';
import {useState} from 'react';
import {useUser} from "@/context/UserContext";

export default function LoginForm({ onSuccess }) {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState('');
    const { fetchUser } = useUser();

    const handleSubmit = async (e) => {
        e.preventDefault();
        setError('');
        setLoading(true);
        try {
            const apiBase = process.env.NEXT_PUBLIC_API_URL;
            const res = await fetch(`${apiBase}/users/login`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ email, password }),
            });
            const data = await res.json();
            if (res.ok && data.token) {
                localStorage.setItem('token', data.token);
                await fetchUser();
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
                name="password"
                type="password"
                className="input input-bordered w-full px-3 py-2 rounded-md border-gray-300"
                value={password}
                required
                onChange={e => setPassword(e.target.value)}
                placeholder="Пароль"
                autoComplete="current-password"
            />
            {error && <div className="text-red-500 text-sm">{error}</div>}
            <button
                type="submit"
                disabled={loading}
                className="w-full bg-blue-600 hover:bg-blue-700 text-white py-2 rounded font-semibold transition-colors disabled:opacity-60"
            >
                {loading ? 'Завантаження...' : 'Увійти'}
            </button>
        </form>
    );
}
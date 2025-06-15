'use client';
import LoginForm from '@/components/Auth/LoginForm';
import {useAuthRedirect} from "@/hooks/useAuthRedirect";

export default function LoginPage() {
    const { user, loading } = useAuthRedirect();

    if (loading) {
        return (
            <div className="w-full min-h-screen flex items-center justify-center text-gray-600 text-xl">
                Завантаження...
            </div>
        );
    }

    if (!user) {
        return (
        <div className="min-h-screen flex flex-col justify-center items-center bg-gray-50">
            <div className="mb-6 text-center">
                <h1 className="text-3xl font-bold mb-2">Вхід в аккаунт</h1>
            </div>
            <LoginForm />
            <div className="mt-8">
                <span className="text-gray-600 text-sm">
                    Немає аккаунта?{' '}
                    <a href="/register" className="text-blue-600 underline">
                        Зареєструватися
                    </a>
                </span>
            </div>
        </div>
    );
    }
}
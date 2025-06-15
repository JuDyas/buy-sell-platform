'use client';
import RegisterForm from '@/components/Auth/RegisterForm';
import {useAuthRedirect} from "@/hooks/useAuthRedirect";

export default function RegisterPage() {
    const { user, loading } = useAuthRedirect({requireAuth: false});

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
                    <h1 className="text-3xl font-bold mb-2">Реєстрація</h1>
                </div>
                <RegisterForm />
                <div className="mt-8">
        <span className="text-gray-600 text-sm">
          Вже є аккаунт?{' '}
            <a href="/login" className="text-blue-600 underline">
            Увійти
          </a>
        </span>
                </div>
            </div>
        );
    }
}
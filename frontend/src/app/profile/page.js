'use client';
import UserProfileForm from '@/components/User/UserProfileForm';

export default function ProfilePage() {
    return (
        <div className="min-h-[80vh] flex flex-col justify-center items-center bg-[#f2f4f5]">
            <UserProfileForm />
        </div>
    );
}
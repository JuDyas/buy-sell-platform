'use client';

import { useParams } from 'next/navigation';
import AdvertDetails from '@/components/Advert/AdvertDetails';

export default function AdvertPage() {
    const params = useParams();
    const advertId = Array.isArray(params?.id) ? params.id[0] : params.id;

    return (
        <AdvertDetails advertId={advertId} />
    );
}
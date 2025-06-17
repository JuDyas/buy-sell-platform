import SearchBar from '@/components/Advert/SearchBar';
import AdvertsGrid from '@/components/Advert/AdvertGrid';

export default function Page() {
    return (
        <div className="min-h-[83vh] 2xl:min-h-[87vh]">
            <SearchBar />
            <AdvertsGrid />
        </div>
    );
}
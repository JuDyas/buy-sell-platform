import AdvertsGrid from "@/components/Advert/AdvertGrid";
import SearchBar from "@/components/Advert/SearchBar";

export default function CategoryPage({ params }) {
    const { id } = params;
    return (
        <div className="min-h-[83vh] 2xl:min-h-[87vh]">
            <SearchBar />
            <AdvertsGrid categoryId={id} />
        </div>
    );
}
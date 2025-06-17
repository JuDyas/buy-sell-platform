import AdvertsGrid from "@/components/Advert/AdvertGrid";

export default function CategoryPage({ params }) {
    const { id } = params;
    return (
        <div className="min-h-[85vh]">
            <AdvertsGrid categoryId={id} />
        </div>
    );
}
import Image from "next/image";

const SOCIALS = [
    {
        href: "https://www.linkedin.com/in/denys-holovkin-287752302",
        icon: "/icons/linkedin.svg",
        alt: "LinkedIn",
        label: "LinkedIn",
    },
    {
        href: "https://github.com/JuDyas",
        icon: "/icons/github.svg",
        alt: "GitHub",
        label: "GitHub",
    },
    {
        href: "https://discord.com/",
        icon: "/icons/discord.svg",
        alt: "Discord",
        label: "Discord",
    },
];

export default function Footer() {
    return (
        <footer className="border-t">
            <div className="container mx-auto px-4 py-4 flex items-center justify-between">
        <span className="text-sm text-gray-500 select-none">
          © 2025 Denys Holovkin
        </span>
                <div className="flex items-center gap-4">
                    {SOCIALS.map(({ href, icon, alt, label }) => (
                        <a
                            key={href}
                            href={href}
                            target="_blank"
                            rel="noopener noreferrer"
                            aria-label={label}
                        >
                            <Image src={icon} alt={alt} width={28} height={28} />
                        </a>
                    ))}
                </div>
            </div>
        </footer>
    );
}
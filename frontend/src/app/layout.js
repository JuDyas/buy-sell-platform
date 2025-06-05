import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";
import Header from "@/components/Layout/Header"
import Footer from "@/components/Layout/Footer";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata = {
  title: "GoSell",
  description: "Виставляйте оголошення та купуйте вигідно!",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body className="antialiased">
      <Header />
        {children}
      <Footer />
      </body>
    </html>
  );
}

import "./globals.css";
import Header from "@/components/Layout/Header"
import Footer from "@/components/Layout/Footer";
import {UserProvider} from "@/context/UserContext";

export const metadata = {
  title: "GoSell",
  description: "Виставляйте оголошення та купуйте вигідно!",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body className="antialiased bg-[#f2f4f5]">
        <UserProvider>
            <Header />
                {children}
            <Footer />
        </UserProvider>
      </body>
    </html>
  );
}

import Header from '@/components/header/Header'
import './globals.css'
import type { Metadata } from 'next'
import Footer from '@/components/footer/Footer'

export const metadata: Metadata = {
  title: 'הטכנולוג | קהילות ופורומים טכנולוגיים',
  description: 'בואו לתכנת, זה כיף! עם הקהילה הגדולה במדינה',
}

export default async function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
        <body dir="rtl" className="m-0 p-0 text-white bg-purple-950">
          <div className="w-3/4 container mx-auto">
            <div className="font-alef flex flex-col items-start my-6 gap-12">
              <Header/>
              {children}
              <Footer/>
            </div>
          </div>
        </body>
    </html>
  )
}

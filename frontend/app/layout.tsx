import './globals.css'
import Header from '@/components/pages/layout/Header'
import Footer from '@/components/pages/layout/Footer'

export default async function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
        <body dir="rtl" className="m-0 p-0 bg-primary text-words noam-regular">
            <div className="w-11/12 container mx-auto flex flex-col gap-6">
                <Header/>
                {children}
                <Footer/>
            </div>
        </body>
    </html>
  )
}

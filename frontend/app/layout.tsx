import './globals.css'
import Header from '@/components/pages/layout/Header'
import Footer from '@/components/pages/layout/Footer'
import Script from 'next/script'

export default async function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
        <Script async src={`https://www.googletagmanager.com/gtag/js?id=${process.env.GA_MEASUREMENT_ID}`}></Script>
        <Script id="google-analytics">
          {`
            window.dataLayer = window.dataLayer || [];
            function gtag(){dataLayer.push(arguments);}
            gtag('js', new Date());
            gtag('config', '${process.env.GA_MEASUREMENT_ID}');
          `}
        </Script  >
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

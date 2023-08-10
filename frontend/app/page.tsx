import type { Metadata } from 'next'
import Hero from "@/components/pages/index/Hero"

export const revalidate = 36000;

export const metadata: Metadata = {
  metadataBase: new URL("https://hatechnolog.com"),
  title: 'הטכנולוג - קהילות ופורומים טכנולוגיים',
  description: `ללמוד לתכנת לבד, להתקבל להייטק או פשוט להנות מטכנולוגיה - זה המקום עבור כולם. בואו לתכנת, זה כיף! עם הקהילה הטכנולוגית הגדולה במדינה'`,
  publisher: "הטכנולוג",
  alternates: {
      canonical: `/`
  },
  openGraph: {
      title: `הטכנולוג - קורסים`,
      description: `ללמוד לתכנת לבד, להתקבל להייטק או פשוט להנות מטכנולוגיה - זה המקום עבור כולם. בואו לתכנת, זה כיף! עם הקהילה הטכנולוגית הגדולה במדינה'`,
      url: "https://hatechnolog.com",
      siteName: "הטכנולוג",
      images: {
          url: "https://hatechnolog.fra1.cdn.digitaloceanspaces.com/technolog_logo.png",
          secureUrl: "https://hatechnolog.fra1.cdn.digitaloceanspaces.com/technolog_logo.png",
          alt: "Hatechnolog Logo",
          width: 256,
          height: 256,
      }
  },
  twitter: {
      site: "https://hatechnolog.com",
      creator: "@hatechnolog",
      description: `ללמוד לתכנת לבד, להתקבל להייטק או פשוט להנות מטכנולוגיה - זה המקום עבור כולם. בואו לתכנת, זה כיף! עם הקהילה הטכנולוגית הגדולה במדינה'`,
      title: 'הטכנולוג - קהילות ופורומים טכנולוגיים',
      images: {
          url: "https://hatechnolog.fra1.cdn.digitaloceanspaces.com/technolog_logo.png",
          secureUrl: "https://hatechnolog.fra1.cdn.digitaloceanspaces.com/technolog_logo.png",
          alt: "Hatechnolog Logo",
          width: 256,
          height: 256,
      }
  }
}

export default function HomePage() {
  return (
    <Hero/>
  )
}

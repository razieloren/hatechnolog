import type { Metadata } from 'next'
import Hero from "@/components/pages/index/Hero"

export const metadata: Metadata = {
  title: 'הטכנולוג - קהילות ופורומים טכנולוגיים',
  description: 'בואו לתכנת, זה כיף! עם הקהילה הטכנולוגית הגדולה במדינה',
}

export default function HomePage() {
  return (
    <Hero/>
  )
}

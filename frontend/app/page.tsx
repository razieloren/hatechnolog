import Header from "./components/Header"
import { Alef } from "next/font/google"
import Navbar from "./components/Navbar"
import LatestStats from "./components/LatestStats"
import Hero from "./components/Hero"

const ALEF_FONT = Alef({subsets: ['hebrew'], weight: '400'})

export default function Home() {
  return (
    <main className={`main ${ALEF_FONT.className}`}>
      <Header/>
      <Navbar/>
      <Hero/>
      <LatestStats/>
    </main>
  )
}

import Header from "./components/Header"
import { Alef } from "next/font/google"
import Navbar from "./components/Navbar"
import LatestStats from "./components/LatestStats"

const ALEF_FONT = Alef({subsets: ['hebrew'], weight: '400'})

export default function Home() {
  return (
    <div className={ALEF_FONT.className}>
      <Header/>
      <Navbar/>
      <LatestStats/>
    </div>
  )
}

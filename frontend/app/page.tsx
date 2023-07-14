import Header from "./components/Header"
import { Alef } from "next/font/google"
import Navbar from "./components/Navbar"

const ALEF_FONT = Alef({subsets: ['hebrew'], weight: '400'})

export default function Home() {
  return (
    <div className={ALEF_FONT.className}>
      <Header/>
      <Navbar/>
    </div>
  )
}

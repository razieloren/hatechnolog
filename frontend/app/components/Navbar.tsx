import Image from 'next/image'
import mainLogo from '../../public/technolog_logo.png'

export default function Navbar() {
    return (
        <nav className="nav">
            <ul className="nav-items">
                <li>בואו לתכנת, זה כיף</li>
                <li>דעתי על דברים</li>
                <li>מחסן</li>
            </ul>
        </nav>
    )
}
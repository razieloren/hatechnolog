import Image from 'next/image'
import mainLogo from '../../public/technolog_logo.png'

export default function Header() {
    return (
        <section className="section">
            <div className="header-main">
                <Image className="header-main-logo" src={mainLogo} alt="Hatechnolog Logo"/>
                <h1 className="header-main-title">הטכנולוג</h1>
            </div>
        </section>
    )
}
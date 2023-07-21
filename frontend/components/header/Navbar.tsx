import NavUser from "./NavUser";

export default function Navbar() {
    return (
        <nav className="responsive-flex justify-between">
            <div className="flex sm:gap-10 gap-4 sm:text-xl text-l font-bold">
                <button className="nav-button">
                    בלוג
                </button>
                <button className="nav-button">
                    קרמה
                </button>
                <button className="nav-button">
                    צרו קשר
                </button>
            </div>
            <NavUser/>
        </nav>
    )
}
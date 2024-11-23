const Navbar = ({ items, selectedItem=0 }) => {
    const selected = selectedItem;

    return (
    <ul className="nav nav-underline nav-justified mb-3">
        {
            items.map((item, index) => <li
                key={index}
                className="nav-item"
                onClick={item.onClick}
                >
                    <a href="#"
                        className={selected === index ? "nav-link active": "nav-link"}
                        aria-current={selected === index ? "page" : false}
                    >{item.name}</a>
            </li>)
        }
    </ul>
    );
}

export default Navbar;
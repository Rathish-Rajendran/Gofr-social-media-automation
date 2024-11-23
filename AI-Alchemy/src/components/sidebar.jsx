const SideBar = ({ items, selectedItem=0 }) => {
    const selected = selectedItem

    return (
    <nav className="nav flex-column align-self-start nav-pills">
        {
            items.map((item, index) => <a key={index}
                href="#"
                className={selected === index ? "nav-link active m-1" : "nav-link"}
                aria-current={selected === index ? "page" : false}
                onClick={() => item.onClick(index)}
                >
                    {item.name}
                </a>
            )
        }
    </nav>
    );
}

export default SideBar;
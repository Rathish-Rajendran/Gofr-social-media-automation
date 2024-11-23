import { useState } from "react";

const SideBar = ({ items }) => {
    const [selected, setSelected] = useState(0)

    return (
    <nav className="nav flex-column align-self-start nav-pills">
        {
            items.map((item, index) => <a key={index}
                href="#"
                className={selected === index ? "nav-link active m-1" : "nav-link"}
                aria-current={selected === index ? "page" : false}
                onClick={() => setSelected(index)}
                >
                    {item}
                </a>
            )
        }
    </nav>
    );
}

export default SideBar;
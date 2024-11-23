import { useState } from "react";

const Navbar = ({ items }) => {
    const[selected, setSelected] = useState(0)

    return (
    <ul class="nav nav-underline nav-justified mb-3">
        {
            items.map((item, index) => <li
                key={index}
                className="nav-item"
                onClick={() => setSelected(index)}
                >
                    <a href="#"
                        className={selected === index ? "nav-link active": "nav-link"}
                        aria-current={selected === index ? "page" : false}
                    >{item}</a>
            </li>)
        }
    </ul>
    );
}

export default Navbar;
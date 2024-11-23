const SideBar = () => {
    return (
    <nav className="nav flex-column align-self-start nav-pills">
        <a className="nav-link" aria-current="page" href="#">Active</a>
        <a className="nav-link active" href="#">Link</a>
        <a className="nav-link" href="#">Link</a>
        <a className="nav-link disabled" aria-disabled="true">Disabled</a>
    </nav>
    );
}

export default SideBar;
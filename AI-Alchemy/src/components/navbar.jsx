const Navbar = () => {
    return (
    <ul class="nav nav-tabs nav-justified mb-3">
        <li class="nav-item">
          <a class="nav-link" aria-current="page" href="#">Active</a>
        </li>
        <li class="nav-item">
          <a class="nav-link active" href="#">Link</a>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="#">Link</a>
        </li>
        <li class="nav-item">
          <a class="nav-link disabled" aria-disabled="true">Disabled</a>
        </li>
    </ul>
    );
}

export default Navbar;
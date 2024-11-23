import { useState } from 'react'
import "bootstrap/dist/css/bootstrap.css"

function App() {
  const [count, setCount] = useState(0)

  return (
    <div className="container-fluid">
      <div className="container-fluid">
        <ul class="nav nav-tabs mb-3">
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
      </div>
      <div className="row">
        <nav className="nav flex-column col-2 align-self-start nav-pills">
          <a className="nav-link" aria-current="page" href="#">Active</a>
          <a className="nav-link active" href="#">Link</a>
          <a className="nav-link" href="#">Link</a>
          <a className="nav-link disabled" aria-disabled="true">Disabled</a>
        </nav>
        <h1 className='col-6'>Hello world</h1>
      </div>
    </div>
  )
}

export default App

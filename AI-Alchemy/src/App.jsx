import "bootstrap/dist/css/bootstrap.css"
import Navbar from './components/navbar'
import SideBar from './components/sidebar'

function App() {
  const navbarItems = [
    "Approve to send",
    "Analytics"
  ]

  const sidebarItems = [
    "LinkedIn",
    "Twitter / X",
    "Mails"
  ]

  return (
    <div className="container-fluid">
      <div className="container-fluid row">
        <div className="col-2 container-fluid">
          <SideBar items={sidebarItems} />
        </div>
        <div className="col container-fluid">
          <Navbar items={navbarItems} />
          <h1>Hello World!!!</h1>
        </div>
      </div>
    </div>
  )
}

export default App

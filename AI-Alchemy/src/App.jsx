import "bootstrap/dist/css/bootstrap.css"
import Navbar from './components/navbar'
import SideBar from './components/sidebar'
import ContentPage from "./components/contentpage"
import { useState } from "react"

function App() {
  // showCards is a bool controlling if we should show
  // the cards or the analytics
  const[showCards, setShowCards] = useState(true)
  const navbarItems = [
    { name: "Approve to send", onClick: () => setShowCards(true)},
    { name: "Analytics", onClick: () => setShowCards(false) }
  ]

  const sidebarItems = [
    "LinkedIn",
    "Twitter / X",
    "Mails"
  ]

  // Some dummy data for LinkedIn
  const linkedInContents = [
    {
      heading: "Heading 1",
      body: "Body 1, abcdefghij"
    },
    {
      heading: "Heading 2",
      body: "Body 2, abcdefghij"
    },
    {
      heading: "Heading 2",
      body: "Body 2, abcdefghij"
    },
  ]

  return (
    <div className="container-fluid">
      <div className="container-fluid row">
        <div className="col-2 container-fluid">
          <SideBar items={sidebarItems} />
        </div>
        <div className="col container-fluid">
          <Navbar items={navbarItems} selectedItem={showCards ? 0 : 1} />
          <ContentPage showItems={showCards} items={linkedInContents} />
        </div>
      </div>
    </div>
  )
}

export default App

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

  const [selectedSidebarItem, setSelectedSidebarItem] = useState(0)
  const sidebarItems = [
    { name: "LinkedIn", onClick: setSelectedSidebarItem },
    { name: "Twitter / X", onClick: setSelectedSidebarItem },
    { name: "Mails", onClick: setSelectedSidebarItem }
  ]

  // Some dummy data for LinkedIn
  const linkedInContents = [
    {
      heading: "Heading LinkedIn 1",
      body: "Body 1, abcdefghij"
    },
    {
      heading: "Heading LinkedIn 2",
      body: "Body 2, abcdefghij"
    },
    {
      heading: "HeadingLinkedIn 2",
      body: "Body 2, abcdefghij"
    },
  ]

  // Some dummy data for Twitter
  const twitterContents = [
    {
      heading: "Heading Twitter / X 1",
      body: "Body 1, abcdefghij"
    },
    {
      heading: "Heading Twitter / X 2",
      body: "Body 2, abcdefghij"
    },
    {
      heading: "HeadingTwitter / X 2",
      body: "Body 2, abcdefghij"
    },
  ]

  // Some mail mock data
  const mailContents = []

  const getItemsToShow = () => {
    if (selectedSidebarItem === 0) return linkedInContents;
    if (selectedSidebarItem === 1) return twitterContents;
    return mailContents;
  }

  return (
    <div className="container-fluid">
      <div className="container-fluid row">
        <div className="col-2 container-fluid">
          <SideBar items={sidebarItems} selectedItem={selectedSidebarItem} />
        </div>
        <div className="col container-fluid">
          <Navbar items={navbarItems} selectedItem={showCards ? 0 : 1} />
          <ContentPage showItems={showCards}
            items={getItemsToShow()}
          />
        </div>
      </div>
    </div>
  )
}

export default App

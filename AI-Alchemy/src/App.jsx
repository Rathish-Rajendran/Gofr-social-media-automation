import "bootstrap/dist/css/bootstrap.css"
import Navbar from './components/navbar'
import SideBar from './components/sidebar'
import ContentPage from "./components/contentpage"
import { useEffect, useState } from "react"

const BACKEND_URL="http://localhost:8000"

const sendTwitterPostRequest = async (heading, body) => {
  const content = `${heading}\n\n${body}`; // Combine heading and body
  try {
    const response = await fetch(BACKEND_URL + "/tweet", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ content }),
    });

    if (response.ok) {
      console.log("Content sent successfully:", content);
    } else {
      console.error("Failed to send content:", response.statusText);
    }
  } catch (error) {
    console.error("Error while sending content:", error);
  }
}

const sendMail = async (from, body, subject) => {
  const content = {
    content: {
      from, body, subject
    }
  }; // Combine heading and body
  console.log(JSON.stringify(content))
  try {
    const response = await fetch(BACKEND_URL + "/googleGroupReply", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(content),
    });

    if (response.ok) {
      console.log("Content sent successfully:", content);
    } else {
      console.error("Failed to send content:", response.statusText);
    }
  } catch (error) {
    console.error("Error while sending content:", error);
  }
}

// let isMounted = true; // To prevent setting state after unmount

function App() {
  const [mailData, setMailData] = useState([])
  const [displayMail, setDisplayMail] = useState(false);

  useEffect(() => {
    const fetchData = async () => {

      try {
        const response = await fetch(BACKEND_URL + "/googleGroup");
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        const result = await response.json();
        const data = JSON.parse(result["data"])["output"]
        data.forEach(element => {
          element.onSend = () => sendMail(element.from, element.body, element.subject)
        });
        setMailData(data);
      } catch (err) {
          console.log(err)
      }
    };

    fetchData();

  }, []); // Empty dependency array ensures it only runs once

  // showCards is a bool controlling if we should show
  // the cards or the analytics
  const[showCards, setShowCards] = useState(true)
  const navbarItems = [
    { name: "Approve to send", onClick: () => setShowCards(true)},
    { name: "Analytics", onClick: () => setShowCards(false) }
  ]

  const [selectedSidebarItem, setSelectedSidebarItem] = useState(0)
  const sidebarItems = [
    { name: "LinkedIn", onClick: (index) => {
      setSelectedSidebarItem(index)
      setDisplayMail(false)
    } },
    { name: "Twitter / X", onClick: (index) => {
      setSelectedSidebarItem(index)
      setDisplayMail(false)
    } },
    { name: "Mails", onClick: (index) => {
        setSelectedSidebarItem(index)
        setDisplayMail(true)
      }
    }
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
      body: "Body 1, abcdefghij",
      onSend: () => sendTwitterPostRequest("Heading Twitter / X 1", "Body 1, abcdefghij")
    },
    {
      heading: "Heading Twitter / X 2",
      body: "Body 2, abcdefghij",
      onSend: () => sendTwitterPostRequest("Heading Twitter / X 2", "Body 2, abcdefghij")
    },
    {
      heading: "HeadingTwitter / X 2",
      body: "Body 2, abcdefghij",
      onSend: () => sendTwitterPostRequest("Heading Twitter / X 2", "Body 2, abcdefghij")
    },
  ]

  const getItemsToShow = () => {
    if (selectedSidebarItem === 0) return linkedInContents;
    if (selectedSidebarItem === 1) return twitterContents;
    return mailData;
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
            isMail={displayMail}
          />
        </div>
      </div>
    </div>
  )
}

export default App

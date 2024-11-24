import Markdown from "react-markdown";

const ContentPage = ({ items, showItems = true, isMail = false }) => {
  return (
    <div className="container-fluid row g-4 mt-4">
      {/* Show the card items */}
      {showItems &&
        items.map((item, index) => (
          <div key={index} className="card text-bg-dark m-1 col-md-5">
            <div className="card-body">
              <h5 className="card-title">
                {isMail ? item.subject : item.heading}
              </h5>
              <div className="card-text">
                <Markdown>{item.body}</Markdown>
              </div>
              <a href="#" className="btn btn-primary" onClick={item.onSend}>
                Send
              </a>
            </div>
          </div>
        ))}
      {showItems && items.length === 0 && (
        <h1>
          GoFr has gone hunting for trends and issues...
          <br />
          Please wait
        </h1>
      )}

      {/* Show the analytics */}
      {!showItems && <h1>No analytics at the moment</h1>}
    </div>
  );
};

export default ContentPage;

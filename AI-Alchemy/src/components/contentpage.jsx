const ContentPage = ({ items, showItems=true }) => {
    return (<div className="container-fluid row g-4 mt-4">
        { /* Show the card items */ }
        { showItems &&
            items.map((item, index) => (
                <div key={index} className="card text-bg-dark m-1 col-md-5">
                    <div className="card-body">
                    <h5 className="card-title">{item.heading}</h5>
                    <p className="card-text">{item.body}</p>
                    <a href="#" className="btn btn-primary">Approve</a>
                    </div>
                </div>
            ))
        }

        { /* Show the analytics */}
        {!showItems && <h1>No analytics at the moment</h1>}
    </ div>);
}

export default ContentPage;
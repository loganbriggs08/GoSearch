function WelcomeComponent() {
    return (
        <div id="Welcome" onContextMenu={(e) => { e.preventDefault() }}>
            <div id="header" className="header-div">
                <h1>Welcome 👋</h1>
                <p>Hello There.. Thanks for installing GoSearch, Please Read Below.</p>
                <hr className="rounded"/>
            </div>
            
            <div className="main-text">
                <h2>Information:</h2>
                <p>In order for you to use GoSearch we first need to cache your system </p>
            </div>
        </div>
        );
}

export default WelcomeComponent;
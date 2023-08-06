import {CloseApp, SetPage} from "../wailsjs/go/main/App";

function WelcomeComponent() {

    function closeApp() {CloseApp()}
    function setPage(page: string) {SetPage(page)}

    return (
        <div id="Welcome" onContextMenu={(e) => { e.preventDefault() }}>
            <div id="header" className="header-div">
                <h1>Welcome 👋</h1>
                <p>Hello There.. Thanks for installing GoSearch, Please Read Below.</p>
                <hr className="rounded"/>
            </div>

            <div className="information-container">
                <h2>Information:</h2>
                <p>In order for you to use GoSearch we first need to cache your system so we can provide you with file results at a super speed,
                    this means we need consent from you to do this as this could take a little time. We also need to add that to keep our cache as
                    speedy as possible we need to run GoSearch in the background to keep caching your files.</p>
            </div>

            <div className="continue-button-container">
                <button className="cancel-button" onClick={() => closeApp()}>Cancel</button>
                <button className="continue-button" onClick={() => setPage("Setup")}>Continue</button>
            </div>
        </div>
        );
}

export default WelcomeComponent;
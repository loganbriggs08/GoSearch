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
                <h2>Disclaimer:</h2>
                <p>By clicking the Continue button you agree to let us cache the locations of your files in order to ensure GoSearch
                functions as intended, you also agree to allow GoSearch to actively search for new files to be cached, this cache is
                stored on your system.</p>
            </div>

            <div className="continue-button-container">
                <button className="cancel-button" onClick={() => closeApp()}>Cancel</button>
                <button className="continue-button" onClick={() => setPage("Setup")}>Continue</button>
            </div>
        </div>
        );
}

export default WelcomeComponent;
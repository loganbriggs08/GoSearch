import {CloseApp} from "../wailsjs/go/main/App";

function SetupComponent() {
    return (
        <div id="Welcome" onContextMenu={(e) => { e.preventDefault() }}>
            <div id="header" className="header-div">
                <h1>Please Sit Tight... 😊</h1>
                <p>We are currently setting up, this could take a while so maybe get a coffee.</p>
                <hr className="rounded"/>
            </div>

            <div className="information-container">
                <h2>Information:</h2>
                <p>We are currently caching the location of all your files, this makes it easier for us to provide you
                    with faster and more accurate search results, however this could take a little bit of time so please
                    be patient and take your time to do something else. It is also important to add that we do not send
                    your files anywhere and our cache stays on your system.

                    <br/><br/>Please don't turn off your computer.</p>
            </div>
        </div>
        );
}

export default SetupComponent;
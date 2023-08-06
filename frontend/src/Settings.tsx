import { ClearCache, SetPage } from "../wailsjs/go/main/App"

function SettingsComponent() {
    function deleteCache() {
        ClearCache().then((response: boolean) => {
            if (response) {
                SetPage("Search")
            }
        })
    }

    function reloadApp() {
        SetPage("Search")
    }

    return (
        <div id="Settings" onContextMenu={(e) => { e.preventDefault() }}>
            <div id="header" className="header-div">
                <h1>Settings:</h1>
                <p>The Settings below are automatically saved after being updated.</p>
                <hr className="rounded"/>
            </div>

            <div className="main-settings">
                <div className="dropdown-container">
                    <h2>Select a Theme:</h2>
                    <p>Select a theme from the dropdown below.</p>

                </div>

                <div className="dangerous-actions-container">
                    <h2>Dangerous Actions:</h2>
                    <button className="clear-cache-button" onClick={() => deleteCache()}>Clear Cache</button>
                    <button className="restart-app-button" onClick={() => reloadApp()}>Restart App</button>
                </div>
            </div>
        </div>
        );
}

export default SettingsComponent;
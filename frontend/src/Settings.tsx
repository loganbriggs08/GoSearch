import React, {useState} from "react";
import { ClearCache, SetPage } from "../wailsjs/go/main/App"

function SettingsComponent() {
    const [selectedOption, setSelectedOption] = useState('');

    let themes = ["Blue Theme", "Cherry Blossom Theme", "Lavender Theme", "Mint Theme"]

    const handleOptionChange = (e: any) => {
        setSelectedOption(e.target.value);

        console.log(e.target.value)
    };
    
    function deleteCache() {
        ClearCache().then((response: boolean) => {
            if (response) {
                SetPage("Search")
            }
        })
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
                    <select value={selectedOption} onChange={handleOptionChange}>
                        <option value="Default Theme">Default Theme</option>
                        {themes.map((theme) => (
                            <option value={theme}>{theme}</option>
                        ))}

                    </select>
                </div>

                <div className="dangerous-actions-container">
                    <h2>Dangerous Actions:</h2>
                    <button onClick={() => deleteCache()}>Clear Cache</button>
                </div>
            </div>
        </div>
        );
}

export default SettingsComponent;
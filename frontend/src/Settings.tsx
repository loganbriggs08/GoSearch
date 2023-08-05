import React, {useState} from "react";

function SettingsComponent() {
    const [selectedOption, setSelectedOption] = useState('');

    const handleOptionChange = (e: any) => {
        setSelectedOption(e.target.value);
    };
    
    return (
        <div id="Settings" onContextMenu={(e) => { e.preventDefault() }}>
            <div id="header" className="header-div">
                <h1>Settings:</h1>
                <p>The Settings below are automatically saved after being updated.</p>
                <hr class="rounded"/>
            </div>

            <div className="dropdown-container">
                <h2>Select a Theme:</h2>
                <p>Select a theme for GoSearch from the dropdown below.</p>
                <select value={selectedOption} onChange={handleOptionChange}>
                    <option value="">Select a Theme...</option>
                    <option value="option1">Option 1</option>
                    <option value="option2">Option 2</option>
                    <option value="option3">Option 3</option>
                </select>
            </div>

            <div id="settings-footer" className="settings-footer-div">
                <h1>Made with ❤️ by Katsu.</h1>
            </div>
        </div>
        );
}

export default SettingsComponent;
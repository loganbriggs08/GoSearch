import React, {useState} from "react";

function SettingsComponent() {
    
    return (
        <div id="App" onContextMenu={(e) => { e.preventDefault() }}>
            <div id="header" className="header-div">
                <h1>GoSearch Settings:</h1>
                <p>The Settings below are automatically saved after being updated.</p>
            </div>
            
            <div className="dropdown">
                
            </div>

            <div id="footer" className="footer-div">
                <h1>Settings are saved Automatically.</h1>
            </div>
        </div>
        );
}

export default SettingsComponent;
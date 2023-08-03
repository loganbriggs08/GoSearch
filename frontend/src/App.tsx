import './App.css';
import React, { useState, useEffect } from "react";
import { Search, HandleButtonClickEvent } from "../wailsjs/go/main/App";

export interface FileReturnStruct {
    Name:         string;
    Location:     string;
    IconLocation: string;
    Favorite:     boolean;
    Visits:       number;
}

function App() {
    const [results, setResults] = useState<FileReturnStruct[]>([]);
    let debounceTimer: number | null = null;

    useEffect(() => {
        fetchSearchResults("");
        }, []);

    function fetchSearchResults(query: string) {
        if (debounceTimer !== null) {
            clearTimeout(debounceTimer);
        }

        debounceTimer = window.setTimeout(() => {
            Search(query).then((response: FileReturnStruct[] | null) => {
                if (response === null || (response.length === 1 && response[0].Name === "")) {
                    setResults([]);
                } else {
                    setResults(response);
                }
            });
            }, 1000);
    }

    const handleButtonClick = (result: FileReturnStruct) => {
        HandleButtonClickEvent(result);
        setResults([]);
    }

    return (
        <div id="App">
            <div id="input" className="search-box">
                <input id="name" className="search-input" onChange={(e) => fetchSearchResults(e.target.value)} name="input" type="text" spellCheck="false" placeholder="Search..." />
            </div>

            <div id="results" className="results-div">
                {results.map((result) => (
                    <div key={result.Name} id="resultcard" className="result-card">
                        <button onClick={() => handleButtonClick(result)}>
                            <img src="https://cdn.discordapp.com/attachments/759689232326328323/1136057337015185559/Google_Chrome_icon_February_2022.svg.png" alt="result icon" />
                            <h1>{result.Name}</h1>

                            <button>Click Me</button>
                        </button>
                    </div>
                ))}
            </div>

            <div id="footer" className="footer-div">
                {results.length === 0 && <h1>No Results found</h1>}
                {results.length === 1 && <h1>{results.length} Result found</h1>}
                {results.length <= 3 && results.length !== 0 && results.length !== 1 && <h1>{results.length} Results found</h1>}
                {results.length >= 4 && <h1>{results.length} Results found - Scroll to see all results</h1>}
            </div>
        </div>
        );
}

export default App;
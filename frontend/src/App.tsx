import './App.css';
import React, { useState, useEffect } from "react";
import { Search, HandleButtonClickEvent } from "../wailsjs/go/main/App";

export interface RecommendedAppStruct {
    Name: string;
    Location: string;
    Visits: number;
}

function App() {
    const [results, setResults] = useState<RecommendedAppStruct[]>([]);

    function fetchSearchResultsSearch(query: any) {
        Search(query).then((response: RecommendedAppStruct[]) => {
            if (response.length === 1 && response[0].Name === "") {
                setResults([]);
            } else {
                setResults(response)
            }
        });
    }

    const UpdateSearch = (e: any) => fetchSearchResultsSearch(e.target.value);

    useEffect(() => {
        UpdateSearch({ target: { value: "" } });
        }, [])

    const handleButtonClick = (result: object) => {
        HandleButtonClickEvent(result)

        useEffect(() => {
            UpdateSearch({ target: { value: "" } });
            }, [])

    }

    return (
        <div id="App">
            <div id="input" className="search-box">
                <input id="name" className="search-input" onChange={UpdateSearch} name="input" type="text" spellCheck="false" placeholder="Search..." />
            </div>

            <div id="results" className="results-div">
                {results.map((result) => (
                    <div key={result.Name} id="resultcard" className="result-card">
                        <button onClick={() => handleButtonClick(result)}>
                            <img src="https://cdn.discordapp.com/attachments/759689232326328323/1136057337015185559/Google_Chrome_icon_February_2022.svg.png" alt="result icon" />
                            <h1>{result.Name}</h1>
                        </button>
                    </div>
                ))}
            </div>

            <div id="footer" className="footer-div">
                {results.length === 0 && <h1>No Results found</h1>}
                {results.length === 1 && <h1>{results.length} Result found</h1>}
                {results.length <= 3 && results.length != 0 && results.length != 1 && <h1>{results.length} Results found</h1>}
                {results.length >= 4 && <h1>{results.length} Results found - Scroll to see all results</h1>}
            </div>
        </div>
    );
}

export default App;
import './App.css';
import {Search} from "../wailsjs/go/main/App";

function App() {
    let apps = ["Google Chrome", "Brave Browser", "FireFox"];

    function printSearch(name: any) {
        Search(name)
    }
    const UpdateSearch = (e: any) => printSearch(e.target.value);

    return (
        <div id="App">
            <div id="input" className="search-box">
                <input id="name" className="search-input" onChange={UpdateSearch} name="input" type="text" spellCheck="false" placeholder="Search..." />
            </div>

            <div id="results" className="results-div">
                {apps.map((item, index) => (
                    <div id="resultcard" className="result-card">
                        <button>
                            <img src="https://cdn.discordapp.com/attachments/759689232326328323/1136057337015185559/Google_Chrome_icon_February_2022.svg.png" alt="result icon"/>
                            <h1>{item}</h1>
                        </button>
                    </div>
                    ))}
            </div>

            <div id="footer" className="footer-div">
                <h1>No Results found</h1>
            </div>
        </div>
    )
}

export default App

import './App.css';
import {Search} from "../wailsjs/go/main/App";

function App() {

    function printUpdate(name: any) {
        Search(name)
    }
    const updateName = (e: any) => printUpdate(e.target.value);

    return (
        <div id="App">
            <div id="input" className="search-box">
                <input id="name" className="search-input" onChange={updateName} name="input" type="text" spellCheck="false" placeholder="Search..." />
            </div>

            <div id="results" className="results-div">

            </div>

            <div id="footer" className="footer-div">
                <h1>No Results found</h1>
            </div>
        </div>
    )
}

export default App

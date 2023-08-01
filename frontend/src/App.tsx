import {useState} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import {Search} from "../wailsjs/go/main/App";

function App() {
    const [name, setName] = useState('');
    const updateName = (e: any) => setName(e.target.value);

    function greet() {
        Search(name);
    }

    return (
        <div id="App">
            <div id="input" className="input-box">
                <input id="name" className="input" onChange={updateName} name="input" type="text"/>
                <button className="btn" onClick={greet}>Greet</button>
            </div>
        </div>
    )
}

export default App

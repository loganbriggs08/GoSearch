import './themes/default-theme.css';
import SearchComponent from './Search'
import SettingsComponent from './Settings'
import {GetCurrentPage} from "../wailsjs/go/main/App"

function App() {
  return (
    <SettingsComponent/>
  );
}

export default App;
import React, { useEffect, useState } from 'react';
import './themes/default-theme.css';
import SearchComponent from './Search';
import SettingsComponent from './Settings';
import { GetCurrentPage } from "../wailsjs/go/main/App";

function App() {
  const [currentPage, setCurrentPage] = useState('');

  useEffect(() => {
    GetCurrentPage().then((currentpage: string) => {
      setCurrentPage(currentpage);
    });
  }, []);

  return (
    <div>
      {currentPage === "Search" && <SearchComponent />}
      {currentPage === "Settings" && <SettingsComponent />}
    </div>
    );
}

export default App;
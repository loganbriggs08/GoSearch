import { GetCurrentPage } from "../wailsjs/go/main/App";
import React, { useEffect, useState } from 'react';
import SearchComponent from './Search';
import SettingsComponent from './Settings';
import WelcomeComponent from "./Welcome";
import SetupComponent from "./Setup";
import "./default-theme.css";

function App() {
  const [currentPage, setCurrentPage] = useState('');

  useEffect(() => {
    GetCurrentPage().then((currentpage: string) => {
      setCurrentPage(currentpage);
    });
  }, []);

  return (
    <div>
      {currentPage === "Welcome" && <WelcomeComponent />}
      {currentPage === "Setup" && <SetupComponent />}
      {currentPage === "Search" && <SearchComponent />}
      {currentPage === "Settings" && <SettingsComponent />}
    </div>
    );
}

export default App;
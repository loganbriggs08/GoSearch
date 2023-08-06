import React, { useEffect, useState } from 'react';
import { GetCurrentPage } from "../wailsjs/go/main/App";
import { useTheme } from './ThemeContext';
import SearchComponent from './Search';
import SettingsComponent from './Settings';
import WelcomeComponent from "./Welcome";
import SetupComponent from "./Setup";

const App: React.FC = () => {
  const {theme} = useTheme();
  const [currentPage, setCurrentPage] = useState('');

  useEffect(() => {
    GetCurrentPage().then((currentpage: string) => {
      setCurrentPage(currentpage);
    });
  }, []);

  return (
    <div>
      <link rel="stylesheet" href={`/src/themes/${theme}.css`} />

      {currentPage === "Welcome" && <WelcomeComponent />}
      {currentPage === "Setup" && <SetupComponent />}
      {currentPage === "Search" && <SearchComponent />}
      {currentPage === "Settings" && <SettingsComponent />}
    </div>
    );
}

export default App;
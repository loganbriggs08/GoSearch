import React, { createContext, useState, useContext, ReactNode, useEffect } from 'react';
import {CurrentTheme} from "../wailsjs/go/main/App";

type Theme = 'blue-theme' | 'cherry-blossom-theme' | 'lavender-theme' | 'mint-theme' | 'default-theme';

interface ThemeContextProps {
    theme: Theme;
    setTheme: React.Dispatch<React.SetStateAction<Theme>>;
}

const ThemeContext = createContext<ThemeContextProps | undefined>(undefined);

const ThemeProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
    const [theme, setTheme] = useState<Theme>('default-theme');
    const [ourSetTheme, setOurSetTheme] = useState('')
    
    useEffect(() => {
        const setThemeBasedOnResponse = (response: any) => {
            setOurSetTheme(response);

            if (response === "Default Theme") {
                setTheme("default-theme");
            } else if (response === "Blue Theme") {
                setTheme("blue-theme");
            } else if (response === "Cherry Blossom Theme") {
                setTheme("cherry-blossom-theme");
            } else if (response === "Lavender Theme") {
                setTheme("lavender-theme");
            } else if (response === "Mint Theme") {
                setTheme("mint-theme");
            }
        };

        CurrentTheme().then((response: any) => {
            setThemeBasedOnResponse(response);
        });

        }, []);
    
    return (
        <ThemeContext.Provider value={{ theme, setTheme }}>
            {children}
        </ThemeContext.Provider>
        );
};

export const useTheme = () => {
    const context = useContext(ThemeContext);
    if (!context) {
        throw new Error('useTheme must be used within a ThemeProvider');
    }
    return context;
};

export default ThemeProvider;
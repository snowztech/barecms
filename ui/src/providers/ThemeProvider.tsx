import React, { useEffect, useState } from "react";
import { ThemeContext } from "@/contexts/ThemeContext";
import {
  Theme,
  DEFAULT_THEME,
  applyTheme,
  getStoredTheme,
  setStoredTheme,
  getNextTheme,
} from "@/lib/theme";

interface ThemeProviderProps {
  children: React.ReactNode;
}

export const ThemeProvider: React.FC<ThemeProviderProps> = ({ children }) => {
  const [theme, setTheme] = useState<Theme>(DEFAULT_THEME);

  useEffect(() => {
    // Get saved theme or default to light
    const savedTheme = getStoredTheme();
    const initialTheme = savedTheme || DEFAULT_THEME;

    setTheme(initialTheme);
    applyTheme(initialTheme);
  }, []);

  const toggleTheme = () => {
    const newTheme = getNextTheme(theme);
    setTheme(newTheme);
    applyTheme(newTheme);
    setStoredTheme(newTheme);
  };

  const value = {
    theme,
    toggleTheme,
  };

  return (
    <ThemeContext.Provider value={value}>{children}</ThemeContext.Provider>
  );
};

export default ThemeProvider;

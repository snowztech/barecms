import React from "react";
import { Moon, Sun } from "lucide-react";
import { useTheme } from "@/hooks/useTheme";

const ThemeToggle: React.FC = () => {
  const { theme, toggleTheme } = useTheme();

  return (
    <button
      onClick={toggleTheme}
      className="theme-toggle-btn p-2 w-10 h-10 rounded-bare border border-bare-200 hover:border-bare-300 hover:bg-bare-50 active:bg-bare-100 transition-all duration-200 flex items-center justify-center"
      aria-label="Toggle theme"
    >
      {theme === "barecms" ? (
        <Moon
          size={16}
          className="text-bare-600 transition-colors duration-200"
        />
      ) : (
        <Sun
          size={16}
          className="text-bare-300 transition-colors duration-200"
        />
      )}
    </button>
  );
};

export default ThemeToggle;

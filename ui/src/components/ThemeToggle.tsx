import React from "react";
import { Moon, Sun } from "lucide-react";
import { useTheme } from "@/providers/ThemeProvider";

const ThemeToggle: React.FC = () => {
  const { theme, toggleTheme } = useTheme();

  return (
    <button
      onClick={toggleTheme}
      className="p-2 w-10 h-10 rounded-bare border border-bare-300 hover:border-bare-400 hover:bg-bare-50 transition-all duration-200 flex items-center justify-center"
      aria-label="Toggle theme"
    >
      {theme === "barecms" ? (
        <Moon size={16} className="text-bare-600" />
      ) : (
        <Sun size={16} className="text-bare-300" />
      )}
    </button>
  );
};

export default ThemeToggle;

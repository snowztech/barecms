import React, { useState, useEffect } from "react";
import { Moon, Sun } from "lucide-react";

const ThemeToggle: React.FC = () => {
  const [theme, setTheme] = useState<"barecms" | "barecms-dark">("barecms");

  useEffect(() => {
    // Check for saved theme preference or default to light mode
    const savedTheme = localStorage.getItem("theme") as
      | "barecms"
      | "barecms-dark"
      | null;
    if (savedTheme) {
      setTheme(savedTheme);
      document.documentElement.setAttribute("data-theme", savedTheme);
    }
  }, []);

  const toggleTheme = () => {
    const newTheme = theme === "barecms" ? "barecms-dark" : "barecms";
    setTheme(newTheme);
    document.documentElement.setAttribute("data-theme", newTheme);
    localStorage.setItem("theme", newTheme);
  };

  return (
    <button
      onClick={toggleTheme}
      className="btn btn-bare btn-sm p-2 w-10 h-10"
      aria-label="Toggle theme"
    >
      {theme === "barecms" ? <Moon size={16} /> : <Sun size={16} />}
    </button>
  );
};

export default ThemeToggle;

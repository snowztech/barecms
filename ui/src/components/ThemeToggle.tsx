import React, { useState, useEffect } from "react";
import { Moon, Sun } from "lucide-react";

const ThemeToggle: React.FC = () => {
  const [theme, setTheme] = useState<"barecms" | "barecms-dark">("barecms");
  const [mounted, setMounted] = useState(false);

  useEffect(() => {
    setMounted(true);
    // Check for saved theme preference or default to light mode
    const savedTheme = localStorage.getItem("theme") as
      | "barecms"
      | "barecms-dark"
      | null;

    const initialTheme = savedTheme || "barecms";
    setTheme(initialTheme);

    // Set theme on both html and body elements to ensure compatibility
    document.documentElement.setAttribute("data-theme", initialTheme);
    document.body.setAttribute("data-theme", initialTheme);
  }, []);

  const toggleTheme = () => {
    const newTheme = theme === "barecms" ? "barecms-dark" : "barecms";
    setTheme(newTheme);

    // Update both elements
    document.documentElement.setAttribute("data-theme", newTheme);
    document.body.setAttribute("data-theme", newTheme);

    localStorage.setItem("theme", newTheme);
  };

  // Prevent hydration mismatch by not rendering until mounted
  if (!mounted) {
    return (
      <div className="w-10 h-10 p-2 rounded-bare border border-bare-300 flex items-center justify-center">
        <div className="w-4 h-4"></div>
      </div>
    );
  }

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

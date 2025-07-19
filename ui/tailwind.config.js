/** @type {import('tailwindcss').Config} */
import daisyui from "daisyui";

export default {
  content: ["./src/**/*.{html,js,ts,tsx}"],
  plugins: [daisyui],
  theme: {
    extend: {
      fontFamily: {
        'sans': ['Inter', 'system-ui', 'sans-serif'],
        'display': ['Space Grotesk', 'system-ui', 'sans-serif'],
        'mono': ['IBM Plex Mono', 'Menlo', 'Monaco', 'monospace'],
      },
      colors: {
        // Custom semantic colors for BareCMS
        'bare': {
          50: '#fafafa',
          100: '#f5f5f5',
          200: '#e5e5e5',
          300: '#d4d4d4',
          400: '#a3a3a3',
          500: '#737373',
          600: '#525252',
          700: '#404040',
          800: '#262626',
          900: '#171717',
          950: '#0a0a0a',
        },
        'accent': {
          50: '#f0f9ff',
          100: '#e0f2fe',
          200: '#bae6fd',
          300: '#7dd3fc',
          400: '#38bdf8',
          500: '#0ea5e9',
          600: '#0284c7',
          700: '#0369a1',
          800: '#075985',
          900: '#0c4a6e',
        }
      },
      spacing: {
        '18': '4.5rem',
        '88': '22rem',
      },
      borderRadius: {
        'bare': '0.5rem',
      },
      boxShadow: {
        'bare': '0 1px 3px 0 rgba(0, 0, 0, 0.1), 0 1px 2px 0 rgba(0, 0, 0, 0.06)',
        'bare-lg': '0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06)',
      }
    },
  },
  daisyui: {
    themes: [
      {
        barecms: {
          "primary": "#0ea5e9",
          "primary-focus": "#0284c7",
          "primary-content": "#ffffff",

          "secondary": "#737373",
          "secondary-focus": "#525252",
          "secondary-content": "#ffffff",

          "accent": "#0ea5e9",
          "accent-focus": "#0284c7",
          "accent-content": "#ffffff",

          "neutral": "#404040",
          "neutral-focus": "#262626",
          "neutral-content": "#ffffff",

          "base-100": "#ffffff",
          "base-200": "#fafafa",
          "base-300": "#f5f5f5",
          "base-content": "#171717",

          "info": "#0ea5e9",
          "info-content": "#ffffff",

          "success": "#22c55e",
          "success-content": "#ffffff",

          "warning": "#f59e0b",
          "warning-content": "#ffffff",

          "error": "#ef4444",
          "error-content": "#ffffff",

          "--rounded-box": "0.5rem",
          "--rounded-btn": "0.375rem",
          "--rounded-badge": "1.9rem",
          "--animation-btn": "0.2s",
          "--animation-input": "0.2s",
          "--btn-text-case": "none",
          "--btn-focus-scale": "0.98",
          "--border-btn": "1px",
          "--tab-border": "1px",
          "--tab-radius": "0.375rem",
        },
        "barecms-dark": {
          "primary": "#38bdf8",
          "primary-focus": "#0ea5e9",
          "primary-content": "#0a0a0a",

          "secondary": "#a3a3a3",
          "secondary-focus": "#737373",
          "secondary-content": "#0a0a0a",

          "accent": "#38bdf8",
          "accent-focus": "#0ea5e9",
          "accent-content": "#0a0a0a",

          "neutral": "#d4d4d4",
          "neutral-focus": "#a3a3a3",
          "neutral-content": "#0a0a0a",

          "base-100": "#0a0a0a",
          "base-200": "#171717",
          "base-300": "#262626",
          "base-content": "#fafafa",

          "info": "#38bdf8",
          "info-content": "#0a0a0a",

          "success": "#4ade80",
          "success-content": "#0a0a0a",

          "warning": "#fbbf24",
          "warning-content": "#0a0a0a",

          "error": "#f87171",
          "error-content": "#0a0a0a",

          "--rounded-box": "0.5rem",
          "--rounded-btn": "0.375rem",
          "--rounded-badge": "1.9rem",
          "--animation-btn": "0.2s",
          "--animation-input": "0.2s",
          "--btn-text-case": "none",
          "--btn-focus-scale": "0.98",
          "--border-btn": "1px",
          "--tab-border": "1px",
          "--tab-radius": "0.375rem",
        }
      },
      "light", "dark"
    ],
  },
};
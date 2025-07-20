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
          50: '#f8faf9',
          100: '#f1f5f2',
          200: '#e2ebe4',
          300: '#c9d6cc',
          400: '#a8b8ab',
          500: '#859289',
          600: '#6b7770',
          700: '#56615b',
          800: '#474f4a',
          900: '#3c423e',
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
        'bare-dark': '0 1px 3px 0 rgba(0, 0, 0, 0.4), 0 1px 2px 0 rgba(0, 0, 0, 0.3)',
        'bare-lg-dark': '0 4px 6px -1px rgba(0, 0, 0, 0.4), 0 2px 4px -1px rgba(0, 0, 0, 0.3)',
      }
    },
  },
  daisyui: {
    themes: [
      {
        barecms: {
          "primary": "#6b7770",
          "primary-focus": "#56615b",
          "primary-content": "#ffffff",

          "secondary": "#737373",
          "secondary-focus": "#525252",
          "secondary-content": "#ffffff",

          "accent": "#859289",
          "accent-focus": "#6b7770",
          "accent-content": "#ffffff",

          "neutral": "#404040",
          "neutral-focus": "#262626",
          "neutral-content": "#ffffff",

          "base-100": "#ffffff",
          "base-200": "#fafafa",
          "base-300": "#f5f5f5",
          "base-content": "#171717",

          "info": "#6b7770",
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
          "primary": "#a8b8ab",
          "primary-focus": "#c9d6cc",
          "primary-content": "#0a0a0a",

          "secondary": "#d4d4d4",
          "secondary-focus": "#e5e5e5",
          "secondary-content": "#0a0a0a",

          "accent": "#c9d6cc",
          "accent-focus": "#e2ebe4",
          "accent-content": "#0a0a0a",

          "neutral": "#e5e5e5",
          "neutral-focus": "#f5f5f5",
          "neutral-content": "#0a0a0a",

          "base-100": "#0a0a0a",
          "base-200": "#171717",
          "base-300": "#262626",
          "base-content": "#f5f5f5",

          "info": "#a8b8ab",
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
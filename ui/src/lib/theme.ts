export type Theme = "barecms" | "barecms-dark";

export const THEME_STORAGE_KEY = "theme";

export const DEFAULT_THEME: Theme = "barecms";

export const THEME_ATTRIBUTE = "data-theme";

export const applyTheme = (theme: Theme): void => {
  document.documentElement.setAttribute(THEME_ATTRIBUTE, theme);
  document.body.setAttribute(THEME_ATTRIBUTE, theme);
};

export const getStoredTheme = (): Theme | null => {
  if (typeof window === 'undefined') return null;
  return localStorage.getItem(THEME_STORAGE_KEY) as Theme | null;
};

export const setStoredTheme = (theme: Theme): void => {
  if (typeof window === 'undefined') return;
  localStorage.setItem(THEME_STORAGE_KEY, theme);
};

export const getNextTheme = (currentTheme: Theme): Theme => {
  return currentTheme === "barecms" ? "barecms-dark" : "barecms";
};
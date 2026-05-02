import { createContext } from "react";

export interface ThemeContextType {
  theme: string;
}

const ThemeLight = "light";
const ThemeDark = "dark";

const ThemeContext = createContext<ThemeContextType>({
  theme: ThemeDark,
});

export default ThemeContext;

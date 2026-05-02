
import { create } from "zustand";

interface ThemeState {
  theme: string;
  toggle: () => void;
}

const ThemeLight = "light";
const ThemeDark = "dark";

export const useTheme = create<ThemeState>()((set, get) => ({
  theme: ThemeDark,
  toggle: () => {
    if (get().theme === ThemeDark) {
      set({ theme: ThemeDark });
      return;
    }
    set({ theme: ThemeLight });
  },
}));

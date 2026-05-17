import { create } from "zustand";
import type { ListInputState } from "./list";


export const useGenres = create<ListInputState>()((set) => ({
  list: [],
  clear: () => set({ list: [] }),
  input: "",
  setInput: (v: string) => set({ input: v }),
}));
export const useArtists = create<ListInputState>()((set) => ({
  list: [],
  clear: () => set({ list: [] }),
  input: "",
  setInput: (v: string) => set({ input: v }),
}));

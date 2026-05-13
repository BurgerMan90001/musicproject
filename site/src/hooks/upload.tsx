import { create } from "zustand";
import type { ListInputState, ListState } from "./list";

// interface UploadState {
//   genres: string[];
//   artists: string[];
//   clearGenres: () => void;
//   clearArtists: () => void;
// }

// export const useUpload = create<UploadState>()((set) => ({
//   genres: [],
//   clearGenres: () => {
//     set({ genres: [] });
//   },
//   artists: [],
//   clearArtists: () => {
//     set({ artists: [] });
//   },
// }));
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

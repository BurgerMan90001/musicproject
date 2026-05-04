import type { Song } from "../types/song.types";
import { create } from "zustand";

interface PlayerState {
  volume: string;
  paused: boolean;
  collapsed: boolean;
  queue: Song[];
  empty(): boolean;

  change: (by: string) => void;
  mute: () => void;

  togglePause: () => void;
  toggleCollapsed: () => void;
}
export const usePlayerStore = create<PlayerState>()((set, get) => ({
  volume: "0",
  paused: true,
  collapsed: false,
  queue: [],
  empty: (): boolean => {
    return get().queue[0] === undefined;
  },
  change: (by) => set({ volume: by }),
  mute: () => set({ volume: "0" }),
  togglePause: () =>
    set((state) => ({
      paused: !state.paused,
    })),
  toggleCollapsed: () =>
    set((state) => ({
      collapsed: !state.collapsed,
    })),
}));
// export const initialVolume: string = "0";
// export const initialQueue: Song[] = [];
// export { PlayerContext, type PlayerContextType, usePlayer };

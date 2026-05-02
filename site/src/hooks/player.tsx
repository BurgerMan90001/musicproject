
import type { Song } from "../types/song.types";
import { create } from "zustand";

interface VolumeState {
  volume: string;
  change: (by: string) => void;
  mute: () => void;
}

export const useVolume = create<VolumeState>()((set) => ({
  volume: "0",
  change: (by) => set({ volume: by }),
  mute: () => set({ volume: "0" }),
}));

interface SongQueueState {
  queue: Song[];
  empty(): boolean;
}

export const useSongQueue = create<SongQueueState>()((_, get) => ({
  queue: [],
  empty: (): boolean => {
    return get().queue[0] === undefined;
  },
}));

interface PlayState {
  paused: boolean;
  toggle: () => void;
}

export const usePlayStore = create<PlayState>()((set) => ({
  paused: true,
  toggle: () =>
    set((state) => ({
      paused: !state.paused,
    })),
}));

// export const initialVolume: string = "0";
// export const initialQueue: Song[] = [];
// export { PlayerContext, type PlayerContextType, usePlayer };

import type { Song } from "../types/song.types";
import { create } from "zustand";

interface PlayerState {
  // volume: string;
  // paused: boolean;
  collapsed: boolean;
  audio: HTMLAudioElement;
  queue: Song[];
  empty(): boolean;

  changeVolume: (by: number) => void;

  togglePause: () => void;
  toggleCollapsed: () => void;
}
export const usePlayerStore = create<PlayerState>()((set, get) => ({
  collapsed: false,
  audio: new Audio("https://storage.songsled.com/song18.mp3"),
  queue: [
    {
      id: "1",
      albumId: "1",
      name: "Test",
      genres: "Pop, Rock",
      artists: "Pop Smoke",
      streams: 123,
      duration: 123,
      creationDate: "2007-03-24",
      audio: "https://storage.songsled.com/audio/89e9eb5b",
    },
  ],
  empty: (): boolean => {
    return get().queue[0] === undefined;
  },
  changeVolume: (by) => {
    get().audio.volume = by;
  },
  // toggleMute: () => set({})
  togglePause: () => {
    if (get().audio.paused) {
      get().audio.play();

      // set({ paused: false });
      return;
    }
    get().audio.pause();
  },
  toggleCollapsed: () =>
    set((state) => ({
      collapsed: !state.collapsed,
    })),
}));
// export const initialVolume: string = "0";
// export const initialQueue: Song[] = [];
// export { PlayerContext, type PlayerContextType, usePlayer };

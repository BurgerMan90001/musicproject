import { createRef, type RefObject } from "react";
import type { Song } from "../types/song.types";
import { create } from "zustand";

interface PlayerState {
  audio: RefObject<HTMLAudioElement | null>;
  progressBar: RefObject<HTMLInputElement | null>;
  progress: number;
  duration: number;
  setDuration: (d: number) => void;
  setSong: (s: Song) => void;
  // queue: Song[];
  song?: Song;
}
export const usePlayerStore = create<PlayerState>()((set) => ({
  audio: createRef<HTMLAudioElement>(),
  progressBar: createRef<HTMLInputElement>(),
  progress: 0,
  duration: 0,
  setDuration: (n) => {
    set({ duration: n });
  },
  setSong: (s) => {
    set({ song: s });
  },

  // setAudio: (a: string) => {

  //   get().audio.current?.src = a;
  // },
  song: undefined,
}));

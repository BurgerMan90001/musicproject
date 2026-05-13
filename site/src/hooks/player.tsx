import { createRef, type RefObject } from "react";
import type { Song } from "../types/song.types";
import { create } from "zustand";

interface PlayerState {
  audio: RefObject<HTMLAudioElement | null>;
  progressBar: RefObject<HTMLInputElement | null>;
  progress: number;
  duration: number;
  setDuration: (d: number) => void;
  queue: Song[];
}
export const usePlayerStore = create<PlayerState>()((set) => ({
  audio: createRef<HTMLAudioElement>(),
  progressBar: createRef<HTMLInputElement>(),
  progress: 0,
  duration: 0,
  setDuration: (n) => {
    set({ duration: n });
  },
  queue: [
    {
      id: "1",
      albumId: "1",
      name: "8bitasdasdasdasd",
      genres: "Pop, Rock",
      artists: "Bossa",
      streams: 123,
      duration: 123,
      creationDate: "2007-03-24",
      audio: "https://storage.songsled.com/audio/8bit%20Bossa.mp3",
    },
  ],
}));

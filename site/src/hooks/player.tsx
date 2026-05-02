import { createContext, useContext } from "react";
import type { Song } from "../types/song.types";

export interface PlayerContextType {
  song?: Song;
  paused: boolean;
  volume: number;
  getVolume(): number;
  isEmpty(): boolean;
}
export const PlayerContext = createContext<PlayerContextType>({
  volume: 0,
  paused: false,
  getVolume() {
    if (this.volume > 100) {
      return 100;
    }
    if (this.volume < 0) {
      return 0;
    }
    return this.volume;
  },
  isEmpty() {
    return this.song == null;
  },
});

export const usePlayer = (): PlayerContextType => {
  const context = useContext(PlayerContext);

  return context;
};

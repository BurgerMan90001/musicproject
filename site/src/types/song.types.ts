export interface Song {
  id: string;
  albumId?: string;
  name: string;
  genres: string[];
  artists: string[];
  streams: number;
  duration: number;
  creationDate: string;
  cover?: string;
  audio: string;
}

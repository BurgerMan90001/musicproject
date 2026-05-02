import { useContext, createContext, useCallback } from "react";
import type { Song } from "../../types/song.types";
import EmptySong from "./Placeholder";
import PlaceholderSongIcon from "./Placeholder";

interface SongListContextType {
  songs: Song[];
}
const SongListContext = createContext<SongListContextType>({
  songs: [
    {
      name: "asdas",
      genre: "",
      streams: 0,
      duration: 0,

      url: "#",
    },
  ],
});

const SongList = () => {
  const { songs }: { songs: Song[] } = useContext(SongListContext);
  const test = useCallback(() => {
    for (var i = 0; i < 10; i++) {
      songs.push({
        name: "asdas",
        genre: "",
        streams: 0,
        duration: 0,

        url: "#",
      });
    }
  }, []);
  test()

  return (
    <ul className="">
      {songs.map((song) => {
        var songImage;
        if (song.image) {
          songImage = <img src={song.image} />;
        } else {
          songImage = (
            <div className="image image-64 bg-color-body-darker color-text-invert">
              <PlaceholderSongIcon />
            </div>
          );
        }
        return (
          <li className="display-flex padding-xxs gap-xxs">
            {songImage}
            <summary>{song.name}</summary>
          </li>
        );
      })}
    </ul>
  );
};
const SongQueue = () => {
  return (
    <>
      <nav className="bg-color-body-dark ">
        <SongList />
      </nav>
    </>
  );
};

export default SongQueue;

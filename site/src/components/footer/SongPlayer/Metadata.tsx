import { usePlayerStore } from "../../../hooks/player";
import { useEffect, useState } from "react";

const Metadata = () => {
  const player = usePlayerStore();

  const [name, setName] = useState("No songs playing");
  const [artists, setArtists] = useState<string[]>();

  useEffect(() => {
    if (!player.song) {
      return;
    }
    if (player.song) {
      setName(player.song.name);
    } else {
      setName("No songs playing");
    }
    if (player.song.artists) {
      setArtists(player.song.artists);
    }
    //  else {
    //   setArtists([""]);
    // }
  }, [player.song]);

  return (
    <div className="padding-xxs display-flex flex-column">
      <span>{name}</span>
      <span>{artists}</span>
    </div>
  );
};

export default Metadata;

import { usePlayerStore } from "../../../hooks/player";
import { SongPlaceholderIcon } from "../../svg/Placeholder";

const SongList = () => {
  const queue = usePlayerStore((state) => state.queue);

  return (
    <ul className="">
      {queue.map((song) => {
        var songImage;
        if (song.image) {
          songImage = <img src={song.image} />;
        } else {
          songImage = (
            <div className="image image-64 bg-color-body-darker color-text-invert">
              <SongPlaceholderIcon />
            </div>
          );
        }
        return (
          <li className="display-flex padding-xxs gap-xxs">
            {songImage}
            {song.name && <summary>{song.name}</summary>}
          </li>
        );
      })}
    </ul>
  );
};
const SongQueue = () => {
  return (
    <nav className="bg-color-body-dark">
      <SongList />
    </nav>
  );
};

export default SongQueue;

import { usePlayerStore } from "../../../hooks/player";
import { SongPlaceholderSvg } from "../Svg";

const SongList = () => {
  const queue = usePlayerStore((state) => state.queue);

  if (queue.length === 0) {
    return <div className="bg-color-body-dark">asddasd</div>;
  }
  return (
    <ul className="">
      {queue.map((song) => {
        var songImage;
        if (song.image) {
          songImage = <img src={song.image} />;
        } else {
          songImage = <SongPlaceholderSvg />;
        }
        return (
          <li className="display-flex padding-xxs gap-xxs">
            <div className="image image-64 bg-color-body-darker color-text-invert">
              {songImage}
            </div>

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

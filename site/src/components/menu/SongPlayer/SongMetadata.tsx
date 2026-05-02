// import { usePlayer } from "../../../hooks/player";

import { useSongQueue } from "../../../hooks/player";

const SongMetadata = () => {
  const queue = useSongQueue();

  if (!queue.empty()) {
    return (
      <div className="padding-xxs">
        <summary>{queue.queue[0].name}</summary>
        <span>By: {queue.queue[0].artist}</span>
      </div>
    );
  }
  return <div className="padding-xxs"></div>;
};

export default SongMetadata;

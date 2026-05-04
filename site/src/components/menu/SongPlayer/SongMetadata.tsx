import { usePlayerStore } from "../../../hooks/player";

const SongMetadata = () => {
  const queue = usePlayerStore();

  if (!queue.empty()) {
    return (
      <div className="padding-xxs">
        <summary>{queue.queue[0].name}</summary>
        <span>By: {queue.queue[0].artist}</span>
      </div>
    );
  }
  return (
    <div className="padding-xxs">
      <summary>No songs playing</summary>
      {/* <span>By: {queue.queue[0].artist}</span> */}
    </div>
  );
};

export default SongMetadata;

import { usePlayerStore } from "../../../hooks/player";

const Metadata = () => {
  const queue = usePlayerStore();

  var title = "No songs playing";
  var artists = "";
  if (!queue.empty()) {
    title = queue.queue[0].name;
    artists = queue.queue[0].artists;
  }
  return (
    <div className="padding-xxs">
      <summary>{title}</summary>
      <span className="color-text-subtle">{artists}</span>
    </div>
  );
};

export default Metadata;

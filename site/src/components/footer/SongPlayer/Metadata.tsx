import { Link } from "react-router";
import { usePlayerStore } from "../../../hooks/player";

const Metadata = () => {
  const player = usePlayerStore();

  var title = <span>No songs playing</span>;
  var artists;
  if (!player.queue || !player.queue[0]) {
    return;
  }
  if (player.queue[0].name) {
    title = (
      <summary>
        <Link to="#">{player.queue[0].name}</Link>
      </summary>
    );
  }
  if (player.queue[0].artists) {
    artists = (
      <span className="color-text-subtle">
        <Link to="#">{player.queue[0].artists}</Link>
      </span>
    );
  }
  return (
    <div className="padding-xxs display-flex flex-column">
      {title}
      {artists}
    </div>
  );
};

export default Metadata;

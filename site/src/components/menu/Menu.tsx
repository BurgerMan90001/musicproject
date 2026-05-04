import SongPlayer from "./SongPlayer/SongPlayer";
import SongQueue from "./SongQueue/SongQueue";

const Menu = () => {
  return (
    <aside className="height-full scroll-vertical border-radius-lg">
      <SongPlayer />
      <SongQueue />
    </aside>
  );
};

export default Menu;

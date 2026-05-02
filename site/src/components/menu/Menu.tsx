import SongPlayer from "./SongPlayer";
import SongQueue from "./SongQueue";

const Menu = () => {
  return (
    <aside className="layout-menu height-full scroll-vertical border-radius-lg">
      <SongPlayer />
      <SongQueue />
      {/* <div id="test" className="position-absolute bg-color-body-dark">asdasda</div> */}
      {/* <button className="button-clear gap-xxs font-size-lg display-flex align-items-center font-weight-bold padding-xs">
                <img className="icon" src={test} />
                <span>Test</span>
              </button> */}
    </aside>
  );
};

export default Menu;

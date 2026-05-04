import { PlayButton, VolumeButton } from "../Button";
import ProgressBar from "./ProgressBar";
import SongImage from "./SongImage";
import SongMetadata from "./SongMetadata";
import VolumeControl from "./VolumeControl";
import { usePlayerStore } from "../../../hooks/player";

const SongPlayer = () => {
  const player = usePlayerStore();

  var image;
  var content;
  if (player.collapsed) {
    image = (
      <div className="image image-64 bg-color-body-darker color-text-invert">
        <SongImage />
      </div>
    );
    content = (
      <div className="display-flex align-items-center bg-color-body-darker grid-template-rows-1fr-auto">
        <span></span>
        <nav className="display-flex">
          {/* <ProgressBar /> */}

          <PlayButton />
          <VolumeButton />

          <VolumeControl />
          <ProgressBar />
        </nav>
        {/* <div className="margin-auto">
          <ProgressBar />
        </div> */}
      </div>
    );
  } else {
    image = (
      <div className="image image-128 bg-color-body-darker color-text-invert">
        <SongImage />
      </div>
    );
    content = (
      <div className="display-flex flex-column bg-color-body-darker ">
        <SongMetadata />
        <ProgressBar />
        <nav className="display-flex">
          <PlayButton />
          <VolumeButton />

          <VolumeControl />
        </nav>
      </div>
    );
  }
  return (
    <div className="display-flex">
      <button
        onClick={() => player.toggleCollapsed()}
        popoverTarget="my-popover"
        className="button-clear"
      >
        {image}
      </button>

      {content}
    </div>
  );
};

export default SongPlayer;

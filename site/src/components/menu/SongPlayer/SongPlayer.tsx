import { PlayButton } from "./Button";

import SongMetadata from "./Metadata";
import VolumeControl from "./VolumeControl";
import { usePlayerStore } from "../../../hooks/player";
import { SongPlaceholderSvg } from "../Svg";

const SongImage = () => {
  const queue = usePlayerStore((state) => state.queue);

  if (queue[0] && queue[0].image) {
    return <img src={queue[0].image} />;
  }
  return <SongPlaceholderSvg />;
};
const ProgressBar = () => {
  return (
    <div className="gap-xxs width-300">
      <span>0:00 / 0:00</span>
      <div>
        <input
          id="progressBar"
          aria-label="Progress Bar"
          className="slider bg-color-body"
          type="range"
        />
      </div>
    </div>
  );
};

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
      <div className="display-flex align-items-center grid-template-rows-1fr-auto">
        <nav className="display-flex bg-color-body-dark">
          <PlayButton />

          <VolumeControl />
        </nav>
      </div>
    );
  } else {
    image = (
      <div className="image image-128 bg-color-body-darker color-text-invert">
        <SongImage />
      </div>
    );
    content = (
      <div className="display-grid">
        <SongMetadata />
        <nav className="display-flex">
          <PlayButton />
          <VolumeControl />
        </nav>
      </div>
    );
  }
  return (
    <div className="display-flex">
      <button
        onClick={() => player.toggleCollapsed()}
        // popoverTarget="my-"
        className="button-clear popover popover-top"
      >
        {image}
        {/* <div className="popover-content">asd</div> */}
      </button>
      {content}

      <ProgressBar />
    </div>
  );
};

export default SongPlayer;

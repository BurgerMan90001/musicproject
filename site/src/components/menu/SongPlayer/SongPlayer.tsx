import { Button } from "../Button";
import { PlayIcon, VolumeIcon } from "../Icons";
import SongImage from "./SongImage";
import SongMetadata from "./SongMetadata";
import VolumeControl from "./VolumeControl";

const SongPlayer = () => {
  return (
    <div className="display-flex">
      <button popoverTarget="my-popover" className="button-clear">
        <div className="image image-128 bg-color-body-darker color-text-invert">
          <SongImage />
        </div>
      </button>

      <div className="display-grid bg-color-body-darker grid-template-rows-1fr-auto">
        <SongMetadata />
        <nav className="display-flex">
          <Button>
            <PlayIcon />
          </Button>
          <Button>
            <VolumeIcon />
          </Button>
          <VolumeControl />
        </nav>
        {/* <ProgressBar /> */}
      </div>
    </div>
  );
};

export default SongPlayer;

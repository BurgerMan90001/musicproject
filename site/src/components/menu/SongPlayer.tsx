import { useContext } from "react";

import { PlayerContext } from "../../hooks/player";
import { Button, PlayIcon, VolumeIcon } from "./Button";

import PlaceholderSongIcon from "./Placeholder";

const SongPlayer = () => {
  const player = useContext(PlayerContext);

  var songContent;
  if (player.isEmpty()) {
    songContent = (
      <div className="image image-128 bg-color-body-darker color-text-invert">
        <PlaceholderSongIcon />
      </div>
    );
  }
  return (
    <div className="bg-color-body-dark">
      <div className="display-flex">
        <button className="button-clear">{songContent}</button>

        <div className="display-grid grid-template-rows-1fr-auto">
          <div className="padding-xxs">
            <summary>Testing Testing Testing Testing</summary>
            <span>By: Testing Testing</span>
          </div>

          <nav className="display-flex align-items-center bg-color-body-darker">
            <Button>
              <PlayIcon />
            </Button>
            <Button>
              <VolumeIcon />
            </Button>

            <input className="slider" type="range" />
          </nav>
        </div>
      </div>
      <div className="display-flex flex-column align-items-center padding-xxs gap-xxs bg-color-body-darker ">
        <span>0:00 / 0:00</span>

        <input className="slider" type="range" />
      </div>
    </div>
  );
};

export default SongPlayer;

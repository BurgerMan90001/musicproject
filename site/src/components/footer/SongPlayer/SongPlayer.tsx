import { PlayButton } from "./Button";

import SongMetadata from "./Metadata";
import VolumeControl from "./VolumeControl";
import { usePlayerStore } from "../../../hooks/player";

import { SongPlaceholderSvg } from "../Svg";
import { useEffect } from "react";

// const SongImage = () => {
//   const queue = usePlayerStore((state) => state.queue);

//   if (queue[0] && queue[0].image) {
//     return <img src={queue[0].image} />;
//   }

//   return <SongPlaceholderSvg />;
// };
const formatTime = (time: number): string => {
  const minutes = Math.floor(time / 60)
    .toString()
    .padStart(2, "0");
  const seconds = Math.floor(time % 60)
    .toString()
    .padStart(2, "0");

  return `${minutes}:${seconds}`;
};
const ProgressBar = () => {
  const player = usePlayerStore();

  useEffect(() => {
    
  }, [player.audio])
  const onChange = () => {
    if (player.audio.current && player.progressBar.current) {
      player.audio.current.currentTime = Number(
        player.progressBar.current.value,
      );
    }
  };
  // uesEffect(() => {
  //   console.log("ASDASDASDasda");
  //   // player.audio.currentTime = value;
  // }, [player.audio.currentTime]);
  // const test = player.audio.onchange;
  // player.audio.showPopover
  // const test = player?.audio?.onplaying((e) => {});
  // player.progress.current?.value
  return (
    <div className="gap-xxs width-300">
      <span>
        {formatTime(player.progress)} / {formatTime(player.duration)}
      </span>
      <div>
        <input
          id="progressBar"
          aria-label="Progress Bar"
          className="slider bg-color-body"
          ref={player.progressBar}
          type="range"
          defaultValue={0}
          onChange={onChange}
        />
      </div>
    </div>
  );
};

const SongPlayer = () => {
  const player = usePlayerStore();

  const onLoadedMetadata = () => {
    const seconds = player.audio.current?.duration;
    if (seconds) {
      player.setDuration(seconds);
    }
    if (player.progressBar.current && seconds) {
      player.progressBar.current.max = seconds.toString();
    }
  };

  return (
    <div className="display-flex position-relative justifiy-content-center ">
      {player.queue[0] && player.queue[0].audio && (
        <audio
          onLoadedMetadata={onLoadedMetadata}
          src={player.queue[0].audio}
          ref={player.audio}
        ></audio>
      )}
      <div className="position-absolute left-0 padding-inline-xs padding-block-xxs">
        <ProgressBar />
      </div>

      <div className="display-flex">
        <SongMetadata />
        <div className="image image-64 bg-color-body-darker color-text-invert">
          {player.queue[0] && player.queue[0].image ? (
            <img src={player.queue[0].image} />
          ) : (
            <SongPlaceholderSvg />
          )}
        </div>

        <div className="display-flex align-items-center grid-template-rows-1fr-auto">
          <nav className="display-flex bg-color-body-dark">
            <PlayButton />
            <VolumeControl />
          </nav>
        </div>
      </div>
    </div>
  );
};

export default SongPlayer;

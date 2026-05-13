import { useEffect, useState } from "react";
import { usePlayerStore } from "../../../hooks/player";
import { MutedSvg, VolumeMediumSvg } from "../Svg";

const VolumeControl = () => {
  const player = usePlayerStore();
  const [volume, setVolume] = useState(0);
  const [muted, setMuted] = useState<boolean>(false);

  useEffect(() => {
    if (player.audio.current) {
      player.audio.current.volume = volume / 100;
      player.audio.current.muted = muted;
    }
  }, [volume, player.audio, muted]);
  const VolumeButton = () => {
    return (
      <button
        onClick={() => setMuted((prev) => !prev)}
        className="image image-48 button-clear color-text-invert"
      >
        {muted || volume < 5 ? <MutedSvg /> : <VolumeMediumSvg />}
      </button>
    );
  };
  return (
    <>
      <VolumeButton />
      <div className="display-flex align-items-center padding-inline-xxs ">
        <input
          className="bg-color-body slider"
          aria-label="Volume Slider"
          type="range"
          min={0}
          max={100}
          value={volume}
          onChange={(e) => setVolume(Number(e.target.value))}
        />
      </div>
    </>
  );
};

export default VolumeControl;

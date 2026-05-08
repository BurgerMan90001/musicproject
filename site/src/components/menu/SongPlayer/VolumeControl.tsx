import { useEffect, useState, type JSX } from "react";
import { usePlayerStore } from "../../../hooks/player";
import { MutedSvg, VolumeMediumSvg } from "../Svg";

const VolumeControl = () => {
  const player = usePlayerStore();
  const [value, setValue] = useState(0);
  const [svg, setSvg] = useState<JSX.Element>(<MutedSvg />);
  useEffect(() => {
    if (player.audio.volume < 0.1) {
      setSvg(<MutedSvg />);
      return;
    }
    setSvg(<VolumeMediumSvg />);
  }, [player.audio.volume]);

  const VolumeButton = () => {
    return (
      <button
        onClick={() => {
        }}
        className="image image-48 button-clear color-text-invert"
      >
        {svg}
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
          value={value}
          onChange={(e) => {
            player.audio.volume = parseInt(e.target.value) / 100;
            setValue(parseInt(e.target.value));
          }}
        />
      </div>
    </>
  );
};

export default VolumeControl;

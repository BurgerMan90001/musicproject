import { useCallback, useEffect, useState, type JSX } from "react";
import { usePlayerStore } from "../../../hooks/player";
import { MutedSvg, PauseSvg, PlaySvg, VolumeMediumSvg } from "../Svg";

const PlayButton = () => {
  const player = usePlayerStore();
  const [svg, setSvg] = useState<JSX.Element>(<PauseSvg />);

  return (
    <button
      onClick={() => {
        player.togglePause();
        if (player.audio.paused) {
          setSvg(<PauseSvg />);
          return;
        }

        setSvg(<PlaySvg />);
      }}
      className="image image-48 button-clear color-text-invert"
    >
      {svg}
    </button>
  );
};

export { PlayButton };

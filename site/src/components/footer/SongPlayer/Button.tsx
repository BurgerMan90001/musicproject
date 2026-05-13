import { useEffect, useState } from "react";
import { usePlayerStore } from "../../../hooks/player";
import { PauseSvg, PlaySvg } from "../Svg";

const PlayButton = () => {
  const player = usePlayerStore();
  const [paused, setPaused] = useState<boolean>(false);

  useEffect(() => {
    if (paused) {
      player.audio.current?.play();
      return;
    }
    player.audio.current?.pause();
  }, [player.audio, paused]);

  return (
    <button
      onClick={() => setPaused((prev) => !prev)}
      className="image image-48 button-clear color-text-invert"
    >
      {paused ? <PlaySvg /> : <PauseSvg />}
    </button>
  );
};

export { PlayButton };

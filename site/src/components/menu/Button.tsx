import { usePlayerStore } from "../../hooks/player";

const VolumeButton = () => {
  const player = usePlayerStore();

  var icon;
  if (player.volume === "0") {
    icon = (
      <svg
        fill="currentColor"
        xmlns="http://www.w3.org/2000/svg"
        viewBox="0 0 24 24"
      >
        <path d="M12,4L9.91,6.09L12,8.18M4.27,3L3,4.27L7.73,9H3V15H7L12,20V13.27L16.25,17.53C15.58,18.04 14.83,18.46 14,18.7V20.77C15.38,20.45 16.63,19.82 17.68,18.96L19.73,21L21,19.73L12,10.73M19,12C19,12.94 18.8,13.82 18.46,14.64L19.97,16.15C20.62,14.91 21,13.5 21,12C21,7.72 18,4.14 14,3.23V5.29C16.89,6.15 19,8.83 19,12M16.5,12C16.5,10.23 15.5,8.71 14,7.97V10.18L16.45,12.63C16.5,12.43 16.5,12.21 16.5,12Z" />
      </svg>
    );
  } else {
    icon = (
      <svg
        fill="currentColor"
        xmlns="http://www.w3.org/2000/svg"
        viewBox="0 0 24 24"
      >
        <path d="M5,9V15H9L14,20V4L9,9M18.5,12C18.5,10.23 17.5,8.71 16,7.97V16C17.5,15.29 18.5,13.76 18.5,12Z" />
      </svg>
    );
  }

  return (
    <button
      onClick={player.mute}
      className="image image-48 button-clear color-text-invert"
    >
      {icon}
    </button>
  );
};

const PlayButton = () => {
  const player = usePlayerStore();
  var icon;
  if (player.paused) {
    icon = (
      <svg
        fill="currentColor"
        xmlns="http://www.w3.org/2000/svg"
        viewBox="0 0 24 24"
      >
        <path d="M14,19H18V5H14M6,19H10V5H6V19Z" />
      </svg>
    );
  } else {
    icon = (
      <svg
        fill="currentColor"
        xmlns="http://www.w3.org/2000/svg"
        viewBox="0 0 24 24"
      >
        <path d="M8,5.14V19.14L19,12.14L8,5.14Z" />
      </svg>
    );
  }
  return (
    <button
      onClick={player.togglePause}
      className="image image-48 button-clear color-text-invert"
    >
      {icon}
    </button>
  );
};

export { VolumeButton, PlayButton };

import { type JSX } from "react";
import { usePlayer } from "../../hooks/player";

const PlayIcon = () => {
  const player = usePlayer();
  
  if (player.paused) {
    return (
      <svg
        fill="currentColor"
        xmlns="http://www.w3.org/2000/svg"
        viewBox="0 0 24 24"
      >
        <path d="M14,19H18V5H14M6,19H10V5H6V19Z" />
      </svg>
    );
  }
  return (
    <svg
      fill="currentColor"
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 24 24"
    >
      <path d="M14,19H18V5H14M6,19H10V5H6V19Z" />
    </svg>
  );
};

const VolumeIcon = () => {
  const player = usePlayer();

  return (
    <svg
      fill="currentColor"
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 24 24"
    >
      <path d="M5,9V15H9L14,20V4L9,9M18.5,12C18.5,10.23 17.5,8.71 16,7.97V16C17.5,15.29 18.5,13.76 18.5,12Z" />
    </svg>
  );
};
const Button = ({ children }: { children: JSX.Element }) => {
  return (
    <button className="image image-48 button-clear color-text">
      {children}
    </button>
  );
};

export { Button, PlayIcon, VolumeIcon };

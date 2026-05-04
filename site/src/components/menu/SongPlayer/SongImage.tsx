import { usePlayerStore } from "../../../hooks/player";
import { SongPlaceholderIcon } from "../../svg/Placeholder";

const SongImage = () => {
  const queue = usePlayerStore((state) => state.queue);

  if (queue[0] && queue[0].image) {
    return queue[0].image;
  }
  return <SongPlaceholderIcon />;
};

export default SongImage;

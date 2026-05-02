import { useVolume } from "../../../hooks/player";

const VolumeControl = () => {
  const volume = useVolume();

  return (
    <div className="display-flex align-items-center padding-inline-xxs ">
      <input
        className="bg-color-body slider"
        aria-label="Volume Slider"
        type="range"
        value={volume.volume}
        onChange={(e) => {
          volume.change(e.target.value);
        }}
      />
    </div>
  );
};

export default VolumeControl;

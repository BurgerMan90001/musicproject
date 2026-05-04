const ProgressBar = () => {
  return (
    <div className="display-flex flex-column align-items-center padding-inline-xxs gap-xxs bg-color-body-darker layout-footer width-300">
      <span>0:00 / 0:00</span>
      <input
          id="progressBar"
          aria-label="Progress Bar"
          className="slider"
          type="range"
        />
    </div>
  );
};

export default ProgressBar;

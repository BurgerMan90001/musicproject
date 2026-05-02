import type { JSX } from "react";

const OauthButton = ({ name, icon }: { name: string; icon: JSX.Element }) => {
  return (
    <button className="display-flex button-filled align-items-center">
      {icon}
      <span className="color-text-invert font-size-md font-weight-semibold padding-xxs">
        {name}
      </span>
    </button>
  );
};

export default OauthButton;

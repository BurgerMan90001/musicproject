import type { JSX } from "react";
import { API_URL } from "../../../config/env";

const Button = ({
  name,
  path,
  icon,
}: {
  name: string;
  path: string;
  icon?: JSX.Element;
}) => {
  const action = async () => {
    window.location.href = `${API_URL}${path}`;
  };
  if (icon) {
    return (
      <button
        onClick={action}
        className="display-flex button-filled align-items-center"
      >
        {icon}
        <span className="color-text-invert font-size-md font-weight-semibold padding-xxs">
          {name}
        </span>
      </button>
    );
  }
  return (
    <div className="bg-color-body-darker">
      <button onClick={action} className="display-flex width-full align-items-center">
        {icon}
        <span className="color-text font-size-md padding-xxs">{name}</span>
      </button>
    </div>
  );
};

export default Button;

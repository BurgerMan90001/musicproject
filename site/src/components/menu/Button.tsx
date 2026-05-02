import { type JSX } from "react";

const Button = ({ children }: { children: JSX.Element }) => {
  return (
    <button className="image image-48 button-clear color-text-invert">
      {children}
    </button>
  );
};

export { Button };

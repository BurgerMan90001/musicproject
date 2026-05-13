import { Link } from "react-router";

export interface NavItem {
  name: string;
  to: string;
  disabled?: boolean;
}
interface Popover {
  title: string;
  buttons: NavItem[];
}
export const Popover = (popover: Popover) => {
  const popoverContent = (
    <>
      <summary>{popover.title}</summary>
      <div className="popover-content bg-color-body-dark display-flex flex-column box-shadow font-weight-normal">
        {popover.buttons.map((n) => {
          return (
            <Link
              key={n.name}
              to={n.to}
              className="button-clear padding-xxs font-size-sm width-150"
            >
              {n.name}
            </Link>
          );
        })}
      </div>
    </>
  );
  return <div className="popover padding-xxs">{popoverContent}</div>;
};

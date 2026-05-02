import { useContext } from "react";
import { Link } from "react-router";
import AuthContext from "../../hooks/auth";
interface NavItem {
  name: string;
  to?: string;
  disabled?: boolean;
}
interface Popover {
  to?: string;
  title: string;
  buttons: NavItem[];
}
const Popover = (popover: Popover) => {
  const popoverContent = (
    <>
      <summary>{popover.title}</summary>
      <div className="popover-content bg-color-body-dark display-flex flex-column box-shadow font-weight-normal">
        {popover.buttons.map((n) => {
          if (n.to) {
            return (
              <Link
                to={n.to}
                className="button-clear padding-xxs font-size-sm width-150"
              >
                {n.name}
              </Link>
            );
          }

          return (
            <button className="button-clear padding-xxs font-size-sm width-150">
              {n.name}
            </button>
          );
        })}
      </div>
    </>
  );
  if (popover.to) {
    return (
      <Link to={popover.to} className="popover padding-xxs">
        {popoverContent}
      </Link>
    );
  }
  return <div className="popover padding-xxs">{popoverContent}</div>;
};
function Navbar() {
  const auth = useContext(AuthContext);

  const create: NavItem[] = [
    { name: "Upload Song", to: "/upload?type=uploadSong" },
  ];

  const playlists: NavItem[] = [];
  if (playlists.length == 0) {
    playlists[0] = { name: "New", to: "/playlists?type=new" };
  }

  const signupButton = { name: "Signup", to: "/signup" };
  const loginButton = { name: "Login", to: "/login" };
  const settingsButton = { name: "Settings", to: "/settings" };

  const profile: NavItem[] = [signupButton, loginButton, settingsButton];

  return (
    <nav
      className="display-flex gap-xxs bg-color-body-dark border-top padding-inline-xs font-weight-bold"
      id="navbar"
    >
      <Link className="button-clear padding-xxs" to="/">
        Home
      </Link>

      <Popover to="/playlists" title="Playlists" buttons={playlists} />

      {/* Check if authenticated  */}
      <Popover to="/profiles" title="Profile" buttons={profile} />

      {/* Check if authenticated  */}

      <Popover title="Create" buttons={create} />
    </nav>
  );
}

export default Navbar;

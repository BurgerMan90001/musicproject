import { Link } from "react-router";
import { useAuthStore } from "../../hooks/auth";
import { Popover, type NavItem } from "./Popover";

function Navbar() {
  const auth = useAuthStore();

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

  var profile: NavItem[] = [signupButton, loginButton, settingsButton];
  if (!auth.authenticated()) {
    profile = [signupButton, loginButton, settingsButton];
  } else {
    if (auth.user) {
      profile = [
        { name: "tsets", to: "/" },
        { name: "Logout", to: "/logout" },
      ];
    }
  }

  return (
    <nav
      className="display-flex gap-xxs bg-color-body-dark border-top padding-inline-xs font-weight-bold"
      id="navbar"
    >
      <Link className="button-clear padding-xxs" to="/">
        Home
      </Link>

      <Popover title="Playlists" buttons={playlists} />

      {/* Check if authenticated  */}

      <Popover title="Profile" buttons={profile} />

      {/* Check if authenticated  */}
      <Popover title="Create" buttons={create} />
    </nav>
  );
}

export default Navbar;

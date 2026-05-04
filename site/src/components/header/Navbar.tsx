import { Link } from "react-router";
import { useAuthStore } from "../../hooks/auth";
import { Popover, type NavItem } from "./Popover";
const create: NavItem[] = [
  { name: "Dashboard", to: "/create" },
  { name: "Upload Song", to: "/create/upload?type=song" },
];

const playlists: NavItem[] = [
  { name: "Discovery", to: "/playlists/discovery" },
  { name: "Library", to: "/playlists" },
  { name: "New", to: "/playlists/new" },
];
function Navbar() {
  const auth = useAuthStore();

  var profile: NavItem[] = [];
  if (!auth.user) {
    profile.push(
      { name: "Signup", to: "/signup" },
      { name: "Login", to: "/login" },
    );
  } else {
    profile.push(
      { name: auth.user.email, to: "/users/" + auth.user.id },
      { name: "Logout", to: "/logout" },
    );
  }
  profile.push({ name: "Settings", to: "/settings" });

  return (
    <nav
      className="display-flex gap-xxs bg-color-body-dark border-top padding-inline-xs font-weight-bold"
      id="navbar"
    >
      <Link className="button-clear padding-xxs" to="/">
        Home
      </Link>

      <Popover title="Playlists" buttons={playlists} />

      <Popover title="Profile" buttons={profile} />

      <Popover title="Create" buttons={create} />
    </nav>
  );
}

export default Navbar;

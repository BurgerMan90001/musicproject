import { Link } from "react-router";
// import { useAuthStore } from "../../hooks/auth";
import { Popover, type NavItem } from "./Popover";
const create: NavItem[] = [
  // { name: "Dashboard", to: "/create" },
  { name: "Upload Song", to: "/create/upload" },
];

const library: NavItem[] = [
  // { name: "Playlists", to: "/playlists" },
  { name: "Songs", to: "/songs" },
  // { name: "New", to: "/playlists/new" },
];
function Navbar() {
  return (
    <nav
      className="display-flex gap-xxs bg-color-body-dark border-top padding-inline-xs font-weight-bold"
      id="navbar"
    >
      <Link className="button-clear padding-xxs" to="/">
        Home
      </Link>

      <Popover title="Library" buttons={library} />

      <Popover title="Create" buttons={create} />
    </nav>
  );
}
{
  /* <Popover title="Profile" buttons={profile} /> */
}
export default Navbar;

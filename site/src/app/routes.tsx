import Dashboard from "./pages/create/Dashboard";
import ErrorPage from "./pages/NotFound";
import Legal from "./pages/legal/Legal";
import Home from "./pages/Home";
import Upload from "./pages/create/Upload";
import Signup from "./pages/auth/Signup";
import Login from "./pages/auth/Login";
import Profile from "./pages/Profile";
import Logout from "./pages/auth/Logout";
import Playlists from "./pages/playlists/Playlists";
import NewPlaylist from "./pages/playlists/NewPlaylist";
import Discovery from "./pages/playlists/Discovery";

const routes = [
  {
    path: "/",
    element: <Home />,
    errorElement: <ErrorPage />,
    // children: [
    //   //{ index: true, element: <Dashboard /> },
    //   //{ path: "dashboard", element: <Dashboard /> },
    //   //{ path: "/profile/:name", element: <Profile /> },
    // ],
  },
  {
    path: "/signup",
    element: <Signup />,
  },
  {
    path: "/login",
    element: <Login />,
  },
  {
    path: "/logout",
    element: <Logout />,
  },
  {
    path: "/profile",
    element: <Profile />,
  },
  {
    path: "/create",
    element: <Dashboard />,
    children: [{ path: "upload", element: <Upload /> }],
  },
  {
    path: "/playlists",
    element: <Playlists />,
    children: [
      { path: "new", element: <NewPlaylist /> },
      { path: "discovery", element: <Discovery /> },
    ],
  },
  {
    path: "/legal",
    element: <Legal />,
  },
];

export default routes;

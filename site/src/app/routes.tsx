import Dashboard from "./pages/create/Create";
import ErrorPage from "./pages/NotFound";
import Legal from "./pages/legal/Legal";
import Home from "./pages/Home";
import Upload from "./pages/create/Upload";
import Signup from "./pages/auth/Signup";
import Login from "./pages/auth/Login";
import Profile from "./pages/Profile";
import Logout from "./pages/auth/Logout";
import Blog from "./pages/blog/Blog";
import Songs from "./pages/library/Songs";

const routes = [
  {
    path: "/",
    element: <Home />,
    errorElement: <ErrorPage />,
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
  // {
  //   path: "/playlists",
  //   element: <Playlists />,
  //   children: [
  //     { path: "new", element: <NewPlaylist /> },
  //     { path: "library", element: <Library /> },
  //   ],
  // },
  {
    path: "/songs",
    element: <Songs />,
  },
  {
    path: "/legal",
    element: <Legal />,
  },
  {
    path: "/blog",
    element: <Blog />,
  },
];

export default routes;

import Dashboard from "./pages/create/Create";
import ErrorPage from "./pages/NotFound";
import Legal from "./pages/legal/Legal";
import Home from "./pages/Home";
import Upload from "./pages/create/Upload";
import Signup from "./pages/auth/Signup";
import Login from "./pages/auth/Login";
import Profile from "./pages/Profile";
import Create from "./pages/create/Create";

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
    path: "/dashboard",
    element: <Dashboard />,
    children: [{ path: "upload", element: <Upload /> }],
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
    path: "/profile",
    element: <Profile />,
  },
  {
    path: "/create",
    element: <Create />,
  },
  {
    path: "/legal",
    element: <Legal />,
  },
];

export default routes;

import App from "../App";
import Dashboard from "./Dashboard";
//import ErrorPage from "./ErrorPage";
import Profile from "./Profile";
import Upload from "./Upload";

const routes = [
  {
    path: "/",
    element: <App />,
    children: [
      { index: true, element: <Dashboard /> },
      { path: "dashboard", element: <Dashboard /> },
      { path: "/profile/:name", element: <Profile /> },
    ],
  },
  {
    path: "/upload",
    element: <Upload />,
  },
];

export default routes;

import App from "../App";
import Dashboard from "./Dashboard";
import ErrorPage from "./ErrorPage";
import Profile from "./profile/Profile";

const routes = [
  {
    path: "/",
    element: <App />,
    errorElement: <ErrorPage />,
    children: [
      { index: true, element: <Dashboard /> },
      { path: "dashboard", element: <Dashboard /> },
      { path: "/profile/:name", element: <Profile />,}
    ],
  },
];

export default routes;

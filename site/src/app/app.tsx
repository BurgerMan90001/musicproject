import "../assets/css/app.css";

import { RouterProvider } from "react-router";
import { StrictMode, useEffect } from "react";
import { createRoot } from "react-dom/client";
import { createBrowserRouter } from "react-router";
import routes from "./routes";
import { COMPANY_NAME } from "../config/env";

const router = createBrowserRouter(routes);
function App() {
  useEffect(() => {
    document.title = `${COMPANY_NAME}`;
  });
  

  return (
    <StrictMode>
      <RouterProvider router={router} />
    </StrictMode>
  );
}

const root = document.getElementById("root")!;

createRoot(root).render(<App />);

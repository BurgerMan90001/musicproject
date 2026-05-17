
import Header from "../../../components/header/Header";
import { Outlet } from "react-router";
function Create() {
  return (
    <div className="layout-holy-grail" id="home">
      <Header />
      <Outlet />

    </div>
  );
}

export default Create;

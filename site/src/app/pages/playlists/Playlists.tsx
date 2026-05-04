import Footer from "../../../components/footer/Footer";
import Header from "../../../components/header/Header";
import { Outlet } from "react-router";
const Playlists = () => {
  return (
    <div className="layout-holy-grail height-full" id="home">
      <Header />
      <Outlet />
      <Footer />
    </div>
  );
};

export default Playlists;

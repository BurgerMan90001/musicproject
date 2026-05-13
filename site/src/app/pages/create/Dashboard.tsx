import Footer from "../../../components/footer/Footer";
import Header from "../../../components/header/Header";
import { Outlet } from "react-router";
function Dashboard() {
  for (let i = 0; i < 10; i++) {}
  return (
    <>
      <div className="layout-holy-grail" id="home">
        <Header />
        <Outlet />
        <Footer />
      </div>
    </>
  );
}

export default Dashboard;

import Footer from "../../../components/footer/Footer";
import Header from "../../../components/header/Header";
import { Outlet } from "react-router";
function Dashboard() {
  for (let i = 0; i < 10; i++) {}
  return (
    <>
      <div className="layout-holy-grail" id="home">
        <Header />
        {/* <aside className="bg-color-body-darker layout-aside width-200">
          <nav>
            <ul>
              <li>adsdasd</li>
              <li>adsdasd</li>
              <li>adsdasd</li>
            </ul>
          </nav>
        </aside> */}
        <Outlet />
        {/* <main className="layout-main margin-auto">
          <div className=" gap-xxs margin-inline-xxxl"></div>
        </main> */}
        <Footer />
      </div>
    </>
  );
}

export default Dashboard;

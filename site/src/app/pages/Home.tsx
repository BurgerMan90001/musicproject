import Header from "../../components/header/Header";
//import { useContext, createContext } from "react";
// import test from "../../assets/images/history.svg";
import SongQueue from "../../components/menu/SongQueue";
import SongPlayer from "../../components/menu/SongPlayer";
import Menu from "../../components/menu/Menu";
function Card({ title }: { title: string }) {
  // fetch("", {
  //   method: "GET",
  // });
  return (
    <button className="card bg-color-body-medium padding-lg">
      <p className="text-align-center">{title}</p>
    </button>
  );
}

function Home() {
  return (
    <>
      <div className="layout-holy-grail" id="home">
        <Header />
        <Menu />

        <main className="layout-main height-full scroll-vertical">
          <div className="gap-xxs margin-inline-xxxl display-flex flex-column">
      
            {/* <Card title="Genre" />
            <Card title="Genre" />
            <Card title="Genre" />
            <Card title="Genre" />
            <Card title="Genre " /> */}
          </div>
        </main>
        <footer className="layout-footer bg-color-body-dark flex-0">
          <span>asdasd</span>
        </footer>
      </div>
    </>
  );
}

export default Home;

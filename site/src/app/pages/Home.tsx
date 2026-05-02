import Header from "../../components/header/Header";
//import { useContext, createContext } from "react";
// import test from "../../assets/images/history.svg";
import SongQueue from "../../components/menu/SongQueue/SongQueue";
import SongPlayer from "../../components/menu/SongPlayer/SongPlayer";
import Menu from "../../components/menu/Menu";
import Media from "../../components/Media";
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
      <div className="layout-holy-grail height-full" id="home">
        <Header />
        <Menu />

        <main className="layout-main">
          <div className="gap-xxs margin-inline-xxxl display-flex flex-column">
            {/* <Card title="Genre" />
            <Card title="Genre" />
            <Card title="Genre" />
            <Card title="Genre" />
            <Card title="Genre " /> */}
            <Media />
          </div>
        </main>

        <footer className="layout-footer bg-color-body-dark">
          <span>asdasd</span>
        </footer>
      </div>
    </>
  );
}

export default Home;

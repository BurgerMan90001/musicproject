import Header from "../../components/header/Header";

import Media from "../../components/Media";
import Footer from "../../components/footer/Footer";
// function Card({ title }: { title: string }) {
// fetch("", {
//   method: "GET",
// });
// return (
//   <button className="card bg-color-body-medium padding-lg">
//     <p className="text-align-center">{title}</p>
//   </button>
// );
// }

function Home() {
  return (
    <>
      <div className="layout-holy-grail height-full" id="home">
        <Header />
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
        <Footer />
      </div>
    </>
  );
}

export default Home;

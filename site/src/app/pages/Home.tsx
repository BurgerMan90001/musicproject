import Header from "../../components/header/Header";
import { Link } from "react-router";
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

const Home = () => {
  return (
    <div className="height-full" id="home">
      <Header />
      <main className="display-flex flex-column layout-main  scroll-vertical gap-lg">
        <section className="bg-color-primary-light color-text-invert padding-xs gap-xs">
          <div className="display-flex align-items-center justifiy-content-center gap-lg">
            <div className="display-flex flex-column padding-sm gap-xs width-300 margin-inline-xxl">
              <h1 className="align-content-center line-height-sm font-weight-semibold">
                Songsled. A free music sharing platform.
              </h1>
              <p className="color-text-invert-subtle ">
                Upload and listen to music for free. No signup required.
              </p>
            </div>
            <div className="image image-192 margin-inline-xxl"></div>
          </div>
        </section>

        <section className="display-flex gap-xs bg-color-body-dark padding-xxl margin-block-lg">
          <div className="flex-1">
            <span className="font-weight-semibold margin-inline-xs"></span>
          </div>
          <div className="display-flex flex-1 line-height-md gap-xs">
            <div className="width-200">
              <h2>Its open source</h2>
              <p className="color-text-subtle">
                Feel free to repurpose it for something, but make sure to credit
                or ask me.
              </p>
            </div>

            <Link
              to="https://github.com/BurgerMan90001/songsled.com"
              target="_blank"
              rel="noopener noreferrer"
              className="button-success padding-xxs font-weight-semibold"
            >
              Github Repository
            </Link>
          </div>
          <div className="flex-1 font-weight-semibold">
            <span
              // to="/blog"
              className=" margin-inline-xs"
            >
              Blog
            </span>
            <div className="display-flex flex-column color-text-primary-light text-decoration-underline ">
              <Link to="/blog">Post</Link>
              <Link to="/blog">Post</Link>
            </div>
          </div>
        </section>
        <section className="display-flex gap-xs border-top border-bottom bg-color-body-dark padding-xs">
          <div>
            <h2 className="color-text">Explore songs</h2>
            <p className="color-text-subtle">Songs</p>
          </div>

          <Link
            className="display-flex align-items-center padding-xxs button-primary "
            to="/songs"
          >
            <span className="font-weight-semibold margin-inline-xs">Songs</span>
          </Link>
        </section>
        <section className="display-flex gap-xs bg-color-body-dark margin-block-lg padding-xs">
          <h2 className="color-text">Try uploading something!</h2>

          <Link
            className="display-flex font-weight-semibold  align-items-center padding-inline-sm button-primary "
            to="/create/upload"
          >
            <span className="">Upload</span>
          </Link>
        </section>
      </main>
      <footer className="padding-lg font-size-md  bg-color-body-dark">
        <div className="display-flex margin-inline-xl">
          <div className="color-text-subtle">
            <span>Contact me at paulcasigay@gmail.com for anything.</span>
          </div>
          <div className="margin-inline-xxxl">
            <Link className="button-clear" to="/blog">
              Blog
            </Link>
          </div>
        </div>
      </footer>
    </div>
  );
};

export default Home;

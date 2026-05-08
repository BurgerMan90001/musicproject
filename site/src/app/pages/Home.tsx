import Header from "../../components/header/Header";

// import Media from "../../components/Media";
import Footer from "../../components/footer/Footer";
import fetchApi from "../../lib/api";
// import test from "../../assets/images/cool-pic-128.jpg";
import { Link } from "react-router";
import { useEffect, useState } from "react";
import type { Song } from "../../types/song.types";
// import { usePlayerStore } from "../../hooks/player";
import { SongPlaceholderSvg } from "../../components/menu/Svg";
import { NetworkError } from "../../lib/error";
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
const Media = ({ song, size }: { song: Song; size: number }) => {
  // const player = usePlayerStore();

  var image = <SongPlaceholderSvg />;
  if (song.image) {
    image = <img src={song.image}></img>;
  }
  return (
    <div
      className={
        "test image color-text-invert bg-color-body-dark image-" + size
      }
    >
      <Link className="" to="#">
        {image}
      </Link>
    </div>
  );
};
// const test = async () => {
//   const res = await fetchApi(
//     "/v1/songs",
//     {
//       method: "GET",
//     },
//     {
//       n: String(10),
//     },
//   );
//   if (res) {
//     const json = await res.json();
//   }
// };
const n = "10";
const Home = () => {
  const [songs, _] = useState<Song[]>([]);
  const [error, setError] = useState<string>("");

  var topContent;

  useEffect(() => {
    fetchApi(
      "/v1/songs",
      {
        method: "GET",
      },
      {
        n: n,
      },
    )
      .then((res) => {
        if (!res) {
          setError(NetworkError);
          return;
        }
        return res?.json();
      })
      .catch((e) => setError(e));
  }, [songs]);

  if (songs) {
    topContent = songs.map((song: Song) => {
      return <Media song={song} size={192} />;
    });
  }
  if (songs.length == 0) {
    topContent = <h1 className="color-text-invert"> Nothing to show...</h1>;
  }
  if (error) {
    topContent = <h1 className="color-text-danger">{error}</h1>;
  }

  return (
    <div className="layout-holy-grail height-full" id="home">
      <Header />
      <main className="display-flex flex-column layout-main padding-lg scroll-vertical gap-lg">
        <section className="display-grid bg-color-body-darker padding-xs gap-xs">
          {/* <summary className="display-flex font-size-lg flex-column width-200">Title</summary> */}
          {topContent}
        </section>
        {/* <div className="margin-xs display-flex flex-column gap-md">
          <section className="margin-inline-xxxl ">
            <h2 className="">Browse by Genre</h2>
            <div className="display-flex margin-block-xs">
              <Media size={128} />
            </div>
          </section>
          <section className="bg-color-body-darker padding-xs margin-inline-xxxl">
            <h2 className="">Community Playlists</h2>
            <div className="display-flex margin-block-xs gap-xs">
              <Media size={128} />
              <Media size={128} />
              <Media size={128} />
            </div>
          </section>
          <section className="bg-color-body-darker padding-xs gap-xxs margin-inline-xxxl ">
            <h2 className="">Featured Artists</h2>
            <div className="display-flex">
              <Media size={128} />
            </div>
          </section>
        </div> */}
      </main>
      <Footer />
    </div>
  );
};

export default Home;

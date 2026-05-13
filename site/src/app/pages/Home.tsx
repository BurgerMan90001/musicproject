import Header from "../../components/header/Header";

import fetchApi from "../../lib/api";
import { Link } from "react-router";
import { useEffect, useState } from "react";
import type { Song } from "../../types/song.types";
import { SongPlaceholderSvg } from "../../components/footer/Svg";
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
        return res.json();
      })
      .then((json) => {
        console.log(json);
      })
      .catch((e) => setError(e));
  }, [songs]);

  if (songs) {
    topContent = songs.map((song: Song) => {
      return <Media song={song} size={192} />;
    });
  }
  if (songs.length === 0) {
    topContent = <h2 className="color-text"> Nothing to show...</h2>;
  }
  if (error) {
    topContent = <h2 className="color-text-danger">{error}</h2>;
  }

  return (
    <div className="layout-holy-grail height-full" id="home">
      <Header />
      <main className="display-flex flex-column layout-main  scroll-vertical gap-lg">
        <section className="bg-color-primary-light color-text-invert padding-xs gap-xs">
          <div className="display-flex align-items-center justifiy-content-center gap-lg">
            <div className="width-300 text-align-center padding-xs margin-inline-xxl">
              <h1 className="font-size-lg font-weight-semibold">Songsled</h1>
              <p className="color-text-invert-subtle ">
                Free uploads!! No signup required!!
              </p>
              <p> Email me at paulcasigay@gmail.com for anything</p>
            </div>
            <div className="image image-192 margin-inline-xxl"></div>
          </div>
        </section>
        <div className="bg-color-body-dark">{topContent}</div>
      </main>
      {/* <Footer /> */}
    </div>
  );
};

export default Home;

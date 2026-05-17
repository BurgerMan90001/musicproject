import { useEffect, useState } from "react";
import type { Song } from "../../../types/song.types";
import { SongPlaceholderSvg } from "../../../components/footer/Svg";
import fetchApi from "../../../lib/api";
import { NetworkError } from "../../../lib/error";
import Header from "../../../components/header/Header";
import { usePlayerStore } from "../../../hooks/player";
import Footer from "../../../components/footer/Footer";

const n = "40";

const Card = ({ song, size }: { song: Song; size: number }) => {
  const player = usePlayerStore();
  const [cover, setCover] = useState(<SongPlaceholderSvg />);

  useEffect(() => {
    if (song.cover) {
      setCover(<img src={song.cover} />);
    }
  }, []);

  return (
    <article className="position-relative" key={song.id}>
      <span className="font-weight-semibold">{song.name}</span>
      <div
        onClick={() => player.setSong(song)}
        className={
          "button button-clear image position-relative bg-color-body-darker color-text-invert image-" +
          size
        }
      >
        {cover}
      </div>
    </article>
  );
};

const Songs = () => {
  const [songs, setSongs] = useState<Song[]>();
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string>("");

  useEffect(() => {
    const fetch = async () =>
      await fetchApi(
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
          setSongs(json);
        })
        .catch((e) => setError(e.message))
        .finally(() => setLoading(false));
    fetch();
  }, []);
  var content;

  if (error) {
    content = <h2 className="color-text-danger">{error}</h2>;
  }
  if (loading) {
    content = <p>Loading songs...</p>;
  }

  if (songs) {
    content = (
      <div className="grid-item-container">
        {songs.map((song) => {
          return <Card song={song} size={128} key={song.id} />;
        })}
      </div>
    );
  }

  return (
    <div className="layout-holy-grail">
      <Header />
      <main className="layout-main ">{content}</main>
      <Footer />
    </div>
  );
};

export default Songs;
// var profile: NavItem[] = [];
// if (!auth.user) {
//   profile.push(
//     { name: "Signup", to: "/signup", disabled: true },
//     { name: "Login", to: "/login", disabled: true },
//   );
// } else {
//   profile.push(
//     { name: auth.user.email, to: "/users/" + auth.user.id },
//     { name: "Logout", to: "/logout" },
//   );
// }
// profile.push({ name: "Settings", to: "/settings" });

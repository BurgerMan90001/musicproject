import { useCallback, useState, type JSX } from "react";
import fetchApi from "../../../lib/api";
import type { Song } from "../../../types/song.types";
import { useNavigate } from "react-router";
import { escapeHTML } from "../../../lib/html";

const Input = ({
  label,
  children,
}: {
  label?: string;
  children: JSX.Element;
}) => {
  return (
    <>
      <div className="">
        {label && <label htmlFor="">{label}</label>}
        <div className="display-flex bg-color-body-dark border">{children}</div>
      </div>
    </>
  );
};
function Upload() {
  const [uploading, setUploading] = useState<boolean>(false);
  const [error, setError] = useState<string>("");
  const [audioFile, setAudioFile] = useState<File>();
  const [imageFile, setImageFile] = useState<File>();
  const [imagePreview, setImagePreview] = useState<string>("");

  const navigate = useNavigate();
  const metadata: Song = {
    id: "",
    name: "",
    genres: "",
    artists: "",
    creationDate: "",
    streams: 0,
    duration: 0,
    image: "",
    audio: "",
  };

  const uploadImage = useCallback(() => {
    if (!imageFile) {
      return;
    }
    const formData = new FormData();
    formData.append("file", imageFile);

    fetchApi("/v1/images/covers", {
      method: "PUT",
      body: formData,
      headers: {
        "Content-Type": imageFile.type,
        "Content-Length": String(imageFile.size),
      },
    })
      .then((res) => {
        if (!res) {
          setError("A network error was encountered");
          return;
        }
        return res.json();
      })
      .then((json) => {
        metadata.image = json.href;
      })
      .catch((e) => setError(e));
  }, [audioFile]);

  const uploadAudio = useCallback(() => {
    if (!audioFile) {
      setError("No audio file selected");
      return;
    }

    fetchApi(
      "/v1/audio",
      {
        method: "PUT",
        headers: {
          "Content-Type": audioFile.type,
        },
      },
      { filename: audioFile.name },
    )
      .then((res) => {
        if (!res || !res.ok) {
          setError("A network error was encountered");
          return;
        }
        return res.json();
      })
      .then((json) => {
        const formData = new FormData();
        formData.append("file", audioFile);

        fetch(json.links[0].href, {
          method: "PUT",
          body: formData,
          headers: {
            "Content-Type": audioFile.type,
            "Content-Length": String(audioFile.size),
          },
        }).then((res) => {
          if (!res.ok) {
            setError("A network error was encountered");
            return;
          }
          metadata.audio = json.href;
        });
      })
      .catch((e) => setError(e))
      .finally(() => {
        setUploading(false);
      });
  }, [audioFile]);

  if (uploading) {
    return <main>Loading</main>;
  }

  return (
    <main className="layout-main display-flex scroll-vertical">
      <form
        onSubmit={() => {
          uploadImage();
          uploadAudio();
          if (!error) {
            navigate("/create");
          }
        }}
        className="display-flex flex-column bg-color-body-medium font-size-md padding-xl margin-0-auto gap-xs"
      >
        <h1 className="font-weight-bold border-bottom">Upload Song</h1>

        <Input label="Song Name">
          <input
            type="text"
            id="songName"
            aria-label="Song name"
            // required={true}
            placeholder="My song"
            className="flex-1 font-size-sm padding-xs "
          />
        </Input>
        <div className="display-flex gap-xxs">
          <Input label="Genres">
            <input
              type="search"
              id="genres"
              aria-label="Genres"
              placeholder="Pop"
              // required={true}
              className="flex-1 font-size-sm padding-xs"
            />
          </Input>
          <Input label="Creation Date">
            <input
              type="text"
              id="creationDate"
              aria-label="Creation Date"
              placeholder="YYYY-MM-DD"
              className="flex-1 font-size-sm padding-xs "
            />
          </Input>
        </div>

        {/* <Input label="In Album">
          <input
            type="radio"
            id="file"
            // onChange={onFileChange}
            className="bg-color-body-dark"
          />
        </Input> */}
        {/* <div className="display "> */}
        <div className="display-flex gap-xxs margin-0-auto">
          <div className="">
            <Input label="Cover File">
              <input
                type="file"
                id="cover"
                accept="image/*"
                onChange={(event) => {
                  if (event.target.files) {
                    const file = event.target.files[0];
                    setImagePreview(URL.createObjectURL(file));
                    setImageFile(file);
                  }
                }}
                className="bg-color-body-dark padding-xxs color-text-subtle"
              />
            </Input>
            {imagePreview && (
              <>
                <span>Preview</span>
                <img src={escapeHTML(imagePreview)}></img>
              </>
            )}
          </div>

          <div>
            <Input label="Audio">
              <input
                type="file"
                id="audio"
                accept="audio/*"
                onChange={(event) => {
                  if (event.target.files) {
                    const file = event.target.files[0];
                    // setImagePreview(URL.createObjectURL(file));
                    setAudioFile(file);
                  }
                }}
                // required
                className="bg-color-body-dark padding-xxs color-text-subtle"
              />
            </Input>
          </div>
        </div>
        <span className="color-text-danger">{error}</span>
        <button className="button-primary padding-xxs">Submit</button>
      </form>
    </main>
  );
}

export default Upload;

import { useState } from "react";
import fetchApi from "../../../lib/api";
import { escapeHTML } from "../../../lib/html";
import { Input, InputList } from "./Input";
import { useArtists, useGenres } from "../../../hooks/upload";
import { SubmitButton } from "./Button";

function Upload() {
  const [uploading, setUploading] = useState<boolean>(false);
  const [error, setError] = useState<string>("");
  const [audioFile, setAudioFile] = useState<File>();
  const [imageFile, setImageFile] = useState<File>();
  const genres = useGenres();
  const artists = useArtists();

  const [imagePreview, setImagePreview] = useState<string>("");

  const uploadFile = async (endpoint: string, file: File) => {
    return fetchApi(
      endpoint,
      {
        method: "PUT",
        headers: {
          "Content-Type": file.type,
        },
      },
      { filename: file.name },
    )
      .then((res) => {
        if (!res || !res.ok) {
          setError("A network error was encountered");
          return;
        }
        return res.json();
      })
      .then(async (json) => {
        return fetch(json.links[0].href, {
          method: "PUT",
          body: file,
          headers: {
            "Content-Type": file.type,
          },
        })
          .then((res) => {
            if (!res.ok) {
              setError("A network error was encountered");
              return;
            }
            return json.location;
          })
          .catch((e) => console.log(e));
      })
      .catch((e) => console.log(e))
      .finally(() => {
        setUploading(false);
      });
  };

  const onSubmit = async (e: React.SubmitEvent<HTMLFormElement>) => {
    e.preventDefault();

    if (!audioFile) {
      setError("No audio file selected");
      return;
    }

    const form = e.target;
    const formData = new FormData(form);

    setUploading(true);
    const audio = await uploadFile("/v1/audio/songs", audioFile);
    formData.append("audio", audio);

    if (imageFile) {
      const image = await uploadFile("/v1/images/covers", imageFile);
      console.log(imageFile.type);
      formData.append("cover", image);
    }

    await fetchApi("/v1/songs", {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
      },
      body: JSON.stringify({
        name: formData.get("name"),
        artists: artists.list,
        genres: genres.list,
        creationDate: formData.get("creationDate"),

        cover: formData.get("cover"),
        audio: formData.get("audio"),
      }),
    })
      .then((res) => {
        if (!res || !res.ok) {
          setError("A network error was encountered");
          return;
        }
        return res.headers.get("location");
      })

      .catch((e) => console.log(e));
  };

  if (uploading) {
    return (
      <main className="layout-main display-flex scroll-vertical">
        Uploading
      </main>
    );
  }
  const ImagePreview = ({ file }: { file: string }) => {
    if (file === "") {
      return;
    }

    return (
      <div>
        <span>Preview</span>
        <div className="image image-128">
          <img src={escapeHTML(file)}></img>
        </div>
      </div>
    );
  };

  return (
    <main className="layout-main display-flex scroll-vertical">
      <form
        onSubmit={onSubmit}
        className="display-flex flex-column bg-color-body-medium font-size-md padding-xl margin-0-auto gap-xs"
      >
        <h1 className="font-weight-bold border-bottom">Upload Song</h1>
        <Input>
          <input
            type="text"
            name="name"
            aria-label="Song name"
            required={true}
            placeholder="My song"
            className="flex-1 font-size-sm padding-xs "
          />
        </Input>
        <div className="display-flex gap-xxs">
          <InputList label="Genres" placeholder="Song genre" state={genres} />
          <InputList
            label="Artists"
            placeholder="Song artists"
            state={artists}
          />
          <Input label="Creation Date">
            <input
              type="date"
              name="creationDate"
              aria-label="Creation Date"
              className="flex-1 font-size-sm padding-xs "
            />
          </Input>
        </div>

        <div className="display-flex gap-xxs margin-0-auto">
          <div>
            <Input label="Cover Image">
              <input
                type="file"
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

            <ImagePreview file={imagePreview} />
          </div>
          <Input label="Audio">
            <input
              type="file"
              accept="audio/*"
              required
              className="bg-color-body-dark padding-xxs color-text-subtle"
              onChange={(event) => {
                if (event.target.files) {
                  const file = event.target.files[0];
                  setAudioFile(file);
                }
              }}
            />
          </Input>
        </div>
        <span className="color-text-danger">{error}</span>
        <SubmitButton />
      </form>
    </main>
  );
}

export default Upload;

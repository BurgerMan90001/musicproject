import { useState, type JSX } from "react";
import fetchApi from "../../../lib/api";

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
  const [audioFile, setAudioFile] = useState<File>();
  const [_, setImageFile] = useState<File>();

  // Location for song upload endpoint
  // const [location, setLocation] = useState<string>();
  // const onFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
  //   if (event.target.files && event.target.files[0]) {
  //     const file = event.target.files[0];
  //     setSelectedFile(file);
  //   }
  // };
  // const onAudioFileChange = () => {
  //   if (audioFile) {
  //     uploadFile(audioFile, "/v1/songs");
  //   }
  // };
  const uploadFile = async (file: File, location: string) => {
    try {
      const formData = new FormData();
      formData.append("file", file);

      const res = await fetchApi(location, {
        method: "PUT",
        body: formData,
        headers: {
          "Content-Type": file.type,
        },
      });
      console.log(res?.body);

      // setSelectedFile(undefined);
      return;
    } catch (e) {
      console.log(e);
    }
  };

  const onSubmit = async () => {
    // const song: Song = {
    //   id: "",
    //   name: "",
    //   genres: "",
    //   artists: "",
    //   creationDate: "",
    //   streams: 0,
    //   duration: 0,
    //   image: "",
    //   url: "",
    // };

    // setLocation("");
    const formData = new FormData();
    if (audioFile) {
      formData.append("file", audioFile);
    }

    // alert(String(audioFile?.size));

    const res = await fetchApi("/v1/images/covers", {
      method: "PUT",
      // body: formData,
      headers: {
        // mode: "cors",
        // "Content-Length": "0",
        // "Orgin": "http://localhost:5173",
        // headers.append('Access-Control-Allow-Credentials' 'true');
        // Orgin: "https://songsled.com",
      },
    });

    console.log(res);
  };
  // interface input {
  //   type: string;
  //   id?: string;
  //   label?: string;
  //   required?: boolean;
  // }

  return (
    <main className="layout-main display-flex scroll-vertical">
      <form
        // onSubmit={onSubmit}
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
          {/* <Input label="Filename">
            <input
              type="text"
              id="filename"
              aria-label="Filename"
              placeholder="my_song.mp3"
              required={true}
              className="flex-1 font-size-sm padding-xs"
            />
          </Input> */}
        </div>

        {/* <Input label="In Album">
          <input
            type="radio"
            id="file"
            // onChange={onFileChange}
            className="bg-color-body-dark"
          />
        </Input> */}
        <div className="display-flex gap-xxs margin-0-auto">
          <Input label="Cover File">
            <input
              type="file"
              id="cover"
              onChange={(event) => {
                if (event.target.files) {
                  const file = event.target.files[0];
                  setImageFile(file);
                }
              }}
              className="bg-color-body-dark padding-xxs color-text-subtle"
            />
          </Input>

          <Input label="Audio">
            <input
              type="file"
              id="audio"
              onChange={(event) => {
                if (event.target.files) {
                  const file = event.target.files[0];
                  setAudioFile(file);
                }
              }}
              // required
              className="bg-color-body-dark padding-xxs color-text-subtle"
            />
          </Input>
        </div>

        <button
          type={"button"}
          onClick={onSubmit}
          className="button-primary padding-xxs"
        >
          Submit
        </button>
      </form>
    </main>
  );
}

export default Upload;

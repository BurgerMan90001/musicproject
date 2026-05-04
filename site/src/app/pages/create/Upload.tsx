import React, { useState } from "react";
import type { Song } from "../../../types/song.types";
import api from "../../../lib/api";
function Upload() {
  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  // Location for song upload endpoint
  const [location, setLocation] = useState<string>();
  const onFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.files && event.target.files[0]) {
      const file = event.target.files[0];

      setSelectedFile(file);
    }
  };

  const uploadFile = () => {
    if (selectedFile && location) {
      const formData = new FormData();
      formData.append("file", selectedFile);

      fetch(location, {
        method: "PUT",
        body: formData,
        headers: {
          "Content-Type": selectedFile.type,
        },
      })
        .then((response) => response.json())
        .then((json) => console.log(json));

      setSelectedFile(null);
    }
  };
  const uploadMetadata = async (song: Song) => {
    // interface response {

    // }
    const res = await api("/v1/upload/songs", {
      method: "POST",
      body: JSON.stringify(song),
      headers: {
        "Content-Type": "application/json",
      },
    });
    setLocation("");
    console.log(res);
    //   .then((response) => {
    //     // Save the location header url
    //     const url = response.headers.get("Location");
    //     if (url) {
    //       setLocation(url);
    //     }
    //     console.log(url);
    //     console.log(response.headers);
    //   })
    //   .then((json) => console.log(json));
  };
  const onSubmit = () => {
    const song: Song = {
      name: "",
      genre: "",
      artist: "",
      streams: 0,
      duration: 0,
      image: "",
      url: "",
    };
    uploadMetadata(song);

    uploadFile();
  };

  return (
    <main className="layout-main display-flex">
      <form className="display-flex flex-column bg-color-body-darker font-size-md padding-xxl margin-auto">
        <h1 className="font-weight-bold">Upload Song</h1>

        <input type="text" id="genre" aria-label="Genre" required={true} />
        <label htmlFor="name">Song Name</label>
        <input type="text" id="name" aria-label="Song name" required={true} />

        <label htmlFor="file">Song File</label>
        <input type="file" id="file" onChange={onFileChange}></input>

        <button type="submit" onClick={onSubmit} className="bg-black-25">
          Submit
        </button>
      </form>
    </main>
  );
}

export default Upload;

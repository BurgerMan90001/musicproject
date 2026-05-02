import React, { useState } from "react";
import type { Song } from "../../../types/song.types";
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
  const uploadMetadata = (song: Song) => {
    fetch("http://server:8081/v1/upload/songs", {
      method: "POST",
      body: JSON.stringify(song),
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then((response) => {
        // Save the location header url
        const url = response.headers.get("Location");
        if (url) {
          setLocation(url);
        }

        console.log(url);
        console.log(response.headers);
      })
      .then((json) => console.log(json));
  };
  const onSubmit = () => {
    const test: Song = {
      name: "",
      genre: "",
      streams: 0,
      duration: 0,
      image: "",
      url: "",
    };
    uploadMetadata(test);

    uploadFile();
  };

  return (
    <form className="margin-auto">
      <div className="bg-black-25 font-weight-bold">Upload Song</div>

      <input type="text" name="genre" required={true} />
      <input type="text" name="name" required={true} />

      <label htmlFor="file">Song File</label>
      <input type="file" name="file" onChange={onFileChange}></input>

      <button type="submit" onClick={onSubmit} className="bg-black-25">
        Submit
      </button>
    </form>
  );
}

export default Upload;

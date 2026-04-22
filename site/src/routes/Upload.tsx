import React, { useState } from "react";

function Upload() {
  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  function onFileChange(event: React.ChangeEvent<HTMLInputElement>) {
    if (event.target.files && event.target.files[0]) {
      const file = event.target.files[0];
      setSelectedFile(file);
    }
  }

  function uploadFile(location: string) {
    if (selectedFile && location) {
      fetch(location, {
        method: "PUT",
        headers: {
          "Content-Type": "",
        },
      })
        .then((response) => response.json())
        .then((json) => console.log(json));
    }
  }
  function onMetadataUpload() {
    fetch("http://server:8081/v1/upload/songs", {
      method: "POST",
      body: JSON.stringify({
        name: "",
        genre: "",
        image: "",
        filename: "",
      }),
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then((response) => {
        // Save the location header url
        const url = response.headers.get("Location");

        console.log(url);
        console.log(response.headers);
        if (url) {
          uploadFile("");
        }
      })
      .then((json) => console.log(json));
  }
  return (
    <>
      <h1 className="bg-black">File Upload</h1>
      <div>
        <input type="file" onChange={onFileChange}></input>
        <button onClick={onMetadataUpload}>test</button>
      </div>
    </>
  );
}

export default Upload;

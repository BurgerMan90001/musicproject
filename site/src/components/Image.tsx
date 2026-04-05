import { useState, useEffect } from "react";

function useImgURL(URL: string) {
  const [imgURL, setImgURL] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetch(URL)
      .then((response) => {
        if (!response.ok) {
          throw new Error(`HTTP error: Status ${response.status}`);
        }
        return response.json();
      })
      .then((response) => {
        setImgURL(response.image.original.url);
      })
      .catch((error) => setError(error))
      .finally(() => setLoading(false));
  }, []);

  return { imgURL, error, loading };
}
function Image({ URL }: { URL: string }) {
  const { imgURL, error, loading } = useImgURL(URL);

  if (loading) {
    return <p>Loading...</p>;
  }
  if (error) {
    return <p>A network error was encountered</p>;
  }
  return imgURL && <img src={imgURL} alt="placeholder text" />;
}
export default Image;

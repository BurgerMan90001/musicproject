import { useState, useEffect } from "react";

function useImgURL(URL: string) {
  const [imgUrl, setImgUrl] = useState(null);
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
        setImgUrl(response.image.original.url);
      })
      .catch((error) => setError(error))
      .finally(() => setLoading(false));
  }, []);

  return { imgUrl, error, loading };
}
function Image({ URL }: { URL: string }) {
  const { imgUrl, error, loading } = useImgURL(URL);

  if (loading) {
    return <p>Loading...</p>;
  }
  if (error) {
    return <p>A network error was encountered</p>;
  }
  return imgUrl && <img src={imgUrl} alt="placeholder text" />;
}
export default Image;

import { useState, useEffect } from "react";
import { HTTPError } from "../lib/error";

const useImgUrl = (url: string) => {
  const [imgUrl, setImgUrl] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetch(url)
      .then((response) => {
        if (!response.ok) {
          HTTPError(response);
        }
        return response.json();
      })
      .then((response) => {
        setImgUrl(response.url);
      })
      .catch((error) => setError(error))
      .finally(() => setLoading(false));
  }, []);

  return { imgUrl, error, loading };
};
const Image = (url: string) => {
  const { imgUrl, error, loading } = useImgUrl(url);

  if (loading) {
    return <p>Loading...</p>;
  }
  if (error) {
    return <p>A network error was encountered</p>;
  }
  return imgUrl && <img src={imgUrl} alt="placeholder text" />;
};
export default Image;

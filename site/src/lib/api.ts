import { API_URL } from "../config/env";

// Starts a request at the API_URL path should start with /{version}/...
async function fetchApi(
  path: string,
  init?: RequestInit,
  params?: Record<string, string>,
): Promise<Response | undefined> {
  try {
    var url = `${API_URL}${path}`;
    if (params) {
      url += "?" + new URLSearchParams(params).toString();
    }
    const res = await fetch(url, init);

    if (!res.ok) {
      throw new Error(
        `HTTP error: Status ${res.status} ${JSON.stringify(res.body)}`,
      );
    }
    return res;
  } catch (error) {
    console.log(error);
  }
}

export default fetchApi;

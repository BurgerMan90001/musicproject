import { API_URL } from "../config/env";

// Starts a request at the API_URL
async function fetchApi(path: string, init?: RequestInit): Promise<Response> {
  const res = await fetch(`${API_URL}${path}`, init);

  return res;
}

export default fetchApi;

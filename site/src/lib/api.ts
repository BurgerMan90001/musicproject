import { API_URL } from "../config/env";
import { HTTPError } from "./error";

async function api<T>(path: string, init?: RequestInit): Promise<T> {
  const res = await fetch(`${API_URL}${path}`, init);

  if (!res.ok) {
    HTTPError(res);
  }

  return (await res.json()) as T;
}

export default api;

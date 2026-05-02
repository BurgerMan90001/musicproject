interface ApiError {
  code: number;
  message: string;
  details: string;
  ok: boolean;
}

const HTTPError = (r: Response) => {
  throw new Error(`HTTP error: Status ${r.status} ${JSON.stringify(r.body)}`);
};

export { HTTPError, type ApiError };

const HTTPError = (r: Response) => {
  throw new Error(`HTTP error: Status ${r.status} ${JSON.stringify(r.body)}`);
};

export { HTTPError };

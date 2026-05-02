// import { render, screen } from "@testing-library/react";
// import App from "../src/App";
// import Header from "../src/components/Header";

import { describe, it, expect } from "vitest";

describe("App component", () => {
  it("renders correct heading", () => {
    //render(<App />);
    // using regex with the i flag allows simpler case-insensitive comparison
    //expect(screen.getByRole("heading").textContent).toMatch(/our first test/i);
  });
});

describe("Goob", () => {
  it("tes", () => {
    //render(<Header />);

    expect(1 + 1).toEqual(2);
    //expect(screen.getByRole("button").textContent).toMatch(/upload/i);
  });
});

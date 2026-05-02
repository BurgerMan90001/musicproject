import { apiUrl, companyName } from "./env";
import { describe, it, expect } from "vitest";

describe("config defined", () => {
  it("apiUrl", () => {
    expect(apiUrl).toBeDefined();
  });

  it("company name", () => {
    expect(companyName).toBeDefined();
  });
});

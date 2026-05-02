import { API_URL, COMPANY_NAME } from "./env";
import { describe, it, expect } from "vitest";

describe("config defined", () => {
  it("apiUrl", () => {
    expect(API_URL).toBeDefined();
  });

  it("company name", () => {
    expect(COMPANY_NAME).toBeDefined();
  });
});

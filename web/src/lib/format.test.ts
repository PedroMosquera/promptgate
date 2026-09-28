import { describe, expect, it } from "vitest";
import { formatDuration, truncate } from "./format";

describe("formatDuration", () => {
  it("renders sub-minute durations in seconds", () => {
    expect(formatDuration(15675)).toBe("15.7s");
  });

  it("renders longer durations as minutes and seconds", () => {
    expect(formatDuration(548000)).toBe("9m 8s");
  });
});

describe("truncate", () => {
  it("leaves short strings alone", () => {
    expect(truncate("abc", 10)).toBe("abc");
  });

  it("marks strings it shortened", () => {
    expect(truncate("abcdefghijkl", 5)).toBe("abcde...");
  });
});

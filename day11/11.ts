#!/usr/bin/env node
import { readFileSync } from "fs";

const MAX_GAP = 3;

function main() {
  let lines: string[];

  try {
    const data = readFileSync(process.argv[2], "utf8");
    lines = data.split("\n");
  } catch (err) {
    console.error(err);
    return;
  }

  let seatMap: string[][] = [];

  for (const l of lines) {
    seatMap.push(l.split(""));
  }

  // If a seat is empty (L) and there are no occupied seats adjacent to it,
  // the seat becomes occupied.
  // If a seat is occupied (#) and four or more seats adjacent to it are
  // also occupied, the seat becomes empty.
  // Otherwise, the seat's state does not change.

  for (let i = 0; i < seatMap.length; i++) {
    for (let j = 0; j < seatMap[i].length; j++) {
      switch (seatMap[i][j]) {
        case "L":
          console.log("got L");
          if (occupiedCheck(i, j, seatMap) === 0) {
            seatMap[i][j] = "#";
          }
          break;
        case "#":
          if (occupiedCheck(i, j, seatMap) >= 4) {
            seatMap[i][j] = "L";
          }
          break;
        case ".":
          console.log("got .");
          break;
      }
    }
  }

  console.log(lines);
}

// Return number of occupied seats adjacent to the provided seat
function occupiedCheck(row: number, col: number, lines: string[][]): number {
  return 8;
}

main();

#!/usr/bin/env node
import { readFileSync } from "fs";

// XMAS starts by transmitting a preamble of 25 numbers.
// After that, each number you receive should be the sum of any two of the 25
// immediately previous numbers. The two numbers will have different values,
// and there might be more than one such pair.

function main() {
  let lines: string[];

  try {
    const data = readFileSync(process.argv[2], "utf8");
    lines = data.split("\n");
  } catch (err) {
    console.error(err);
    return;
  }

  let numbers: number[] = [];

  for (const l of lines) {
    numbers.push(+l);
  }

  for (let i = 0; i + 25 < numbers.length; i++) {
    if (!sumExists(numbers.slice(i, i + 25), numbers[i + 25])) {
      console.log(numbers[i + 25]);
      break;
    }
  }
}

function sumExists(arr: number[], sum: number): boolean {
  let m: { [key: string]: boolean } = {};
  for (const num of arr) {
    if (sum - num in m) {
      return true;
    }
    m[num] = true;
  }
  return false;
}

main();

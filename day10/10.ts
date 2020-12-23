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

  let numbers: number[] = [];

  for (const l of lines) {
    if (l) {
      numbers.push(+l);
    }
  }

  numbers.sort((a, b) => a - b);
  console.log(numbers);

  let gapCounts: { [key: number]: number } = { 0: 0, 1: 0, 2: 0, 3: 0 };

  for (let i = 0; i < numbers.length; i++) {
    const prev = i > 0 ? numbers[i - 1] : 0;
    const curr = numbers[i];

    if (curr - prev > MAX_GAP) {
      break;
    }
    if (curr - prev <= MAX_GAP) {
      gapCounts[curr - prev] += 1;
    }
  }

  gapCounts[3]++;

  console.log(gapCounts[1] * gapCounts[3]);
}

main();

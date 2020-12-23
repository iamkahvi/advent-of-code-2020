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
  numbers = [0, ...numbers, numbers[numbers.length - 1] + MAX_GAP];

  console.log(numbers);

  let gapCounts: { [key: number]: number } = { 0: 0, 1: 0, 2: 0, 3: 0 };

  let start = 0;
  let count = 1;

  for (let i = 0; i < numbers.length; i++) {
    const prev = i > 0 ? numbers[i - 1] : 0;
    const curr = numbers[i];

    if (curr - prev > MAX_GAP) {
      break;
    }

    gapCounts[curr - prev] += 1;

    if (curr - prev === MAX_GAP) {
      if (i - start >= MAX_GAP) {
        count *= findPathsRecursive(0, numbers.slice(start, i));
      }
      start = i;
    }
  }

  console.log(gapCounts[1] * gapCounts[3]);
  console.log(count);
}

function findPathsRecursive(ind: number, arr: number[]): number {
  if (ind == arr.length - 1) {
    return 1;
  }

  let count = 0;

  let i = 1;
  while (ind + i <= arr.length && arr[ind + i] - arr[ind] <= MAX_GAP) {
    count += findPathsRecursive(ind + i, arr);
    i++;
  }

  return count;
}

main();

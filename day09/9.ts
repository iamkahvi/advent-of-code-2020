#!/usr/bin/env node
import { readFileSync } from "fs";

// XMAS starts by transmitting a preamble of 25 numbers.
// After that, each number you receive should be the sum of any two of the 25
// immediately previous numbers. The two numbers will have different values,
// and there might be more than one such pair.

const LENGTH = 25;

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

  let invalidNum = 0;
  for (let i = 0; i + LENGTH < numbers.length; i++) {
    if (!sumExists(numbers.slice(i, i + LENGTH), numbers[i + LENGTH])) {
      invalidNum = numbers[i + LENGTH];
      break;
    }
  }

  const [start, end] = slidingWindowSum(numbers, invalidNum);

  let max = 0;
  let min = Number.MAX_SAFE_INTEGER;
  for (const n of numbers.slice(start, end)) {
    if (n > max) {
      max = n;
    }
    if (n < min) {
      min = n;
    }
  }
  console.log(max + min);
}

// Find indexes of subarray that sums to the number provided, or return [0,0]
function slidingWindowSum(arr: number[], desiredSum: number): [number, number] {
  let sum = 0;

  for (let i = 0; i < arr.length; i++) {
    sum = 0;
    let j = i;

    while (sum < desiredSum) {
      sum += arr[j];
      if (sum === desiredSum) {
        return [i, j + 1];
      }
      if (sum > desiredSum) {
        break;
      }
      j++;
    }
  }
  return [0, 0];
}

// Return true if any two numbers in the array sum to the number provided
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

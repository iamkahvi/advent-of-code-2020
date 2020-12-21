#!/usr/bin/env node
"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var fs_1 = require("fs");
// XMAS starts by transmitting a preamble of 25 numbers.
// After that, each number you receive should be the sum of any two of the 25
// immediately previous numbers. The two numbers will have different values,
// and there might be more than one such pair.
var LENGTH = 25;
function main() {
    var lines;
    try {
        var data = fs_1.readFileSync(process.argv[2], "utf8");
        lines = data.split("\n");
    }
    catch (err) {
        console.error(err);
        return;
    }
    var numbers = [];
    for (var _i = 0, lines_1 = lines; _i < lines_1.length; _i++) {
        var l = lines_1[_i];
        numbers.push(+l);
    }
    var invalidNum = 0;
    for (var i = 0; i + LENGTH < numbers.length; i++) {
        if (!sumExists(numbers.slice(i, i + LENGTH), numbers[i + LENGTH])) {
            invalidNum = numbers[i + LENGTH];
            break;
        }
    }
    var _a = slidingWindowSum(numbers, invalidNum), start = _a[0], end = _a[1];
    var max = 0;
    var min = Number.MAX_SAFE_INTEGER;
    for (var _b = 0, _c = numbers.slice(start, end); _b < _c.length; _b++) {
        var n = _c[_b];
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
function slidingWindowSum(arr, desiredSum) {
    var sum = 0;
    for (var i = 0; i < arr.length; i++) {
        sum = 0;
        var j = i;
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
function sumExists(arr, sum) {
    var m = {};
    for (var _i = 0, arr_1 = arr; _i < arr_1.length; _i++) {
        var num = arr_1[_i];
        if (sum - num in m) {
            return true;
        }
        m[num] = true;
    }
    return false;
}
main();

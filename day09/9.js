#!/usr/bin/env node
"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var fs_1 = require("fs");
// XMAS starts by transmitting a preamble of 25 numbers.
// After that, each number you receive should be the sum of any two of the 25
// immediately previous numbers. The two numbers will have different values,
// and there might be more than one such pair.
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
    for (var i = 0; i + 25 < numbers.length; i++) {
        if (!sumExists(numbers.slice(i, i + 25), numbers[i + 25])) {
            console.log(numbers[i + 25]);
            break;
        }
    }
}
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

#!/usr/bin/env node
"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var fs_1 = require("fs");
var MAX_GAP = 3;
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
        if (l) {
            numbers.push(+l);
        }
    }
    numbers.sort(function (a, b) { return a - b; });
    console.log(numbers);
    var gapCounts = { 0: 0, 1: 0, 2: 0, 3: 0 };
    for (var i = 0; i < numbers.length; i++) {
        var prev = i > 0 ? numbers[i - 1] : 0;
        var curr = numbers[i];
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

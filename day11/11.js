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
    var seatMap = [];
    for (var _i = 0, lines_1 = lines; _i < lines_1.length; _i++) {
        var l = lines_1[_i];
        seatMap.push(l.split(""));
    }
    // If a seat is empty (L) and there are no occupied seats adjacent to it,
    // the seat becomes occupied.
    // If a seat is occupied (#) and four or more seats adjacent to it are
    // also occupied, the seat becomes empty.
    // Otherwise, the seat's state does not change.
    for (var i = 0; i < seatMap.length; i++) {
        for (var j = 0; j < seatMap[i].length; j++) {
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
function occupiedCheck(row, col, lines) {
    return 8;
}
main();

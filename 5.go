package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

// BPass : An airplane boarding pass
type BPass struct {
	row, column, id int
}

// FBFBBFF = 0101100 = 44
// Every seat also has a unique seat ID: multiply the row by 8, then add the column.
// In this example, the seat has ID 44 * 8 + 5 = 357.

func main() {
	fileString := os.Args[1]
	file, err := os.Open(fileString)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	passes := make([]BPass, 0)

	for scanner.Scan() {
		line := scanner.Text()
		bp := parseLine(line)
		passes = append(passes, bp)
	}

	MaxID := 128*8 + 8
	ids := make([]int, MaxID)

	for i := 0; i < len(passes); i++ {
		p := passes[i]
		ids[p.id] = 1
	}

	fmt.Println(ids)

	missingID := 0
	for id, found := range ids {
		if found == 0 && id > 0 && id < MaxID {
			if ids[id-1] == 1 && ids[id+1] == 1 {
				fmt.Println("Found it!", id)
				missingID = id
				break
			}
		}
	}
	fmt.Println(missingID)
}

func parseLine(s string) BPass {
	row := 0
	col := 0

	for i, c := range s {
		if i <= 6 && c == 'B' {
			pow := float64(6 - i)
			row += int(math.Pow(2, pow))
		} else if c == 'R' {
			pow := float64(9 - i)
			col += int(math.Pow(2, pow))
		}
	}

	id := row*8 + col

	return BPass{row, col, id}
}

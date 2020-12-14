package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type slope struct {
	right, down int
}

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

	slopes := []slope{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	forest := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		forest = append(forest, line)
	}

	count := 1
	for i := 0; i < len(slopes); i++ {
		trees := countTrees(slopes[i].right, slopes[i].down, forest)
		fmt.Println(trees)
		count *= trees
	}

	fmt.Println(count)

}

func countTrees(right int, down int, forest []string) int {
	encountered := 0
	width := len(forest[0])
	height := len(forest)
	// fmt.Println(width, height)

	row := 0
	column := 0
	for row < height {
		if forest[row][column] == '#' {
			encountered++
			// fmt.Println(row, forest[row], encountered, column)
		} else {
			// fmt.Println(row, forest[row])
		}
		row += down
		column = column + right
		if column > width-1 {
			column = column % width
		}
	}
	return encountered
}

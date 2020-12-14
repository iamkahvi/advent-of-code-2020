package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// Using the above example again, the three entries that sum to 2020 are 979, 366, and 675.
// Multiplying them together produces the answer, 241861950.

// In your expense report, what is the product of the three entries that sum to 2020?

func main() {
	fileString := os.Args[1]
	file, err := os.Open(fileString)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// ^(\d+-\d*) (\D): (.*)$
	e := regexp.MustCompile(`^(\d+)-(\d+) (\D): (.*)`)

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	valid := 0

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		matches := e.FindAllStringSubmatch(line, -1)

		start, _ := strconv.Atoi(matches[0][1])
		end, _ := strconv.Atoi(matches[0][2])
		letter := matches[0][3][0]
		password := matches[0][4]

		if end-1 >= len(password) {
			continue
		}

		if (password[start-1] == letter) && (password[end-1] != letter) {
			valid++
			fmt.Println(password)
		} else if (password[start-1] != letter) && (password[end-1] == letter) {
			valid++
			fmt.Println(password)
		}
	}
	fmt.Println(valid)

}

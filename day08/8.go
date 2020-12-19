package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Operations in the input file
const (
	ACC = iota
	JMP = iota
	NOP = iota
)

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

	lines := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	visited := make(map[int]bool)
	count, i := 0, 0

	for {
		if _, ok := visited[i]; ok {
			break
		} else {
			visited[i] = true
		}

		command, num := parseLine(lines[i])

		switch command {
		case ACC:
			count += num
			i++
		case JMP:
			i += num
		case NOP:
			i++
		default:
			fmt.Println("Command not recognized")
			break
		}
	}

	fmt.Println(count)
}

func parseLine(s string) (command int, num int) {
	num, _ = strconv.Atoi(s[5:])

	if s[4] == '-' {
		num *= -1
	}

	switch s[:3] {
	case "acc":
		return ACC, num
	case "nop":
		return NOP, num
	case "jmp":
		return JMP, num
	default:
		return 0, 0
	}
}

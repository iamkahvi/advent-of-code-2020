package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Type of operations in the input file
const (
	ACC = iota
	JMP = iota
	NOP = iota
)

// Op : An operation
type Op struct {
	command int
	num     int
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

	operations := make([]Op, 0)

	for scanner.Scan() {
		line := scanner.Text()
		o, _ := parseLine(line)

		operations = append(operations, o)
	}

	visited := make(map[int]bool)
	i, count := 0, 0

	// The program is supposed to terminate by attempting to execute an instruction
	// immediately after the last instruction in the file
	// i == len(operations)

	for {
		if i == len(operations) {
			break
		}

		op := operations[i]
		// fmt.Println(len(visited))

		if _, ok := visited[i]; ok {
			fmt.Println("Something went wrong", i)
			break
		}
		visited[i] = true

		switch op.command {
		case ACC:
			// fmt.Println(i+1, ": ACC", op)
			count += op.num
			i++
		case JMP:
			if ok, count := terminates(i+1, &operations, visited, count); ok {
				fmt.Println(i+1, ": NOP, ", count, op)
				return
			}
			// fmt.Println(i+1, ": JMP", op)

			i += op.num
		case NOP:
			if ok, count := terminates(i+op.num, &operations, visited, count); ok {
				fmt.Println(i+1, ": JMP, ", count, op)
				return
			}
			// fmt.Println(i+1, ": JMP", op)

			i++
		default:
			fmt.Println("Command not recognized")
			return
		}
	}

	fmt.Println(count)
}

func terminates(i int, operations *[]Op, prevVisited map[int]bool, count int) (bool, int) {
	visited := make(map[int]bool)
	for k, v := range prevVisited {
		visited[k] = v
	}
	for {
		if i == len(*operations) {
			return true, count
		}

		if _, ok := visited[i]; ok {
			return false, count
		}
		visited[i] = true

		op := (*operations)[i]

		switch op.command {
		case ACC:
			count += op.num
			i++
		case JMP:
			i += op.num
		case NOP:
			i++
		default:
			fmt.Println("Command not recognized")
			break
		}
	}
}

func parseLine(s string) (Op, bool) {
	num, _ := strconv.Atoi(s[5:])

	if s[4] == '-' {
		num *= -1
	}

	switch s[:3] {
	case "acc":
		return Op{ACC, num}, true
	case "nop":
		return Op{NOP, num}, true
	case "jmp":
		return Op{JMP, num}, true
	default:
		return Op{0, 0}, false
	}
}

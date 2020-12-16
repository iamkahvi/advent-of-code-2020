package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Bag : A bag with a color
type Bag struct {
	color    string
	contents map[string]int
}

// CanHoldBag : Can this bag hold some other bag
func (b *Bag) CanHoldBag(color string, bags map[string]*Bag) bool {
	if len(b.contents) == 0 {
		return false
	}
	for c := range b.contents {
		if c == color {
			return true
		}
		if b, ok := bags[c]; ok {
			if b.CanHoldBag(color, bags) {
				return true
			}
		}
	}
	return false
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

	bags := make(map[string]*Bag)

	for scanner.Scan() {
		line := scanner.Text()

		color, contents := parseLine(line)
		bag := Bag{color, contents}

		bags[bag.color] = &bag
	}

	fmt.Println("done parsing")

	count := 0
	for _, bag := range bags {
		if bag.CanHoldBag("shiny gold", bags) {
			count++
		}
	}

	fmt.Println("count:", count)
}

// Need to parse:
// vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
// into:
// "vibrant plum", {"faded blue": 5, "dotted black": 6}

func parseLine(s string) (string, map[string]int) {
	m := make(map[string]int)
	words := strings.Split(s, " ")
	bagColor := strings.Join(words[0:2], " ")

	i := 4
	for i < len(words) {
		num, err := strconv.Atoi(words[i])
		i++
		if err != nil {
			break
		}
		color := strings.Join(words[i:i+2], " ")
		m[color] = num
		i += 3
	}

	return bagColor, m
}

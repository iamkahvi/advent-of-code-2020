package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Group : A group of custom declaration forms
type Group struct {
	answers []Answer
}

// Answer : Answers to question on a form
type Answer map[byte]bool

// AddAnswer : Add an answer to a group
func (g *Group) AddAnswer(a Answer) []Answer {
	g.answers = append(g.answers, a)
	return g.answers
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

	groups := make([]*Group, 0)

	for scanner.Scan() {
		line := scanner.Text()
		m := make([]Answer, 0)
		g := Group{m}

		for scanner.Scan() {
			if line == "" {
				break
			}
			ans := parseLine(line)
			fmt.Println(line)
			g.AddAnswer(ans)
			line = scanner.Text()
		}

		groups = append(groups, &g)
		fmt.Println(g)
	}

	fmt.Println(len(groups))

}

func parseLine(s string) Answer {
	m := make(map[byte]bool)

	for i := 0; i < len(s); i++ {
		m[s[i]] = true
	}

	return m
}

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

// FindCount : Calculate a group's answer "count"
func (g *Group) FindCount() int {
	qsAnswered := make(map[byte]bool)
	for _, ans := range g.answers {
		for k := range ans {
			qsAnswered[k] = true
		}
	}

	return len(qsAnswered)
}

// FindEveryoneCount : Identify the questions to which everyone answered "yes"
func (g *Group) FindEveryoneCount() int {
	m := make(map[byte]int)
	for _, ans := range g.answers {
		for k := range ans {
			m[k]++
		}
	}

	count := 0
	for _, v := range m {
		if v == len(g.answers) {
			count++
		}
	}

	return count
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
		if line != "" {
			g := Group{make([]Answer, 0)}

			for {
				ans := parseLine(line)
				g.AddAnswer(ans)

				scanner.Scan()
				line = scanner.Text()
				if line == "" {
					break
				}
			}
			groups = append(groups, &g)
			// fmt.Println(g)
			// fmt.Println(g.FindEveryoneCount())
			// fmt.Println("\n")
		}
	}

	totalCount := 0
	for _, g := range groups {
		totalCount += g.FindEveryoneCount()
	}
	fmt.Println(totalCount)

}

func parseLine(s string) Answer {
	m := make(map[byte]bool)

	for i := 0; i < len(s); i++ {
		m[s[i]] = true
	}

	return m
}

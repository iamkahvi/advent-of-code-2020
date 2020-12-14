package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	scanner := bufio.NewScanner(file)

	// Map for storing complements
	m := make(map[int]int)

	// Build array of inputs
	nums := make([]int, 0)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		nums = append(nums, num)
	}

	fmt.Println(nums)

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		for k, v := range m {
			if v < num {
				continue
			} else if v > num {
				for j := i + 1; j < len(nums); j++ {
					if k+num+nums[j] == 2020 {
						fmt.Println("Got It!")
						fmt.Println(k, num, nums[j])
						fmt.Println(k * num * nums[j])
						i = len(nums)
						break
					}
				}
			} else if v == num {
				fmt.Println("Got It!")
				fmt.Println(k, num)
			}
		}
		complement := 2020 - num
		m[num] = complement
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

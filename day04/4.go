package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

	passports := make([]string, 0)
	data := ""

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			passports = append(passports, data)
			data = ""
		} else {
			data += line
		}
	}

	valid := len(passports)
	fmt.Println(valid)

	fields := []string{`byr:(19[2-9][0-9]|200[0-2])`, `iyr:(201[0-9]|2020)`, `eyr:(202[0-9]|2030)`,
		`hgt:(1[5-8][0-9]|19[0-3])cm`, `hcl:#[0-9a-f]{6}`, `ecl:(amb|blu|brn|gry|grn|hzl|oth)`,
		`pid:\d{9}`, `cid:`}

	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	// `byr:(19[2-9][0-9]|200[0-2])`
	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	// `iyr:(201[0-9]|2020)`
	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	// `eyr:(202[0-9]|2030)`
	// hgt (Height) - a number followed by either cm or in:
	//     If cm, the number must be at least 150 and at most 193.
	//     If in, the number must be at least 59 and at most 76.
	// `hgt:(1[5-8][0-9]|19[0-3])cm`
	// `hgt:(59|6[0-9]|7[0-6])in`
	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	// `hcl:#[0-9a-f]{6}`
	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	// `ecl:(amb|blu|brn|gry|grn|hzl|oth)`
	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	// `pid:\d{9}`
	// cid (Country ID) - ignored, missing or not.
	// `cid:`

	for i := 0; i < len(passports); i++ {
		passport := passports[i]
		for j := 0; j < len(fields); j++ {
			e := regexp.MustCompile(fields[j])
			if !e.Match([]byte(passport)) {
				if fields[j] == `cid:` {
					continue
				} else if fields[j] == `hgt:(1[5-8][0-9]|19[0-3])cm` {
					matched, _ := regexp.Match(`hgt:(59|6[0-9]|7[0-6])in`, []byte(passport))
					if !matched {
						// fmt.Println(passport)
						// fmt.Println(e.String())
						// fmt.Println(`hgt:(59|6[0-9]|7[0-6])in`)
						// fmt.Println("\n")
						valid--
						break
					} else {
						continue
					}
				} else {
					valid--
					break
				}
			} else {
				if j == 6 {
					fmt.Println(passport)
					fmt.Println(e.String())
					fmt.Println("\n")
				}
			}
		}
	}

	fmt.Println(valid)

}

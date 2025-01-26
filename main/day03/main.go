package main

import (
	"fmt"
	"regexp"

	"github.com/nitely/adventofgo2024/main/utils"
)

func muls(s string) int {
	result := 0
	re := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	matches := re.FindAllStringSubmatch(s, -1)
	for _, match := range matches {
		result += utils.Atoi(match[1]) * utils.Atoi(match[2])
	}
	return result
}

func part1(s string) {
	result := muls(s)
	fmt.Println(result)
}

func part2(s string) {
	result := 0
	re := regexp.MustCompile(`(?s)(^|do\(\)).+?(don't\(\)|$)`)
	matches := re.FindAllString(s, -1)
	for _, match := range matches {
		result += muls(match)
	}
	fmt.Println(result)
}

func main() {
	s := utils.MustReadFile("input.txt")
	part1(s)
	part2(s)
	fmt.Println("ok")
}

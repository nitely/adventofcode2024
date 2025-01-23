package main

import (
	"adventofgo2024/main/utils"
	"fmt"
	"strings"
)

type Direction int

const (
	up Direction = iota
	down
	left
	right
	diaglup
	diagldown
	diagrup
	diagrdown
)

func next(i, j int, dir Direction) (int, int) {
	switch dir {
	case up:
		i -= 1
	case down:
		i += 1
	case left:
		j -= 1
	case right:
		j += 1
	case diagldown:
		i, j = next(i, j, down)
		i, j = next(i, j, left)
	case diaglup:
		i, j = next(i, j, up)
		i, j = next(i, j, left)
	case diagrdown:
		i, j = next(i, j, down)
		i, j = next(i, j, right)
	case diagrup:
		i, j = next(i, j, up)
		i, j = next(i, j, right)
	default:
		panic(fmt.Sprintf("Direction not covered %v", dir))
	}
	return i, j
}

func xmasCheck(i, j int, s []string, dir Direction, word string) int {
	for k := range word {
		if i < 0 || j < 0 || i >= len(s) || j >= len(s[0]) {
			return 0
		}
		if s[i][j] != word[k] {
			return 0
		}
		i, j = next(i, j, dir)
	}
	return 1
}

func xmasCount(i, j int, s []string) int {
	result := 0
	dirs := []Direction{up, down, left, right, diagldown, diaglup, diagrdown, diagrup}
	for _, dir := range dirs {
		result += xmasCheck(i, j, s, dir, "XMAS")
	}
	return result
}

func part1(s []string) {
	result := 0
	for i, line := range s {
		for j := range line {
			result += xmasCount(i, j, s)
		}
	}
	fmt.Println(result)
}

func xmasCount2(i, j int, s []string) int {
	const word = "MAS"
	if xmasCheck(i, j, s, diagrdown, word)+xmasCheck(i+2, j+2, s, diaglup, word) > 0 {
		if xmasCheck(i, j+2, s, diagldown, word)+xmasCheck(i+2, j, s, diagrup, word) > 0 {
			return 1
		}
	}
	return 0
}

func part2(s []string) {
	result := 0
	for i, line := range s {
		for j := range line {
			result += xmasCount2(i, j, s)
		}
	}
	fmt.Println(result)
}

func main() {
	s := utils.MustReadFile("input.txt")
	s = strings.Trim(s, "\n")
	lines := strings.Split(s, "\n")
	part1(lines)
	part2(lines)
	fmt.Println("ok")
}

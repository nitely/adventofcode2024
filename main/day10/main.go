package main

import (
	"fmt"

	"github.com/nitely/adventofgo2024/main/utils"
)

type pos struct {
	y int
	x int
}

var directions = []pos{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

// part2 can use memoization to run in linear time, but meh
func walk(theMap [][]int, prev, x, y int, reached map[pos]bool) int {
	if y < 0 || y >= len(theMap) || x < 0 || x >= len(theMap[0]) {
		return 0
	}
	if prev+1 != theMap[y][x] {
		return 0
	}
	if reached[pos{y, x}] {
		return 0
	}
	if theMap[y][x] == 9 {
		//reached[pos{y, x}] = true  // part 1
		return 1
	}
	result := 0
	for _, n := range directions {
		result += walk(theMap, theMap[y][x], x+n.x, y+n.y, reached)
	}
	return result
}

func part1and2(theMap [][]int) {
	result := 0
	for y := range theMap {
		for x := range theMap[0] {
			if theMap[y][x] == 0 {
				reached := make(map[pos]bool)
				result += walk(theMap, -1, x, y, reached)
			}
		}
	}
	fmt.Println(result)
}

func main() {
	lines := utils.ByteLines("input.txt")
	theMap := utils.MakeMatrix(len(lines), len(lines[0]), 0)
	for i := range lines {
		utils.Assert(len(lines[i]) > 0)
		for j := range lines[i] {
			theMap[i][j] = int(lines[i][j]) - int('0')
		}
	}
	//fmt.Println(theMap)
	part1and2(theMap)
	fmt.Println("ok")
}

package main

import (
	"adventofgo2024/main/utils"
	"fmt"
)

type pos struct {
	i, j int
}

const (
	up = iota
	right
	down
	left
)

func part1(theMap [][]byte, i, j int) {
	steps := 0
	dirs := []pos{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	curr := up
	visited := make(map[pos]struct{})
	pi, pj := i, j
	for i >= 0 && j >= 0 && i < len(theMap) && j < len(theMap[0]) {
		if theMap[i][j] == '#' {
			curr = (curr + 1) % 4
			i, j = pi, pj
		}
		if _, ok := visited[pos{i, j}]; !ok {
			steps++
			visited[pos{i, j}] = struct{}{}
		}
		pi, pj = i, j
		i += dirs[curr].i
		j += dirs[curr].j
	}
	fmt.Println(steps)
}

type posDir struct {
	i, j, dir int
}

func walk(theMap [][]byte, i, j int) error {
	dirs := []pos{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	curr := up
	visited := make(map[posDir]struct{})
	pi, pj := i, j
	for i >= 0 && j >= 0 && i < len(theMap) && j < len(theMap[0]) {
		if theMap[i][j] == '#' {
			curr = (curr + 1) % 4
			i, j = pi, pj
			if _, ok := visited[posDir{i, j, curr}]; ok {
				return fmt.Errorf("loop found")
			}
			visited[posDir{i, j, curr}] = struct{}{}
		}
		pi, pj = i, j
		i += dirs[curr].i
		j += dirs[curr].j
	}
	return nil
}

// Try putting an object in every free cell;
// and check if it forms a loop.
// Runs in quadratic time.
func part2(theMap [][]byte, i, j int) {
	result := 0
	for i2 := range theMap {
		for j2 := range theMap[i2] {
			if theMap[i2][j2] == '.' {
				theMap[i2][j2] = '#'
				err := walk(theMap, i, j)
				if err != nil {
					result++
				}
				theMap[i2][j2] = '.'
			}
		}
	}
	fmt.Println(result)
}

func main() {
	theMap := utils.ByteLines("input.txt")
	utils.Assert(len(theMap[len(theMap)-1]) > 0)
	si, sj := 0, 0
	for i := range theMap {
		for j := range theMap[i] {
			if theMap[i][j] == '^' {
				si = i
				sj = j
			}
		}
	}
	part1(theMap, si, sj) // 5129
	part2(theMap, si, sj) // 1888
	fmt.Println("ok")
}

package main

import (
	"fmt"
	"strings"

	"github.com/nitely/adventofgo2024/main/utils"
)

func part1(rules map[int]map[int]bool, pages [][]int) {
	result := 0
	for _, page := range pages {
		pageOk := true
		for i, num := range page {
			for _, n := range page[i+1:] {
				if !rules[num][n] {
					pageOk = false
				}
			}
		}
		if pageOk {
			result += page[len(page)/2]
		}
	}
	fmt.Println(result)
}

func part2(rules map[int]map[int]bool, pages [][]int) {
	result := 0
	for _, page := range pages {
		pageOk := true
		for i, num := range page {
			for _, n := range page[i+1:] {
				if !rules[num][n] {
					pageOk = false
				}
			}
		}
		if !pageOk {
			//sort.Slice(page, func(i, j int) bool {
			//	return rules[page[i]][page[j]]
			//})
			page2 := utils.Sorted(page, func(i, j int) bool {
				return rules[i][j]
			})
			result += page2[len(page2)/2]
		}
	}
	fmt.Println(result)
}

func main() {
	s := utils.MustReadFile("input.txt")
	lines := strings.Split(s, "\n")
	rules := make(map[int]map[int]bool)
	pages := make([][]int, 0)
	for _, line := range lines {
		switch {
		case strings.Contains(line, "|"):
			parts := strings.Split(line, "|")
			utils.Assert(len(parts) == 2)
			a := utils.Atoi(parts[0])
			b := utils.Atoi(parts[1])
			if rules[a] == nil {
				rules[a] = make(map[int]bool)
			}
			rules[a][b] = true
		case line != "":
			parts := strings.Split(line, ",")
			utils.Assert(len(parts) > 1)
			pages = append(pages, []int{})
			ln := len(pages) - 1
			for _, n := range parts {
				pages[ln] = append(pages[ln], utils.Atoi(n))
			}
		}
	}
	part1(rules, pages)
	part2(rules, pages)
	fmt.Println("ok")
}

package main

import (
	"fmt"

	"github.com/nitely/adventofgo2024/main/utils"
)

type pos struct {
	i, j int
}

func antennas(theMap [][]byte) map[byte][]pos {
	result := make(map[byte][]pos)
	for i := range theMap {
		for j := range theMap[i] {
			if theMap[i][j] != '.' {
				result[theMap[i][j]] = append(result[theMap[i][j]], pos{i, j})
			}
		}
	}
	return result
}

func part1(theMap [][]byte) {
	theAntennas := antennas(theMap)
	antinodes := make(map[pos]struct{})
	for freq, ps := range theAntennas {
		for i := range ps {
			for j := range ps {
				di, dj := ps[j].i-ps[i].i, ps[j].j-ps[i].j
				idi, jdj := ps[i].i-di, ps[i].j-dj
				if idi >= 0 && idi < len(theMap) && jdj >= 0 && jdj < len(theMap[0]) {
					if theMap[idi][jdj] != freq {
						antinodes[pos{idi, jdj}] = struct{}{}
					}
				}
			}
		}
	}
	fmt.Println(len(antinodes))
}

func part2(theMap [][]byte) {
	theAntennas := antennas(theMap)
	antinodes := make(map[pos]struct{})
	for _, ps := range theAntennas {
		for i := range ps {
			for j := range ps {
				if j == i {
					continue
				}
				di, dj := ps[j].i-ps[i].i, ps[j].j-ps[i].j
				idi, jdj := ps[i].i-di, ps[i].j-dj
				for idi >= 0 && idi < len(theMap) && jdj >= 0 && jdj < len(theMap[0]) {
					antinodes[pos{idi, jdj}] = struct{}{}
					idi -= di
					jdj -= dj
				}
				// the antenna itself is antinode
				antinodes[pos{ps[i].i, ps[i].j}] = struct{}{}
			}
		}
	}
	fmt.Println(len(antinodes))
}

func main() {
	theMap := utils.ByteLines("input.txt")
	utils.Assert(len(theMap[len(theMap)-1]) > 0)
	part1(theMap) // 357
	part2(theMap)
	fmt.Println("ok")
}

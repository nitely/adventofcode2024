package main

import (
	"fmt"

	"github.com/nitely/adventofgo2024/main/utils"
)

func part1(s string) {
	blocks := make([]int, 0)
	n := 0
	for i := range s {
		num := int(s[i] - '0')
		for range num {
			if i%2 == 0 {
				blocks = append(blocks, n)
			} else {
				blocks = append(blocks, -1)
			}
		}
		if i%2 == 0 {
			n++
		}
	}
	i := 0
	j := len(blocks) - 1
	for i < j {
		if blocks[i] != -1 {
			i++
		}
		if blocks[j] == -1 {
			j--
		}
		if blocks[i] == -1 && blocks[j] != -1 {
			blocks[i] = blocks[j]
			blocks[j] = -1
		}
	}
	result := 0
	for i := range blocks {
		if blocks[i] == -1 {
			break
		}
		result += i * blocks[i]
	}
	//fmt.Println(blocks)
	fmt.Println(result)
}

type block struct {
	pos, size int
}

// input is 20k len, so quadratic time is ok
func part2(s string) {
	blocks := make([]block, 0)
	slots := make([]block, 0)
	pos := 0
	for i := range s {
		size := int(s[i] - '0')
		if i%2 == 0 {
			blocks = append(blocks, block{pos, size})
		} else if size > 0 {
			slots = append(slots, block{pos, size})
		}
		pos += size
	}
	//fmt.Println(blocks)
	//fmt.Println(slots)
	for i := len(blocks) - 1; i >= 0; i-- {
		for j := range slots {
			if slots[j].pos > blocks[i].pos {
				break
			}
			if slots[j].size >= blocks[i].size {
				blocks[i].pos = slots[j].pos
				slots[j].pos += blocks[i].size
				slots[j].size -= blocks[i].size
				break
			}
		}
	}
	result := 0
	for idNum, bk := range blocks {
		for i := range bk.size {
			result += (i + bk.pos) * idNum
		}
	}
	//fmt.Println(blocks)
	//fmt.Println(slots)
	fmt.Println(result)
}

func main() {
	line := utils.MustReadFile("input.txt")
	part1(line) // 6288599492129
	part2(line) // 6321896265143
	fmt.Println("ok")
}

package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/nitely/adventofgo2024/main/utils"
)

func calc(target int, nums []int, i, acc int) bool {
	if acc == target && i == len(nums) {
		return true
	}
	if acc > target || i >= len(nums) {
		return false
	}
	return calc(target, nums, i+1, acc*nums[i]) ||
		calc(target, nums, i+1, acc+nums[i])
}

func concat(a, b int) int {
	return utils.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
}

func calc2(target int, nums []int, i, acc int) bool {
	if acc == target && i == len(nums) {
		return true
	}
	if acc > target || i >= len(nums) {
		return false
	}
	return calc2(target, nums, i+1, acc*nums[i]) ||
		calc2(target, nums, i+1, acc+nums[i]) ||
		calc2(target, nums, i+1, concat(acc, nums[i]))
}

type record struct {
	testValue int
	nums      []int
}

func part1(s []record) {
	result := 0
	for _, rec := range s {
		if calc(rec.testValue, rec.nums, 1, rec.nums[0]) {
			result += rec.testValue
		}
	}
	fmt.Println(result)
}

func part2(s []record) {
	result := 0
	for _, rec := range s {
		if calc2(rec.testValue, rec.nums, 1, rec.nums[0]) {
			result += rec.testValue
		}
	}
	fmt.Println(result)
}

func main() {
	s := utils.MustReadFile("input.txt")
	records := make([]record, 0)
	re := regexp.MustCompile(`([0-9]+): (.+)`)
	matches := re.FindAllStringSubmatch(s, -1)
	for _, match := range matches {
		testValue := utils.Atoi(match[1])
		nums := utils.AtoiN(strings.Split(match[2], " "))
		records = append(records, record{testValue, nums})
	}
	part1(records)
	part2(records) // 61561126043536
	fmt.Println("ok")
}

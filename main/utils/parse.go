package utils

import (
	"strconv"
)

func Atoi(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}

func AtoiN(s []string) []int {
	result := make([]int, len(s))
	for i, x := range s {
		val, err := strconv.Atoi(x)
		if err != nil {
			panic(err)
		}
		result[i] = val
	}
	return result
}

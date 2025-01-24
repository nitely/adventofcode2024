package utils

import "sort"

func Assert(condition bool) {
	if !condition {
		panic("assertion")
	}
}

func Sorted[T any](original []T, less func(i, j T) bool) []T {
	// return sorted copy instead of in-place sort
	copySlice := make([]T, len(original))
	copy(copySlice, original)
	sort.Slice(copySlice, func(i, j int) bool {
		return less(copySlice[i], copySlice[j])
	})
	return copySlice
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

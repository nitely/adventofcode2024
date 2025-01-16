package utils

import (
	"log"
	"strconv"
)

func Atoi(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

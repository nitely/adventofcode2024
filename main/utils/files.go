package utils

import (
	"log"
	"os"
)

func MustReadFile(s string) string {
	data, err := os.ReadFile(s)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

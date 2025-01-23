package utils

import (
	"bufio"
	"os"
)

func MustReadFile(s string) string {
	data, err := os.ReadFile(s)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func ByteLines(s string) [][]byte {
	file, err := os.Open(s)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	result := make([][]byte, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := make([]byte, len(scanner.Bytes()))
		copy(line, scanner.Bytes())
		result = append(result, line)
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	return result
}

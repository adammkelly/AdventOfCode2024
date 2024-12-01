package aoc

import (
	"os"
)

func OpenFile(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	str := string(file)

	return str
}

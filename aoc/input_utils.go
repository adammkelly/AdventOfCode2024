package aoc

import (
	"strconv"
	"strings"
)

func SplitToLines(input string) []string {
	return strings.Split(input, "\n")
}

func SplitElementsOnLine(lines []string) [][]string {
	collection := [][]string{}
	for i := range lines {
		line := SplitLineToElements(lines[i])
		collection = append(collection, line)
	}
	return collection
}

func SplitLineToElements(line string) []string {
	return strings.Fields(line)
}

func StringToInt(elem string) int {
	intElem, err := strconv.Atoi(elem)
	if err != nil {
		panic(err)
	}
	return intElem
}

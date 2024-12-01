package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/adammkelly/AdventOfCode2024/aoc"
)

func get_slices(input string) ([]int, []int) {
	s := strings.Fields(input)
	sliceLeft := []int{}
	sliceRight := []int{}

	for idx, elem := range s {
		intElem, err := strconv.Atoi(elem)
		if err != nil {
			panic(err)
		}
		if idx%2 == 0 {
			sliceLeft = append(sliceLeft, intElem)
		} else {
			sliceRight = append(sliceRight, intElem)
		}
	}
	sort.Ints(sliceLeft)
	sort.Ints(sliceRight)

	return sliceLeft, sliceRight
}

func part_01(input string) int {
	var total_distance int = 0
	sliceLeft, sliceRight := get_slices(input)
	for i := range sliceLeft {
		calc := (sliceLeft[i] - sliceRight[i])
		if calc < 0 {
			calc = (sliceRight[i] - sliceLeft[i])
		}
		total_distance += calc
	}
	fmt.Println("Part 1 answer: ", total_distance)
	return total_distance
}

func part_02(input string) int {
	var total_distance int = 0
	occurances := make(map[int]int)
	sliceLeft, sliceRight := get_slices(input)
	for i := range sliceRight {
		current_val := 0
		val := sliceRight[i]
		_, ok := occurances[val]
		if ok {
			current_val = occurances[val]
		}
		occurances[val] = current_val + 1
	}

	for i := range sliceLeft {
		val := sliceLeft[i]
		rightVal, ok := occurances[val]
		if ok {
			total_distance += val * rightVal
		}
	}
	fmt.Println("Part 2 answer: ", total_distance)
	return total_distance
}

func main() {
	fmt.Println("hello world")
	input := aoc.OpenFile("input.txt")

	part_01(input)
	part_02(input)
}

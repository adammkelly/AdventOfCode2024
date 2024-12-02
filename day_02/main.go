package main

import (
	"fmt"

	"github.com/adammkelly/AdventOfCode2024/aoc"
)

var DIRECTION_INCREASE int = 1
var DIRECTION_DECREASE int = 2
var DIRECTION_NOT_SET int = 3

type conditionalFunc func(int, int, int) bool

// True means safe, False is not safe
func iter_levels(s []string, condFn conditionalFunc) bool {
	direction := DIRECTION_NOT_SET
	for idx, elem := range s {
		next := s[idx+1]
		intElem := aoc.StringToInt(elem)
		intNext := aoc.StringToInt(next)
		if direction == DIRECTION_NOT_SET {
			if intElem > intNext {
				direction = DIRECTION_DECREASE
			} else if intElem < intNext {
				direction = DIRECTION_INCREASE
			} else {
				return false
			}
			break
		}
	}

	for idx, elem := range s {
		if idx+1 == len(s) {
			break
		}
		next := s[idx+1]
		intElem := aoc.StringToInt(elem)
		intNext := aoc.StringToInt(next)

		if condFn(intElem, intNext, direction) {
			return false
		}
	}
	return true
}

func part1_cmp(a int, b int, direction int) bool {

	p := 0
	if direction == DIRECTION_INCREASE {
		p = b - a
	} else if direction == DIRECTION_DECREASE {
		p = a - b
	} else {
		panic("Direction not set!")
	}
	if p > 0 && p <= 3 {
		return false
	}
	return true
}

func part_01(input string) int {
	var safe_levels int = 0
	lines := aoc.SplitToLines(input)
	line_elems := aoc.SplitElementsOnLine(lines)

	for _, line := range line_elems {
		isSafe := iter_levels(line, part1_cmp)
		if isSafe {
			safe_levels++
		}
	}
	fmt.Println("Part 1 answer: ", safe_levels)
	return safe_levels
}

func part_02(input string) int {
	var safe_levels_with_dampener int = 0
	lines := aoc.SplitToLines(input)
	line_elems := aoc.SplitElementsOnLine(lines)

	for _, line := range line_elems {
		perms := [][]string{}
		var safe_levels int = 0

		if iter_levels(line, part1_cmp) {
			safe_levels_with_dampener++
			continue
		}

		for idx := range line {
			destination := []string{}
			for idl, le := range line {
				if idl == idx {
					continue
				}
				destination = append(destination, le)
			}
			perms = append(perms, destination)
		}

		for _, perm := range perms {
			isSafe := iter_levels(perm, part1_cmp)
			if isSafe {
				safe_levels++
				break
			}
		}
		if safe_levels == 1 {
			safe_levels_with_dampener++
		}
	}
	fmt.Println("Part 2 answer: ", safe_levels_with_dampener)
	return safe_levels_with_dampener
}

func main() {
	input := aoc.OpenFile("input.txt")

	part_01(input)
	part_02(input)
}

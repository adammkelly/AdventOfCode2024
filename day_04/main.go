package main

import (
	"fmt"
	"strings"

	"github.com/adammkelly/AdventOfCode2024/aoc"
)

func get_char_map(lines []string) [][]string {
	chars := make([][]string, len(lines))
	for y, line := range lines {
		_chars := strings.Split(line, "")
		chars[y] = make([]string, len(_chars))
		copy(chars[y], _chars)
	}
	return chars
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func LookingForLetter(chars [][]string, cur_point aoc.Point, level int, dir string) int {
	var total int = 0
	word := []string{"X", "M", "A", "S"}
	level = level + 1
	if level == 0 && chars[cur_point.Y][cur_point.X] != word[level] {
		return 0
	} else if level == 3 {
		return 1
	}
	var neighbors map[string]aoc.Point = cur_point.Neighbors()

	for direction, point := range neighbors {
		if point.X < 0 || point.Y < 0 || len(chars)-1 < point.Y || len(chars[point.Y])-1 < point.X {
			// Off the edge
			continue
		}
		if level != 0 && dir != direction {
			// Not a valid path so no point checking
			continue
		}
		element := chars[point.Y][point.X]

		if element == word[level+1] {
			total += LookingForLetter(chars, point, level, direction)
		}
	}

	return total
}

func LookingForX_MAS(chars [][]string, cur_point aoc.Point, level int, dir string) int {
	var total int = 0
	word := []string{"A", "M", "S"}
	valid_directions := []string{aoc.TOP_LEFT, aoc.TOP_RIGHT}
	level = level + 1
	if level == 0 && chars[cur_point.Y][cur_point.X] != word[level] {
		return 0
	}
	var neighbors map[string]aoc.Point = cur_point.Neighbors()

	for direction, point := range neighbors {
		if point.X < 0 || point.Y < 0 || len(chars)-1 < point.Y || len(chars[point.Y])-1 < point.X {
			// Off the edge
			continue
		}
		if level != 0 && dir != direction {
			// Not a valid path so no point checking
			continue
		}
		if !contains(valid_directions, direction) {
			continue
		}

		var inverse aoc.Point
		if direction == aoc.TOP_RIGHT {
			inverse = cur_point.Next_neighbor_by_direction(aoc.BOTTOM_LEFT)
		} else {
			inverse = cur_point.Next_neighbor_by_direction(aoc.BOTTOM_RIGHT)
		}
		if inverse.X < 0 || inverse.Y < 0 || len(chars)-1 < inverse.Y || len(chars[inverse.Y])-1 < inverse.X {
			// Off the edge
			continue
		}
		element := chars[point.Y][point.X]
		op := chars[inverse.Y][inverse.X]

		if element == word[level+1] { // M
			if op == word[level+2] {
				total += 1
			}
		} else if element == word[level+2] { // S
			if op == word[level+1] {
				total += 1
			}
		}
	}
	if total == 2 {
		return 1
	}
	return 0
}

func part_01(input string) int {
	var total int = 0
	lines := aoc.SplitToLines(input)
	chars_arr := get_char_map(lines)

	for lidx, line_arr := range chars_arr {
		for cidx, _ := range line_arr {
			cur := aoc.Point{X: cidx, Y: lidx}
			total += LookingForLetter(chars_arr, cur, -1, "")
		}
	}
	fmt.Println("Part 1 answer: ", total)
	return total
}

func part_02(input string) int {
	var total int = 0
	lines := aoc.SplitToLines(input)
	chars_arr := get_char_map(lines)

	for lidx, line_arr := range chars_arr {
		for cidx, _ := range line_arr {
			cur := aoc.Point{X: cidx, Y: lidx}
			total += LookingForX_MAS(chars_arr, cur, -1, "")
		}
	}
	fmt.Println("Part 2 answer: ", total)
	return total
}

func main() {
	input := aoc.OpenFile("input.txt")

	part_01(input)
	part_02(input)
}

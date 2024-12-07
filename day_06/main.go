package main

import (
	"fmt"
	"strings"
	"time"

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

func get_vistor_map(lines []string) [][][]int {
	_lines := make([][][]int, len(lines))
	for y, line := range lines {
		_chars := strings.Split(line, "")
		_lines[y] = make([][]int, len(_chars))
		for x, _ := range _chars {
			_lines[y][x] = make([]int, 4)
		}
	}
	return _lines
}

var DIRECTIONS = map[string]string{
	"^": aoc.TOP,
	">": aoc.RIGHT,
	"<": aoc.TOP,
	"v": aoc.LEFT,
}

var ROTATIONS = map[string]string{
	aoc.TOP:    aoc.RIGHT,
	aoc.RIGHT:  aoc.BOTTOM,
	aoc.LEFT:   aoc.TOP,
	aoc.BOTTOM: aoc.LEFT,
}

var ROTATIONS_ACC = map[string]int{
	aoc.TOP:    0,
	aoc.RIGHT:  1,
	aoc.LEFT:   2,
	aoc.BOTTOM: 3,
}

func outOfBounds(chars [][]string, point aoc.Point) bool {
	if point.X < 0 || point.Y < 0 || len(chars)-1 < point.Y || len(chars[point.Y])-1 < point.X {
		return true
	}
	return false
}

func LookingForExit(chars [][]string, cur_point aoc.Point, direction string, vistor_map [][]string) (aoc.Point, int) {
	var steps int = 0
	var final_p aoc.Point
	var neighbors map[string]aoc.Point = cur_point.Neighbors()
	var rotatation string = direction
	_, guard := locateGuard(chars)

	next_element := cur_point.Next_neighbor_by_direction(rotatation)
	if outOfBounds(chars, next_element) {
		final_p = cur_point
		return final_p, steps
	}

	var peek string
	for {
		peek = chars[next_element.Y][next_element.X]
		if peek == "#" {
			// fmt.Println("ROTATE", rotatation, "->", ROTATIONS[direction])
			rotatation = ROTATIONS[rotatation]
			next_element = cur_point.Next_neighbor_by_direction(rotatation)
		} else {
			break
		}
	}

	for dir, point := range neighbors {
		if dir != rotatation {
			continue
		}
		if outOfBounds(chars, point) {
			final_p = point
			return final_p, steps
		}

		seen_before := vistor_map[point.Y][point.X]
		if seen_before != "X" && point != guard {
			vistor_map[point.Y][point.X] = "X"
			steps++
		}

		_p, _s := LookingForExit(chars, point, rotatation, vistor_map)
		steps += _s
		final_p = _p
		return final_p, steps
	}

	return final_p, steps
}

func LookingForStuck(chars [][]string, cur_point aoc.Point, direction string, vistor_map [][][]int) bool {
	var neighbors map[string]aoc.Point = cur_point.Neighbors()
	var rotatation string = direction

	next_element := cur_point.Next_neighbor_by_direction(rotatation)
	if outOfBounds(chars, next_element) {
		return false
	}

	var peek string
	for {
		peek = chars[next_element.Y][next_element.X]
		if peek == "#" {
			// fmt.Println("ROTATE", rotatation, "->", ROTATIONS[direction])
			rotatation = ROTATIONS[rotatation]
			vistor_map[next_element.Y][next_element.X][ROTATIONS_ACC[rotatation]]++
			next_element = cur_point.Next_neighbor_by_direction(rotatation)
		} else {
			break
		}
	}

	for dir, point := range neighbors {
		if dir != rotatation {
			continue
		}
		if outOfBounds(chars, point) {
			return false
		}

		vistor_map[point.Y][point.X][ROTATIONS_ACC[rotatation]]++
		seen_before := vistor_map[point.Y][point.X][ROTATIONS_ACC[rotatation]]
		if seen_before > 100 {
			return true
		}

		return LookingForStuck(chars, point, rotatation, vistor_map)
	}

	return false
}

func locateGuard(chars_arr [][]string) (string, aoc.Point) {
	for lidx, line_arr := range chars_arr {
		for cidx, char := range line_arr {
			if char != "^" && char != "<" && char != ">" && char != "v" {
				continue
			}

			cur := aoc.Point{X: cidx, Y: lidx}
			return DIRECTIONS[char], cur
		}
	}
	panic("Shouldn't get here")
}

func part_01(input string) int {
	start := time.Now()
	var total int = 0
	lines := aoc.SplitToLines(input)
	chars_arr := get_char_map(lines)
	vistor_map := get_char_map(lines)

	direction, pos := locateGuard(chars_arr)
	_, steps := LookingForExit(chars_arr, pos, direction, vistor_map)
	total = steps
	duration := time.Since(start)
	fmt.Println("Part 1 answer: ", total, "- Time Taken:", duration.Seconds())
	return total
}

func part_02(input string) int {
	start := time.Now()
	var total int = 0
	lines := aoc.SplitToLines(input)
	chars_arr := get_char_map(lines)

	direction, pos := locateGuard(chars_arr)
	visit_map := get_char_map(lines)
	_, _ = LookingForExit(chars_arr, pos, direction, visit_map)

	for lidx, line_arr := range chars_arr {

		for cidx, char := range line_arr {
			vistor_map := get_vistor_map(lines)
			_chars_arr := get_char_map(lines)

			if char != "." {
				continue
			}
			if lidx == pos.Y && cidx == pos.X {
				// Cannot place on guard
				continue
			}

			if visit_map[lidx][cidx] != "X" {
				continue
			}
			_chars_arr[lidx][cidx] = "#"

			if LookingForStuck(_chars_arr, pos, direction, vistor_map) {
				total++
			}
		}
	}

	duration := time.Since(start)
	fmt.Println("Part 2 answer: ", total, "- Time Taken:", duration.Seconds())
	return total
}

func main() {
	input := aoc.OpenFile("input.txt")

	part_01(input)
	part_02(input)
}

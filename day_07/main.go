package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/adammkelly/AdventOfCode2024/aoc"
)

var OP_MULT string = "MULTIPLY"
var OP_ADD string = "ADD"
var OP_CON string = "CON"

func calc_op(a int, b int, op string) int {
	var calc int
	if op == OP_ADD {
		calc = a + b
	} else if op == OP_CON {
		s := aoc.IntToString(a) + aoc.IntToString(b)
		calc = aoc.StringToInt(s)
	} else {
		calc = a * b
	}
	return calc
}

func contains_expected(expected int, lines []int) bool {
	for _, line := range lines {
		if line == expected {
			return true
		}
	}
	return false
}

func get_calc_list(lines []string) []map[int][]int {
	map_list := make([]map[int][]int, 0)
	for _, line := range lines {
		s := strings.Split(line, ":")
		expected := aoc.StringToInt(s[0])
		new_map := make(map[int][]int, len(strings.Fields(s[1])))
		for _, sc := range strings.Fields(s[1]) {
			new_map[expected] = append(new_map[expected], aoc.StringToInt(sc))
		}
		map_list = append(map_list, new_map)
	}
	return map_list
}

func walk_arr(calc_list []int, ops []string, prev_values []int, idx int) []int {
	prev := make([]int, len(prev_values)*len(ops))

	if len(calc_list) == idx {
		return prev_values
	}

	next_val := calc_list[idx]
	loop_id := 0

	for _, _pv := range prev_values {
		for _, o := range ops {
			prev[loop_id] = calc_op(_pv, next_val, o)
			loop_id++
		}
	}

	return walk_arr(calc_list, ops, prev, idx+1)
}

func perform_ops(calc_list []int, ops []string) []int {
	prev_values := make([]int, len(ops))

	for idx, o := range ops {
		prev_values[idx] = calc_op(calc_list[0], calc_list[1], o)
	}

	if len(calc_list) == 2 {
		return prev_values
	}

	return walk_arr(calc_list, ops, prev_values, 2)
}

func part_01(input string) int {
	start := time.Now()
	var total int = 0
	lines := aoc.SplitToLines(input)
	calcs := get_calc_list(lines)
	ops := []string{OP_MULT, OP_ADD}

	for _, nl := range calcs {
		for expected, to_be_cal := range nl {
			ret_list := perform_ops(to_be_cal, ops)
			// fmt.Println(expected, "---", ret_list)

			if contains_expected(expected, ret_list) {
				total += expected
			}

		}
	}

	duration := time.Since(start)
	fmt.Println("Part 1 answer: ", total, "- Time Taken:", duration.Seconds())
	return total
}

func part_02(input string) int {
	start := time.Now()
	var total int = 0
	lines := aoc.SplitToLines(input)
	calcs := get_calc_list(lines)
	ops := []string{OP_MULT, OP_ADD, OP_CON}

	for _, nl := range calcs {
		for expected, to_be_cal := range nl {
			ret_list := perform_ops(to_be_cal, ops)
			// fmt.Println(expected, "---", ret_list)

			if contains_expected(expected, ret_list) {
				total += expected
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

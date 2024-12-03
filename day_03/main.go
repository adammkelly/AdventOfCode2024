package main

import (
	"fmt"

	"github.com/adammkelly/AdventOfCode2024/aoc"
)

var FIND_MUL_REGEX string = `mul\(\d{1,3}[[:punct:]]\d{1,3}\)`
var mul_r = aoc.CompileRegex(FIND_MUL_REGEX)

var FIND_NUMBERS_REGEX string = `(\d{1,3})`
var mul_r_num = aoc.CompileRegex(FIND_NUMBERS_REGEX)

var COLLECT_PART2_REGEX string = `(mul\(\d{1,3}[[:punct:]]\d{1,3}\))|(don't\(\))|(do\(\))`
var part_2_collection = aoc.CompileRegex(COLLECT_PART2_REGEX)

func find_mul(input string) []string {
	return mul_r.FindAllString(input, -1)
}

func find_numbers_in_mul(input string) []string {
	return mul_r_num.FindAllString(input, -1)
}

func find_mul_part2(input string) []string {
	return part_2_collection.FindAllString(input, -1)
}

func part_01(input string) int {
	var total int = 0
	all_muls := find_mul(input)

	for _, mul := range all_muls {
		items := find_numbers_in_mul(mul)
		total += aoc.StringToInt(items[0]) * aoc.StringToInt(items[1])
	}
	fmt.Println("Part 1 answer: ", total)
	return total
}

func part_02(input string) int {
	var total int = 0
	var enabled bool = true
	all_muls := find_mul_part2(input)

	for _, mul := range all_muls {
		if mul == "do()" {
			enabled = true
			continue
		} else if mul == "don't()" {
			enabled = false
			continue
		} else {
			// Not do or dont, process
			if !enabled {
				continue
			}

		}
		items := find_numbers_in_mul(mul)
		total += aoc.StringToInt(items[0]) * aoc.StringToInt(items[1])
	}
	fmt.Println("Part 2 answer: ", total)
	return total
}

func main() {
	input := aoc.OpenFile("input.txt")

	part_01(input)
	part_02(input)
}

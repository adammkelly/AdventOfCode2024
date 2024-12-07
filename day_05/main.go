package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/adammkelly/AdventOfCode2024/aoc"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func getOrdersAndRules(lines []string) ([]string, []string) {
	var rules, page_orders []string
	for _, line := range lines {
		if strings.Contains(line, "|") {
			rules = append(rules, line)
		} else {
			page_orders = append(page_orders, line)
		}
	}
	return rules, page_orders
}

func part_01(input string) int {
	var total int = 0
	lines := strings.Fields(input)
	rules, page_orders := getOrdersAndRules(lines)

	var invalid bool = false
	var next_page string
	for _, update := range page_orders {
		pages := strings.Split(update, ",")
		for i, page := range pages {
			invalid = false
			if i != len(pages)-1 {
				next_page = pages[i+1]
			}

			cmp := fmt.Sprintf("%s|%s", next_page, page)
			for _, rule := range rules {
				if cmp == rule {
					invalid = true
					break
				}
			}

			if invalid {
				break
			}
		}

		if !invalid {
			total += aoc.StringToInt(pages[(len(pages)-1)/2])
		}

	}
	fmt.Println("Part 1 answer: ", total)
	return total
}

func part_02(input string) int {
	var total int = 0
	lines := strings.Fields(input)
	rules, page_orders := getOrdersAndRules(lines)

	var invalid bool = false
	var next_page string
	var page_orders_invalid [][]string
	for _, update := range page_orders {
		pages := strings.Split(update, ",")
		for i, page := range pages {
			invalid = false
			if i != len(pages)-1 {
				next_page = pages[i+1]
			}

			cmp := fmt.Sprintf("%s|%s", next_page, page)
			for _, rule := range rules {
				if cmp == rule {
					invalid = true
					break
				}
			}

			if invalid {
				break
			}
		}
		if invalid {
			s := strings.Split(update, ",")
			page_orders_invalid = append(page_orders_invalid, s)
		}
	}

	for _, ivp := range page_orders_invalid {
		sort.SliceStable(ivp, func(i, j int) bool {
			page := ivp[i]
			next_page := ivp[j]
			cmp := fmt.Sprintf("%s|%s", next_page, page)

			return !stringInSlice(cmp, rules)
		})
	}
	for _, page := range page_orders_invalid {
		total += aoc.StringToInt(page[(len(page))/2])
	}

	fmt.Println("Part 2 answer: ", total)
	return total
}

func main() {
	input := aoc.OpenFile("input.txt")

	part_01(input)
	part_02(input)
}

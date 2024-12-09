package main

import (
	"fmt"
	"strings"

	"github.com/adammkelly/AdventOfCode2024/aoc"
)

type Frequency struct {
	freq             []rune
	known_nodes      []aoc.Point
	known_anodes     []aoc.Point
	known_anodes_oob []aoc.Point
}

func get_char_map(lines []string) [][]string {
	chars := make([][]string, len(lines))
	for y, line := range lines {
		line = strings.Trim(line, "\n\r")
		_chars := strings.Split(line, "")
		chars[y] = make([]string, len(_chars))
		copy(chars[y], _chars)
	}
	return chars
}

func outOfBounds(chars [][]string, point aoc.Point) bool {
	if point.X < 0 || point.Y < 0 || len(chars)-1 < point.Y || len(chars[point.Y])-1 < point.X {
		return true
	}
	return false
}

func lookup_freq(list []*Frequency, id []rune) *Frequency {
	for _, f := range list {
		if string(f.freq) == string(id) {
			return f
		}
	}
	return nil
}

func points_match(p aoc.Point, p2 aoc.Point) bool {
	if p.X == p2.X && p.Y == p2.Y {
		return true
	}
	return false
}

func known_anode(freq_list []*Frequency, freq *Frequency, new_point aoc.Point) bool {
	for _, pi_node := range freq.known_anodes {
		if points_match(pi_node, new_point) {
			return true
		}
	}
	for _, pi_node := range freq.known_anodes_oob {
		if points_match(pi_node, new_point) {
			return true
		}
	}
	for _, _freq := range freq_list {
		if string(_freq.freq) == "." || freq == _freq {
			continue
		}
		for _, pi_node := range _freq.known_anodes {
			if points_match(pi_node, new_point) {
				return true
			}
		}
		for _, pi_node := range _freq.known_anodes_oob {
			if points_match(pi_node, new_point) {
				return true
			}
		}
	}
	return false
}

func get_real_coord(cur_point aoc.Point, offset aoc.Point) aoc.Point {
	X := cur_point.X + offset.X
	Y := cur_point.Y + offset.Y
	return aoc.Point{X, Y}
}

func build_freq_list(chars_arr [][]string) (freq_list []*Frequency) {
	freq_list = make([]*Frequency, 0)

	for lidx, line_arr := range chars_arr {
		for cidx, c := range line_arr {
			runeArray := []rune(c)
			cur := aoc.Point{X: cidx, Y: lidx}
			freq := lookup_freq(freq_list, runeArray)
			if freq == nil {
				freq = new(Frequency)
				freq.freq = runeArray
				freq.known_nodes = make([]aoc.Point, 0)
				freq.known_anodes = make([]aoc.Point, 0)
				freq.known_anodes_oob = make([]aoc.Point, 0)
				freq_list = append(freq_list, freq)
			}
			freq.known_nodes = append(freq.known_nodes, cur)
		}
	}
	return freq_list
}

func build_anodes(chars_arr [][]string, freq_list []*Frequency) {
	for _, freq := range freq_list {
		if string(freq.freq) == "." {
			continue
		}
		for _, pi_node := range freq.known_nodes {
			for _, pi_view := range freq.known_nodes {
				if pi_node == pi_view {
					continue
				}
				anode1Y := pi_node.Y - pi_view.Y
				anode1X := pi_node.X - pi_view.X
				anode2Y := pi_view.Y - pi_node.Y
				anode2X := pi_view.X - pi_node.X

				anodep1 := aoc.Point{X: anode1X, Y: anode1Y}
				anodep2 := aoc.Point{X: anode2X, Y: anode2Y}
				// anode2 == pivew
				//anode1 == pinode
				anodep1_coords := get_real_coord(pi_node, anodep1)
				anodep2_coords := get_real_coord(pi_view, anodep2)
				if !known_anode(freq_list, freq, anodep1_coords) {
					if outOfBounds(chars_arr, anodep1_coords) {
						freq.known_anodes_oob = append(freq.known_anodes_oob, anodep1_coords)
					} else {
						freq.known_anodes = append(freq.known_anodes, anodep1_coords)
					}
				}
				if !known_anode(freq_list, freq, anodep2_coords) {
					if outOfBounds(chars_arr, anodep2_coords) {
						freq.known_anodes_oob = append(freq.known_anodes_oob, anodep2_coords)
					} else {
						freq.known_anodes = append(freq.known_anodes, anodep2_coords)
					}
				}
			}
		}
	}
}

func build_anodes2(chars_arr [][]string, freq_list []*Frequency) {
	for _, freq := range freq_list {
		if string(freq.freq) == "." {
			continue
		}
		for _, pi_node := range freq.known_nodes {
			for _, pi_view := range freq.known_nodes {
				if pi_node == pi_view {
					continue
				}
				anode1Y := pi_node.Y - pi_view.Y
				anode1X := pi_node.X - pi_view.X
				anode2Y := pi_view.Y - pi_node.Y
				anode2X := pi_view.X - pi_node.X

				anodep1 := aoc.Point{X: anode1X, Y: anode1Y}
				anodep1_coords := get_real_coord(pi_node, anodep1)
				for !outOfBounds(chars_arr, anodep1_coords) {
					// anode2 == pivew
					//anode1 == pinode
					if !known_anode(freq_list, freq, anodep1_coords) {
						if outOfBounds(chars_arr, anodep1_coords) {
							freq.known_anodes_oob = append(freq.known_anodes_oob, anodep1_coords)
						} else {
							freq.known_anodes = append(freq.known_anodes, anodep1_coords)
						}
					}
					anodep1_coords = get_real_coord(anodep1_coords, anodep1)
				}

				anodep2 := aoc.Point{X: anode2X, Y: anode2Y}
				anodep2_coords := get_real_coord(pi_view, anodep2)
				for !outOfBounds(chars_arr, anodep2_coords) {
					if !known_anode(freq_list, freq, anodep2_coords) {
						if outOfBounds(chars_arr, anodep2_coords) {
							freq.known_anodes_oob = append(freq.known_anodes_oob, anodep2_coords)
						} else {
							freq.known_anodes = append(freq.known_anodes, anodep2_coords)
						}
					}
					anodep2_coords = get_real_coord(anodep2_coords, anodep2)
				}
				if !known_anode(freq_list, freq, pi_view) {
					freq.known_anodes = append(freq.known_anodes, pi_view)
				}
			}
			if !known_anode(freq_list, freq, pi_node) {
				freq.known_anodes = append(freq.known_anodes, pi_node)
			}
		}
	}
}

func part_01(input string) int {
	var total int = 0
	lines := aoc.SplitToLines(input)
	chars_arr := get_char_map(lines)
	freq_list := build_freq_list(chars_arr)
	build_anodes(chars_arr, freq_list)

	for _, freq := range freq_list {
		if string(freq.freq) == "." {
			continue
		}
		total += len(freq.known_anodes)
	}
	fmt.Println("Part 1 answer: ", total)
	return total
}

func part_02(input string) int {
	var total int = 0
	lines := aoc.SplitToLines(input)
	chars_arr := get_char_map(lines)
	freq_list := build_freq_list(chars_arr)
	build_anodes2(chars_arr, freq_list)

	for _, freq := range freq_list {
		if string(freq.freq) == "." {
			continue
		}
		total += len(freq.known_anodes)
	}
	fmt.Println("Part 2 answer: ", total)
	return total
}

func main() {
	input := aoc.OpenFile("input.txt")

	part_01(input)
	part_02(input)
}

package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/adammkelly/AdventOfCode2024/aoc"
)

type Block struct {
	id    int
	value string
}

func is_free_space(id int) bool {
	return (id%2 == 1)
}

func is_block(id int) bool {
	return (id%2 == 0)
}

func get_disk_map(lines string) []string {
	chars := strings.Split(lines, "")
	return chars
}

func generate_block(blocks []string) []*Block {
	blk_list := make([]*Block, 0)
	cur_id := 0
	for pos, block_num := range blocks {
		block_num_int := aoc.StringToInt(block_num)
		if is_free_space(pos) {
			for range block_num_int {

				blk := new(Block)
				blk.value = "."
				blk.id = pos
				blk_list = append(blk_list, blk)
			}

		} else if is_block(pos) {
			for range block_num_int {
				blk := new(Block)
				blk.value = aoc.IntToString(cur_id)
				blk.id = pos
				blk_list = append(blk_list, blk)
			}
			cur_id++
		}
	}
	return blk_list
}

func reverse_slice[T comparable](s []T) {
	sort.SliceStable(s, func(i, j int) bool {
		return i > j
	})
}

func defrag_blk(blks []*Block) []*Block {
	last_seen := 0
	dots_found := 0
	for blk_pos, blk := range blks {
		if blk.value == "." {
			dots_found++
			for rblk_pos := len(blks) - 1; rblk_pos > 0; rblk_pos-- {
				// fmt.Println(blks[rblk_pos].value, "->>", rblk_pos)
				rblk := blks[rblk_pos]
				last_seen = rblk_pos
				if rblk.value != "." {
					break
				}
			}
			// fmt.Println(blks[last_seen].value, "->", blks[blk_pos].value, blk_pos, last_seen)
			if last_seen+1 == blk_pos {
				break
			}
			blks[blk_pos] = blks[last_seen]
			blks[last_seen] = blk
		}
	}
	return blks
}

func generate_checksum(blks []*Block) int {
	var accum int = 0
	for px, blk := range blks {
		if is_free_space(blk.id) {
			continue
		}
		accum += px * aoc.StringToInt(blk.value)
	}
	return accum
}

func part_01(input string) int {
	var total int = 0
	chars_arr := get_disk_map(input)
	block_pat := generate_block(chars_arr)
	defrag := defrag_blk(block_pat)
	total = generate_checksum(defrag)
	fmt.Println("Part 1 answer: ", total)
	return total
}

func part_02(input string) int {
	var total int = 0
	fmt.Println("Part 2 answer: ", total)
	return total
}

func main() {
	input := aoc.OpenFile("input.txt")

	part_01(input)
	part_02(input)
}

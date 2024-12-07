package main

import (
	"testing"
)

var example1 string = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func TestPart1(t *testing.T) {
	want := 41
	res := part_01(example1)

	if res != want {
		t.Fatalf("part_01 expected %d got: %d", want, res)
	}
}

func TestPart2(t *testing.T) {
	want := 6
	res := part_02(example1)

	if res != want {
		t.Fatalf("part_02 expected %d got: %d", want, res)
	}
}

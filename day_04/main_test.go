package main

import (
	"testing"
)

var example1 string = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func TestPart1(t *testing.T) {
	want := 18
	res := part_01(example1)

	if res != want {
		t.Fatalf("part_01 expected %d got: %d", want, res)
	}
}

var example2 string = `.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........`

func TestPart2(t *testing.T) {
	want := 9
	res := part_02(example2)

	if res != want {
		t.Fatalf("part_02 expected %d got: %d", want, res)
	}
}

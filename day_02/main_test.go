package main

import (
	"testing"
)

var example1 string = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestPart1(t *testing.T) {
	want := 2
	res := part_01(example1)

	if res != want {
		t.Fatalf("part_01 expected %d got: %d", want, res)
	}
}

func TestPart2(t *testing.T) {
	want := 4
	res := part_02(example1)

	if res != want {
		t.Fatalf("part_02 expected %d got: %d", want, res)
	}
}

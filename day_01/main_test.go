package main

import (
	"testing"
)

var example1 string = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestPart1(t *testing.T) {
	want := 11
	res := part_01(example1)

	if res != want {
		t.Fatalf("part_01 expected %d got: %d", want, res)
	}
}

var example2 string = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestPart2(t *testing.T) {
	want := 31
	res := part_02(example2)

	if res != want {
		t.Fatalf("part_02 expected %d got: %d", want, res)
	}
}

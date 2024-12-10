package main

import (
	"testing"
)

var example1 string = `2333133121414131402`

func TestPart1(t *testing.T) {
	want := 1928
	res := part_01(example1)

	if res != want {
		t.Fatalf("part_01 expected %d got: %d", want, res)
	}
}

func TestPart2(t *testing.T) {
	want := 1928
	res := part_02(example1)

	if res != want {
		t.Fatalf("part_02 expected %d got: %d", want, res)
	}
}

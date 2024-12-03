package main

import (
	"testing"
)

var example1 string = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
var example2 string = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

func TestPart1(t *testing.T) {
	want := 161
	res := part_01(example1)

	if res != want {
		t.Fatalf("part_01 expected %d got: %d", want, res)
	}
}

func TestPart2(t *testing.T) {
	want := 48
	res := part_02(example2)

	if res != want {
		t.Fatalf("part_02 expected %d got: %d", want, res)
	}
}

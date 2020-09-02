package main

import (
	"goleetcode/Others"
	"testing"
)

func TestFindMagicIndex(t *testing.T) {
	tests := []struct {
		nums     []int
		excepted int
	}{
		{[]int{1, 2, 3, 4}, -1},
		{[]int{0, 2, 3, 4}, 0},
		{[]int{1, 2, 3, 4}, -1},
		{[]int{1, 2, 3, 4}, -1},
		{[]int{1, 2, 3, 4}, -1},
	}

	for _, test := range tests {
		if actual := Others.FindMagicIndex(test.nums); actual != test.excepted {
			t.Error("error")
		}
	}
}

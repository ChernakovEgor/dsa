package main

import (
	"cmp"
	"fmt"
	"testing"
)

func binary_search[T cmp.Ordered](haystack []T, needle T) int {
	lo := 0
	hi := len(haystack)
	if hi == 0 {
		return -1
	}
	for {
		m := (hi + lo) / 2
		if haystack[m] == needle {
			return m
		} else if haystack[m] > needle {
			hi = m
		} else {
			lo = m + 1
		}

		if lo >= hi {
			break
		}
	}
	return -1
}

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		haystack []int
		needle   int
		want     int
	}{
		{[]int{}, 10, -1},
		{[]int{10}, 10, 0},
		{[]int{1, 5}, 1, 0},
		{[]int{1, 5}, 5, 1},
		{[]int{1, 2, 3, 5, 8, 10, 12}, 10, 5},
		{[]int{1, 2, 3, 5, 8, 10, 12}, 1, 0},
		{[]int{1, 2, 3, 5, 8, 10, 12}, 11, -1},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.haystack), func(t *testing.T) {
			got := binary_search(test.haystack, test.needle)
			if got != test.want {
				t.Errorf("got %v want %v", got, test.want)
			}
		})
	}
}

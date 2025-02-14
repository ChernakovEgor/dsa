package binarysearch

import (
	"cmp"
	"fmt"
	"reflect"
	"testing"
)

func bubbleSort[T cmp.Ordered](arr []T) []T {
	for i := range arr {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{5, 2, 3, 7, 8, 1, 4, 6}, []int{1, 2, 3, 4, 5, 6, 7, 8}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.arr), func(t *testing.T) {
			got := bubbleSort(test.arr)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v want %v", got, test.want)
			}
		})
	}
}

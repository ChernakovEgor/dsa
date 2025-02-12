package binarysearch

import "math"
import "fmt"
import "testing"

func twoCrystalBalls(breaks []bool) int {
	step := int(math.Sqrt(float64(len(breaks))))
	i := step
	for ; i < len(breaks); i += step {
		if breaks[i] {
			break
		}
	}

	i -= step
	for j := i; j <= i+step && j < len(breaks); j++ {
		if breaks[j] {
			return j
		}
	}

	return -1
}

func TestTwoBalls(t *testing.T) {
	tests := []struct {
		breaks []bool
		want   int
	}{
		{[]bool{}, -1},
		{[]bool{false, false, false, true, true, true}, 3},
		{[]bool{true, true, true, true, true, true}, 0},
		{[]bool{false, false, false, false, true}, 4},
		{[]bool{false, false, false, false, false}, -1},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := twoCrystalBalls(test.breaks)
			if got != test.want {
				t.Errorf("got %v want %v", got, test.want)
			}
		})
	}
}

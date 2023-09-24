package lc

import (
	"testing"
)

func largestAltitude(gain []int) int {
	var currentAltitude, maxAltitude int
	for _, g := range gain {
		currentAltitude += g
		maxAltitude = max(currentAltitude, maxAltitude)
	}
	return maxAltitude
}

func TestLargestAltitude(t *testing.T) {
	tests := []struct {
		gain []int
		want int
	}{
		{[]int{-5, 1, 5, 0, -7}, 1},
		{[]int{-4, -3, -2, -1, 4, 3, 2}, 0},
	}

	for _, tt := range tests {
		got := largestAltitude(tt.gain)
		if got != tt.want {
			t.Errorf("largestAltitude(%v) = %v; want %v", tt.gain, got, tt.want)
		}
	}
}

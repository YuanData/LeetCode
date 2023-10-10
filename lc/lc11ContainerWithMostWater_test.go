package lc

import (
	"testing"
)

func maxArea(height []int) int {
	maxArea := 0
	left, right := 0, len(height)-1
	for left < right {
		currentArea := (right - left) * min(height[left], height[right])
		maxArea = max(currentArea, maxArea)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func TestMaxArea(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
	}

	for _, test := range tests {
		got := maxArea(test.height)
		if got != test.want {
			t.Errorf("maxArea(%v) = %v; want %v", test.height, got, test.want)
		}
	}
}

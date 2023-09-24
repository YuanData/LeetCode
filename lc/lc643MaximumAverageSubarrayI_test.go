package lc

import (
	"testing"
)

func findMaxAverage(nums []int, k int) float64 {
	var currSum int
	maxSum := float64(-1 << 63) // equivalent to -infinity in float64
	for i, num := range nums {
		currSum += num // adding the current number to the sum
		if i >= k {    // when i is equal to or greater than k, subtract the leftmost number in the window from the sum
			currSum -= nums[i-k]
		}
		if i >= k-1 { // if the window size is k, update maxSum
			maxSum = max(maxSum, float64(currSum))
		}
	}
	return maxSum / float64(k)
}

func TestFindMaxAverage(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want float64
	}{
		{[]int{1, 12, -5, -6, 50, 3}, 4, 12.75},
		{[]int{5}, 1, 5.0},
		{[]int{-1}, 1, -1.0},
	}

	for _, tt := range tests {
		got := findMaxAverage(tt.nums, tt.k)
		if got != tt.want {
			t.Errorf("findMaxAverage(%v, %v) = %v; want %v", tt.nums, tt.k, got, tt.want)
		}
	}
}

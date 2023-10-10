package lc

import (
	"sort"
	"testing"
)

func maxOperationsSort(nums []int, k int) int {
	sort.Ints(nums)               // 先排序
	left, right := 0, len(nums)-1 // 使用左右指針
	operations := 0

	for left < right {
		currentSum := nums[left] + nums[right]

		// 如果當前和等於k，則找到一對，左右指針都要移動
		if currentSum == k {
			operations++
			left++
			right--
		} else if currentSum < k { // 如果當前和小於k，則左指針向右移動，使得和增加
			left++
		} else { // 如果當前和大於k，則右指針向左移動，使得和減少
			right--
		}
	}

	return operations
}

func TestMaxOperationsSort(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 2, 3, 4}, 5, 2},
		{[]int{3, 1, 3, 4, 3}, 6, 1},
		{[]int{2, 5, 4, 4, 1, 3, 4, 4, 1, 4, 4, 1, 2, 1, 2, 2, 3, 2, 4, 2}, 3, 4},
		// 更多的測試案例...
	}

	for _, test := range tests {
		got := maxOperationsSort(test.nums, test.k)
		if got != test.want {
			t.Errorf("maxOperations(%v, %v) = %v; want %v", test.nums, test.k, got, test.want)
		}
	}
}

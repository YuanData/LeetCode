package lc

import (
	"fmt"
	"reflect"
	"testing"
)

func productExceptSelf(nums []int) []int {
	n := len(nums)      // 設定數組的長度變數
	left, right := 1, 1 // 設定左側、右側乘積的變數，初始都是1

	result := make([]int, n) // 初始化結果數組
	for i := range result {
		result[i] = 1
	}

	// 同時從左和從右遍歷數組
	for i := 0; i < n; i++ {
		// 先將當前索引左側的乘積與結果數組中的元素相乘
		result[i] *= left
		left *= nums[i] // 更新左側乘積

		// 先將當前索引右側的乘積與結果數組中的元素相乘
		result[n-i-1] *= right
		right *= nums[n-i-1] // 更新右側乘積
	}

	return result // 返回結果數組
}

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}

	for _, test := range tests {
		result := productExceptSelf(test.nums)
		if !reflect.DeepEqual(result, test.expected) {
			fmt.Printf("For %v, expected %v but got %v\n", test.nums, test.expected, result)
		}
	}
}

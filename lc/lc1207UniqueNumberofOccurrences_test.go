package lc

import (
	"testing"
)

func uniqueOccurrences(arr []int) bool {
	// 创建一个 map 来存储每个元素出现的次数
	countMap := make(map[int]int)
	for _, num := range arr {
		countMap[num]++
	}

	// 使用另一个 map 来检查各个出现次数的唯一性
	uniqueCountMap := make(map[int]bool)
	for _, count := range countMap {
		if uniqueCountMap[count] {
			return false // 如果已经存在，说明出现次数不唯一，返回 false
		}
		uniqueCountMap[count] = true
	}

	// 所有出现次数都是唯一的
	return true
}

func TestUniqueOccurrences(t *testing.T) {
	testCases := []struct {
		arr    []int
		expect bool
	}{
		{
			arr:    []int{1, 2, 2, 1, 1, 3},
			expect: true,
		},
		{
			arr:    []int{1, 2},
			expect: false,
		},
		{
			arr:    []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0},
			expect: true,
		},
		{
			arr:    []int{},
			expect: true,
		},
		{
			arr:    []int{5},
			expect: true,
		},
		{
			arr:    []int{1, 1, 1, 2, 2, 2, 3, 3, 3},
			expect: false,
		},
		{
			arr:    []int{-1, -1, -1, -2, -2, -2},
			expect: false,
		},
	}

	for _, testCase := range testCases {
		result := uniqueOccurrences(testCase.arr)
		if result != testCase.expect {
			t.Errorf("Input: %v, Expected: %v, Got: %v", testCase.arr, testCase.expect, result)
		}
	}
}

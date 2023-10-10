package lc

import (
	"testing"
)

func maxOperationsMap(nums []int, k int) int {

	// 初始化答案為0
	ans := 0
	n := len(nums)

	// 如果數組的大小小於或等於1，則直接返回答案，因為沒有對可以形成
	if n <= 1 {
		return ans
	}

	// 創建一個hashmap來存儲數組中每個數字的出現次數
	hashmap := make(map[int]int)

	for _, num := range nums {

		// 如果當前數字大於或等於k，則我們可以跳過它，因為它不能與其他數字形成和為k的對
		if num >= k {
			continue
		}

		// 計算當前數字的complement，也就是k減去當前數字
		counterpart := k - num

		// 檢查hashmap中是否存在這個complement
		if val, exists := hashmap[counterpart]; exists && val > 0 {
			ans++ // 如果存在，則我們找到一個對，答案增加1

			// 更新hashmap中complement的計數
			if val-1 == 0 {
				delete(hashmap, counterpart) // 如果計數為0，則從hashmap中刪除該complement
			} else {
				hashmap[counterpart]--
			}

		} else { // 如果complement不存在，則更新當前數字的計數
			hashmap[num]++
		}
	}

	// 返回答案
	return ans
}

func TestMaxOperationsMap(t *testing.T) {
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
		got := maxOperationsMap(test.nums, test.k)
		if got != test.want {
			t.Errorf("maxOperations(%v, %v) = %v; want %v", test.nums, test.k, got, test.want)
		}
	}
}

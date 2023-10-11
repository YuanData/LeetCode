package lc

import (
	"reflect"
	"testing"
)

// asteroidCollision 處理小行星碰撞的情況並返回碰撞後的小行星狀態
func asteroidCollision(asteroids []int) []int {
	// stack 用於儲存目前還未發生碰撞的小行星
	var stack []int

	// 遍歷每個小行星
	for _, a := range asteroids {
		// 若堆疊為空或小行星向右移動（代表其不會與堆疊內的小行星碰撞）
		if len(stack) == 0 || a > 0 {
			stack = append(stack, a)
		} else {
			// 若當前小行星向左移動，則需要檢查其是否會與堆疊頂部的小行星碰撞
			for len(stack) > 0 && stack[len(stack)-1] > 0 && absIntG(stack[len(stack)-1]) < absIntG(a) {
				// 若堆疊頂部的小行星會被破壞，則移除它
				stack = stack[:len(stack)-1]
			}

			// 若堆疊為空或堆疊頂部的小行星向左移動（代表其不會與當前小行星碰撞）
			if len(stack) == 0 || stack[len(stack)-1] < 0 {
				stack = append(stack, a)
			} else if stack[len(stack)-1] == -a { // 若堆疊頂部的小行星與當前小行星大小相同，兩者都會被破壞
				stack = stack[:len(stack)-1]
			}
		}
	}

	// 返回碰撞後的小行星狀態
	return stack
}

func TestAsteroidCollision(t *testing.T) {
	tests := []struct {
		input  []int
		expect []int
	}{
		{input: []int{5, 10, -5}, expect: []int{5, 10}},
		{input: []int{8, -8}, expect: []int{}},
		{input: []int{10, 2, -5}, expect: []int{10}},
	}

	for _, test := range tests {
		result := asteroidCollision(test.input)
		if !reflect.DeepEqual(result, test.expect) {
			t.Errorf("For input %v, expected %v, but got %v.", test.input, test.expect, result)
		}
	}
}

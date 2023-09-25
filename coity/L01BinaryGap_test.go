package coity

import (
	"strconv"
	"testing"
)

func binaryGap(N int64) int {
	// 將數字 N 轉換為二進制表示
	binaryRepresentation := strconv.FormatInt(int64(N), 2)

	// 初始化兩個變量，一個用於存儲當前的間隙長度，另一個用於存儲最大的間隙長度
	currentGapLength := 0
	maxGapLength := 0

	// 遍歷二進制表示中的每一位，找出所有的二進制間隙
	for _, bit := range binaryRepresentation {
		if bit == '0' {
			// 如果當前位是 0，則增加當前間隙長度
			currentGapLength++
		} else {
			// 如果當前位是 1，則檢查當前間隙長度是否大於最大間隙長度，如果是，則更新最大間隙長度
			if currentGapLength > maxGapLength {
				maxGapLength = currentGapLength
			}
			// 重置當前間隙長度為 0
			currentGapLength = 0
		}
	}

	// 返回最大間隙長度
	return maxGapLength
}

func TestBinaryGap(t *testing.T) {
	tests := []struct {
		N    int64
		want int
	}{
		{1041, 5},
		{15, 0},
		{32, 0},
		{888, 1},
	}

	for _, tt := range tests {
		got := binaryGap(tt.N)
		if got != tt.want {
			t.Errorf("binaryGap(%v) = %v; want %v", tt.N, got, tt.want)
		}
	}
}

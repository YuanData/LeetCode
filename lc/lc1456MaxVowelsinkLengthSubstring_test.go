package lc

import "testing"

func maxVowels(s string, k int) int {
	vowels := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	count := 0

	// 使用for迴圈來計算第一個長度為k的子字串中的母音字母數量
	for _, char := range s[:k] {
		if _, exists := vowels[char]; exists {
			count++
		}
	}

	maxCount := count

	// 滑動窗口
	for i := k; i < len(s); i++ {
		// 如果左邊的字母是母音，則減少count
		if _, exists := vowels[rune(s[i-k])]; exists {
			count--
		}
		// 如果右邊的字母是母音，則增加count
		if _, exists := vowels[rune(s[i])]; exists {
			count++
		}
		// 更新最大母音數量
		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}

func TestMaxVowels(t *testing.T) {
	tests := []struct {
		s      string
		k      int
		result int
	}{
		{"abciiidef", 3, 3},
		{"aeiou", 2, 2},
		{"leetcode", 3, 2},
	}

	for _, test := range tests {
		got := maxVowels(test.s, test.k)
		if got != test.result {
			t.Errorf("For s=%s, k=%d; expected %d but got %d", test.s, test.k, test.result, got)
		}
	}
}

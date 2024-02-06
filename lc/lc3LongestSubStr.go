package lc

func lenLongestSubStr(s string) int {

	charMap := make(map[rune]int)

	start, maxLen := 0, 0

	for end, char := range s {
		if idx, found := charMap[char]; found && idx >= start {
			start = idx + 1
		}

		curLen := end - start + 1
		if curLen > maxLen {
			maxLen = curLen
		}
		charMap[char] = end
	}

	return maxLen
}

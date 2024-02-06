package lc

func lenLongestSubStr(s string) int {

	charMap := make(map[rune]int)
	// sotres the last idx where char was last seen.

	start, maxLen := 0, 0
	// start is the start of window
	// maxLen is the len of the longest non-repeating substr

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

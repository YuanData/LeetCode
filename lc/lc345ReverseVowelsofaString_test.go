package lc

import (
	"testing"
)

func TestReverseVowels(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{"hello", "holle"},
		{"leetcode", "leotcede"},
		{"aeiou", "uoiea"},
		{"", ""},       // empty string
		{"bcd", "bcd"}, // string with no vowels
		{"Aa", "aA"},   // case sensitivity
	}

	for _, testCase := range testCases {
		actualOutput := reverseVowels(testCase.input)
		if actualOutput != testCase.output {
			t.Errorf("For input %q, expected %q, but got %q", testCase.input, testCase.output, actualOutput)
		}
	}
}

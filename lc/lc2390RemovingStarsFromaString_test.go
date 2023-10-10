package lc

import "testing"

func removeStars(s string) string {
	stack := []rune{}

	for _, char := range s {
		if char == '*' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1] // Pop the top element
			}
		} else {
			stack = append(stack, char)
		}
	}

	return string(stack)
}
func TestRemoveStars(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{"leet**cod*e", "lecoe"},
		{"erase*****", ""},
	}

	for _, test := range tests {
		result := removeStars(test.input)
		if result != test.expect {
			t.Errorf("For input '%s', expected '%s', but got '%s'", test.input, test.expect, result)
		}
	}
}

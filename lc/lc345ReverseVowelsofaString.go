package lc

import "unicode"

func isVowel(char rune) bool {
	// converting the char to lower case to ensure case insensitivity
	char = unicode.ToLower(char)
	// checking whether the given char is a vowel
	return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u'
}

func reverseVowels(inputString string) string {
	runes := []rune(inputString)                // converting the string to a slice of runes
	strLength := len(runes)                     // storing the length of the runes slice
	leftPointer, rightPointer := 0, strLength-1 // initializing the pointers

	// looping until the left pointer is less than the right pointer
	for leftPointer < rightPointer {
		if leftPointer < rightPointer && !isVowel(runes[leftPointer]) {
			// finding the index of the first vowel from the left
			leftPointer += 1
		} else if rightPointer > leftPointer && !isVowel(runes[rightPointer]) {
			// finding the index of the first vowel from the right
			rightPointer -= 1
		} else {
			// swapping the vowels
			runes[leftPointer], runes[rightPointer] = runes[rightPointer], runes[leftPointer]

			// moving the pointers inward
			leftPointer += 1
			rightPointer -= 1
		}
	}
	return string(runes) // converting the slice of runes back to a string
}

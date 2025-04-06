//Given a string s, reverse only all the vowels in the string and return it.
//The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.
//
//Example 1:
//Input: s = "IceCreAm"
//Output: "AceCreIm"
//
//Explanation:
//The vowels in s are ['I', 'e', 'e', 'A']. On reversing the vowels, s becomes "AceCreIm".
//
//Example 2:
//Input: s = "leetcode"
//Output: "leotcede"
//
//Constraints:
//1 <= s.length <= 3 * 105
//s consist of printable ASCII characters.

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("r1: %s\n", reverseVowels("IceCreAm"))
	fmt.Printf("r2: %s\n", reverseVowels("leetcode"))
}

var target = map[rune]struct{}{
	'a': {},
	'e': {},
	'i': {},
	'o': {},
	'u': {},
	'A': {},
	'E': {},
	'I': {},
	'O': {},
	'U': {},
}

func reverseVowels(s string) string {
	// check edge case
	if len(s) == 1 {
		return s
	}

	startPtr := 0
	endPtr := len(s) - 1

	resource := []rune(s)

	for {
		sIdx := findTargetFromBegin(resource, startPtr)
		eIdx := findTargetFromEnd(resource, endPtr)

		// check is found
		if sIdx == -1 || eIdx == -1 {
			return string(resource)
		}

		// check is overlapped
		if sIdx >= eIdx {
			return string(resource)
		}

		// swap
		resource[eIdx], resource[sIdx] = resource[sIdx], resource[eIdx]

		startPtr, endPtr = sIdx+1, eIdx-1
	}
}

func findTargetFromBegin(r []rune, startIdx int) int {
	for i := startIdx; i < len(r); i++ {
		if _, ok := target[r[i]]; ok {
			return i
		}
	}

	return -1
}

func findTargetFromEnd(r []rune, endIdx int) int {
	for i := endIdx; i > 0; i-- {
		if _, ok := target[r[i]]; ok {
			return i
		}
	}

	return -1
}

func reverseVowelsV2(s string) string {
	// Define the set of vowels
	vowels := "aeiouAEIOU"
	// Convert the string to a slice of runes to handle UTF-8 characters
	runes := []rune(s)

	// Use two pointers to swap vowels
	left, right := 0, len(runes)-1
	for left < right {
		// Move left pointer to the next vowel
		for left < right && !strings.ContainsRune(vowels, runes[left]) {
			left++
		}
		// Move right pointer to the previous vowel
		for left < right && !strings.ContainsRune(vowels, runes[right]) {
			right--
		}
		// Swap the vowels at left and right pointers
		if left < right {
			runes[left], runes[right] = runes[right], runes[left]
			left++
			right--
		}
	}

	// Convert the slice of runes back to a string
	return string(runes)
}

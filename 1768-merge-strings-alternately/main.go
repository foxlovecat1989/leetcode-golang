package main

import (
	"fmt"
	"strings"
)

//You are given two strings word1 and word2. Merge the strings by adding letters in alternating order, starting with word1.
//If a string is longer than the other, append the additional letters onto the end of the merged string.
//Return the merged string.
//
//Example 1:
//Input: word1 = "abc", word2 = "pqr"
//Output: "apbqcr"
//Explanation: The merged string will be merged as so:
//word1:  a   b   c
//word2:    p   q   r
//merged: a p b q c r
//
//Example 2:
//Input: word1 = "ab", word2 = "pqrs"
//Output: "apbqrs"
//Explanation: Notice that as word2 is longer, "rs" is appended to the end.
//word1:  a   b
//word2:    p   q   r   s
//merged: a p b q   r   s
//
//Example 3:
//Input: word1 = "abcd", word2 = "pq"
//Output: "apbqcd"
//Explanation: Notice that as word1 is longer, "cd" is appended to the end.
//word1:  a   b   c   d
//word2:    p   q
//merged: a p b q c   d
//
//Constraints:
//1 <= word1.length, word2.length <= 100
//word1 and word2 consist of lowercase English letters.

func main() {
	result1 := mergeAlternately("abc", "pqr")
	fmt.Printf("result1: %s", result1)

	result2 := mergeAlternately("ab", "pqrs")
	fmt.Printf("result2: %s\n", result2)

	result3 := mergeAlternately("abcd", "pq")
	fmt.Printf("result3: %s\n", result3)
}

func mergeAlternately(word1 string, word2 string) string {
	// check word1 is greater than or equal to word2
	isWord1LenGreaterThanOrEqualWord2 := true
	if len(word2) > len(word1) {
		isWord1LenGreaterThanOrEqualWord2 = false
	}

	// get bigger one's remaining part
	var remainingPart []rune
	var minLen int
	if isWord1LenGreaterThanOrEqualWord2 {
		minLen = len(word2)
		remainingPart = []rune(word1)[minLen:]
	} else {
		minLen = len(word1)
		remainingPart = []rune(word2)[minLen:]
	}

	// merge two string
	var result strings.Builder
	for i := 0; i < minLen; i++ {
		_, err := result.WriteRune(rune(word1[i]))
		if err != nil {
			fmt.Printf("write rune occur an err: %#v", err)
			return ""
		}

		_, err = result.WriteRune(rune(word2[i]))
		if err != nil {
			fmt.Printf("write rune occur an err: %#v", err)
			return ""
		}
	}

	// add bigger one's remaining part
	for _, r := range remainingPart {
		result.WriteRune(r)
	}

	return result.String()
}

func mergeAlternatelyGPT(word1 string, word2 string) string {
	var result []byte
	i, j := 0, 0
	// Traverse through both strings until we reach the end of both
	for i < len(word1) || j < len(word2) {
		if i < len(word1) {
			result = append(result, word1[i])
			i++
		}
		if j < len(word2) {
			result = append(result, word2[j])
			j++
		}
	}

	return string(result)
}

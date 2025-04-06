//Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
//A subsequence of a string is a new string that is formed from the original string
//by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters.
//(i.e., "ace" is a subsequence of "abcde" while "aec" is not).
//
//Example 1:
//Input: s = "abc", t = "ahbgdc"
//Output: true
//
//Example 2:
//Input: s = "axc", t = "ahbgdc"
//Output: false
//
//Constraints:
//0 <= s.length <= 100
//0 <= t.length <= 104
//s and t consist only of lowercase English letters.

package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Printf("r1: %t\n", isSubsequence("abc", "ahbgdc"))
	//fmt.Printf("r2: %t\n", isSubsequence("axc", "ahbgdc"))
	//fmt.Printf("r3: %t\n", isSubsequence("acb", "ahbgdc"))
	fmt.Printf("r3: %t\n", isSubsequence("aaaaaa", "bbaaaa"))
}

func isSubsequence(s string, t string) bool {
	// check edge case, when s = "", t = ""
	if len(t) == 0 && len(s) == 0 {
		return true
	}
	// edge case, when s = "abc", t = ""
	if len(t) == 0 {
		return false
	}

	target := t
	for _, r := range []rune(s) {
		idx := strings.Index(target, string(r))

		if idx == -1 {
			return false
		}

		// remaining str
		target = target[idx+1:]
	}

	return true
}

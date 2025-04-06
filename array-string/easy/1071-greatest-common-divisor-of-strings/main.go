package main

import (
	"fmt"
	"strings"
)

//For two strings s and t, we say "t divides s" if and only if s = t + t + t + ... + t + t
//(i.e., t is concatenated with itself one or more times).
//Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.
//
//Example 1:
//Input: str1 = "ABCABC", str2 = "ABC"
//Output: "ABC"
//
//Example 2:
//Input: str1 = "ABABAB", str2 = "ABAB"
//Output: "AB"
//
//Example 3:
//Input: str1 = "LEET", str2 = "CODE"
//Output: ""
//
//Constraints:
//1 <= str1.length, str2.length <= 1000
//str1 and str2 consist of English uppercase letters.

func main() {
	r1 := gcdOfStrings("ABCABC", "ABC")
	fmt.Printf("r1: %s\n", r1)

	r2 := gcdOfStrings("ABABAB", "ABAB")
	fmt.Printf("r2: %s\n", r2)

	r3 := gcdOfStrings("LEET", "CODE")
	fmt.Printf("r3: %s\n", r3)

	r4 := gcdOfStrings("TAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX")
	fmt.Printf("r4: %s\n", r4)

	r5 := gcdOfStrings("AAAAAAAAA", "AAACCC")
	fmt.Printf("r5: %s\n", r5)
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

func gcdOfStrings(str1 string, str2 string) string {
	gcdLength := gcd(len(str1), len(str2))
	commonStr := str1[:gcdLength]

	if strings.Repeat(commonStr, len(str1)/gcdLength) == str1 &&
		strings.Repeat(commonStr, len(str2)/gcdLength) == str2 {
		return commonStr
	}

	return ""
}

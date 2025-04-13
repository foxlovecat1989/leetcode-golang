//Given an array of integers arr, return true if the number of occurrences
//of each value in the array is unique or false otherwise.
//
//Example 1:
//Input: arr = [1,2,2,1,1,3]
//Output: true
//Explanation: The value 1 has 3 occurrences, 2 has 2 and 3 has 1. No two values have the same number of occurrences.
//
//Example 2:
//Input: arr = [1,2]
//Output: false
//
//Example 3:
//Input: arr = [-3,0,1,-3,1,1,1,-3,10,0]
//Output: true
//
//
//Constraints:
//1 <= arr.length <= 1000
//-1000 <= arr[i] <= 1000

package main

import "fmt"

func main() {
	fmt.Printf("r1: %v\n", uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
	fmt.Printf("r2: %v\n", uniqueOccurrences([]int{1, 2}))
	fmt.Printf("r3: %v\n", uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}))
}

func uniqueOccurrences(arr []int) bool {
	dataMap := make(map[int]int)
	for _, a := range arr {
		if v, ok := dataMap[a]; ok {
			dataMap[a] = v + 1
		} else {
			dataMap[a] = 1
		}
	}

	timesSet := make(map[int]struct{})
	for _, v := range dataMap {
		timesSet[v] = struct{}{}
	}

	return len(dataMap) == len(timesSet)
}

func uniqueOccurrences2(arr []int) bool {
	// Count the frequency of each number
	freq := make(map[int]int)
	for _, num := range arr {
		freq[num]++
	}

	// Track which frequencies we've seen
	seen := make(map[int]bool)
	for _, count := range freq {
		if seen[count] {
			return false
		}

		seen[count] = true
	}

	return true
}

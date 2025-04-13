//Given two 0-indexed integer arrays nums1 and nums2, return a list answer of size 2 where:
//answer[0] is a list of all distinct integers in nums1 which are not present in nums2.
//answer[1] is a list of all distinct integers in nums2 which are not present in nums1.
//Note that the integers in the lists may be returned in any order.
//
//Example 1:
//Input: nums1 = [1,2,3], nums2 = [2,4,6]
//Output: [[1,3],[4,6]]
//Explanation:
//For nums1, nums1[1] = 2 is present at index 0 of nums2,
//whereas nums1[0] = 1 and nums1[2] = 3 are not present in nums2. Therefore, answer[0] = [1,3].
//For nums2, nums2[0] = 2 is present at index 1 of nums1,
//whereas nums2[1] = 4 and nums2[2] = 6 are not present in nums1. Therefore, answer[1] = [4,6].
//
//Example 2:
//Input: nums1 = [1,2,3,3], nums2 = [1,1,2,2]
//Output: [[3],[]]
//
//Explanation:
//For nums1, nums1[2] and nums1[3] are not present in nums2. Since nums1[2] == nums1[3], their value is only included once and answer[0] = [3].
//Every integer in nums2 is present in nums1. Therefore, answer[1] = [].
//
//Constraints:
//1 <= nums1.length, nums2.length <= 1000
//-1000 <= nums1[i], nums2[i] <= 1000

package main

import "fmt"

func main() {
	fmt.Printf("r1: %v", findDifference([]int{1, 2, 3}, []int{2, 4, 6}))
	fmt.Printf("r1: %v", findDifference([]int{1, 2, 3, 3}, []int{1, 1, 2, 2}))
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	set1 := make(map[int]struct{})
	set2 := make(map[int]struct{})
	for _, n := range nums1 {
		set1[n] = struct{}{}
	}
	for _, n := range nums2 {
		set2[n] = struct{}{}
	}

	r1 := make([]int, 0)
	duplicated := make(map[int]struct{})
	for k := range set1 {
		if _, ok := set2[k]; !ok {
			r1 = append(r1, k)
		} else {
			duplicated[k] = struct{}{}
		}
	}
	r2 := make([]int, 0)
	for k := range set2 {
		if _, ok := duplicated[k]; !ok {
			r2 = append(r2, k)
		}
	}

	return [][]int{r1, r2}
}

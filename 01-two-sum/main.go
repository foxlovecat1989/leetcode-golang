package main

import "fmt"

//Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
//You can return the answer in any order.
//
//
//Example 1:
//Input: nums = [2,7,11,15], target = 9
//Output: [0,1]
//Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
//
//Example 2:
//Input: nums = [3,2,4], target = 6
//Output: [1,2]
//
//Example 3:
//Input: nums = [3,3], target = 6
//Output: [0,1]
//
//
//Constraints:
//2 <= nums.length <= 104
//-109 <= nums[i] <= 109
//-109 <= target <= 109
//Only one valid answer exists.

func main() {
	//sum1 := twoSum([]int{2, 7, 11, 15}, 9)
	//fmt.Printf("sum1: %#v\n", sum1)
	//
	//sum2 := twoSum([]int{3, 2, 4}, 6)
	//fmt.Printf("sum2: %#v\n", sum2)
	//
	//sum3 := twoSum([]int{3, 3}, 6)
	//fmt.Printf("sum3: %#v\n", sum3)

	sum4 := twoSum2([]int{2, 5, 5, 11}, 10)
	fmt.Printf("sum4: %#v\n", sum4)
}

func twoSum1(nums []int, target int) []int {
	// deal with only two elements
	if len(nums) == 2 {
		return []int{0, 1}
	}

	for i := 0; i < len(nums); i++ {
		for k := i + 1; k < len(nums); k++ {
			diff := target - nums[i]
			if nums[k] == diff {
				return []int{i, k}
			}
		}
	}

	return nil
}

// better solution
func twoSum2(nums []int, target int) []int {
	// deal with only two elements
	if len(nums) == 2 {
		return []int{0, 1}
	}

	m := make(map[int]int) // Store number and its index
	for i, num := range nums {
		diff := target - num
		if index, found := m[diff]; found {
			return []int{index, i}
		}
		m[num] = i
	}

	return nil
}

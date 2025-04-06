//You have a long flowerbed in which some of the plots are planted, and some are not.
//However, flowers cannot be planted in adjacent plots.
//Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty, and an integer n,
//return true if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule and false otherwise.
//
//Example 1:
//Input: flowerbed = [1,0,0,0,1], n = 1
//Output: true
//
//Example 2:
//Input: flowerbed = [1,0,0,0,1], n = 2
//Output: false
//
//Constraints:
//1 <= flowerbed.length <= 2 * 104
//flowerbed[i] is 0 or 1.
//There are no two adjacent flowers in flowerbed.
//0 <= n <= flowerbed.length

package main

import "fmt"

func main() {
	fmt.Printf("r1: %t\n", canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
	fmt.Printf("r2: %t\n", canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	// check edge case
	if n == 0 {
		return true
	}

	// add two vacancies to begin and end
	flowerbed = append([]int{0}, flowerbed...)
	flowerbed = append(flowerbed, 0)

	count := 3
	for _, f := range flowerbed {
		// minus 1, if is empty
		if f == 0 {
			count--
		}
		// reset count, if is not empty
		if f == 1 {
			count = 3
		}

		if count == 0 {
			// planted
			n -= 1
			// there is one more position is empty, so count reset to 2
			count = 2
			// check is all planted
			if n == 0 {
				return true
			}
		}
	}

	return false
}

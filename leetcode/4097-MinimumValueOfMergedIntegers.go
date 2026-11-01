package main

// 4097. Minimum Value of Merged Integers
// You are given an integer array nums of even length n.

// You have to merge the elements of nums in pairs by merging nums[0] with nums[1], nums[2] with nums[3], and so on, until nums[n - 2] with nums[n - 1].

// Return the minimum value among all the resulting merged integers.

// To merge two integers, concatenate all the decimal digits of the first integer followed by all the decimal digits of the second integer.

// Example 1:
// Input: nums = [23,55,42,21]
// Output: 2355
// Explanation:
// - Merging nums[0] = 23 and nums[1] = 55 gives 2355.
// - Merging nums[2] = 42 and nums[3] = 21 gives 4221.
// - The minimum value among the merged numbers is min(2355, 4221) = 2355.

// Example 2:
// Input: nums = [34,7,11,11]
// Output: 347
// Explanation:
// - Merging nums[0] = 34 and nums[1] = 7 gives 347.
// - Merging nums[2] = 11 and nums[3] = 11 gives 1111.
// - The minimum value among the merged numbers is min(347, 1111) = 347.

// Example 3:
// Input: nums = [9,20]
// Output: 920
// Explanation:
// - Merging nums[0] = 9 and nums[1] = 20 gives 920.
// - There is only one merged number, so the minimum value is 920.

// Constraints:
//     - 2 <= n == nums.length <= 100
//     - n is even
//     - 1 <= nums[i] <= 1000

import (
	"fmt"
	"strconv"
)

func minMergedInteger(nums []int) int {
	res := 1 << 61
	for i := 0; i < len(nums); i += 2 {
		x, y := nums[i], nums[i+1]
		for v := nums[i+1]; v > 0; v /= 10 {
			x *= 10
		}
		res = min(res, x+y)
	}
	return res
}

func minMergedInteger1(nums []int) int {
	res, n := int64(1<<61), len(nums)
	for i := 0; i < n; i += 2 {
		v, _ := strconv.ParseInt(strconv.Itoa(nums[i])+strconv.Itoa(nums[i+1]), 10, 64)
		res = min(res, v)
	}
	return int(res)
}

func main() {
	// Example 1:
	// Input: nums = [23,55,42,21]
	// Output: 2355
	// Explanation:
	// - Merging nums[0] = 23 and nums[1] = 55 gives 2355.
	// - Merging nums[2] = 42 and nums[3] = 21 gives 4221.
	// - The minimum value among the merged numbers is min(2355, 4221) = 2355.
	fmt.Println(minMergedInteger([]int{23, 55, 42, 21})) // 2355
	// Example 2:
	// Input: nums = [34,7,11,11]
	// Output: 347
	// Explanation:
	// - Merging nums[0] = 34 and nums[1] = 7 gives 347.
	// - Merging nums[2] = 11 and nums[3] = 11 gives 1111.
	// - The minimum value among the merged numbers is min(347, 1111) = 347.
	fmt.Println(minMergedInteger([]int{34, 7, 11, 11})) // 347
	// Example 3:
	// Input: nums = [9,20]
	// Output: 920
	// Explanation:
	// - Merging nums[0] = 9 and nums[1] = 20 gives 920.
	// - There is only one merged number, so the minimum value is 920.
	fmt.Println(minMergedInteger([]int{9, 20})) // 920

	fmt.Println(minMergedInteger([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})) // 12
	fmt.Println(minMergedInteger([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})) // 21

	fmt.Println(minMergedInteger1([]int{23, 55, 42, 21}))                // 2355
	fmt.Println(minMergedInteger1([]int{34, 7, 11, 11}))                 // 347
	fmt.Println(minMergedInteger1([]int{9, 20}))                         // 920
	fmt.Println(minMergedInteger1([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})) // 12
	fmt.Println(minMergedInteger1([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})) // 21
}

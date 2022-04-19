package main

import "fmt"

/**
136. Single Number
Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
You must implement a solution with a linear runtime complexity and use only constant extra space.

Constraints:

	1 <= nums.length <= 3 * 10^4
	-3 * 10^4 <= nums[i] <= 3 * 10^4
	Each element in the array appears twice except for one element which appears only once.

Example 1:

	Input: nums = [2,2,1]
	Output: 1

Example 2:

	Input: nums = [4,1,2,1,2]
	Output: 4

Example 3:

	Input: nums = [1]
	Output: 1

# 解题思路

	给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
	找出那个只出现了一次的元素。
	要求算法时间复杂度是线性的，
	并且不使用额外的辅助空间。

	 x^x = 0
 */

func singleNumber(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result ^= nums[i] // x^x = 0
	}
	return result
}

func main() {
	fmt.Printf("singleNumber([]int{2,2,1}) = %v\n",singleNumber([]int{2,2,1})) // 1
	fmt.Printf("singleNumber([]int{4,1,2,1,2}) = %v\n",singleNumber([]int{4,1,2,1,2})) // 4
	fmt.Printf("singleNumber([]int{1}) = %v\n",singleNumber([]int{1})) // 1
}

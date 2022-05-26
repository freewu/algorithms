package main

import "fmt"

/**
260. Single Number III
Given an integer array nums, in which exactly two elements appear only once and all the other elements appear exactly twice.
Find the two elements that appear only once. You can return the answer in any order.
You must write an algorithm that runs in linear runtime complexity and uses only constant extra space.

Example 1:

	Input: nums = [1,2,1,3,2,5]
	Output: [3,5]
	Explanation:  [5, 3] is also a valid answer.

Example 2:

	Input: nums = [-1,0]
	Output: [-1,0]

Example 3:

	Input: nums = [0,1]
	Output: [1,0]

Constraints:

	2 <= nums.length <= 3 * 10^4
	-2^31 <= nums[i] <= 2^31 - 1
	Each integer in nums will appear twice, only two integers will appear once.

给定一个整数数组 nums，其中恰好有两个元素只出现一次，
其余所有元素均出现两次。 找出只出现一次的那两个元素
 */

func singleNumber(nums []int) []int {
	diff := 0
	// 先做一次异或处理 diff 就是 两个出现一次的 异或 值
	for _, num := range nums {
		diff ^= num
	}
	// Get its last set bit (lsb)
	diff &= -diff
	res := []int{0, 0} // this array stores the two numbers we will return
	for _, num := range nums {
		if (num & diff) == 0 { // the bit is not set
			res[0] ^= num
		} else { // the bit is set
			res[1] ^= num
		}
	}
	return res
}

// best solution
func singleNumberBest(nums []int) []int {
	if len(nums) == 2 {
		return nums
	}
	xor := nums[0]
	for i := 1; i < len(nums); i++ {
		xor = xor ^ nums[i]
	}
	k := 0
	for {
		if xor & 1 == 0 {
			xor = xor >> 1
			k++
		} else {
			break
		}
	}
	num1, num2 := 0, 0
	for _, n := range nums {
		if (n >> k) & 1 == 1 {
			num1 = num1 ^ n
		} else {
			num2 = num2 ^ n
		}
	}
	return []int{num1, num2}
}

func main() {
	fmt.Printf("singleNumber([]int{1,2,1,3,2,5}) = %v\n",singleNumber([]int{1,2,1,3,2,5})) // [3,5]
	fmt.Printf("singleNumber([]int{-1,0}) = %v\n",singleNumber([]int{-1,0})) // [-1,0]
	fmt.Printf("singleNumber([]int{0,1}) = %v\n",singleNumber([]int{0,1})) // [0,1]

	fmt.Printf("singleNumberBest([]int{1,2,1,3,2,5}) = %v\n",singleNumberBest([]int{1,2,1,3,2,5})) // [3,5]
	fmt.Printf("singleNumberBest([]int{-1,0}) = %v\n",singleNumberBest([]int{-1,0})) // [-1,0]
	fmt.Printf("singleNumberBest([]int{0,1}) = %v\n",singleNumberBest([]int{0,1})) // [0,1]
}

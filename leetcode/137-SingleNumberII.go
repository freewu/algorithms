package main

import "fmt"

/**
137. Single Number II
Given an integer array nums where every element appears three times except for one, which appears exactly once. Find the single element and return it.
You must implement a solution with a linear runtime complexity and use only constant extra space.

Constraints:

	1 <= nums.length <= 3 * 10^4
	-2^31 <= nums[i] <= 2^31 - 1
	Each element in nums appears exactly three times except for one element which appears once.

Example 1:

	Input: nums = [2,2,3,2]
	Output: 3

Example 2:

	Input: nums = [0,1,0,1,0,1,99]
	Output: 99

# 解题思路
	除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。
	要求算法时间复杂度是线性的，
	并且不使用额外的辅助空间。

	& ^twos  为了能做到三进制，出现 3 次就清零
 */

func singleNumberII(nums []int) int {
	ones, twos := 0, 0
	for i := 0; i < len(nums); i++ {
		ones = (ones ^ nums[i]) & ^twos
		twos = (twos ^ nums[i]) & ^ones
		fmt.Printf("nums[%v] = %v,ones = %v,twos = %v\n",i,nums[i],ones,twos)
	}
	return ones
}

// best solution
func singleNumberBest(nums []int) int {
	ones, twos := 0, 0
	for _, num := range nums {
		ones = (ones ^ num) & ^twos
		twos = (twos ^ num) & ^ones
	}
	return ones
}

// 在数组中每个元素都出现 5 次，找出只出现 1 次的数。

// 解法一
func singleNumberIIIII(nums []int) int {
	na, nb, nc := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		nb = nb ^ (nums[i] & na)
		na = (na ^ nums[i]) & ^nc
		nc = nc ^ (nums[i] & ^na & ^nb)
	}
	return na & ^nb & ^nc
}

// 解法二
func singleNumberIIIII1(nums []int) int {
	twos, threes, ones := 0xffffffff, 0xffffffff, 0
	for i := 0; i < len(nums); i++ {
		threes = threes ^ (nums[i] & twos)
		twos = (twos ^ nums[i]) & ^ones
		ones = ones ^ (nums[i] & ^twos & ^threes)
	}
	return ones
}

func main() {
	fmt.Printf("singleNumberII([]int{ 2,2,3,2 }) = %v\n",singleNumberII([]int{ 2,2,3,2 })) // 3
	fmt.Printf("singleNumberII([]int{ 0,1,0,1,0,1,99 }) = %v\n",singleNumberII([]int{ 0,1,0,1,0,1,99 })) // 99

	fmt.Printf("singleNumberBest([]int{ 2,2,3,2 }) = %v\n",singleNumberBest([]int{ 2,2,3,2 })) // 3
	fmt.Printf("singleNumberBest([]int{ 0,1,0,1,0,1,99 }) = %v\n",singleNumberBest([]int{ 0,1,0,1,0,1,99 })) // 99
}

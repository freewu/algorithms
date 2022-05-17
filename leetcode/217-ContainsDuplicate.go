package main

import "fmt"

/**
217. Contains Duplicate
Given an integer array nums, return true if any value appears at least twice in the array,
and return false if every element is distinct.

Example 1:

	Input: nums = [1,2,3,1]
	Output: true
Example 2:

	Input: nums = [1,2,3,4]
	Output: false

Example 3:

	Input: nums = [1,1,1,3,3,4,3,2,4,2]
	Output: true


Constraints:

	1 <= nums.length <= 105
	-109 <= nums[i] <= 109
 */

func containsDuplicate(nums []int) bool {
	record := make(map[int]bool, len(nums)) // 声明一个map 元素为 key 如果存 直接返回 true 不存，保存到map里
	for _, n := range nums {
		if _, found := record[n]; found { // 存在重复数
			return true
		}
		record[n] = true // 第一次出 n 这值
	}
	return false
}

func main() {
	fmt.Printf("containsDuplicate([]int{ 1,2,3,1 }) = %v\n",containsDuplicate([]int{ 1,2,3,1 })) // true
	fmt.Printf("containsDuplicate([]int{ 1,2,3,4 }) = %v\n",containsDuplicate([]int{ 1,2,3,4 })) // false
	fmt.Printf("containsDuplicate([]int{ 1,1,1,3,3,4,3,2,4,2 }) = %v\n",containsDuplicate([]int{ 1,1,1,3,3,4,3,2,4,2 })) // true
}
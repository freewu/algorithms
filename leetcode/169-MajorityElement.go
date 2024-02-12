package main

import "fmt"

/**
169. Majority Element
Given an array nums of size n, return the majority element.
The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

Example 1:

	Input: nums = [3,2,3]
	Output: 3

Example 2:

	Input: nums = [2,2,1,1,1,2,2]
	Output: 2

Constraints:

	n == nums.length
	1 <= n <= 5 * 10^4
	-10^9 <= nums[i] <= 10^9

Follow-up: Could you solve the problem in linear time and in O(1) space?

# 解题思路

	给定一个大小为 n 的数组，找到其中的众数。
	众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。你可以假设数组是非空的，并且给定的数组总是存在众数。

 */

// 解法一 时间复杂度 O(n) 空间复杂度 O(1)
func majorityElement(nums []int) int {
	res, count := 0, 0 // 默认第一的值就是返回值 只有一个的话 直接返回了
	for i := 0; i < len(nums); i++ {
		if count == 0 { // 如果累加到 0 重新赋值
			res, count = nums[i], 1
		} else {
			if nums[i] == res {
				count++ // 如果还是自己 普累加
			} else {
				count-- // 如果不是自己就减去 1
			}
		}
	}
	return res
}

// 解法二 时间复杂度 O(n) 空间复杂度 O(n)
func majorityElement1(nums []int) int {
	m := make(map[int]int) // 声明一个map来保存数量
	l := len(nums) / 2
	for _, v := range nums {
		m[v]++
		if m[v] > l { // 如果统计到的数值多于一半了直接返回
			return v
		}
	}
	return 0
}

// 思路和解法1一样
func majorityElement2(nums []int) int {
	major, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			major = nums[i]
			count++
		} else if major == nums[i] {
			count++
		} else {
			count--
		}
	}
	return major
}

func main() {
	fmt.Printf("majorityElement1([]int{ 3,2,3 }) = %v\n",majorityElement1([]int{ 3,2,3 })) // 3
	fmt.Printf("majorityElement1([]int{ 2,2,1,1,1,2,2 }) = %v\n",majorityElement1([]int{ 2,2,1,1,1,2,2 })) // 2

	fmt.Printf("majorityElement([]int{ 3,2,3 }) = %v\n",majorityElement([]int{ 3,2,3 })) // 3
	fmt.Printf("majorityElement([]int{ 2,2,1,1,1,2,2 }) = %v\n",majorityElement([]int{ 2,2,1,1,1,2,2 })) // 2

	fmt.Printf("majorityElement2([]int{ 3,2,3 }) = %v\n",majorityElement2([]int{ 3,2,3 })) // 3
	fmt.Printf("majorityElement2([]int{ 2,2,1,1,1,2,2 }) = %v\n",majorityElement2([]int{ 2,2,1,1,1,2,2 })) // 2

}
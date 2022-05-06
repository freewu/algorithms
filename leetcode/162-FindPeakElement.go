package main

import (
	"fmt"
	"math"
)

/**
162. Find Peak Element
A peak element is an element that is strictly greater than its neighbors.
Given an integer array nums, find a peak element, and return its index. If the array contains multiple peaks, return the index to any of the peaks.
You may imagine that nums[-1] = nums[n] = -∞.
You must write an algorithm that runs in O(log n) time.

Constraints:

	1 <= nums.length <= 1000
	-2^31 <= nums[i] <= 2^31 - 1
	nums[i] != nums[i + 1] for all valid i.

Example 1:

	Input: nums = [1,2,3,1]
	Output: 2
	Explanation: 3 is a peak element and your function should return the index number 2.

Example 2:

	Input: nums = [1,2,1,3,5,6,4]
	Output: 5
	Explanation: Your function can return either index number 1 where the peak element is 2, or index number 5 where the peak element is 6.

# 解题思路

	给出一个数组，数组里面存在多个“山峰”，(山峰的定义是，下标 i 比 i-1、i+1 位置上的元素都要大)，找到这个“山峰”，并输出其中一个山峰的下标。
	num[i] > num[i - 1] & num[i] > num[i + 1] return i

 */

// 解法一 二分
func findPeakElement(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return 0
	}
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low) >> 1
		if (mid == len(nums)-1 && nums[mid-1] < nums[mid]) || (mid > 0 && nums[mid-1] < nums[mid] && (mid <= len(nums)-2 && nums[mid+1] < nums[mid])) || (mid == 0 && nums[1] < nums[0]) {
			return mid
		}
		if mid > 0 && nums[mid-1] < nums[mid] {
			low = mid + 1
		}
		if mid > 0 && nums[mid-1] > nums[mid] {
			high = mid - 1
		}
		if mid == low {
			low++
		}
		if mid == high {
			high--
		}
	}
	return -1
}

// 解法二 二分
func findPeakElement1(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low) >> 1
		// 如果 mid 较大，则左侧存在峰值，high = m，如果 mid + 1 较大，则右侧存在峰值，low = mid + 1
		if nums[mid] > nums[mid+1] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func findPeakElementBest(nums []int) int {
	return find(nums, 0, len(nums) - 1)
}

func find(n []int, i, j int) int {
	mid := (i + j) / 2
	left, right := math.MinInt32, math.MinInt32
	if mid > 0 {
		left = n[mid - 1]
	}
	if mid < len(n) - 1 {
		right = n[mid + 1]
	}
	if n[mid] > left && n[mid] > right {
		return mid
	}
	if n[mid] < left && n[mid] > right {
		return find(n, i, mid - 1)
	}
	return find(n, mid + 1, j)
}

func main() {
	fmt.Printf("findPeakElement([]int{ 1,2,3,1 }) = %v\n",findPeakElement([]int{ 1,2,3,1 })) // 2
	fmt.Printf("findPeakElement([]int{ 1,2,1,3,5,6,4 }) = %v\n",findPeakElement([]int{ 1,2,1,3,5,6,4 })) // 5

	fmt.Printf("findPeakElement1([]int{ 1,2,3,1 }) = %v\n",findPeakElement1([]int{ 1,2,3,1 })) // 2
	fmt.Printf("findPeakElement1([]int{ 1,2,1,3,5,6,4 }) = %v\n",findPeakElement1([]int{ 1,2,1,3,5,6,4 })) // 5

	fmt.Printf("findPeakElementBest([]int{ 1,2,3,1 }) = %v\n",findPeakElementBest([]int{ 1,2,3,1 })) // 2
	fmt.Printf("findPeakElementBest([]int{ 1,2,1,3,5,6,4 }) = %v\n",findPeakElementBest([]int{ 1,2,1,3,5,6,4 })) // 5
}
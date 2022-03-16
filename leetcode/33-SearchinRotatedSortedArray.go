package main

import "fmt"

/**
33. Search in Rotated Sorted Array

Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
You are given a target value to search. If found in the array return its index, otherwise return -1.
You may assume no duplicate exists in the array.
Your algorithm’s runtime complexity must be in the order of O(log n).

Example 1:

	Input: nums = [4,5,6,7,0,1,2], target = 0
	Output: 4

Example 2:

	Input: nums = [4,5,6,7,0,1,2], target = 3
	Output: -1
*/

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > nums[low] { // 在数值大的一部分区间里
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else if nums[mid] < nums[high] { // 在数值小的一部分区间里
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			if nums[low] == nums[mid] {
				low++
			}
			if nums[high] == nums[mid] {
				high--
			}
		}
	}
	return -1
}
// best solution
func searchBest(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums) - 1
	for l < r {
		mid := l + (r - l) / 2
		if nums[mid] > nums[r] { // 先找到中的点
			l = mid + 1
		} else {
			r = mid
		}
	}
	fmt.Printf("l = %v\n",l) // 找到了 两段数组的中间
	rot := l
	l, r = 0, len(nums) - 1
	if nums[rot] <= target && target <= nums[r] { // 确定值在哪段数组上
		l = rot
	} else {
		r = rot - 1
	}
	for l <= r {
		mid := l + (r - l) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

func main() {
	fmt.Printf("search([]int{4,5,6,7,0,1,2},0)  = %v\n",search([]int{4,5,6,7,0,1,2},0))
	fmt.Printf("search([]int{4,5,6,7,0,1,2},3)  = %v\n",search([]int{4,5,6,7,0,1,2},3))
	fmt.Printf("searchBest([]int{4,5,6,7,0,1,2},0)  = %v\n",searchBest([]int{4,5,6,7,0,1,2},0))
	fmt.Printf("searchBest([]int{4,5,6,7,0,1,2},3)  = %v\n",searchBest([]int{4,5,6,7,0,1,2},3))
}
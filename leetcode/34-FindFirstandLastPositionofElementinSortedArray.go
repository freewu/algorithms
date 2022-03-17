package main

import "fmt"

/**
34. Find First and Last Position of Element in Sorted Array

Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.
Your algorithm’s runtime complexity must be in the order of O(log n).
If the target is not found in the array, return [-1, -1].

Example 1:

	Input: nums = [5,7,7,8,8,10], target = 8
	Output: [3,4]

Example 2:

	Input: nums = [5,7,7,8,8,10], target = 6
	Output: [-1,-1]

给定一个按照升序排列的整数数组 nums，和一个目标值 target。
找出给定目标值在数组中的开始位置和结束位置。你的算法时间复杂度必须是 O(log n) 级别。
如果数组中不存在目标值，返回 [-1, -1]
 */

func searchRange(nums []int, target int) []int {
	return []int{searchFirstEqualElement(nums, target), searchLastEqualElement(nums, target)}

}

// 二分查找第一个与 target 相等的元素，时间复杂度 O(logn)
func searchFirstEqualElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if (mid == 0) || (nums[mid-1] != target) { // 找到第一个与 target 相等的元素
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

// 二分查找最后一个与 target 相等的元素，时间复杂度 O(logn)
func searchLastEqualElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if (mid == len(nums)-1) || (nums[mid+1] != target) { // 找到最后一个与 target 相等的元素
				return mid
			}
			low = mid + 1
		}
	}
	return -1
}

// best solution
func searchRangeBest(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	return []int{findLeft(nums, target), findRight(nums, target)}
}

func findLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left+right)/2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left >= 0 && left < len(nums) && nums[left] == target {
		return left
	}
	return -1

}

func findRight(nums []int, target int) int {
	left, right := 0, len(nums) - 1
	for left <= right {
		mid := (left+right)/2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if right >= 0 && right < len(nums) && nums[right] == target {
		return right
	}
	return -1
}

func main() {
	fmt.Printf("searchRange([]int{5,7,7,8,8,10},8) = %v\n",searchRange([]int{5,7,7,8,8,10},8))
	fmt.Printf("searchRange([]int{5,7,7,8,8,10},6) = %v\n",searchRange([]int{5,7,7,8,8,10},6))
	fmt.Printf("searchRangeBest([]int{5,7,7,8,8,10},8) = %v\n",searchRangeBest([]int{5,7,7,8,8,10},8))
	fmt.Printf("searchRangeBest([]int{5,7,7,8,8,10},6) = %v\n",searchRangeBest([]int{5,7,7,8,8,10},6))
}

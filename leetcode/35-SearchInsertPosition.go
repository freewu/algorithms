package main

/*
Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.

Example 1:
Input: [1,3,5,6], 5
Output: 2

Example 2:
Input: [1,3,5,6], 2
Output: 1

Example 3:
Input: [1,3,5,6], 7
Output: 4

Example 4:
Input: [1,3,5,6], 0
Output: 0
*/

import (
	"fmt"
)

// n loop every item
func searchInsert1(nums []int, target int) int {
	var l = len(nums)
	var t = 0
	for i := 0; i < l; i++ {
		if nums[i] == target {
			return i
		}
		if nums[i] > target {
			return t
		}
		t++
	}
    return l
}

// logn binary search 
func searchInsert(nums []int, target int) int {
	var l = len(nums) - 1
	var s = 0

	if nums[s] >= target {
		return s
	}
	if nums[l] == target {
		return l
	}
	if nums[l] < target {
		return l + 1
	}

	for {
		if s > l {
			break;
		}
		var m = (s + l) / 2
		if nums[m] == target {
			return m
		}
		if nums[m] < target {
			s = m + 1
		} else {
			l = m - 1
		}
	}
	return s
}


func main() {
	fmt.Println(searchInsert([]int{1,3,5,6},5)) // 2
	fmt.Println(searchInsert([]int{1,3,5,6},2)) // 1
	fmt.Println(searchInsert([]int{1,3,5,6},7)) // 4
	fmt.Println(searchInsert([]int{1,3,5,6},0)) // 0
}
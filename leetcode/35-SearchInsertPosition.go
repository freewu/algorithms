package main

/*
35. Search Insert Position

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
			break
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

func searchInsert2(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1 // 找到中间位置
		if nums[mid] >= target { // 如果中间位置 >= 目标值,  结束位置 设置为中间值  low -- mid -- high 说明目标值 存在 low -- mid 段中
			high = mid - 1
		} else {
			// mid == len(nums)-1) 如果 中间值 是最后一位 说明 target 最大 要从最后一位插入
            // (nums[mid+1] >= target) 中间值 >=  目标值
			//m := mid + 1
            // fmt.Printf("nums[mid+1] = %v\n",nums[m])
			fmt.Printf("mid = %v\n",mid)
			fmt.Printf("target = %v\n",target)
			if (mid == len(nums)-1) || (nums[mid+1] >= target) {
				return mid + 1
			}
			low = mid + 1
		}
	}
	return 0
}

// best solution
func searchInsertBest(nums []int, target int) int {
	//  Return right or left if target not found
	right := len(nums) - 1
	left  := 0
	for left <= right {
		middle := left + (right - left)/2
		if nums[middle] == target {
			return middle
		}
		if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func main() {
	fmt.Printf("searchInsert([]int{1,3,5,6},5) = %v\n",searchInsert([]int{1,3,5,6},5)) // 2
	fmt.Printf("searchInsert([]int{1,3,5,6},2) = %v\n",searchInsert([]int{1,3,5,6},2)) // 1
	fmt.Printf("searchInsert([]int{1,3,5,6},7) = %v\n",searchInsert([]int{1,3,5,6},7)) // 4
	fmt.Printf("searchInsert([]int{1,3,5,6},0) = %v\n",searchInsert([]int{1,3,5,6},0)) // 0
	fmt.Printf("searchInsert1([]int{1,3,5,6},5) = %v\n",searchInsert1([]int{1,3,5,6},5)) // 2
	fmt.Printf("searchInsert1([]int{1,3,5,6},2) = %v\n",searchInsert1([]int{1,3,5,6},2)) // 1
	fmt.Printf("searchInsert1([]int{1,3,5,6},7) = %v\n",searchInsert1([]int{1,3,5,6},7)) // 4
	fmt.Printf("searchInsert1([]int{1,3,5,6},0) = %v\n",searchInsert1([]int{1,3,5,6},0)) // 0
	fmt.Printf("searchInsert2([]int{1,3,5,6},5) = %v\n",searchInsert2([]int{1,3,5,6},5)) // 2
	fmt.Printf("searchInsert2([]int{1,3,5,6},2) = %v\n",searchInsert2([]int{1,3,5,6},2)) // 1
	fmt.Printf("searchInsert2([]int{1,3,5,6},7) = %v\n",searchInsert2([]int{1,3,5,6},7)) // 4
	fmt.Printf("searchInsert2([]int{1,3,5,6},0) = %v\n",searchInsert2([]int{1,3,5,6},0)) // 0

	fmt.Printf("searchInsertBest([]int{1,3,5,6},5) = %v\n",searchInsertBest([]int{1,3,5,6},5)) // 2
	fmt.Printf("searchInsertBest([]int{1,3,5,6},2) = %v\n",searchInsertBest([]int{1,3,5,6},2)) // 1
	fmt.Printf("searchInsertBest([]int{1,3,5,6},7) = %v\n",searchInsertBest([]int{1,3,5,6},7)) // 4
	fmt.Printf("searchInsertBest([]int{1,3,5,6},0) = %v\n",searchInsertBest([]int{1,3,5,6},0)) // 0
}
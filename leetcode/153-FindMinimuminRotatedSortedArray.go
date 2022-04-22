package main

import "fmt"

/**
153. Find Minimum in Rotated Sorted Array
Suppose an array of length n sorted in ascending order is rotated between 1 and n times. For example, the array nums = [0,1,2,4,5,6,7] might become:

[4,5,6,7,0,1,2] if it was rotated 4 times.
[0,1,2,4,5,6,7] if it was rotated 7 times.
Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].

Given the sorted rotated array nums of unique elements, return the minimum element of this array.
You must write an algorithm that runs in O(log n) time.

Constraints:

	n == nums.length
	1 <= n <= 5000
	-5000 <= nums[i] <= 5000
	All the integers of nums are unique.
	nums is sorted and rotated between 1 and n times.

Example 1:

	Input: nums = [3,4,5,1,2]
	Output: 1
	Explanation: The original array was [1,2,3,4,5] rotated 3 times.

Example 2:

	Input: nums = [4,5,6,7,0,1,2]
	Output: 0
	Explanation: The original array was [0,1,2,4,5,6,7] and it was rotated 4 times.

Example 3:

	Input: nums = [11,13,15,17]
	Output: 11
	Explanation: The original array was [11,13,15,17] and it was rotated 4 times.

解题思路 #
	给出一个原本从小到大排序过的数组，但是在某一个分割点上，把数组切分后的两部分对调位置，数值偏大的放到了数组的前部。
	求这个数组中最小的元素。
 */

// 解法一 二分 O(logN)
func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		if nums[low] < nums[high] {
			return nums[low]
		}
		mid := low + (high-low) >> 1 // 取中间值
		if nums[mid] >= nums[low] { // 如果 mid  >= low 把 low 移动到 mid 后一位,还在后面说明在
			low = mid + 1
		} else { // 否则说明在前面
			high = mid
		}
	}
	return nums[low]
}

// 解法二 二分
func findMin1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if nums[len(nums)-1] > nums[0] {
		return nums[0]
	}
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[low] < nums[high] {
			return nums[low]
		}
		if (mid == len(nums)-1 && nums[mid-1] > nums[mid]) || (mid < len(nums)-1 && mid > 0 && nums[mid-1] > nums[mid] && nums[mid] < nums[mid+1]) {
			return nums[mid]
		}
		if nums[mid] > nums[low] && nums[low] > nums[high] { // mid 在数值大的一部分区间里
			low = mid + 1
		} else if nums[mid] < nums[low] && nums[low] > nums[high] { // mid 在数值小的一部分区间里
			high = mid - 1
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

// 解法三 暴力 O(N)
func findMin2(nums []int) int {
	min := nums[0]
	for _, num := range nums[1:] {
		if min > num {
			min = num
		}
	}
	return min
}

// best solution
func findMinBest(nums []int) int {
	low := 0
	high := len(nums) - 1
	for low < high {
		mid := low + (high - low) / 2
		fmt.Printf("mid: %v, low: %v,high: %v\n",mid,low,high)
		fmt.Printf("nums[high]: %v, nums[mid]: %v\n",nums[high],nums[mid])
		if nums[high] < nums[mid] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return nums[low]
}

func main() {
	fmt.Printf("findMin([]int{ 3,4,5,1,2 }) = %v\n",findMin([]int{ 3,4,5,1,2 })) // 1
	fmt.Printf("findMin([]int{ 4,5,6,7,0,1,2 }) = %v\n",findMin([]int{ 4,5,6,7,0,1,2 })) // 0
	fmt.Printf("findMin([]int{ 11,13,15,17 }) = %v\n",findMin([]int{ 11,13,15,17 })) // 11

	fmt.Printf("findMin1([]int{ 3,4,5,1,2 }) = %v\n",findMin1([]int{ 3,4,5,1,2 })) // 1
	fmt.Printf("findMin1([]int{ 4,5,6,7,0,1,2 }) = %v\n",findMin1([]int{ 4,5,6,7,0,1,2 })) // 0
	fmt.Printf("findMin1([]int{ 11,13,15,17 }) = %v\n",findMin1([]int{ 11,13,15,17 })) // 11

	fmt.Printf("findMin2([]int{ 3,4,5,1,2 }) = %v\n",findMin2([]int{ 3,4,5,1,2 })) // 1
	fmt.Printf("findMin2([]int{ 4,5,6,7,0,1,2 }) = %v\n",findMin2([]int{ 4,5,6,7,0,1,2 })) // 0
	fmt.Printf("findMin2([]int{ 11,13,15,17 }) = %v\n",findMin2([]int{ 11,13,15,17 })) // 11

	fmt.Printf("findMinBest([]int{ 3,4,5,1,2 }) = %v\n",findMinBest([]int{ 3,4,5,1,2 })) // 1
	fmt.Printf("findMinBest([]int{ 4,5,6,7,0,1,2 }) = %v\n",findMinBest([]int{ 4,5,6,7,0,1,2 })) // 0
	fmt.Printf("findMinBest([]int{ 11,13,15,17 }) = %v\n",findMinBest([]int{ 11,13,15,17 })) // 11
}
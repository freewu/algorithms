package main

// 154. Find Minimum in Rotated Sorted Array II
// Suppose an array of length n sorted in ascending order is rotated between 1 and n times. 
// For example, the array nums = [0,1,4,4,5,6,7] might become:
//     [4,5,6,7,0,1,4] if it was rotated 4 times.
//     [0,1,4,4,5,6,7] if it was rotated 7 times.

// Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].
// Given the sorted rotated array nums that may contain duplicates, return the minimum element of this array.
// You must decrease the overall operation steps as much as possible.

// Example 1:
// Input: nums = [1,3,5]
// Output: 1

// Example 2:
// Input: nums = [2,2,2,0,1]
// Output: 0
 

// Constraints:
//     n == nums.length
//     1 <= n <= 5000
//     -5000 <= nums[i] <= 5000
//     nums is sorted and rotated between 1 and n times.

// Follow up: This problem is similar to Find Minimum in Rotated Sorted Array, but nums may contain duplicates. Would this affect the runtime complexity? How and why?

// # 解题思路
//     153 题的加强版，增加了重复元素的条件。
//     二分搜索，在相等元素上多增加一个判断即可。时间复杂度 O(log n)。

import "fmt"

func findMin(nums []int) int {
    low, high := 0, len(nums)-1
    for low < high {
        if nums[low] < nums[high] {
            return nums[low]
        }
        mid := low + (high-low) >> 1
        if nums[mid] > nums[low] {
            low = mid + 1
        } else if nums[mid] == nums[low] { // 判断是否是相等元素
            low++
        } else {
            high = mid
        }
    }
    return nums[low]
}

// best solution
func findMinBest(nums []int) int {
    low, high := 0, len(nums) - 1
    for low < high {
        mid := low + (high - low) / 2
        if nums[high] < nums[mid] {
            low = mid + 1
        } else if nums[high] == nums[mid] { // 判断是否是相等元素
            high--
        } else {
            high = mid
        }
    }
    return nums[low]
}

func main() {
    fmt.Printf("findMin([]int{ 1,3,5 }) = %v\n",findMin([]int{ 1,3,5 })) // 1
    fmt.Printf("findMin([]int{ 2,2,2,0,1 }) = %v\n",findMin([]int{ 2,2,2,0,1 })) // 0

    fmt.Printf("findMinBest([]int{ 1,3,5 }) = %v\n",findMinBest([]int{ 1,3,5 })) // 1
    fmt.Printf("findMinBest([]int{ 2,2,2,0,1 }) = %v\n",findMinBest([]int{ 2,2,2,0,1 })) // 0
}

package main

// 81. Search in Rotated Sorted Array II
// There is an integer array nums sorted in non-decreasing order (not necessarily with distinct values).
// Before being passed to your function, nums is rotated at an unknown pivot index k (0 <= k < nums.length)
// such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
// For example, [0,1,2,4,4,4,5,6,6,7] might be rotated at pivot index 5 and become [4,5,6,6,7,0,1,2,4,4].
// Given the array nums after the rotation and an integer target, return true if target is in nums, or false if it is not in nums.
// You must decrease the overall operation steps as much as possible.

// Example 1:
// Input: nums = [2,5,6,0,0,1,2], target = 0
// Output: true

// Example 2:
// Input: nums = [2,5,6,0,0,1,2], target = 3
// Output: false

// Constraints:
//     1 <= nums.length <= 5000
//     -10^4 <= nums[i] <= 10^4
//     nums is guaranteed to be rotated at some pivot.
//     -10^4 <= target <= 10^4

import "fmt"

func search(nums []int, target int) bool {
    if len(nums) == 0 {
        return false
    }
    low, high := 0, len(nums) - 1
    for low <= high {
        // 取中间位置
        mid := low + (high-low) >> 1
        if nums[mid] == target { // 如果是中间位置就是目标值 直接返回 true
            return true
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
    return false
}

// best solution
func search1(nums []int, target int) bool {
    low, high := 0, len(nums) - 1
    for low <= high {
        mid := (low + high) / 2
        if nums[mid] == target {
            return true
        }
        if nums[low] == nums[mid] {
            low++
            continue
        }
        if nums[low] <= nums[mid] {
            if nums[low] <= target && target <= nums[mid] {
                high = mid - 1
            } else {
                low = mid + 1
            }
        } else {
            if nums[mid] <= target && target <= nums[high] {
                low = mid + 1
            } else {
                high = mid - 1
            }
        }
    }
    return false
}

func main() {
    // Example 1:
    // Input: nums = [2,5,6,0,0,1,2], target = 0
    // Output: true
    fmt.Printf("search([]int{2,5,6,0,0,1,2},0) = %v\n",search([]int{2,5,6,0,0,1,2},0)) // true
    // Example 2:
    // Input: nums = [2,5,6,0,0,1,2], target = 3
    // Output: false
    fmt.Printf("search([]int{2,5,6,0,0,1,2},3) = %v\n",search([]int{2,5,6,0,0,1,2},3)) // false

    fmt.Printf("search1([]int{2,5,6,0,0,1,2},0) = %v\n",search1([]int{2,5,6,0,0,1,2},0)) // true
    fmt.Printf("search1([]int{2,5,6,0,0,1,2},3) = %v\n",search1([]int{2,5,6,0,0,1,2},3)) // false
}

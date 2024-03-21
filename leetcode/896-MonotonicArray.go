package main

// 896. Monotonic Array
// An array is monotonic if it is either monotone increasing or monotone decreasing.
// An array nums is monotone increasing if for all i <= j, nums[i] <= nums[j]. 
// An array nums is monotone decreasing if for all i <= j, nums[i] >= nums[j].

// Given an integer array nums, return true if the given array is monotonic, or false otherwise.

// Example 1:
// Input: nums = [1,2,2,3]
// Output: true

// Example 2:
// Input: nums = [6,5,4,4]
// Output: true

// Example 3:
// Input: nums = [1,3,2]
// Output: false
 
// Constraints:
//     1 <= nums.length <= 10^5
//     -10^5 <= nums[i] <= 10^5

import "fmt"

func isMonotonic(nums []int) bool {
    if len(nums) < 2 {
        return true
    }
    direction := 0  // 0 means unknown, 1 means increasing, -1 means decreasing
    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[i-1] {  // increasing
            if direction == 0 {
                direction = 1
            } else if direction == -1 { // 递增情况出现 下降 说明没有单调性直接返回 false
                return false
            }
        } else if nums[i] < nums[i-1] {  // decreasing
            if direction == 0 {
                direction = -1
            } else if direction == 1 {  // 递减情况出现 上升 说明没有单调性直接返回 false
                return false
            }
        }
    }
    return true
}

func isMonotonic1(nums []int) bool {
    increasing, decreasing := true, true
    for i := 0; i < len(nums) -1; i++{
        if nums[i] < nums[i+1] { // 如果 小于后面说明不为 递减 decreasing = false
            decreasing = false
        } else if nums[i] > nums[i+1] { // 如果 大于后面说明不为 递增 increasing = false
            increasing = false
        }
    }
    return decreasing || increasing
}

func main() {
    fmt.Println(isMonotonic([]int{1,2,2,3})) // true
    fmt.Println(isMonotonic([]int{6,5,4,4})) // true
    fmt.Println(isMonotonic([]int{1,3,2})) // false

    fmt.Println(isMonotonic1([]int{1,2,2,3})) // true
    fmt.Println(isMonotonic1([]int{6,5,4,4})) // true
    fmt.Println(isMonotonic1([]int{1,3,2})) // false
}
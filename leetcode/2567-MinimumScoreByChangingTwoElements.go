package main

// 2567. Minimum Score by Changing Two Elements
// You are given an integer array nums.
//     1. The low score of nums is the minimum absolute difference between any two integers.
//     2. The high score of nums is the maximum absolute difference between any two integers.
//     3. The score of nums is the sum of the high and low scores.

// Return the minimum score after changing two elements of nums.

// Example 1:
// Input: nums = [1,4,7,8,5]
// Output: 3
// Explanation:
// Change nums[0] and nums[1] to be 6 so that nums becomes [6,6,7,8,5].
// The low score is the minimum absolute difference: |6 - 6| = 0.
// The high score is the maximum absolute difference: |8 - 5| = 3.
// The sum of high and low score is 3.

// Example 2:
// Input: nums = [1,4,3]
// Output: 0
// Explanation:
// Change nums[1] and nums[2] to 1 so that nums becomes [1,1,1].
// The sum of maximum absolute difference and minimum absolute difference is 0.

// Constraints:
//     3 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9

import "fmt"
import "sort"

func minimizeSum(nums []int) int {
    sort.Ints(nums)
    n := len(nums)
    res := nums[n - 3] - nums[0]
    if nums[n - 2] - nums[1] < res {
        res = nums[n-2] - nums[1]
    }
    if nums[n - 1] - nums[2] < res {
        res = nums[n-1] - nums[2]
    }
    return res
}

func minimizeSum1(nums []int) int {
    mx1, mx2, mx3, mn1, mn2, mn3 := -1 << 31, -1 << 31, -1 << 31, 1 << 31, 1 << 31, 1 << 31
    for _, v := range nums {
        if v > mx1 {
            mx3 = mx2 
            mx2 = mx1
            mx1 = v
        } else if v > mx2 {
            mx3 = mx2
            mx2 = v
        } else if v > mx3 {
            mx3 = v
        }
        if v < mn1 {
            mn3 = mn2
            mn2 = mn1
            mn1 = v
        } else if v < mn2 {
            mn3 = mn2
            mn2 = v
        } else if v < mn3 {
            mn3 = v
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    return min(min(mx1 - mn3, mx3 - mn1), mx2 - mn2)
}

func main() {
    // Example 1:
    // Input: nums = [1,4,7,8,5]
    // Output: 3
    // Explanation:
    // Change nums[0] and nums[1] to be 6 so that nums becomes [6,6,7,8,5].
    // The low score is the minimum absolute difference: |6 - 6| = 0.
    // The high score is the maximum absolute difference: |8 - 5| = 3.
    // The sum of high and low score is 3.
    fmt.Println(minimizeSum([]int{1,4,7,8,5})) // 3
    // Example 2:
    // Input: nums = [1,4,3]
    // Output: 0
    // Explanation:
    // Change nums[1] and nums[2] to 1 so that nums becomes [1,1,1].
    // The sum of maximum absolute difference and minimum absolute difference is 0.
    fmt.Println(minimizeSum([]int{1,4,3})) // 0

    fmt.Println(minimizeSum([]int{1,2,3,4,5,6,7,8,9})) // 6
    fmt.Println(minimizeSum([]int{9,8,7,6,5,4,3,2,1})) // 6

    fmt.Println(minimizeSum1([]int{1,4,7,8,5})) // 3
    fmt.Println(minimizeSum1([]int{1,4,3})) // 0
    fmt.Println(minimizeSum1([]int{1,2,3,4,5,6,7,8,9})) // 6
    fmt.Println(minimizeSum1([]int{9,8,7,6,5,4,3,2,1})) // 6
}
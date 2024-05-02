package main

// 2441. Largest Positive Integer That Exists With Its Negative
// Given an integer array nums that does not contain any zeros, find the largest positive integer k such that -k also exists in the array.
// Return the positive integer k. If there is no such integer, return -1.

// Example 1:
// Input: nums = [-1,2,-3,3]
// Output: 3
// Explanation: 3 is the only valid k we can find in the array.

// Example 2:
// Input: nums = [-1,10,6,7,-7,1]
// Output: 7
// Explanation: Both 1 and 7 have their corresponding negative values in the array. 7 has a larger value.

// Example 3:
// Input: nums = [-10,8,6,7,-2,-3]
// Output: -1
// Explanation: There is no a single valid k, we return -1.
 
// Constraints:
//     1 <= nums.length <= 1000
//     -1000 <= nums[i] <= 1000
//     nums[i] != 0

import "fmt"
import "sort"

func findMaxK(nums []int) int {
    sort.Ints(nums)
    l, r := 0, len(nums) - 1
    for l < r {
        t := -1 * nums[l] - nums[r]
        if t == 0 {
            return nums[r]
        } else if t > 0 {
            l++
        } else {
            r--
        }
    }
    return -1
}

func main() {
    // Example 1:
    // Input: nums = [-1,2,-3,3]
    // Output: 3
    // Explanation: 3 is the only valid k we can find in the array.
    fmt.Println(findMaxK([]int{-1,2,-3,3})) // 3
    // Example 2:
    // Input: nums = [-1,10,6,7,-7,1]
    // Output: 7
    // Explanation: Both 1 and 7 have their corresponding negative values in the array. 7 has a larger value.
    fmt.Println(findMaxK([]int{-1,10,6,7,-7,1})) // 7
    // Example 3:
    // Input: nums = [-10,8,6,7,-2,-3]
    // Output: -1
    // Explanation: There is no a single valid k, we return -1.
    fmt.Println(findMaxK([]int{-10,8,6,7,-2,-3})) // -1
}
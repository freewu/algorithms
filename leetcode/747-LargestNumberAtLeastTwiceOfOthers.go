package main

// 747. Largest Number At Least Twice of Others
// You are given an integer array nums where the largest integer is unique.
// Determine whether the largest element in the array is at least twice as much as every other number in the array. 
// If it is, return the index of the largest element, or return -1 otherwise.

// Example 1:
// Input: nums = [3,6,1,0]
// Output: 1
// Explanation: 6 is the largest integer.
// For every other number in the array x, 6 is at least twice as big as x.
// The index of value 6 is 1, so we return 1.

// Example 2:
// Input: nums = [1,2,3,4]
// Output: -1
// Explanation: 4 is less than twice the value of 3, so we return -1.

// Constraints:
//     2 <= nums.length <= 50
//     0 <= nums[i] <= 100
//     The largest element in nums is unique.

import "fmt"

func dominantIndex(nums []int) int {
    res, mx := -1, 0
    for i, v := range nums {
        if v > mx { // 找到最大值
            mx = v
            res = i
        } 
    }
    for i, v := range nums { // 检查最大值是否都是大其它值的2倍
        if res != i {
            if mx < 2 * v {
                return -1
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [3,6,1,0]
    // Output: 1
    // Explanation: 6 is the largest integer.
    // For every other number in the array x, 6 is at least twice as big as x.
    // The index of value 6 is 1, so we return 1.
    fmt.Println(dominantIndex([]int{3,6,1,0})) // 1
    // Example 2:
    // Input: nums = [1,2,3,4]
    // Output: -1
    // Explanation: 4 is less than twice the value of 3, so we return -1.
    fmt.Println(dominantIndex([]int{1,2,3,4})) // -1
}
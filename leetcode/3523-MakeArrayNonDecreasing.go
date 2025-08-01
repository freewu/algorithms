package main

// 3523. Make Array Non-decreasing
// You are given an integer array nums. 
// In one operation, you can select a subarray and replace it with a single element equal to its maximum value.

// Return the maximum possible size of the array after performing zero 
// or more operations such that the resulting array is non-decreasing.

// Example 1:
// Input: nums = [4,2,5,3,5]
// Output: 3
// Explanation:
// One way to achieve the maximum size is:
// Replace subarray nums[1..2] = [2, 5] with 5 → [4, 5, 3, 5].
// Replace subarray nums[2..3] = [3, 5] with 5 → [4, 5, 5].
// The final array [4, 5, 5] is non-decreasing with size 3.

// Example 2:
// Input: nums = [1,2,3]
// Output: 3
// Explanation:
// No operation is needed as the array [1,2,3] is already non-decreasing.

// Constraints:
//     1 <= nums.length <= 2 * 10^5
//     1 <= nums[i] <= 2 * 10^5

import "fmt"

func maximumPossibleSize(nums []int) int {
    res, prev := 0, -1
    for _, v := range nums {
        if v >= prev { 
            prev = v
            res++
        } 
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [4,2,5,3,5]
    // Output: 3
    // Explanation:
    // One way to achieve the maximum size is:
    // Replace subarray nums[1..2] = [2, 5] with 5 → [4, 5, 3, 5].
    // Replace subarray nums[2..3] = [3, 5] with 5 → [4, 5, 5].
    // The final array [4, 5, 5] is non-decreasing with size 3.
    fmt.Println(maximumPossibleSize([]int{4,2,5,3,5})) // 3
    // Example 2:
    // Input: nums = [1,2,3]
    // Output: 3
    // Explanation:
    // No operation is needed as the array [1,2,3] is already non-decreasing.
    fmt.Println(maximumPossibleSize([]int{1,2,3})) // 3

    fmt.Println(maximumPossibleSize([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(maximumPossibleSize([]int{9,8,7,6,5,4,3,2,1})) // 1
}
package main

// 3353. Minimum Total Operations
// Given an array of integers nums, you can perform any number of operations on this array.

// In each operation, you can:
//     1. Choose a prefix of the array.
//     2. Choose an integer k (which can be negative) and add k to each element in the chosen prefix.

// A prefix of an array is a subarray that starts from the beginning of the array and extends to any point within it.

// Return the minimum number of operations required to make all elements in arr equal.

// Example 1:
// Input: nums = [1,4,2]
// Output: 2
// Explanation:
// Operation 1: Choose the prefix [1, 4] of length 2 and add -2 to each element of the prefix. The array becomes [-1, 2, 2].
// Operation 2: Choose the prefix [-1] of length 1 and add 3 to it. The array becomes [2, 2, 2].
// Thus, the minimum number of required operations is 2.

// Example 2:
// Input: nums = [10,10,10]
// Output: 0
// Explanation:
// All elements are already equal, so no operations are needed.

// Constraints:
//     1 <= nums.length <= 10^5
//     -10^9 <= nums[i] <= 10^9

import "fmt"

func minOperations(nums []int) int {
    n := len(nums)
    last, mn, res := nums[n - 1], 0, 0
    for i := n - 2; i >= 0; i-- {
        if nums[i] + mn != last {
            res++
            mn = last - nums[i]
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,4,2]
    // Output: 2
    // Explanation:
    // Operation 1: Choose the prefix [1, 4] of length 2 and add -2 to each element of the prefix. The array becomes [-1, 2, 2].
    // Operation 2: Choose the prefix [-1] of length 1 and add 3 to it. The array becomes [2, 2, 2].
    // Thus, the minimum number of required operations is 2.
    fmt.Println(minOperations([]int{1,4,2})) // 2
    // Example 2:
    // Input: nums = [10,10,10]
    // Output: 0
    // Explanation:
    // All elements are already equal, so no operations are needed.
    fmt.Println(minOperations([]int{10,10,10})) // 0
}
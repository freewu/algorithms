package main

// 3674. Minimum Operations to Equalize Array
// You are given an integer array nums of length n.

// In one operation, choose any subarray nums[l...r] (0 <= l <= r < n) and replace each element in that subarray with the bitwise AND of all elements.

// Return the minimum number of operations required to make all elements of nums equal.

// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:
// Input: nums = [1,2]
// Output: 1
// Explanation:
// Choose nums[0...1]: (1 AND 2) = 0, so the array becomes [0, 0] and all elements are equal in 1 operation.

// Example 2:
// Input: nums = [5,5,5]
// Output: 0
// Explanation:
// nums is [5, 5, 5] which already has all elements equal, so 0 operations are required.

// Constraints:
//     1 <= n == nums.length <= 100
//     1 <= nums[i] <= 10^5

import "fmt"

func minOperations(nums []int) int {
    n := len(nums)
    for i := 1; i < n; i++ {
        if nums[i] != nums[i-1] {
            return 1
        }
    }
    return 0
}

func main() {
    // Example 1:
    // Input: nums = [1,2]
    // Output: 1
    // Explanation:
    // Choose nums[0...1]: (1 AND 2) = 0, so the array becomes [0, 0] and all elements are equal in 1 operation.
    fmt.Println(minOperations([]int{1,2})) // 1
    // Example 2:
    // Input: nums = [5,5,5]
    // Output: 0
    // Explanation:
    // nums is [5, 5, 5] which already has all elements equal, so 0 operations are required.
    fmt.Println(minOperations([]int{5,5,5})) // 0

    fmt.Println(minOperations([]int{1,2,3,4,5,6,7,8,9})) // 1
    fmt.Println(minOperations([]int{9,8,7,6,5,4,3,2,1})) // 1
}
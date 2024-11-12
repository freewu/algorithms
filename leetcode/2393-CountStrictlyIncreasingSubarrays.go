package main

// 2393. Count Strictly Increasing Subarrays
// You are given an array nums consisting of positive integers.

// Return the number of subarrays of nums that are in strictly increasing order.

// A subarray is a contiguous part of an array.

// Example 1:
// Input: nums = [1,3,5,4,4,6]
// Output: 10
// Explanation: The strictly increasing subarrays are the following:
// - Subarrays of length 1: [1], [3], [5], [4], [4], [6].
// - Subarrays of length 2: [1,3], [3,5], [4,6].
// - Subarrays of length 3: [1,3,5].
// The total number of subarrays is 6 + 3 + 1 = 10.

// Example 2:
// Input: nums = [1,2,3,4,5]
// Output: 15
// Explanation: Every subarray is strictly increasing. There are 15 possible subarrays that we can take.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^6

import "fmt"

func countSubarrays(nums []int) int64 {
    res, count := 0, 1
    for i := 1; i < len(nums); i++ {
        if nums[i] <= nums[i - 1] {
            res += count * (count + 1) / 2
            count = 1
        } else {
            count++
        }
    }
    res += (count * (count + 1) / 2)
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [1,3,5,4,4,6]
    // Output: 10
    // Explanation: The strictly increasing subarrays are the following:
    // - Subarrays of length 1: [1], [3], [5], [4], [4], [6].
    // - Subarrays of length 2: [1,3], [3,5], [4,6].
    // - Subarrays of length 3: [1,3,5].
    // The total number of subarrays is 6 + 3 + 1 = 10.
    fmt.Println(countSubarrays([]int{1,3,5,4,4,6})) // 10
    // Example 2:
    // Input: nums = [1,2,3,4,5]
    // Output: 15
    // Explanation: Every subarray is strictly increasing. There are 15 possible subarrays that we can take.
    fmt.Println(countSubarrays([]int{1,2,3,4,5})) // 15
}
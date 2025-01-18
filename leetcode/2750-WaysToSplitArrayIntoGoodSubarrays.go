package main

// 2750. Ways to Split Array Into Good Subarrays
// You are given a binary array nums.

// A subarray of an array is good if it contains exactly one element with the value 1.

// Return an integer denoting the number of ways to split the array nums into good subarrays. 
// As the number may be too large, return it modulo 10^9 + 7.

// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:
// Input: nums = [0,1,0,0,1]
// Output: 3
// Explanation: There are 3 ways to split nums into good subarrays:
// - [0,1] [0,0,1]
// - [0,1,0] [0,1]
// - [0,1,0,0] [1]

// Example 2:
// Input: nums = [0,1,0]
// Output: 1
// Explanation: There is 1 way to split nums into good subarrays:
// - [0,1,0]

// Constraints:
//     1 <= nums.length <= 10^5
//     0 <= nums[i] <= 1

import "fmt"

func numberOfGoodSubarraySplits(nums []int) int {
    res, last1 := 1, -1
    for i, v := range nums {
        if v == 1 {
            if last1 == -1 {
                last1 = i
                continue
            }
            count0 := i - last1
            res, last1 = (res * count0) % 1_000_000_007, i
        }
    }
    if last1 == -1 { return 0 } // 全是 0 的情况
    return res
}

func main() {
    // Example 1:
    // Input: nums = [0,1,0,0,1]
    // Output: 3
    // Explanation: There are 3 ways to split nums into good subarrays:
    // - [0,1] [0,0,1]
    // - [0,1,0] [0,1]
    // - [0,1,0,0] [1]
    fmt.Println(numberOfGoodSubarraySplits([]int{0,1,0,0,1})) // 3
    // Example 2:
    // Input: nums = [0,1,0]
    // Output: 1
    // Explanation: There is 1 way to split nums into good subarrays:
    // - [0,1,0]
    fmt.Println(numberOfGoodSubarraySplits([]int{0,1,0})) // 1

    fmt.Println(numberOfGoodSubarraySplits([]int{0,0,0})) // 0
    fmt.Println(numberOfGoodSubarraySplits([]int{1,1,1})) // 1
}
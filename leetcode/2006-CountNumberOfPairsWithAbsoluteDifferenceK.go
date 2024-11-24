package main

// 2006. Count Number of Pairs With Absolute Difference K
// Given an integer array nums and an integer k, 
// return the number of pairs (i, j) where i < j such that |nums[i] - nums[j]| == k.

// The value of |x| is defined as:
//     x if x >= 0.
//     -x if x < 0.

// Example 1:
// Input: nums = [1,2,2,1], k = 1
// Output: 4
// Explanation: The pairs with an absolute difference of 1 are:
// - [1,2,2,1]
// - [1,2,2,1]
// - [1,2,2,1]
// - [1,2,2,1]

// Example 2:
// Input: nums = [1,3], k = 3
// Output: 0
// Explanation: There are no pairs with an absolute difference of 3.

// Example 3:
// Input: nums = [3,2,1,5,4], k = 2
// Output: 3
// Explanation: The pairs with an absolute difference of 2 are:
// - [3,2,1,5,4]
// - [3,2,1,5,4]
// - [3,2,1,5,4]

// Constraints:
//     1 <= nums.length <= 200
//     1 <= nums[i] <= 100
//     1 <= k <= 99

import "fmt"

func countKDifference(nums []int, k int) int {
    res, n := 0, len(nums)
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i := 0; i < n ; i++ {
        for j := i + 1; j < n ; j++ {
            if abs(nums[i] - nums[j]) == k {
                res++
            }
        }
    }
    return res
}

func countKDifference1(nums []int, k int) int {
    mp := make(map[int]int)
    res := 0
    for _, v := range nums {
        res += mp[v- k] + mp[v +k ]
        mp[v]++
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,2,1], k = 1
    // Output: 4
    // Explanation: The pairs with an absolute difference of 1 are:
    // - [1,2,2,1]
    // - [1,2,2,1]
    // - [1,2,2,1]
    // - [1,2,2,1]
    fmt.Println(countKDifference([]int{1,2,2,1}, 1)) // 4
    // Example 2:
    // Input: nums = [1,3], k = 3
    // Output: 0
    // Explanation: There are no pairs with an absolute difference of 3.
    fmt.Println(countKDifference([]int{1,3}, 3)) // 0
    // Example 3:
    // Input: nums = [3,2,1,5,4], k = 2
    // Output: 3
    // Explanation: The pairs with an absolute difference of 2 are:
    // - [3,2,1,5,4]
    // - [3,2,1,5,4]
    // - [3,2,1,5,4]
    fmt.Println(countKDifference([]int{3,2,1,5,4}, 2)) // 3

    fmt.Println(countKDifference1([]int{1,2,2,1}, 1)) // 4
    fmt.Println(countKDifference1([]int{1,3}, 3)) // 0
    fmt.Println(countKDifference1([]int{3,2,1,5,4}, 2)) // 3
}
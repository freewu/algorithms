package main

// 3202. Find the Maximum Length of Valid Subsequence II
// You are given an integer array nums and a positive integer k.
// A subsequence sub of nums with length x is called valid if it satisfies:
//     (sub[0] + sub[1]) % k == (sub[1] + sub[2]) % k == ... == (sub[x - 2] + sub[x - 1]) % k.

// Return the length of the longest valid subsequence of nums.

// Example 1:
// Input: nums = [1,2,3,4,5], k = 2
// Output: 5
// Explanation:
// The longest valid subsequence is [1, 2, 3, 4, 5].

// Example 2:
// Input: nums = [1,4,2,3,1,4], k = 3
// Output: 4
// Explanation:
// The longest valid subsequence is [1, 4, 1, 4].

// Constraints:
//     2 <= nums.length <= 10^3
//     1 <= nums[i] <= 10^7
//     1 <= k <= 10^3

import "fmt"

func maximumLength(nums []int, k int) int {
    res, n := 0, len(nums)
    dp := make([][]int, k)
    for i := range dp { dp[i] = make([]int, k) }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n; i++ {
        rem := nums[i] % k
        for j := 0; j <k ; j++ {
            dp[rem][j] = max(dp[rem][j], 1 + dp[j][rem])
            res = max(res, dp[rem][j])
        }
    }
    return res
}

func maximumLength1(nums []int, k int) int {
    res := 0
    for i := range nums {
        nums[i] = nums[i] % k
    }
    for i := 0; i < k; i++ {
        arr := make([]int, k)
        for _, v := range nums {
            if i >= v {
                arr[v] = arr[i - v] + 1
            } else {
                arr[v] = arr[k + i - v] + 1
            }
        }
        for _, v := range arr {
            if v > res {
                res = v
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4,5], k = 2
    // Output: 5
    // Explanation:
    // The longest valid subsequence is [1, 2, 3, 4, 5].
    fmt.Println(maximumLength([]int{1,2,3,4,5}, 2)) // 5
    // Example 2:
    // Input: nums = [1,4,2,3,1,4], k = 3
    // Output: 4
    // Explanation:
    // The longest valid subsequence is [1, 4, 1, 4].
    fmt.Println(maximumLength([]int{1,4,2,3,1,4}, 2)) // 4

    fmt.Println(maximumLength1([]int{1,2,3,4,5}, 2)) // 5
    fmt.Println(maximumLength1([]int{1,4,2,3,1,4}, 2)) // 4
}
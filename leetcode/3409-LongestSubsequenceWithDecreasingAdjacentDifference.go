package main

// 3409. Longest Subsequence With Decreasing Adjacent Difference
// You are given an array of integers nums.

// Your task is to find the length of the longest subsequence seq of nums, such that the absolute differences between consecutive elements form a non-increasing sequence of integers. 
// In other words, for a subsequence seq0, seq1, seq2, ..., seqm of nums, |seq1 - seq0| >= |seq2 - seq1| >= ... >= |seqm - seqm - 1|.

// Return the length of such a subsequence.

// Example 1:
// Input: nums = [16,6,3]
// Output: 3
// Explanation: 
// The longest subsequence is [16, 6, 3] with the absolute adjacent differences [10, 3].

// Example 2:
// Input: nums = [6,5,3,4,2,1]
// Output: 4
// Explanation:
// The longest subsequence is [6, 4, 2, 1] with the absolute adjacent differences [2, 2, 1].

// Example 3:
// Input: nums = [10,20,10,19,10,20]
// Output: 5
// Explanation: 
// The longest subsequence is [10, 20, 10, 19, 10] with the absolute adjacent differences [10, 10, 9, 9].

// Constraints:
//     2 <= nums.length <= 10^4
//     1 <= nums[i] <= 300

import "fmt"
import "slices"

func longestSubsequence(nums []int) int {
    res, dp, arr := 0, make([][]int, 301), make([]bool, 301)
    for i := range dp {
        dp[i] = make([]int, 301)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for _, v := range nums {
        for i := 300; i >= 0; i-- {
            if arr[i] {
                diff := abs(v - i)
                dp[v][diff] = max(dp[v][diff], dp[i][diff] + 1)
            }
        }
        for i := 299; i >= 0; i-- {
            dp[v][i] = max(dp[v][i], dp[v][i + 1])
        }
        res = max(res, dp[v][0])
        arr[v] = true
    }
    return res + 1
}

func longestSubsequence1(nums []int) int {
    res, mx := 0, slices.Max(nums)
    diff := mx - slices.Min(nums)
    dp := make([][]int, mx + 1)
    for i := range dp {
        dp[i] = make([]int, diff + 1)
    }
    for _, v := range nums {
        val := 1
        for j := diff; j >= 0; j-- {
            if v - j >= 0 {
                val = max(val, dp[v - j][j] + 1)
            }
            if v + j <= mx {
                val = max(val, dp[v + j][j] + 1)
            }
            dp[v][j] = val
            res = max(res, val)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [16,6,3]
    // Output: 3
    // Explanation: 
    // The longest subsequence is [16, 6, 3] with the absolute adjacent differences [10, 3].
    fmt.Println(longestSubsequence([]int{16,6,3})) // 3
    // Example 2:
    // Input: nums = [6,5,3,4,2,1]
    // Output: 4
    // Explanation:
    // The longest subsequence is [6, 4, 2, 1] with the absolute adjacent differences [2, 2, 1].
    fmt.Println(longestSubsequence([]int{6,5,3,4,2,1})) // 4
    // Example 3:
    // Input: nums = [10,20,10,19,10,20]
    // Output: 5
    // Explanation: 
    // The longest subsequence is [10, 20, 10, 19, 10] with the absolute adjacent differences [10, 10, 9, 9].
    fmt.Println(longestSubsequence([]int{10,20,10,19,10,20})) // 5

    fmt.Println(longestSubsequence([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(longestSubsequence([]int{9,8,7,6,5,4,3,2,1})) // 9

    fmt.Println(longestSubsequence1([]int{16,6,3})) // 3
    fmt.Println(longestSubsequence1([]int{6,5,3,4,2,1})) // 4
    fmt.Println(longestSubsequence1([]int{10,20,10,19,10,20})) // 5
    fmt.Println(longestSubsequence1([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(longestSubsequence1([]int{9,8,7,6,5,4,3,2,1})) // 9
}
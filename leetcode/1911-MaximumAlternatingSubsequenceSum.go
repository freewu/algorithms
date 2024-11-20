package main

// 1911. Maximum Alternating Subsequence Sum
// The alternating sum of a 0-indexed array is defined as the sum of the elements at even indices minus the sum of the elements at odd indices.
//     For example, the alternating sum of [4,2,5,3] is (4 + 5) - (2 + 3) = 4.

// Given an array nums, return the maximum alternating sum of any subsequence of nums 
// (after reindexing the elements of the subsequence).

// A subsequence of an array is a new array generated from the original array by deleting some elements (possibly none) without changing the remaining elements' relative order. 
// For example, [2,7,4] is a subsequence of [4,2,3,7,2,1,4] (the underlined elements), while [2,4,2] is not.

// Example 1:
// Input: nums = [4,2,5,3]
// Output: 7
// Explanation: It is optimal to choose the subsequence [4,2,5] with alternating sum (4 + 5) - 2 = 7.

// Example 2:
// Input: nums = [5,6,7,8]
// Output: 8
// Explanation: It is optimal to choose the subsequence [8] with alternating sum 8.

// Example 3:
// Input: nums = [6,2,1,2,4,5]
// Output: 10
// Explanation: It is optimal to choose the subsequence [6,1,5] with alternating sum (6 + 5) - 1 = 10.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5

import "fmt"

func maxAlternatingSum(nums []int) int64 {
    res, n := nums[0], len(nums)
    even, odd := []int{nums[0]}, []int{}
    for i := 1; i < n; i++ {
        if nums[i] >= even[len(even)-1] {
            res += (nums[i] - even[len(even) - 1])
            even[len(even) - 1] = nums[i]
            continue
        }
        if i != n - 1 && nums[i + 1] - nums[i] > 0 {
            even = append(even, nums[i + 1])
            odd = append(odd, nums[i])
            res += (nums[i + 1] - nums[i])
        }
    }
    return int64(res)
}

func maxAlternatingSum1(nums []int) int64 {
    // 同时算出奇数与偶数的最大交替和
    // 可以定义dp[i][0]表示前i个数中为偶数的最大交替和，dp[i][1]表示前i个数中为奇数的最大交替和
    dp := [2]int{ 0, -1 << 31 }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range nums {
        dp = [2]int{ max(dp[0], dp[1] - v), max(dp[1], dp[0] + v)}
    }
    return int64(dp[1])
}

func main() {
    // Example 1:
    // Input: nums = [4,2,5,3]
    // Output: 7
    // Explanation: It is optimal to choose the subsequence [4,2,5] with alternating sum (4 + 5) - 2 = 7.
    fmt.Println(maxAlternatingSum([]int{4,2,5,3})) // 7
    // Example 2:
    // Input: nums = [5,6,7,8]
    // Output: 8
    // Explanation: It is optimal to choose the subsequence [8] with alternating sum 8.
    fmt.Println(maxAlternatingSum([]int{5,6,7,8})) // 8
    // Example 3:
    // Input: nums = [6,2,1,2,4,5]
    // Output: 10
    // Explanation: It is optimal to choose the subsequence [6,1,5] with alternating sum (6 + 5) - 1 = 10.
    fmt.Println(maxAlternatingSum([]int{6,2,1,2,4,5})) // 10

    fmt.Println(maxAlternatingSum1([]int{4,2,5,3})) // 7
    fmt.Println(maxAlternatingSum1([]int{5,6,7,8})) // 8
    fmt.Println(maxAlternatingSum1([]int{6,2,1,2,4,5})) // 10
}
package main

// 3077. Maximum Strength of K Disjoint Subarrays
// You are given an array of integers nums with length n, and a positive odd integer k.

// Select exactly k disjoint subarrays sub1, sub2, ..., 
// subk from nums such that the last element of subi appears before the first element of sub{i+1} for all 1 <= i <= k-1. 
// The goal is to maximize their combined strength.

// The strength of the selected subarrays is defined as:
//     strength = k * sum(sub1)- (k - 1) * sum(sub2) + (k - 2) * sum(sub3) - ... - 2 * sum(sub{k-1}) + sum(subk)

// where sum(subi) is the sum of the elements in the i-th subarray.

// Return the maximum possible strength that can be obtained from selecting exactly k disjoint subarrays from nums.

// Note that the chosen subarrays don't need to cover the entire array.

// Example 1:
// Input: nums = [1,2,3,-1,2], k = 3
// Output: 22
// Explanation:
// The best possible way to select 3 subarrays is: nums[0..2], nums[3..3], and nums[4..4]. The strength is calculated as follows:
// strength = 3 * (1 + 2 + 3) - 2 * (-1) + 2 = 22

// Example 2:
// Input: nums = [12,-2,-2,-2,-2], k = 5
// Output: 64
// Explanation:
// The only possible way to select 5 disjoint subarrays is: nums[0..0], nums[1..1], nums[2..2], nums[3..3], and nums[4..4]. The strength is calculated as follows:
// strength = 5 * 12 - 4 * (-2) + 3 * (-2) - 2 * (-2) + (-2) = 64

// Example 3:
// Input: nums = [-1,-2,-3], k = 1
// Output: -1
// Explanation:
// The best possible way to select 1 subarray is: nums[0..0]. The strength is -1.

// Constraints:
//     1 <= n <= 10^4
//     -10^9 <= nums[i] <= 10^9
//     1 <= k <= n
//     1 <= n * k <= 10^6
//     k is odd.

import "fmt"

func maximumStrength(nums []int, k int) int64 {
    n := len(nums)
    dp := make([]int64, n + 1)
    max := func(x, y int64) int64 { if x > y { return x; }; return y; }
    for j := 0; j < k; j++ {
        mul := int64(k - j)
        if j % 2 == 1 {
            mul = -mul
        }
        tmp := make([]int64, n + 1)
        tmp[0] = -1 << 61
        for i := 1; i <= n; i++ {
            tmp[i] = max(tmp[i - 1], dp[i - 1]) + int64(nums[i - 1]) * mul
        }
        copy(dp, tmp)
        for i := 1; i <= n; i++ {
            dp[i] = max(dp[i], dp[i - 1])
        }
    }
    return dp[n]
}

func maximumStrength1(nums []int, k int) int64 {
    n := len(nums)
    prefix := make([]int, n + 1)
    for i, v := range nums {
        prefix[i + 1] = prefix[i] + v
    }
    dp := make([][]int, k + 1)
    for i := range dp {
        dp[i] = make([]int, n + 1)
        for j := range dp[i] {
            dp[i][j] = -1 << 61
        }
    }
    for i := range dp[0] {
        dp[0][i] = 0
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    flag := 1
    for i := 0; i < k; i++ {
        suf, coeff := -1 << 61, (i + 1) * flag
        for j := n - 1; j >= 0; j-- {
            suf = max(suf, dp[i][j + 1] + prefix[j + 1] * coeff)
            dp[i + 1][j] = suf - prefix[j] * coeff
            dp[i + 1][j] = max(dp[i + 1][j], dp[i + 1][j + 1])
        }
        flag *= -1
    }
    return int64(dp[k][0])
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,-1,2], k = 3
    // Output: 22
    // Explanation:
    // The best possible way to select 3 subarrays is: nums[0..2], nums[3..3], and nums[4..4]. The strength is calculated as follows:
    // strength = 3 * (1 + 2 + 3) - 2 * (-1) + 2 = 22
    fmt.Println(maximumStrength([]int{1,2,3,-1,2}, 3)) // 22
    // Example 2:
    // Input: nums = [12,-2,-2,-2,-2], k = 5
    // Output: 64
    // Explanation:
    // The only possible way to select 5 disjoint subarrays is: nums[0..0], nums[1..1], nums[2..2], nums[3..3], and nums[4..4]. The strength is calculated as follows:
    // strength = 5 * 12 - 4 * (-2) + 3 * (-2) - 2 * (-2) + (-2) = 64
    fmt.Println(maximumStrength([]int{12,-2,-2,-2,-2}, 5)) // 64
    // Example 3:
    // Input: nums = [-1,-2,-3], k = 1
    // Output: -1
    // Explanation:
    // The best possible way to select 1 subarray is: nums[0..0]. The strength is -1.
    fmt.Println(maximumStrength([]int{-1,-2,-3}, 1)) // -1

    fmt.Println(maximumStrength([]int{1,2,3,4,5,6,7,8,9}, 1)) // 45
    fmt.Println(maximumStrength([]int{9,8,7,6,5,4,3,2,1}, 1)) // 45
    fmt.Println(maximumStrength([]int{-1000000000,-100000000,-10000000,123,234}, 5)) // -4630000012

    fmt.Println(maximumStrength1([]int{1,2,3,-1,2}, 3)) // 22
    fmt.Println(maximumStrength1([]int{12,-2,-2,-2,-2}, 5)) // 64
    fmt.Println(maximumStrength1([]int{-1,-2,-3}, 1)) // -1
    fmt.Println(maximumStrength1([]int{1,2,3,4,5,6,7,8,9}, 1)) // 45
    fmt.Println(maximumStrength1([]int{9,8,7,6,5,4,3,2,1}, 1)) // 45
    fmt.Println(maximumStrength1([]int{-1000000000,-100000000,-10000000,123,234}, 5)) // -4630000012
}
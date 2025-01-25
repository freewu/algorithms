package main

// 3351. Sum of Good Subsequences
// You are given an integer array nums. 
// A good subsequence is defined as a subsequence of nums where the absolute difference between any two consecutive elements in the subsequence is exactly 1.

// Return the sum of all possible good subsequences of nums.

// Since the answer may be very large, return it modulo 10^9 + 7.

// Note that a subsequence of size 1 is considered good by definition.

// Example 1:
// Input: nums = [1,2,1]
// Output: 14
// Explanation:
// Good subsequences are: [1], [2], [1], [1,2], [2,1], [1,2,1].
// The sum of elements in these subsequences is 14.

// Example 2:
// Input: nums = [3,4,5]
// Output: 40
// Explanation:
// Good subsequences are: [3], [4], [5], [3,4], [4,5], [3,4,5].
// The sum of elements in these subsequences is 40.

// Constraints:
//     1 <= nums.length <= 10^5
//     0 <= nums[i] <= 10^5

import "fmt"
import "slices"

func sumOfGoodSubsequences(nums []int) int {
    res, mod := 0, 1_000_000_007
    dp := make([][2]int, slices.Max(nums) + 2)
    for _, v := range nums {
        sum := v
        dp[v][0] += 1
        if v - 1 >= 0 {
            sum = (sum + dp[v - 1][1] + (v * dp[v - 1][0]) % mod) % mod
            dp[v][0] = (dp[v][0] + dp[v - 1][0]) % mod
        }
        sum = (sum + dp[v + 1][1] + v * dp[v + 1][0] % mod) % mod
        dp[v][0] = (dp[v][0] + dp[v + 1][0]) % mod
        dp[v][1] = (dp[v][1] + sum) % mod
    }
    for i := range dp {
        res = (res + dp[i][1]) % mod
    }
    return res
}

func sumOfGoodSubsequences1(nums []int) int {
    res, mod, mx := 0, 1_000_000_007, slices.Max(nums)
    c, s := make([]int, mx + 2), make([]int, mx + 2)
    for _, v := range nums {
        if v > 0 {
            c[v] = (c[v] + c[v-1] + c[v+1] + 1) % mod
            s[v] = (s[v] + (c[v-1]+c[v+1])*v % mod + s[v-1] + s[v+1] + v) % mod
        } else {
            c[v] = (c[v] + c[v+1] + 1) % mod
            s[v] = (s[v] + c[v+1]*v%mod + s[v+1] + v) % mod
        }
    }
    for _, v := range s {
        res = (res + v) % mod
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,1]
    // Output: 14
    // Explanation:
    // Good subsequences are: [1], [2], [1], [1,2], [2,1], [1,2,1].
    // The sum of elements in these subsequences is 14.
    fmt.Println(sumOfGoodSubsequences([]int{1,2,1})) // 14
    // Example 2:
    // Input: nums = [3,4,5]
    // Output: 40
    // Explanation:
    // Good subsequences are: [3], [4], [5], [3,4], [4,5], [3,4,5].
    // The sum of elements in these subsequences is 40.
    fmt.Println(sumOfGoodSubsequences([]int{3,4,5})) // 40

    fmt.Println(sumOfGoodSubsequences([]int{1,2,3,4,5,6,7,8,9})) // 825
    fmt.Println(sumOfGoodSubsequences([]int{9,8,7,6,5,4,3,2,1})) // 825

    fmt.Println(sumOfGoodSubsequences1([]int{1,2,1})) // 14
    fmt.Println(sumOfGoodSubsequences1([]int{3,4,5})) // 40
    fmt.Println(sumOfGoodSubsequences1([]int{1,2,3,4,5,6,7,8,9})) // 825
    fmt.Println(sumOfGoodSubsequences1([]int{9,8,7,6,5,4,3,2,1})) // 825
}
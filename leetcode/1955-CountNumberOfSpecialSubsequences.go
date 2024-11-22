package main

// 1955. Count Number of Special Subsequences
// A sequence is special if it consists of a positive number of 0s, followed by a positive number of 1s, then a positive number of 2s.
//     For example, [0,1,2] and [0,0,1,1,1,2] are special.
//     In contrast, [2,1,0], [1], and [0,1,2,0] are not special.

// Given an array nums (consisting of only integers 0, 1, and 2), return the number of different subsequences that are special. 
// Since the answer may be very large, return it modulo 10^9 + 7.

// A subsequence of an array is a sequence that can be derived from the array by deleting some or no elements without changing the order of the remaining elements. 
// Two subsequences are different if the set of indices chosen are different.

// Example 1:
// Input: nums = [0,1,2,2]
// Output: 3
// Explanation: The special subsequences are bolded [0,1,2,2], [0,1,2,2], and [0,1,2,2].

// Example 2:
// Input: nums = [2,2,0,0]
// Output: 0
// Explanation: There are no special subsequences in [2,2,0,0].

// Example 3:
// Input: nums = [0,1,2,0,1,2]
// Output: 7
// Explanation: The special subsequences are bolded:
// - [0,1,2,0,1,2]
// - [0,1,2,0,1,2]
// - [0,1,2,0,1,2]
// - [0,1,2,0,1,2]
// - [0,1,2,0,1,2]
// - [0,1,2,0,1,2]
// - [0,1,2,0,1,2]

// Constraints:
//     1 <= nums.length <= 10^5
//     0 <= nums[i] <= 2

import "fmt"

func countSpecialSubsequences(nums []int) int {
    n, mod := len(nums), 1_000_000_007
    memo := make([][4]int, n)
    for i := range memo {
        memo[i] = [4]int{ -1, -1, -1, -1 }
    }
    var dp func(i, p int) int
    dp = func(i, p int) int {
        if i >= n {
            if p == 2 { return 1 }
            return 0
        }
        if memo[i][p + 1] != -1 { return memo[i][p + 1] }
        v := 0
        if p + 1 == nums[i] {
            v += dp(i + 1, nums[i]) % mod
        } else if p == nums[i] {
            v += dp(i + 1, p) % mod
        }
        v += dp(i+1, p) % mod
        memo[i][p + 1] = v
        return v
    }
    return dp(0, -1) % mod
}

func countSpecialSubsequences1(nums []int) int {
    start, n, mod := 0, len(nums), 1_000_000_007
    for ; start < n && nums[start] != 0; start++ {}
    if start == n { return 0 }
    dp := [3]int{1, 0, 0}
    nums = nums[start + 1:]
    n = len(nums)
    for i := 0; i < n; i++ {
        dp[nums[i]] = (dp[nums[i]] * 2) % mod
        if nums[i] == 0 {
            dp[0] = (dp[0] + 1) % mod
        } else {
            dp[nums[i]] = (dp[nums[i]] + dp[nums[i] - 1]) % mod
        }
    }
    return dp[2]
}

func main() {
    // Example 1:
    // Input: nums = [0,1,2,2]
    // Output: 3
    // Explanation: The special subsequences are bolded [0,1,2,2], [0,1,2,2], and [0,1,2,2].
    fmt.Println(countSpecialSubsequences([]int{0,1,2,2})) // 3
    // Example 2:
    // Input: nums = [2,2,0,0]
    // Output: 0
    // Explanation: There are no special subsequences in [2,2,0,0].
    fmt.Println(countSpecialSubsequences([]int{2,2,0,0})) // 0
    // Example 3:
    // Input: nums = [0,1,2,0,1,2]
    // Output: 7
    // Explanation: The special subsequences are bolded:
    // - [0,1,2,0,1,2]
    // - [0,1,2,0,1,2]
    // - [0,1,2,0,1,2]
    // - [0,1,2,0,1,2]
    // - [0,1,2,0,1,2]
    // - [0,1,2,0,1,2]
    // - [0,1,2,0,1,2]
    fmt.Println(countSpecialSubsequences([]int{0,1,2,0,1,2})) // 7

    fmt.Println(countSpecialSubsequences1([]int{0,1,2,2})) // 3
    fmt.Println(countSpecialSubsequences1([]int{2,2,0,0})) // 0
    fmt.Println(countSpecialSubsequences1([]int{0,1,2,0,1,2})) // 7
}
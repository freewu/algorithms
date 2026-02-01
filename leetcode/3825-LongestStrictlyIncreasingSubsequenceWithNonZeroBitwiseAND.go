package main

// 3825. Longest Strictly Increasing Subsequence With Non-Zero Bitwise AND
// You are given an integer array nums.

// Return the length of the longest strictly increasing subsequence in nums whose bitwise AND is non-zero. 
// If no such subsequence exists, return 0.

// Example 1:
// Input: nums = [5,4,7]
// Output: 2
// Explanation:
// One longest strictly increasing subsequence is [5, 7]. The bitwise AND is 5 AND 7 = 5, which is non-zero.

// Example 2:
// Input: nums = [2,3,6]
// Output: 3
// Explanation:
// The longest strictly increasing subsequence is [2, 3, 6]. The bitwise AND is 2 AND 3 AND 6 = 2, which is non-zero.

// Example 3:
// Input: nums = [0,1]
// Output: 1
// Explanation:
// One longest strictly increasing subsequence is [1]. The bitwise AND is 1, which is non-zero.

// Constraints:
//     1 <= nums.length <= 10^5
//     0 <= nums[i] <= 10^9​​​​​ ​​

import "fmt"

func longestSubsequence(nums []int) int {
    res := 0
    for i := 0; i < 32; i++ {
        dp := make([]int, 0)
        mask := 1 << i
        for _, n := range nums {
            if (n & mask) == 0 { continue }
            l, r := 0, len(dp)
            for l < r {
                m := (l + r) >> 1
                if dp[m] < n {
                    l = m + 1
                } else {
                    r = m
                }
            }
            if l < len(dp) && dp[l] == n { continue }
            if l == len(dp) {
                dp = append(dp, n)
            } else {
                dp[l] = n
            }
        }
        if len(dp) > res {
            res = len(dp)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [5,4,7]
    // Output: 2
    // Explanation:
    // One longest strictly increasing subsequence is [5, 7]. The bitwise AND is 5 AND 7 = 5, which is non-zero.
    fmt.Println(longestSubsequence([]int{5,4,7})) // 2
    // Example 2:
    // Input: nums = [2,3,6]
    // Output: 3
    // Explanation:
    // The longest strictly increasing subsequence is [2, 3, 6]. The bitwise AND is 2 AND 3 AND 6 = 2, which is non-zero.
    fmt.Println(longestSubsequence([]int{2,3,6})) // 3
    // Example 3:
    // Input: nums = [0,1]
    // Output: 1
    // Explanation:
    // One longest strictly increasing subsequence is [1]. The bitwise AND is 1, which is non-zero.
    fmt.Println(longestSubsequence([]int{0,1})) // 1

    fmt.Println(longestSubsequence([]int{1,2,3,4,5,6,7,8,9})) // 5
    fmt.Println(longestSubsequence([]int{9,8,7,6,5,4,3,2,1})) // 1
}
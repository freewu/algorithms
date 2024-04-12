package main

// 673. Number of Longest Increasing Subsequence
// Given an integer array nums, return the number of longest increasing subsequences.
// Notice that the sequence has to be strictly increasing.

// Example 1:
// Input: nums = [1,3,5,4,7]
// Output: 2
// Explanation: The two longest increasing subsequences are [1, 3, 4, 7] and [1, 3, 5, 7].

// Example 2:
// Input: nums = [2,2,2,2,2]
// Output: 5
// Explanation: The length of the longest increasing subsequence is 1, and there are 5 increasing subsequences of length 1, so output 5.

// Constraints:
//     1 <= nums.length <= 2000
//     -10^6 <= nums[i] <= 10^6

import "fmt"

// dp 
func findNumberOfLIS(nums []int) int {
    dp, c, res, mc := make([]int, len(nums)), make([]int,len(nums)), 0, 0
    for i := 0; i < len(nums); i++ {
        dp[i] = 1
        c[i] = 1
        for j := i-1; j >= 0; j-- {
            if nums[i] > nums[j] {
                if dp[i] < dp[j] + 1 {
                    dp[i] = dp[j] + 1
                    c[i] = c[j]
                } else if dp[i] == dp[j] + 1 {
                    c[i] += c[j]
                }
            }
        }
        if dp[i] > mc {
            mc = dp[i]
            res = c[i]
        } else if dp[i] == mc {
            res += c[i]
        }
    }
    return res  
}

func main() {
    // Explanation: The two longest increasing subsequences are [1, 3, 4, 7] and [1, 3, 5, 7].
    fmt.Println(findNumberOfLIS([]int{1,3,5,4,7})) // 2
    // Explanation: The length of the longest increasing subsequence is 1, and there are 5 increasing subsequences of length 1, so output 5.
    fmt.Println(findNumberOfLIS([]int{2,2,2,2,2})) // 5
}
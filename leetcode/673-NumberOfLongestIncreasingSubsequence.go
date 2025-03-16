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
import "sort"

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

func findNumberOfLIS1(nums []int) int {
    q, s := [][]int{}, [][]int{}
    for _, x := range nums {
        i := sort.Search(len(q), func(i int) bool {
            return q[i][len(q[i]) - 1] >= x
        })
        y := 1
        if i > 0 {
            j := sort.Search(len(q[i-1]), func(j int) bool {
                return q[i - 1][j] < x
            })
            y = s[i - 1][len(s[i - 1]) - 1] - s[i - 1][j]
        }
        if i == len(q) {
            q = append(q, []int{x})
            s = append(s, []int{0, y})
        } else {
            q[i] = append(q[i], x)
            s[i] = append(s[i], y + s[i][len(s[i]) - 1])
        }
    }
    p := s[len(s) - 1]
    return p[len(p) - 1]
}

func main() {
    // Example 1:
    // Input: nums = [1,3,5,4,7]
    // Output: 2
    // Explanation: The two longest increasing subsequences are [1, 3, 4, 7] and [1, 3, 5, 7].
    fmt.Println(findNumberOfLIS([]int{1,3,5,4,7})) // 2
    // Example 2:
    // Input: nums = [2,2,2,2,2]
    // Output: 5
    // Explanation: The length of the longest increasing subsequence is 1, and there are 5 increasing subsequences of length 1, so output 5.
    fmt.Println(findNumberOfLIS([]int{2,2,2,2,2})) // 5

    fmt.Println(findNumberOfLIS([]int{1,2,3,4,5,6,7,8,9})) // 1
    fmt.Println(findNumberOfLIS([]int{9,8,7,6,5,4,3,2,1})) // 9

    fmt.Println(findNumberOfLIS1([]int{1,3,5,4,7})) // 2
    fmt.Println(findNumberOfLIS1([]int{2,2,2,2,2})) // 5
    fmt.Println(findNumberOfLIS1([]int{1,2,3,4,5,6,7,8,9})) // 1
    fmt.Println(findNumberOfLIS1([]int{9,8,7,6,5,4,3,2,1})) // 9
}
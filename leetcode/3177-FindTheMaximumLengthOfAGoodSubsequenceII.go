package main

// 3177. Find the Maximum Length of a Good Subsequence II
// You are given an integer array nums and a non-negative integer k. 
// A sequence of integers seq is called good if there are at most k indices i in the range [0, seq.length - 2] such that seq[i] != seq[i + 1].

// Return the maximum possible length of a good subsequence of nums.

// Example 1:
// Input: nums = [1,2,1,1,3], k = 2
// Output: 4
// Explanation:
// The maximum length subsequence is [1,2,1,1,3].

// Example 2:
// Input: nums = [1,2,3,4,5,1], k = 0
// Output: 2
// Explanation:
// The maximum length subsequence is [1,2,3,4,5,1].

// Constraints:
//     1 <= nums.length <= 5 * 10^3
//     1 <= nums[i] <= 10^9
//     0 <= k <= min(50, nums.length)

import "fmt"

// // dfs
// func maximumLength(nums []int, k int) int {
//     n := len(nums)
//     memo := make([][]int, k + 1)
//     for i := 0; i < k+1; i++ {
//         memo[i] = make([]int, n)
//     }
//     max := func (x, y int) int { if x > y { return x; }; return y; }
//     var dfs func(i, k int) int
//     dfs = func(i, k int) int {
//         res := 0
//         if k < 0 { return 0 }
//         if i == n - 1 { return 1 }
//         if memo[k][i] != 0 { return memo[k][i] }
//         for j := i + 1; j < n; j++ {
//             if nums[i] == nums[j] {
//                 res = max(res, dfs(j, k) + 1)
//             } else {
//                 res = max(res, dfs(j, k-1)+1)
//             }
//         }
//         memo[k][i] = res
//         return res
//     }
//     res := 0
//     for i := 0; i < n; i++ {
//         res = max(res, dfs(i, k))
//     }
//     return dfs(0, k)
// }

func maximumLength(nums []int, k int) int {
    n := len(nums)
    dp := make([][]int, n + 1)
    dp[0] = make([]int, k+1)
    lastIdx, maxSoFar := make(map[int]int, n), make([]int, k+1)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < k+1; i++ { maxSoFar[i] = 1 }
    for i, num := range nums {
        dp[i+1] = make([]int, k+1)
        for j := k; j >= 0; j-- {
            if j == 0 {
                if numDupe, found := lastIdx[num]; found {
                    dp[i+1][j] = dp[numDupe+1][j]+1
                } else {
                    dp[i+1][j]++
                }
            } else if i == 0 {
                dp[i+1][j] = 1
            } else if num == nums[i-1] {
                dp[i+1][j] = max(dp[i][j] + 1, maxSoFar[j-1]+1)
            } else {
                if numDupe, found := lastIdx[num]; found {
                    dp[i+1][j] = dp[numDupe+1][j]
                }

                dp[i+1][j] = max(dp[i+1][j]+1, maxSoFar[j-1]+1)
            }

            maxSoFar[j] = max(maxSoFar[j], dp[i+1][j])
        }
        lastIdx[num] = i
    }
    return maxSoFar[k]
}

func maximumLength1(nums []int, k int) int {
    fs, mx := map[int][]int{},  make([]int, k + 2)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range nums {
        if fs[v] == nil {
            fs[v] = make([]int, k+1)
        }
        f := fs[v]
        for i := k; i >= 0; i-- {
            f[i] = max(f[i], mx[i]) + 1
            mx[i+1] = max(mx[i+1], f[i])
        }
    }
    return mx[k+1]
}

func main() {
    // Example 1:
    // Input: nums = [1,2,1,1,3], k = 2
    // Output: 4
    // Explanation:
    // The maximum length subsequence is [1,2,1,1,3].
    fmt.Println(maximumLength([]int{1,2,1,1,3}, 2)) // 4
    // Input: nums = [1,2,3,4,5,1], k = 0
    // Output: 2
    // Explanation:
    // The maximum length subsequence is [1,2,3,4,5,1].
    fmt.Println(maximumLength([]int{1,2,3,4,5,1}, 0)) // 2

    fmt.Println(maximumLength([]int{23,18,18}, 0)) // 2

    fmt.Println(maximumLength1([]int{1,2,1,1,3}, 2)) // 4
    fmt.Println(maximumLength1([]int{1,2,3,4,5,1}, 0)) // 2
    fmt.Println(maximumLength1([]int{23,18,18}, 0)) // 2
}
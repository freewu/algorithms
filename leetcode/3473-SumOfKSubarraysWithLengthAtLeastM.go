package main

// 3473. Sum of K Subarrays With Length at Least M
// You are given an integer array nums and two integers, k and m.

// Return the maximum sum of k non-overlapping subarrays of nums, where each subarray has a length of at least m.

// Example 1:
// Input: nums = [1,2,-1,3,3,4], k = 2, m = 2
// Output: 13
// Explanation:
// The optimal choice is:
// Subarray nums[3..5] with sum 3 + 3 + 4 = 10 (length is 3 >= m).
// Subarray nums[0..1] with sum 1 + 2 = 3 (length is 2 >= m).
// The total sum is 10 + 3 = 13.

// Example 2:
// Input: nums = [-10,3,-1,-2], k = 4, m = 1
// Output: -10
// Explanation:
// The optimal choice is choosing each element as a subarray. The output is (-10) + 3 + (-1) + (-2) = -10.

// Constraints:
//     1 <= nums.length <= 2000
//     -10^4 <= nums[i] <= 10^4
//     1 <= k <= floor(nums.length / m)
//     1 <= m <= 3

import "fmt"

func maxSum(nums []int, k int, m int) int {
    n, inf := len(nums), 1 << 31
    dp := make([][]int, n + 1)
    for i := range dp {
        dp[i] = make([]int, k + 1)
        for j := range dp[i] {
            if j == 0 {
                dp[i][j] = 0
            } else {
                dp[i][j] = -inf
            }
        }
    }
    prefix := make([]int, n + 1)
    for i := 1; i <= n; i++ {
        prefix[i] = prefix[i - 1] + nums[i - 1]
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for seg := 1; seg <= k; seg++ {
        best := -inf
        for i := seg * m; i <= n; i++ {
            start := i - m
            if start >= 0 {
                best = max(best, dp[start][seg - 1] - prefix[start])
            }
            if i == seg * m {
                dp[i][seg] = prefix[i] + best
            } else {
                dp[i][seg] = max(dp[i-1][seg], prefix[i] + best)
            }
        }
    }
    return dp[n][k]
}

func maxSum1(nums []int, k, m int) int {
    n, inf := len(nums), 1 << 31
    sum, f, d := make([]int, n + 1), make([]int, n + 1), make([]int, n + 1)
    for i, v := range nums {
        sum[i + 1] = sum[i] + v
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i <= k; i++ {
        for j := i * m - m; j < i*m; j++ {
            d[j] = f[j] - sum[j]
            f[j] = -inf // 即使 [0,j) 全选，也没有 i 个长为 m 的子数组
        }
        mx := -inf
        // 左右两边留出足够空间给其他子数组
        for j := i * m; j <= n-(k-i)*m; j++ {
            // mx 表示最大的 f[L] - sum[L]，其中 L 在区间 [(i-1)*m, j-m] 中
            mx = max(mx, d[j-m])
            d[j] = f[j] - sum[j]
            f[j] = max(f[j - 1], mx + sum[j]) // 不选 vs 选
        }
    }
    return f[n]
}

func main() {
    // Example 1:
    // Input: nums = [1,2,-1,3,3,4], k = 2, m = 2
    // Output: 13
    // Explanation:
    // The optimal choice is:
    // Subarray nums[3..5] with sum 3 + 3 + 4 = 10 (length is 3 >= m).
    // Subarray nums[0..1] with sum 1 + 2 = 3 (length is 2 >= m).
    // The total sum is 10 + 3 = 13.
    fmt.Println(maxSum([]int{1,2,-1, 3, 3, 4}, 2, 2)) // 13
    // Example 2:
    // Input: nums = [-10,3,-1,-2], k = 4, m = 1
    // Output: -10
    // Explanation:
    // The optimal choice is choosing each element as a subarray. The output is (-10) + 3 + (-1) + (-2) = -10.
    fmt.Println(maxSum([]int{-10,3,-1,-2}, 4, 1)) // -10

    fmt.Println(maxSum([]int{1,2,3,4,5,6,7,8,9}, 4, 1)) // 45
    fmt.Println(maxSum([]int{9,8,7,6,5,4,3,2,1}, 4, 1)) // 45

    fmt.Println(maxSum1([]int{1,2,-1, 3, 3, 4}, 2, 2)) // 13
    fmt.Println(maxSum1([]int{-10,3,-1,-2}, 4, 1)) // -10
    fmt.Println(maxSum1([]int{1,2,3,4,5,6,7,8,9}, 4, 1)) // 45
    fmt.Println(maxSum1([]int{9,8,7,6,5,4,3,2,1}, 4, 1)) // 45
}
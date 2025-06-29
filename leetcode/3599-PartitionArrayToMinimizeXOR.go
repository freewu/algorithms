package main

// 3599. Partition Array to Minimize XOR
// You are given an integer array nums and an integer k.

// Your task is to partition nums into k non-empty subarrays. 
// For each subarray, compute the bitwise XOR of all its elements.

// Return the minimum possible value of the maximum XOR among these k subarrays.

// Example 1:
// Input: nums = [1,2,3], k = 2
// Output: 1
// Explanation:
// The optimal partition is [1] and [2, 3].
// XOR of the first subarray is 1.
// XOR of the second subarray is 2 XOR 3 = 1.
// The maximum XOR among the subarrays is 1, which is the minimum possible.

// Example 2:
// Input: nums = [2,3,3,2], k = 3
// Output: 2
// Explanation:
// The optimal partition is [2], [3, 3], and [2].
// XOR of the first subarray is 2.
// XOR of the second subarray is 3 XOR 3 = 0.
// XOR of the third subarray is 2.
// The maximum XOR among the subarrays is 2, which is the minimum possible.

// Example 3:
// Input: nums = [1,1,2,3,1], k = 2
// Output: 0
// Explanation:
// The optimal partition is [1, 1] and [2, 3, 1].
// XOR of the first subarray is 1 XOR 1 = 0.
// XOR of the second subarray is 2 XOR 3 XOR 1 = 0.
// The maximum XOR among the subarrays is 0, which is the minimum possible.

// Constraints:
//     1 <= nums.length <= 250
//     1 <= nums[i] <= 10^9
//     1 <= k <= n

import "fmt"

func minXor(nums []int, k int) int {
    n := len(nums)
    // dp[i][j]表示将[0: j]划分为i个子数组的最大XOR的最小值
    dp := make([][]int, k + 1)
    dp[0] = make([]int, n + 1)
    for i := 0; i < n; i++ {
        dp[0][i + 1] = 1 << 31
    }
    for i := 1; i <= k; i++ {
        dp[i] = make([]int, n + 1)
        for j := i; j <= n - k + i; j++ {
            xor, cur := 0, 1 << 31
            for p := j - 1; p >= i - 1; p-- {
                xor ^= nums[p]
                cur = min(cur, max(xor, dp[i - 1][p]))
            }
            dp[i][j] = cur
        }
    }
    return dp[k][n]
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3], k = 2
    // Output: 1
    // Explanation:
    // The optimal partition is [1] and [2, 3].
    // XOR of the first subarray is 1.
    // XOR of the second subarray is 2 XOR 3 = 1.
    // The maximum XOR among the subarrays is 1, which is the minimum possible.
    fmt.Println(minXor([]int{1,2,3}, 2)) // 1
    // Example 2:
    // Input: nums = [2,3,3,2], k = 3
    // Output: 2
    // Explanation:
    // The optimal partition is [2], [3, 3], and [2].
    // XOR of the first subarray is 2.
    // XOR of the second subarray is 3 XOR 3 = 0.
    // XOR of the third subarray is 2.
    // The maximum XOR among the subarrays is 2, which is the minimum possible.
    fmt.Println(minXor([]int{2,3,3,2}, 3)) // 2
    // Example 3:
    // Input: nums = [1,1,2,3,1], k = 2
    // Output: 0
    // Explanation:
    // The optimal partition is [1, 1] and [2, 3, 1].
    // XOR of the first subarray is 1 XOR 1 = 0.
    // XOR of the second subarray is 2 XOR 3 XOR 1 = 0.
    // The maximum XOR among the subarrays is 0, which is the minimum possible.
    fmt.Println(minXor([]int{1,1,2,3,1}, 2)) // 0

    fmt.Println(minXor([]int{1,2,3,4,5,6,7,8,9}, 2)) // 1
    fmt.Println(minXor([]int{9,8,7,6,5,4,3,2,1}, 2)) // 1
}
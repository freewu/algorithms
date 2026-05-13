package main

// 3929. Minimum Partition Score II
// You are given an integer array nums and an integer k.

// Your task is to partition nums into exactly k subarrays and return an integer denoting the minimum possible score among all valid partitions.

// The score of a partition is the sum of the values of all its subarrays.

// The value of a subarray is defined as sumArr * (sumArr + 1) / 2, where sumArr is the sum of its elements.

// Example 1:
// Input: nums = [5,1,2,1], k = 2
// Output: 25
// Explanation:
// We must partition the array into k = 2 subarrays. One optimal partition is [5] and [1, 2, 1].
// The first subarray has sum = 5 and value = 5 * 6 / 2 = 15.
// The second subarray has sum = 1 + 2 + 1 = 4 and value = 4 * 5 / 2 = 10.
// The score of this partition is 15 + 10 = 25, which is the minimum possible score.

// Example 2:
// Input: nums = [1,2,3,4], k = 1
// Output: 55
// Explanation:
// Since we must partition the array into k = 1 subarray, all elements belong to the same subarray: [1, 2, 3, 4].
// This subarray has sum = 1 + 2 + 3 + 4 = 10 and value = 10 * 11 / 2 = 55.​​​​​​​
// The score of this partition is 55, which is the minimum possible score.

// Example 3:
// Input: nums = [1,1,1], k = 3
// Output: 3
// Explanation:
// We must partition the array into k = 3 subarrays. The only valid partition is [1], [1], [1].
// Each subarray has sum = 1 and value = 1 * 2 / 2 = 1.
// The score of this partition is 1 + 1 + 1 = 3, which is the minimum possible score.

// Constraints:
//     1 <= nums.length <= 5 * 10^4
//     1 <= nums[i] <= 10^3
//     1 <= k <= nums.length 

import "fmt"

func minPartitionScore(nums []int, k int) int64 {
    n := len(nums)
    // 前缀和 s[0] = 0, s[i] = nums[0]+...+nums[i-1]
    s := make([]int64, n+1)
    for i := 0; i < n; i++ {
        s[i+1] = s[i] + int64(nums[i])
    }
    // dp[i] = 前 i 个元素分成当前轮次段数的最小代价
    dp := make([]int64, n+1)
    for i := range dp {
        dp[i] = 1 << 61
    }
    dp[0] = 0
    // 分成 k 段，迭代 k 次
    for seg := 1; seg <= k; seg++ {
        nextDp := make([]int64, n+1)
        for i := range nextDp {
            nextDp[i] = 1 << 61
        }
        // 枚举右端点 i，左端点 j
        // 分成 seg 段 => j >= seg-1
        for i := seg; i <= n; i++ {
            for j := seg - 1; j < i; j++ {
                ds := s[i] - s[j]
                cost := ds * (ds + 1) / 2
                if dp[j]+cost < nextDp[i] {
                    nextDp[i] = dp[j] + cost
                }
            }
        }
        dp = nextDp
    }
    return dp[n]
}

func main() {
    // Example 1:
    // Input: nums = [5,1,2,1], k = 2
    // Output: 25
    // Explanation:
    // We must partition the array into k = 2 subarrays. One optimal partition is [5] and [1, 2, 1].
    // The first subarray has sum = 5 and value = 5 * 6 / 2 = 15.
    // The second subarray has sum = 1 + 2 + 1 = 4 and value = 4 * 5 / 2 = 10.
    // The score of this partition is 15 + 10 = 25, which is the minimum possible score.
    fmt.Println(minPartitionScore([]int{5,1,2,1}, 2)) // 25 
    // Example 2:
    // Input: nums = [1,2,3,4], k = 1
    // Output: 55
    // Explanation:
    // Since we must partition the array into k = 1 subarray, all elements belong to the same subarray: [1, 2, 3, 4].
    // This subarray has sum = 1 + 2 + 3 + 4 = 10 and value = 10 * 11 / 2 = 55.​​​​​​​
    // The score of this partition is 55, which is the minimum possible score.
    fmt.Println(minPartitionScore([]int{1,2,3,4}, 1)) // 55 
    // Example 3:
    // Input: nums = [1,1,1], k = 3
    // Output: 3
    // Explanation:
    // We must partition the array into k = 3 subarrays. The only valid partition is [1], [1], [1].
    // Each subarray has sum = 1 and value = 1 * 2 / 2 = 1.
    // The score of this partition is 1 + 1 + 1 = 3, which is the minimum possible score.   
    fmt.Println(minPartitionScore([]int{1,1,1}, 3)) // 3 

    fmt.Println(minPartitionScore([]int{1,2,3,4,5,6,7,8,9}, 2)) // 531 
    fmt.Println(minPartitionScore([]int{9,8,7,6,5,4,3,2,1}, 2)) // 531 
}
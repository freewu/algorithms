package main

// 2297. Jump Game VIII
// You are given a 0-indexed integer array nums of length n. 
// You are initially standing at index 0. 
// You can jump from index i to index j where i < j if:
//     1. nums[i] <= nums[j] and nums[k] < nums[i] for all indexes k in the range i < k < j, or
//     2. nums[i] > nums[j] and nums[k] >= nums[i] for all indexes k in the range i < k < j.

// You are also given an integer array costs of length n where costs[i] denotes the cost of jumping to index i.

// Return the minimum cost to jump to the index n - 1.

// Example 1:
// Input: nums = [3,2,4,4,1], costs = [3,7,6,4,2]
// Output: 8
// Explanation: You start at index 0.
// - Jump to index 2 with a cost of costs[2] = 6.
// - Jump to index 4 with a cost of costs[4] = 2.
// The total cost is 8. It can be proven that 8 is the minimum cost needed.
// Two other possible paths are from index 0 -> 1 -> 4 and index 0 -> 2 -> 3 -> 4.
// These have a total cost of 9 and 12, respectively.

// Example 2:
// Input: nums = [0,1,2], costs = [1,1,1]
// Output: 2
// Explanation: Start at index 0.
// - Jump to index 1 with a cost of costs[1] = 1.
// - Jump to index 2 with a cost of costs[2] = 1.
// The total cost is 2. Note that you cannot jump directly from index 0 to index 2 because nums[0] <= nums[1].

// Constraints:
//     n == nums.length == costs.length
//     1 <= n <= 10^5
//     0 <= nums[i], costs[i] <= 10^5

import "fmt"

func minCost(nums []int, costs []int) int64 {
    n := len(nums)
    dp, mn, mx := make([]int, n), []int{ 0 }, []int{ 0 }
    for i := 1; i < n; i++ {
        v := 10000000001
        for len(mn) > 0 && nums[mn[len(mn)-1]] <= nums[i] {
            top := mn[len(mn)-1]
            mn = mn[:len(mn)-1]
            if dp[top] < v {
                v = dp[top]
            }
        }
        mn = append(mn, i)
        for len(mx) > 0 && nums[mx[len(mx)-1]] > nums[i] {
            top := mx[len(mx)-1]
            mx = mx[:len(mx)-1]
            if dp[top] < v {
                v = dp[top]
            }
        }
        mx = append(mx, i)
        dp[i] = v + costs[i]
    }
    return int64(dp[n-1])
}

func main() {
    // Example 1:
    // Input: nums = [3,2,4,4,1], costs = [3,7,6,4,2]
    // Output: 8
    // Explanation: You start at index 0.
    // - Jump to index 2 with a cost of costs[2] = 6.
    // - Jump to index 4 with a cost of costs[4] = 2.
    // The total cost is 8. It can be proven that 8 is the minimum cost needed.
    // Two other possible paths are from index 0 -> 1 -> 4 and index 0 -> 2 -> 3 -> 4.
    // These have a total cost of 9 and 12, respectively.
    fmt.Println(minCost([]int{3,2,4,4,1}, []int{3,7,6,4,2})) // 8
    // Example 2:
    // Input: nums = [0,1,2], costs = [1,1,1]
    // Output: 2
    // Explanation: Start at index 0.
    // - Jump to index 1 with a cost of costs[1] = 1.
    // - Jump to index 2 with a cost of costs[2] = 1.
    // The total cost is 2. Note that you cannot jump directly from index 0 to index 2 because nums[0] <= nums[1].
    fmt.Println(minCost([]int{0,1,2}, []int{1,1,1})) // 2
}
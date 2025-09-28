package main

// 3693. Climbing Stairs II
// You are climbing a staircase with n + 1 steps, numbered from 0 to n.

// You are also given a 1-indexed integer array costs of length n, where costs[i] is the cost of step i.

// From step i, you can jump only to step i + 1, i + 2, or i + 3. The cost of jumping from step i to step j is defined as: costs[j] + (j - i)2

// You start from step 0 with cost = 0.

// Return the minimum total cost to reach step n.

// Example 1:
// Input: n = 4, costs = [1,2,3,4]
// Output: 13
// Explanation:
// One optimal path is 0 → 1 → 2 → 4
// Jump	Cost Calculation	Cost
// 0 → 1	costs[1] + (1 - 0)2 = 1 + 1	2
// 1 → 2	costs[2] + (2 - 1)2 = 2 + 1	3
// 2 → 4	costs[4] + (4 - 2)2 = 4 + 4	8
// Thus, the minimum total cost is 2 + 3 + 8 = 13

// Example 2:
// Input: n = 4, costs = [5,1,6,2]
// Output: 11
// Explanation:
// One optimal path is 0 → 2 → 4
// Jump	Cost Calculation	Cost
// 0 → 2	costs[2] + (2 - 0)2 = 1 + 4	5
// 2 → 4	costs[4] + (4 - 2)2 = 2 + 4	6
// Thus, the minimum total cost is 5 + 6 = 11

// Example 3:
// Input: n = 3, costs = [9,8,3]
// Output: 12
// Explanation:
// The optimal path is 0 → 3 with total cost = costs[3] + (3 - 0)2 = 3 + 9 = 12

// Constraints:
//     1 <= n == costs.length <= 10^5​​​​​​​
//     1 <= costs[i] <= 10^4

import "fmt"

func climbStairs(n int, costs []int) int {
    dp := make([]int, n + 1)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i <= n; i++ {
        dp[i] = 1 << 31
        for j := max(0, i - 3); j < i; j++ {
            dp[i] = min(dp[i], dp[j] + (i - j) * (i - j) + costs[i - 1])
        }
    }
    return dp[n]
}

func climbStairs1(n int, costs []int) int {
    res, prev2, prev1 := 0, 1 << 31, 1 << 31
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i <= n; i++ {
        next := min(min(prev2 + costs[i - 1] + 3 * 3, prev1 + costs[i - 1] + 2 * 2), res + costs[i - 1] + 1 * 1)
        prev2, prev1, res = prev1, res, next
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 4, costs = [1,2,3,4]
    // Output: 13
    // Explanation:
    // One optimal path is 0 → 1 → 2 → 4
    // Jump	Cost Calculation	Cost
    // 0 → 1	costs[1] + (1 - 0)2 = 1 + 1	2
    // 1 → 2	costs[2] + (2 - 1)2 = 2 + 1	3
    // 2 → 4	costs[4] + (4 - 2)2 = 4 + 4	8
    // Thus, the minimum total cost is 2 + 3 + 8 = 13
    fmt.Println(climbStairs(4, []int{1,2,3,4})) // 13
    // Example 2:
    // Input: n = 4, costs = [5,1,6,2]
    // Output: 11
    // Explanation:
    // One optimal path is 0 → 2 → 4
    // Jump	Cost Calculation	Cost
    // 0 → 2	costs[2] + (2 - 0)2 = 1 + 4	5
    // 2 → 4	costs[4] + (4 - 2)2 = 2 + 4	6
    // Thus, the minimum total cost is 5 + 6 = 11
    fmt.Println(climbStairs(4, []int{5,1,6,2})) // 11
    // Example 3:
    // Input: n = 3, costs = [9,8,3]
    // Output: 12
    // Explanation:
    // The optimal path is 0 → 3 with total cost = costs[3] + (3 - 0)2 = 3 + 9 = 12
    fmt.Println(climbStairs(3, []int{9,8,3})) // 12

    fmt.Println(climbStairs(3, []int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(climbStairs(3, []int{9,8,7,6,5,4,3,2,1})) // 16

    fmt.Println(climbStairs1(4, []int{1,2,3,4})) // 13
    fmt.Println(climbStairs1(4, []int{5,1,6,2})) // 11
    fmt.Println(climbStairs1(3, []int{9,8,3})) // 12
    fmt.Println(climbStairs1(3, []int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(climbStairs1(3, []int{9,8,7,6,5,4,3,2,1})) // 16
}
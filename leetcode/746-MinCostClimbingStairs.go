package main

// 746. Min Cost Climbing Stairs
// You are given an integer array cost where cost[i] is the cost of ith step on a staircase. 
// Once you pay the cost, you can either climb one or two steps.

// You can either start from the step with index 0, or the step with index 1.
// Return the minimum cost to reach the top of the floor.

// Example 1:
// Input: cost = [10,15,20]
// Output: 15
// Explanation: You will start at index 1.
// - Pay 15 and climb two steps to reach the top.
// The total cost is 15.

// Example 2:
// Input: cost = [1,100,1,1,1,100,1,1,100,1]
// Output: 6
// Explanation: You will start at index 0.
// - Pay 1 and climb two steps to reach index 2.
// - Pay 1 and climb two steps to reach index 4.
// - Pay 1 and climb two steps to reach index 6.
// - Pay 1 and climb one step to reach index 7.
// - Pay 1 and climb two steps to reach index 9.
// - Pay 1 and climb one step to reach the top.
// The total cost is 6.
 
// Constraints:
//     2 <= cost.length <= 1000
//     0 <= cost[i] <= 999

import "fmt"

func minCostClimbingStairs(cost []int) int {
    pay := make([]int, len(cost))
    pay[0], pay[1] = cost[0], cost[1]
    min := func (x, y int) int { if x < y { return x; }; return y; }
    // 计算每一步需要的累计费用
    for i := 2 ; i < len(pay); i++ {
        // 走两步 pay[i-2] 走一步 pay[i-1]
        pay[i] = min(pay[i-2], pay[i-1]) + cost[i]
    }
    return min(pay[len(pay)-2], pay[len(pay)-1])
}

// dfs
func minCostClimbingStairs1(cost []int) int {
    n := len(cost)
    dp := make([]int, n + 1)
    for i := 0; i < n+1; i += 1 {
        dp[i] = -1
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    var dfs func(int) int
    dfs = func (i int) int {
        if i < 0 {
            return 0
        } else if i == n {
            return min(dfs(i-1), dfs(i-2))
        }
        res := &dp[i]
        if *res != -1 {
            return dp[i]
        }
        *res = min(dfs(i-1), dfs(i-2)) + cost[i]
        return *res
    }
    return dfs(n)
}

func minCostClimbingStairs2(cost []int) int {
    // 只用三位来保存数据
    dp := make([]int,3)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 2; i <= len(cost); i++ {
        // 走一步的费用: dp[1] + cost[i-1] ,  走两步的费用:  dp[2]+ cost[i-2]
        dp[0] = min(dp[1] + cost[i-1], dp[2]+ cost[i-2])
        dp[2] = dp[1]
        dp[1] = dp[0]
    }
    return dp[1]
}

func minCostClimbingStairs3(cost []int) int {
    n := len(cost)
    dp := make([]int,  n + 1)
    dp[0], dp[1] = 0, 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 2; i <= n; i++ {
        dp[i] = min(dp[i-1] + cost[i-1], dp[i-2] + cost[i-2])
    }
    return dp[n]
}

func main() {
    // Explanation: You will start at index 1.
    // - Pay 15 and climb two steps to reach the top.
    // The total cost is 15.
    fmt.Println(minCostClimbingStairs([]int{10,15,20})) // 15
    // Explanation: You will start at index 0.
    // - Pay 1 and climb two steps to reach index 2.
    // - Pay 1 and climb two steps to reach index 4.
    // - Pay 1 and climb two steps to reach index 6.
    // - Pay 1 and climb one step to reach index 7.
    // - Pay 1 and climb two steps to reach index 9.
    // - Pay 1 and climb one step to reach the top.
    // The total cost is 6.
    fmt.Println(minCostClimbingStairs([]int{1,100,1,1,1,100,1,1,100,1})) // 6

    fmt.Println(minCostClimbingStairs1([]int{10,15,20})) // 15
    fmt.Println(minCostClimbingStairs1([]int{1,100,1,1,1,100,1,1,100,1})) // 6

    fmt.Println(minCostClimbingStairs2([]int{10,15,20})) // 15
    fmt.Println(minCostClimbingStairs2([]int{1,100,1,1,1,100,1,1,100,1})) // 6

    fmt.Println(minCostClimbingStairs3([]int{10,15,20})) // 15
    fmt.Println(minCostClimbingStairs3([]int{1,100,1,1,1,100,1,1,100,1})) // 6
}
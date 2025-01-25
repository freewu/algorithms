package main

// 3418. Maximum Amount of Money Robot Can Earn
// You are given an m x n grid. A robot starts at the top-left corner of the grid (0, 0) 
// and wants to reach the bottom-right corner (m - 1, n - 1). 
// The robot can move either right or down at any point in time.

// The grid contains a value coins[i][j] in each cell:
//     1. If coins[i][j] >= 0, the robot gains that many coins.
//     2. If coins[i][j] < 0, the robot encounters a robber, and the robber steals the absolute value of coins[i][j] coins.

// The robot has a special ability to neutralize robbers in at most 2 cells on its path, 
// preventing them from stealing coins in those cells.

// Note: The robot's total coins can be negative.

// Return the maximum profit the robot can gain on the route.

// Example 1:
// Input: coins = [[0,1,-1],[1,-2,3],[2,-3,4]]
// Output: 8
// Explanation:
// An optimal path for maximum coins is:
// Start at (0, 0) with 0 coins (total coins = 0).
// Move to (0, 1), gaining 1 coin (total coins = 0 + 1 = 1).
// Move to (1, 1), where there's a robber stealing 2 coins. The robot uses one neutralization here, avoiding the robbery (total coins = 1).
// Move to (1, 2), gaining 3 coins (total coins = 1 + 3 = 4).
// Move to (2, 2), gaining 4 coins (total coins = 4 + 4 = 8).

// Example 2:
// Input: coins = [[10,10,10],[10,10,10]]
// Output: 40
// Explanation:
// An optimal path for maximum coins is:
// Start at (0, 0) with 10 coins (total coins = 10).
// Move to (0, 1), gaining 10 coins (total coins = 10 + 10 = 20).
// Move to (0, 2), gaining another 10 coins (total coins = 20 + 10 = 30).
// Move to (1, 2), gaining the final 10 coins (total coins = 30 + 10 = 40).

// Constraints:
//     m == coins.length
//     n == coins[i].length
//     1 <= m, n <= 500
//     -1000 <= coins[i][j] <= 1000

import "fmt"

func maximumAmount(coins [][]int) int {
    dp := [501][501][3]int{}
    for i := range dp {
        for j := range dp[i] {
            dp[i][j] = [3]int{ -1 << 31, -1 << 31, -1 << 31 }
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(c [][]int, i, j int, n int) int
    dfs = func(c [][]int, i, j int, n int) int {
        if i == len(c) || j == len(c[0]) { return -1 << 31 }
        if i == len(c) - 1 && j == len(c[0]) - 1 {
            if n > 0 { return max(0, c[i][j]) }
            return c[i][j]
        }
        if dp[i][j][n] != -1 << 31 { return dp[i][j][n] }
        res := c[i][j] + max(dfs(c, i + 1, j, n), dfs(c, i, j + 1, n))
        if c[i][j] < 0 && n > 0 {
            res = max(res, max(dfs(c, i + 1, j, n - 1), dfs(c, i, j + 1, n - 1)))
        }
        dp[i][j][n] = res
        return res
    }
    return dfs(coins, 0, 0, 2)
}

func maximumAmount1(coins [][]int) int {
    n := len(coins[0])
    dp := make([][3]int, n + 1)
    for i := range dp {
        dp[i] = [3]int{ -1 << 31, -1 << 31, -1 << 31 }
    }
    dp[1] = [3]int{}
    for _, row := range coins {
        for i, v := range row {
            dp[i + 1][2] = max(dp[i][2] + v, dp[i + 1][2] + v, dp[i][1], dp[i + 1][1])
            dp[i + 1][1] = max(dp[i][1] + v, dp[i + 1][1] + v, dp[i][0], dp[i + 1][0])
            dp[i + 1][0] = max(dp[i][0], dp[i + 1][0]) + v
        }
    }
    return dp[n][2]
}

func main() {
    // Example 1:
    // Input: coins = [[0,1,-1],[1,-2,3],[2,-3,4]]
    // Output: 8
    // Explanation:
    // An optimal path for maximum coins is:
    // Start at (0, 0) with 0 coins (total coins = 0).
    // Move to (0, 1), gaining 1 coin (total coins = 0 + 1 = 1).
    // Move to (1, 1), where there's a robber stealing 2 coins. The robot uses one neutralization here, avoiding the robbery (total coins = 1).
    // Move to (1, 2), gaining 3 coins (total coins = 1 + 3 = 4).
    // Move to (2, 2), gaining 4 coins (total coins = 4 + 4 = 8).
    fmt.Println(maximumAmount([][]int{{0,1,-1},{1,-2,3},{2,-3,4}})) // 8
    // Example 2:
    // Input: coins = [[10,10,10],[10,10,10]]
    // Output: 40
    // Explanation:
    // An optimal path for maximum coins is:
    // Start at (0, 0) with 10 coins (total coins = 10).
    // Move to (0, 1), gaining 10 coins (total coins = 10 + 10 = 20).
    // Move to (0, 2), gaining another 10 coins (total coins = 20 + 10 = 30).
    // Move to (1, 2), gaining the final 10 coins (total coins = 30 + 10 = 40).
    fmt.Println(maximumAmount([][]int{{10,10,10},{10,10,10}})) // 40

    fmt.Println(maximumAmount1([][]int{{0,1,-1},{1,-2,3},{2,-3,4}})) // 8
    fmt.Println(maximumAmount1([][]int{{10,10,10},{10,10,10}})) // 40
}
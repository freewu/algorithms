package main

// 3603. Minimum Cost Path with Alternating Directions II
// You are given two integers m and n representing the number of rows and columns of a grid, respectively.

// The cost to enter cell (i, j) is defined as (i + 1) * (j + 1).

// You are also given a 2D integer array waitCost where waitCost[i][j] defines the cost to wait on that cell.

// You start at cell (0, 0) at second 1.

// At each step, you follow an alternating pattern:
//     1. On odd-numbered seconds, you must move right or down to an adjacent cell, paying its entry cost.
//     2. On even-numbered seconds, you must wait in place, paying waitCost[i][j].

// Return the minimum total cost required to reach (m - 1, n - 1).

// Example 1:
// Input: m = 1, n = 2, waitCost = [[1,2]]
// Output: 3
// Explanation:
// The optimal path is:
// Start at cell (0, 0) at second 1 with entry cost (0 + 1) * (0 + 1) = 1.
// Second 1: Move right to cell (0, 1) with entry cost (0 + 1) * (1 + 1) = 2.
// Thus, the total cost is 1 + 2 = 3.

// Example 2:
// Input: m = 2, n = 2, waitCost = [[3,5],[2,4]]
// Output: 9
// Explanation:
// The optimal path is:
// Start at cell (0, 0) at second 1 with entry cost (0 + 1) * (0 + 1) = 1.
// Second 1: Move down to cell (1, 0) with entry cost (1 + 1) * (0 + 1) = 2.
// Second 2: Wait at cell (1, 0), paying waitCost[1][0] = 2.
// Second 3: Move right to cell (1, 1) with entry cost (1 + 1) * (1 + 1) = 4.
// Thus, the total cost is 1 + 2 + 2 + 4 = 9.

// Example 3:
// Input: m = 2, n = 3, waitCost = [[6,1,4],[3,2,5]]
// Output: 16
// Explanation:
// The optimal path is:
// Start at cell (0, 0) at second 1 with entry cost (0 + 1) * (0 + 1) = 1.
// Second 1: Move right to cell (0, 1) with entry cost (0 + 1) * (1 + 1) = 2.
// Second 2: Wait at cell (0, 1), paying waitCost[0][1] = 1.
// Second 3: Move down to cell (1, 1) with entry cost (1 + 1) * (1 + 1) = 4.
// Second 4: Wait at cell (1, 1), paying waitCost[1][1] = 2.
// Second 5: Move right to cell (1, 2) with entry cost (1 + 1) * (2 + 1) = 6.
// Thus, the total cost is 1 + 2 + 1 + 4 + 2 + 6 = 16.

// Constraints:
//     1 <= m, n <= 10^5
//     2 <= m * n <= 10^5
//     waitCost.length == m
//     waitCost[0].length == n
//     0 <= waitCost[i][j] <= 10^5

import "fmt"

func minCost(m int, n int, waitCost [][]int) int64 {
    memo := make([][]int64, m + 1)
    for i := range memo {
        memo[i] = make([]int64, n + 1)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
    min := func (x, y int64) int64 { if x < y { return x; }; return y; }
    var dfs func(i, j int) int64
    dfs = func(i, j int) int64 {
        if i == m - 1 && j == n - 1 { return 0 }
        if i >= m || j >= n         { return 1e15 }
        if memo[i][j] != -1         { return memo[i][j]}

        memo[i][j] = int64(waitCost[i][j]) + min(
            dfs(i + 1, j) + int64((i + 2) * (j + 1)), 
            dfs(i, j + 1) + int64((i + 1) * (j + 2)),
        )
        return memo[i][j]
    }
    // add first entry cost (1) and remove first waitCost
    return dfs(0, 0) + 1 - int64(waitCost[0][0])
}

func minCost1(m, n int, f [][]int) int64 {
    f[0][0] = 1
    f[m-1][n-1] = 0
    for j := 1; j < n; j++ {
        f[0][j] += f[0][j-1] + j + 1
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i < m; i++ {
        f[i][0] += f[i-1][0] + i + 1
        for j := 1; j < n; j++ {
            f[i][j] += min(f[i][j-1], f[i-1][j]) + (i+1)*(j+1)
        }
    }
    return int64(f[m-1][n-1])
}

func main() {
    // Example 1:
    // Input: m = 1, n = 2, waitCost = [[1,2]]
    // Output: 3
    // Explanation:
    // The optimal path is:
    // Start at cell (0, 0) at second 1 with entry cost (0 + 1) * (0 + 1) = 1.
    // Second 1: Move right to cell (0, 1) with entry cost (0 + 1) * (1 + 1) = 2.
    // Thus, the total cost is 1 + 2 = 3.
    fmt.Println(minCost(1, 2, [][]int{{1,2}})) // 3
    // Example 2:
    // Input: m = 2, n = 2, waitCost = [[3,5],[2,4]]
    // Output: 9
    // Explanation:
    // The optimal path is:
    // Start at cell (0, 0) at second 1 with entry cost (0 + 1) * (0 + 1) = 1.
    // Second 1: Move down to cell (1, 0) with entry cost (1 + 1) * (0 + 1) = 2.
    // Second 2: Wait at cell (1, 0), paying waitCost[1][0] = 2.
    // Second 3: Move right to cell (1, 1) with entry cost (1 + 1) * (1 + 1) = 4.
    // Thus, the total cost is 1 + 2 + 2 + 4 = 9.
    fmt.Println(minCost(2, 2, [][]int{{3,5},{2,4}})) // 9
    // Example 3:
    // Input: m = 2, n = 3, waitCost = [[6,1,4],[3,2,5]]
    // Output: 16
    // Explanation:
    // The optimal path is:
    // Start at cell (0, 0) at second 1 with entry cost (0 + 1) * (0 + 1) = 1.
    // Second 1: Move right to cell (0, 1) with entry cost (0 + 1) * (1 + 1) = 2.
    // Second 2: Wait at cell (0, 1), paying waitCost[0][1] = 1.
    // Second 3: Move down to cell (1, 1) with entry cost (1 + 1) * (1 + 1) = 4.
    // Second 4: Wait at cell (1, 1), paying waitCost[1][1] = 2.
    // Second 5: Move right to cell (1, 2) with entry cost (1 + 1) * (2 + 1) = 6.
    // Thus, the total cost is 1 + 2 + 1 + 4 + 2 + 6 = 16.
    fmt.Println(minCost(2, 3, [][]int{{6,1,4},{3,2,5}})) // 16

    fmt.Println(minCost1(1, 2, [][]int{{1,2}})) // 3
    fmt.Println(minCost1(2, 2, [][]int{{3,5},{2,4}})) // 9
    fmt.Println(minCost1(2, 3, [][]int{{6,1,4},{3,2,5}})) // 16
}
package main

// 3393. Count Paths With the Given XOR Value
// You are given a 2D integer array grid with size m x n. 
// You are also given an integer k.

// Your task is to calculate the number of paths you can take from the top-left cell (0, 0) to the bottom-right cell (m - 1, n - 1) satisfying the following constraints:
//     1. You can either move to the right or down. Formally, from the cell (i, j) you may move to the cell (i, j + 1) or to the cell (i + 1, j) if the target cell exists.
//     2. The XOR of all the numbers on the path must be equal to k.

// Return the total number of such paths.

// Since the answer can be very large, return the result modulo 10^9 + 7.

// Example 1:
// Input: grid = [[2, 1, 5], [7, 10, 0], [12, 6, 4]], k = 11
// Output: 3
// Explanation: 
// The 3 paths are:
// (0, 0) → (1, 0) → (2, 0) → (2, 1) → (2, 2)
// (0, 0) → (1, 0) → (1, 1) → (1, 2) → (2, 2)
// (0, 0) → (0, 1) → (1, 1) → (2, 1) → (2, 2)

// Example 2:
// Input: grid = [[1, 3, 3, 3], [0, 3, 3, 2], [3, 0, 1, 1]], k = 2
// Output: 5
// Explanation:
// The 5 paths are:
// (0, 0) → (1, 0) → (2, 0) → (2, 1) → (2, 2) → (2, 3)
// (0, 0) → (1, 0) → (1, 1) → (2, 1) → (2, 2) → (2, 3)
// (0, 0) → (1, 0) → (1, 1) → (1, 2) → (1, 3) → (2, 3)
// (0, 0) → (0, 1) → (1, 1) → (1, 2) → (2, 2) → (2, 3)
// (0, 0) → (0, 1) → (0, 2) → (1, 2) → (2, 2) → (2, 3)

// Example 3:
// Input: grid = [[1, 1, 1, 2], [3, 0, 3, 2], [3, 0, 2, 2]], k = 10
// Output: 0

// Constraints:
//     1 <= m == grid.length <= 300
//     1 <= n == grid[r].length <= 300
//     0 <= grid[r][c] < 16
//     0 <= k < 16

import "fmt"

func countPathsWithXorValue(grid [][]int, k int) int {
    m, n, mod := len(grid), len(grid[0]), 1_000_000_007
    memo := make(map[int64]int)
    var dfs func(row, col, cur int) int
    dfs = func(row, col, cur int) int {
        if row == m - 1 && col == n - 1 {
            if (cur ^ grid[row][col]) == k { return 1 }
            return 0
        }
        key := int64(row) << 32 | int64(col) << 16 | int64(cur)
        if val, ok := memo[key]; ok {
            return val
        }
        paths := 0
        if row + 1 < m {
            paths = (paths + dfs(row+1, col, cur ^ grid[row][col])) % mod
        }
        if col + 1 < n {
            paths = (paths + dfs(row, col+1, cur ^ grid[row][col])) % mod
        }
        memo[key] = paths
        return paths
    }
    return dfs(0, 0, 0)
}

func countPathsWithXorValue1(grid [][]int, k int) int {
    m, n, mod := len(grid), len(grid[0]), 1_000_000_007
    dp := make([][16]int, n)
    dp[0][grid[0][0]] = 1
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i == 0 && j == 0 { continue }
            counts := [16]int{}
            for x := 0; x < 16; x++ {
                if i > 0 { counts[x] = dp[j][x ^ grid[i][j]] }
                if j > 0 { counts[x] = (counts[x] + dp[j-1][x ^ grid[i][j]]) % mod }
            }
            dp[j] = counts
        }
    }
    return dp[n-1][k]
}

func main() {
    // Example 1:
    // Input: grid = [[2, 1, 5], [7, 10, 0], [12, 6, 4]], k = 11
    // Output: 3
    // Explanation: 
    // The 3 paths are:
    // (0, 0) → (1, 0) → (2, 0) → (2, 1) → (2, 2)
    // (0, 0) → (1, 0) → (1, 1) → (1, 2) → (2, 2)
    // (0, 0) → (0, 1) → (1, 1) → (2, 1) → (2, 2)
    fmt.Println(countPathsWithXorValue([][]int{{2,1,5},{7,10,0},{12,6,4}}, 11)) // 3
    // Example 2:
    // Input: grid = [[1, 3, 3, 3], [0, 3, 3, 2], [3, 0, 1, 1]], k = 2
    // Output: 5
    // Explanation:
    // The 5 paths are:
    // (0, 0) → (1, 0) → (2, 0) → (2, 1) → (2, 2) → (2, 3)
    // (0, 0) → (1, 0) → (1, 1) → (2, 1) → (2, 2) → (2, 3)
    // (0, 0) → (1, 0) → (1, 1) → (1, 2) → (1, 3) → (2, 3)
    // (0, 0) → (0, 1) → (1, 1) → (1, 2) → (2, 2) → (2, 3)
    // (0, 0) → (0, 1) → (0, 2) → (1, 2) → (2, 2) → (2, 3)
    fmt.Println(countPathsWithXorValue([][]int{{1,3,3,3},{0,3,3,2},{3,0,1,1}}, 2)) // 5
    // Example 3:
    // Input: grid = [[1, 1, 1, 2], [3, 0, 3, 2], [3, 0, 2, 2]], k = 10
    // Output: 0
    fmt.Println(countPathsWithXorValue([][]int{{1,1,1,2},{3,0,3,2},{3,0,2,2}}, 10)) // 0

    fmt.Println(countPathsWithXorValue1([][]int{{2,1,5},{7,10,0},{12,6,4}}, 11)) // 3
    fmt.Println(countPathsWithXorValue1([][]int{{1,3,3,3},{0,3,3,2},{3,0,1,1}}, 2)) // 5
    fmt.Println(countPathsWithXorValue1([][]int{{1,1,1,2},{3,0,3,2},{3,0,2,2}}, 10)) // 0
}
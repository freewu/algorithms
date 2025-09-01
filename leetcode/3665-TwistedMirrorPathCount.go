package main

// 3665. Twisted Mirror Path Count
// Given an m x n binary grid grid where:
//     1. grid[i][j] == 0 represents an empty cell, and
//     2. grid[i][j] == 1 represents a mirror.

// A robot starts at the top-left corner of the grid (0, 0) and wants to reach the bottom-right corner (m - 1, n - 1). 
// It can move only right or down. 
// If the robot attempts to move into a mirror cell, it is reflected before entering that cell:
//     1. If it tries to move right into a mirror, it is turned down and moved into the cell directly below the mirror.
//     2. If it tries to move down into a mirror, it is turned right and moved into the cell directly to the right of the mirror.

// If this reflection would cause the robot to move outside the grid boundaries, the path is considered invalid and should not be counted.

// Return the number of unique valid paths from (0, 0) to (m - 1, n - 1).

// Since the answer may be very large, return it modulo 10^9 + 7.

// Note: If a reflection moves the robot into a mirror cell, the robot is immediately reflected again based on the direction it used to enter that mirror: if it entered while moving right, it will be turned down; if it entered while moving down, it will be turned right. 
// This process will continue until either the last cell is reached, the robot moves out of bounds or the robot moves to a non-mirror cell.

// Example 1:
// Input: grid = [[0,1,0],[0,0,1],[1,0,0]]
// Output: 5
// Explanation:
// Number	Full Path
// 1	(0, 0) → (0, 1) [M] → (1, 1) → (1, 2) [M] → (2, 2)
// 2	(0, 0) → (0, 1) [M] → (1, 1) → (2, 1) → (2, 2)
// 3	(0, 0) → (1, 0) → (1, 1) → (1, 2) [M] → (2, 2)
// 4	(0, 0) → (1, 0) → (1, 1) → (2, 1) → (2, 2)
// 5	(0, 0) → (1, 0) → (2, 0) [M] → (2, 1) → (2, 2)
// [M] indicates the robot attempted to enter a mirror cell and instead reflected.

// Example 2:
// Input: grid = [[0,0],[0,0]]
// Output: 2
// Explanation:
// Number	Full Path
// 1	(0, 0) → (0, 1) → (1, 1)
// 2	(0, 0) → (1, 0) → (1, 1)

// Example 3:
// Input: grid = [[0,1,1],[1,1,0]]
// Output: 1
// Explanation:
// Number	Full Path
// 1	(0, 0) → (0, 1) [M] → (1, 1) [M] → (1, 2)
// (0, 0) → (1, 0) [M] → (1, 1) [M] → (2, 1) goes out of bounds, so it is invalid.

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     2 <= m, n <= 500
//     grid[i][j] is either 0 or 1.
//     grid[0][0] == grid[m - 1][n - 1] == 0

import "fmt"

func uniquePaths(grid [][]int) int {
    res, m, n, mod := 0, len(grid), len(grid[0]), 1_000_000_007
    visited := make([][][]int, m)
    for i := 0; i < m; i++ {
        visited[i] = make([][]int, n)
        // 需要初始化为 -1，如果初始化为 0 则无法区分是真的无法访问还是能访问但方案数为 0，会超时
        for j := 0; j < n; j++ {
            visited[i][j] = []int{-1, -1}
        }
    }
    directions := [][]int{ []int{0, 1}, []int{1, 0},  }
    var dfs func(i, j, dir int) int
    dfs = func(i, j, dir int) int {
        if i < 0 || i >= m || j < 0 || j >= n {  // 越界
            return 0
        }
        if i == m - 1 && j == n - 1 {  // 走到终点则更新答案
            return 1
        }
        //  如果已经访问过则不重复计算
        if visited[i][j][dir] != -1 {
            return visited[i][j][dir] % mod
        }
        num := 0
        if grid[i][j] == 1 {
             // 如果到达镜子，则只能被反射
            newDir := directions[(dir + 1) % 2]
            num += dfs(i + newDir[0], j + newDir[1], (dir + 1) % 2)
        } else {
            // 如果不是镜子，则分别尝试向右向下
            for di, d := range directions {
                num += dfs(i + d[0], j + d[1], di)
            }
        }
        visited[i][j][dir] = num % mod
        return num % mod
    }
    for di, d := range directions {
        res += dfs(d[0], d[1], di)
    }
    return res % mod
}

func uniquePaths1(grid [][]int) (ans int) {
    m, n, mod := len(grid), len(grid[0]), 1_000_000_007
    dp := make([][][2]int, m + 1)
    for i := range dp {
        dp[i] = make([][2]int, n + 1)
    }
    dp[0][1] = [2]int{1, 1} // 原理见 62 题我的题解
    for i, row := range grid {
        for j, x := range row {
            if x == 0 {
                dp[i+1][j+1][0] = (dp[i+1][j][0] + dp[i][j+1][1]) % mod
                dp[i+1][j+1][1] = dp[i+1][j+1][0]
            } else {
                dp[i+1][j+1][0] = dp[i][j+1][1]
                dp[i+1][j+1][1] = dp[i+1][j][0]
            }
        }
    }
    return dp[m][n][0]
}

func main() {
    // Example 1:
    // Input: grid = [[0,1,0],[0,0,1],[1,0,0]]
    // Output: 5
    // Explanation:
    // Number	Full Path
    // 1	(0, 0) → (0, 1) [M] → (1, 1) → (1, 2) [M] → (2, 2)
    // 2	(0, 0) → (0, 1) [M] → (1, 1) → (2, 1) → (2, 2)
    // 3	(0, 0) → (1, 0) → (1, 1) → (1, 2) [M] → (2, 2)
    // 4	(0, 0) → (1, 0) → (1, 1) → (2, 1) → (2, 2)
    // 5	(0, 0) → (1, 0) → (2, 0) [M] → (2, 1) → (2, 2)
    // [M] indicates the robot attempted to enter a mirror cell and instead reflected.
    fmt.Println(uniquePaths([][]int{{0,1,0},{0,0,1},{1,0,0}})) // 5
    // Example 2:
    // Input: grid = [[0,0],[0,0]]
    // Output: 2
    // Explanation:
    // Number	Full Path
    // 1	(0, 0) → (0, 1) → (1, 1)
    // 2	(0, 0) → (1, 0) → (1, 1)
    fmt.Println(uniquePaths([][]int{{0,0},{0,0}})) // 2
    // Example 3:
    // Input: grid = [[0,1,1],[1,1,0]]
    // Output: 1
    // Explanation:
    // Number	Full Path
    // 1	(0, 0) → (0, 1) [M] → (1, 1) [M] → (1, 2)
    // (0, 0) → (1, 0) [M] → (1, 1) [M] → (2, 1) goes out of bounds, so it is invalid.
    fmt.Println(uniquePaths([][]int{{0,1,1},{1,1,0}})) // 1

    fmt.Println(uniquePaths1([][]int{{0,1,0},{0,0,1},{1,0,0}})) // 5
    fmt.Println(uniquePaths1([][]int{{0,0},{0,0}})) // 2
    fmt.Println(uniquePaths1([][]int{{0,1,1},{1,1,0}})) // 1
}
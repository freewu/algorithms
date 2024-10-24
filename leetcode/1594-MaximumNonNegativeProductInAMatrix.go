package main

// 1594. Maximum Non Negative Product in a Matrix
// You are given a m x n matrix grid. 
// Initially, you are located at the top-left corner (0, 0), 
// and in each step, you can only move right or down in the matrix.

// Among all possible paths starting from the top-left corner (0, 0) 
// and ending in the bottom-right corner (m - 1, n - 1), find the path with the maximum non-negative product. 
// The product of a path is the product of all integers in the grid cells visited along the path.

// Return the maximum non-negative product modulo 10^9 + 7. If the maximum product is negative, return -1.

// Notice that the modulo is performed after getting the maximum product.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/12/23/product1.jpg" />
// Input: grid = [[-1,-2,-3],[-2,-3,-3],[-3,-3,-2]]
// Output: -1
// Explanation: It is not possible to get non-negative product in the path from (0, 0) to (2, 2), so return -1.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/12/23/product2.jpg" />
// Input: grid = [[1,-2,1],[1,-2,1],[3,-4,1]]
// Output: 8
// Explanation: Maximum non-negative product is shown (1 * 1 * -2 * -4 * 1 = 8).

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/12/23/product3.jpg" />
// Input: grid = [[1,3],[0,-4]]
// Output: 0
// Explanation: Maximum non-negative product is shown (1 * 0 * -4 = 0).

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 15
//     -4 <= grid[i][j] <= 4

import "fmt"

func maxProductPath(grid [][]int) int {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    type Pair struct {
        Max, Min int
    }
    n, m := len(grid), len(grid[0])
    dp := make([][]Pair, n)
    for i := range dp {
        dp[i] = make([]Pair, m)
    }
    dp[0][0] = Pair{ grid[0][0], grid[0][0] }
    for j := 1; j < m; j++ {
        dp[0][j] = Pair{ dp[0][j-1].Max * grid[0][j], dp[0][j-1].Min * grid[0][j] }
    }
    for i := 1; i < n; i++ {
        dp[i][0] = Pair{ dp[i-1][0].Max * grid[i][0], dp[i-1][0].Min * grid[i][0] }
    }
    for i := 1; i < n; i++ {
        for j := 1; j < m; j++ {
            p1, p2 := min(dp[i-1][j].Min, dp[i][j-1].Min) * grid[i][j], max(dp[i-1][j].Max, dp[i][j-1].Max) * grid[i][j]
            dp[i][j] = Pair{ max(p1, p2), min(p1, p2) }
        }
    }
    if dp[n-1][m-1].Max < 0 {
        return -1
    }
    return dp[n-1][m-1].Max % 1_000_000_007
}

func maxProductPath1(grid [][]int) int {
    // 在从左上角 (0, 0) 开始到右下角 (m - 1, n - 1) 结束的所有路径中，找出具有 最大非负积 的路径
    // 向下或向右
    // 存储正数和负数
    n, m := len(grid), len(grid[0])
    dpMin, dpMax := make([][]int, n+1), make([][]int, n+1)
    for i := 0; i < n; i++ {
        dpMin[i], dpMax[i] = make([]int, m+1), make([]int, m+1)
        for j := 0; j < m; j++ {
            dpMin[i][j], dpMax[i][j] = 100, -100
        }
    }
    dpMin[0][0], dpMax[0][0] = grid[0][0], grid[0][0]
    for i := 1; i < n; i++ {
        dpMin[i][0], dpMax[i][0] = dpMin[i-1][0] * grid[i][0], dpMax[i-1][0] * grid[i][0]
    }
    for i := 1; i < m; i++ {
        dpMin[0][i], dpMax[0][i] = dpMin[0][i-1] * grid[0][i], dpMax[0][i-1] * grid[0][i]
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < n; i++ {
        for j := 1; j < m; j++ {
            dpMax[i][j] = max(max(dpMax[i-1][j] * grid[i][j], dpMax[i][j-1] * grid[i][j]), max(dpMin[i-1][j] * grid[i][j], dpMin[i][j-1] * grid[i][j]))
            dpMin[i][j] = min(min(dpMax[i-1][j] * grid[i][j], dpMax[i][j-1] * grid[i][j]), min(dpMin[i-1][j] * grid[i][j], dpMin[i][j-1] * grid[i][j]))
        }
    }
    if dpMax[n-1][m-1] < 0 {
        return -1
    }
    return dpMax[n-1][m-1] % 1_000_000_007
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/12/23/product1.jpg" />
    // Input: grid = [[-1,-2,-3],[-2,-3,-3],[-3,-3,-2]]
    // Output: -1
    // Explanation: It is not possible to get non-negative product in the path from (0, 0) to (2, 2), so return -1.
    fmt.Println(maxProductPath([][]int{{-1,-2,-3},{-2,-3,-3},{-3,-3,-2}})) // -1
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/12/23/product2.jpg" />
    // Input: grid = [[1,-2,1],[1,-2,1],[3,-4,1]]
    // Output: 8
    // Explanation: Maximum non-negative product is shown (1 * 1 * -2 * -4 * 1 = 8).
    fmt.Println(maxProductPath([][]int{{1,-2,1},{1,-2,1},{3,-4,1}})) // 8
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/12/23/product3.jpg" />
    // Input: grid = [[1,3],[0,-4]]
    // Output: 0
    // Explanation: Maximum non-negative product is shown (1 * 0 * -4 = 0).
    fmt.Println(maxProductPath([][]int{{1,3},{0,-4}})) // 0

    fmt.Println(maxProductPath1([][]int{{-1,-2,-3},{-2,-3,-3},{-3,-3,-2}})) // -1
    fmt.Println(maxProductPath1([][]int{{1,-2,1},{1,-2,1},{3,-4,1}})) // 8
    fmt.Println(maxProductPath1([][]int{{1,3},{0,-4}})) // 0
}
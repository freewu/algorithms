package main

// 3148. Maximum Difference Score in a Grid
// You are given an m x n matrix grid consisting of positive integers. 
// You can move from a cell in the matrix to any other cell that is either to the bottom or to the right (not necessarily adjacent). 
// The score of a move from a cell with the value c1 to a cell with the value c2 is c2 - c1.
// You can start at any cell, and you have to make at least one move.

// Return the maximum total score you can achieve.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2024/03/14/grid1.png" />
// Input: grid = [[9,5,7,3],[8,9,6,1],[6,7,14,3],[2,5,3,1]]
// Output: 9
// Explanation: We start at the cell (0, 1), and we perform the following moves:
// - Move from the cell (0, 1) to (2, 1) with a score of 7 - 5 = 2.
// - Move from the cell (2, 1) to (2, 2) with a score of 14 - 7 = 7.
// The total score is 2 + 7 = 9.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2024/04/08/moregridsdrawio-1.png" />
// Input: grid = [[4,3,2],[3,2,1]]
// Output: -1
// Explanation: We start at the cell (0, 0), and we perform one move: (0, 0) to (0, 1). The score is 3 - 4 = -1.

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     2 <= m, n <= 1000
//     4 <= m * n <= 10^5
//     1 <= grid[i][j] <= 10^5

import "fmt"

func maxScore(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    res, dp := -1 << 31, make([][]int, m)
    for i, _ := range dp {
        dp[i] = make([]int, n)
    }
    dp[m - 1][n - 1] = grid[m -1][n - 1]
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := m - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            if i < m - 1 {
                dp[i][j] = max(dp[i][j], dp[i + 1][j])
            }
            if j < n - 1 {
                dp[i][j] = max(dp[i][j], dp[i][j + 1])
            }
            dp[i][j] = max(dp[i][j], grid[i][j])
        }
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i < m - 1 {
                res = max(res, dp[i + 1][j] - grid[i][j])
            }
            if j < n - 1 {
                res = max(res, dp[i][j + 1] - grid[i][j])
            }
        }
    }
    return res
}

func maxScore1(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    res, inf := -1 << 31, 1 << 31
    dp := make([][]int, m+1) // use sentinels
    for i := 0; i <= m; i++ {
        dp[i] = make([]int, n+1)
    }
    for i := 0; i <= n; i++ {
        dp[0][i] = inf / 2
    }
    for i := 0; i <= m; i++ {
        dp[i][0] = inf / 2
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            dp[i][j] = min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1]))
            cur := grid[i-1][j-1] - dp[i][j]
            res = max(res, cur)
            dp[i][j] = min(dp[i][j], grid[i-1][j-1])
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2024/03/14/grid1.png" />
    // Input: grid = [[9,5,7,3],[8,9,6,1],[6,7,14,3],[2,5,3,1]]
    // Output: 9
    // Explanation: We start at the cell (0, 1), and we perform the following moves:
    // - Move from the cell (0, 1) to (2, 1) with a score of 7 - 5 = 2.
    // - Move from the cell (2, 1) to (2, 2) with a score of 14 - 7 = 7.
    // The total score is 2 + 7 = 9.
    grid1 := [][]int{
        {9,5,7,3},
        {8,9,6,1},
        {6,7,14,3},
        {2,5,3,1},
    }
    fmt.Println(maxScore(grid1)) // 9
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2024/04/08/moregridsdrawio-1.png" />
    // Input: grid = [[4,3,2],[3,2,1]]
    // Output: -1
    // Explanation: We start at the cell (0, 0), and we perform one move: (0, 0) to (0, 1). The score is 3 - 4 = -1.
    grid2 := [][]int{
        {4,3,2},
        {3,2,1},
    }
    fmt.Println(maxScore(grid2)) // -1

    grid3 := [][]int{
        {9,6},
        {4,2},
    }
    fmt.Println(maxScore(grid3)) // -2

    fmt.Println(maxScore1(grid1)) // 9
    fmt.Println(maxScore1(grid2)) // -1
    fmt.Println(maxScore1(grid3)) // -2
}
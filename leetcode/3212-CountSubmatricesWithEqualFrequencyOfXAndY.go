package main

// 3212. Count Submatrices With Equal Frequency of X and Y
// Given a 2D character matrix grid, where grid[i][j] is either 'X', 'Y', or '.', 
// return the number of submatrices that contain:
//     grid[0][0]
//     an equal frequency of 'X' and 'Y'.
//     at least one 'X'.

// Example 1:
// Input: grid = [["X","Y","."],["Y",".","."]]
// Output: 3
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/06/07/examplems.png" />

// Example 2:
// Input: grid = [["X","X"],["X","Y"]]
// Output: 0
// Explanation:
// No submatrix has an equal frequency of 'X' and 'Y'.

// Example 3:
// Input: grid = [[".","."],[".","."]]
// Output: 0
// Explanation:
// No submatrix has at least one 'X'.

// Constraints:
//     1 <= grid.length, grid[i].length <= 1000
//     grid[i][j] is either 'X', 'Y', or '.'.

import "fmt"

func numberOfSubmatrices(grid [][]byte) int {
    res, m, n := 0, len(grid), len(grid[0])
    prefixX, prefixY := make([][]int, m + 1), make([][]int, m + 1)
    for i := range prefixX {
        prefixX[i], prefixY[i] = make([]int, n + 1), make([]int, n + 1)
    }
    for i, row := range grid {
        for j, v := range row {
            prefixX[i+1][j+1] = prefixX[i+1][j] + prefixX[i][j+1] - prefixX[i][j]
            prefixY[i+1][j+1] = prefixY[i+1][j] + prefixY[i][j+1] - prefixY[i][j]
            switch v {
                case 'X':
                    prefixX[i+1][j+1]++
                case 'Y':
                    prefixY[i+1][j+1]++
            }
            if prefixX[i+1][j+1] > 0 && prefixX[i+1][j+1] == prefixY[i+1][j+1] {
                res++
            }
        }
    }
    return res
}

func numberOfSubmatrices1(grid [][]byte) int {
    if len(grid) == 0 { return 0 }
    type Pair struct { x, y int }
    res, m, n := 0, len(grid), len(grid[0])
    dp := make([]Pair, n)
    for x, y, i := 0, 0, 0; i < n; i++ {
        if grid[0][i] == 'X' {
            x++
        } else if grid[0][i] == 'Y' {
            y++
        }
        dp[i].x, dp[i].y = x, y
        if dp[i].x == dp[i].y && dp[i].x > 0 {
            res++
        }
    }
    for i := 1; i < m; i++ {
        next := make([]Pair, n)
        next[0] = dp[0]
        if grid[i][0] == 'X' {
            next[0].x++
        } else if grid[i][0] == 'Y'  {
            next[0].y++
        }
        if next[0].x == next[0].y && next[0].x > 0 {
            res++
        }
        for j := 1; j < n; j++ {
            x, y := 0, 0
            if grid[i][j] == 'X' {
                x++
            } else if grid[i][j] == 'Y' {
                y++
            }
            next[j].x = x + next[j - 1].x + dp[j].x - dp[j - 1].x
            next[j].y = y + next[j - 1].y + dp[j].y - dp[j - 1].y
            if next[j].x == next[j].y && next[j].x > 0 {
                res++
            }
        }
        dp = next
    }
    return res
}

func main() {
    // Example 1:
    // Input: grid = [["X","Y","."],["Y",".","."]]
    // Output: 3
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/06/07/examplems.png" />
    fmt.Println(numberOfSubmatrices([][]byte{{'X','Y','.'},{'Y','.','.'}})) // 3
    // Example 2:
    // Input: grid = [["X","X"],["X","Y"]]
    // Output: 0
    // Explanation:
    // No submatrix has an equal frequency of 'X' and 'Y'.
    fmt.Println(numberOfSubmatrices([][]byte{{'X','X'},{'X','Y'}})) // 0
    // Example 3:
    // Input: grid = [[".","."],[".","."]]
    // Output: 0
    // Explanation:
    // No submatrix has at least one 'X'.
    fmt.Println(numberOfSubmatrices([][]byte{{'.','.'},{'.','.'}})) // 0

    fmt.Println(numberOfSubmatrices1([][]byte{{'X','Y','.'},{'Y','.','.'}})) // 3
    fmt.Println(numberOfSubmatrices1([][]byte{{'X','X'},{'X','Y'}})) // 0
    fmt.Println(numberOfSubmatrices1([][]byte{{'.','.'},{'.','.'}})) // 0
}
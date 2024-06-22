package main

// 562. Longest Line of Consecutive One in Matrix
// Given an m x n binary matrix mat, return the length of the longest line of consecutive one in the matrix.
// The line could be horizontal, vertical, diagonal, or anti-diagonal.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/04/24/long1-grid.jpg" />
// Input: mat = [[0,1,1,0],[0,1,1,0],[0,0,0,1]]
// Output: 3

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/04/24/long2-grid.jpg" />
// Input: mat = [[1,1,1,1],[0,1,1,0],[0,0,0,1]]
// Output: 4

// Constraints:
//     m == mat.length
//     n == mat[i].length
//     1 <= m, n <= 10^4
//     1 <= m * n <= 10^4
//     mat[i][j] is either 0 or 1.

import "fmt"

func longestLine(mat [][]int) int {
    m, n, res := len(mat), len(mat[0]), 0
    dp := make([][][4]int, m + 1)
    for i := range dp {
        dp[i] = make([][4]int, n+1)
    }
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if mat[i-1][j-1] == 1 {
                dp[i][j][0] = dp[i][j-1][0] + 1 // 水平
                dp[i][j][1] = dp[i-1][j][1] + 1 // 垂直
                dp[i][j][2] = dp[i-1][j-1][2] + 1 // 对角线
                dp[i][j][3] = 1 // 反对角线
                if j + 1 <= n {
                    dp[i][j][3] += dp[i-1][j+1][3] 
                }
                for _, v := range dp[i][j] {
                    if v > res {
                        res = v
                    }
                }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/04/24/long1-grid.jpg" />
    // Input: mat = [[0,1,1,0],[0,1,1,0],[0,0,0,1]]
    // Output: 3
    mat1 := [][]int{
        {0,1,1,0},
        {0,1,1,0},
        {0,0,0,1},
    }
    fmt.Println(longestLine(mat1)) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/04/24/long2-grid.jpg" />
    // Input: mat = [[1,1,1,1],[0,1,1,0],[0,0,0,1]]
    // Output: 4
    mat2 := [][]int{
        {1,1,1,1},
        {0,1,1,0},
        {0,0,0,1},
    }
    fmt.Println(longestLine(mat2)) // 4
}
package main

// 62. Unique Paths
// There is a robot on an m x n grid. 
// The robot is initially located at the top-left corner (i.e., grid[0][0]). 
// The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]). 
// The robot can only move either down or right at any point in time.

// Given the two integers m and n, return the number of possible unique paths that the robot can take to reach the bottom-right corner.
// The test cases are generated so that the answer will be less than or equal to 2 * 109.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2018/10/22/robot_maze.png" />
// Input: m = 3, n = 7
// Output: 28

// Example 2:
// Input: m = 3, n = 2
// Output: 3
// Explanation: From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
// 1. Right -> Down -> Down
// 2. Down -> Down -> Right
// 3. Down -> Right -> Down

// Constraints:
//     1 <= m, n <= 100

import "fmt"

func uniquePaths(m int, n int) int {
    res := make([][]int, m)
    for i := 0; i<m; i++{
        res[i] = make([]int,n)
    }
    for i := 0; i < m; i++ {
        res[i][0] = 1
    }
    for j := 1; j < n; j++ {
        res[0][j] = 1
    }
    for i:= 1; i < m; i++ {
        for j := 1; j < n; j++ {
            // 由于机器人只能向右走和向下走，所以地图的第一行和第一列的走法数都是 1，
            // 地图中任意一点的走法数是 dp[i][j] = dp[i-1][j] + dp[i][j-1]
            res[i][j] = res[i-1][j] + res[i][j-1]
        }
    }
    //fmt.Println(res)
    return res[m-1][n-1]
}

func uniquePaths1(m int, n int) int {
    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, m)
    }
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if i == 0 || j == 0 {
                // 由于机器人只能向右走和向下走，所以地图的第一行和第一列的走法数都是 1，
                dp[i][j] = 1
                continue
            }
            // 地图中任意一点的走法数是 dp[i][j] = dp[i-1][j] + dp[i][j-1]
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
        }
    }
    return dp[n-1][m-1]
}

func main() {
    fmt.Println(uniquePaths(3,7)) // 28
    // Explanation: From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
    // 1. Right -> Down -> Down
    // 2. Down -> Down -> Right
    // 3. Down -> Right -> Down
    fmt.Println(uniquePaths(3,2)) // 3

    fmt.Println(uniquePaths1(3,7)) // 28
    fmt.Println(uniquePaths1(3,2)) // 3
}
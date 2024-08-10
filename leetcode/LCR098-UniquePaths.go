package main

// LCR 098. 不同路径
// 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
// 问总共有多少条不同的路径？

// 示例 1：
// <img src="https://assets.leetcode.com/uploads/2018/10/22/robot_maze.png" />
// 输入：m = 3, n = 7
// 输出：28

// 示例 2：
// 输入：m = 3, n = 2
// 输出：3
// 解释：
// 从左上角开始，总共有 3 条路径可以到达右下角。
// 1. 向右 -> 向下 -> 向下
// 2. 向下 -> 向下 -> 向右
// 3. 向下 -> 向右 -> 向下

// 示例 3：
// 输入：m = 7, n = 3
// 输出：28

// 示例 4：
// 输入：m = 3, n = 3
// 输出：6

// 提示：
//     1 <= m, n <= 100
//     题目数据保证答案小于等于 2 * 10^9

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
package main

// 面试题 08.02. Robot in a Grid LCCI
// Imagine a robot sitting on the upper left corner of grid with r rows and c columns. 
// The robot can only move in two directions, right and down, but certain cells are "off limits" such that the robot cannot step on them. 
// Design an algorithm to find a path for the robot from the top left to the bottom right.

// <img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/22/robot_maze.png" />

// "off limits" and empty grid are represented by 1 and 0 respectively.

// Return a valid path, consisting of row number and column number of grids in the path.

// Example 1:
// Input:
// [
//   [0,0,0],
//   [0,1,0],
//   [0,0,0]
// ]
// Output: [[0,0],[0,1],[0,2],[1,2],[2,2]]

// Note:
//     r, c <= 100

import "fmt"

func pathWithObstacles(obstacleGrid [][]int) [][]int {
    n, m := len(obstacleGrid), len(obstacleGrid[0])
    dp := make([][]bool, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]bool, m)
    }
    dp[0][0] = obstacleGrid[0][0] == 0
    for i := 1; i < n; i++ { // 边界情况单独处理
        if obstacleGrid[i][0] == 0 {
            dp[i][0] = dp[i-1][0]
        }
    }
    for j := 1; j < m; j++ { // 边界情况单独处理
        if obstacleGrid[0][j] == 0{
            dp[0][j] = dp[0][j-1]
        }
    }
    // 因为题目要求只能走下或走右，所以(i, j)是否可达取决于(i-1, j)或者(i, j-1)是否可达
    for i := 1; i < n; i++ {
        for j := 1; j < m; j++ {
            if obstacleGrid[i][j] == 0 {
                dp[i][j] = dp[i-1][j] || dp[i][j-1]
            }
        }
    }
    if !dp[n-1][m-1] {  return [][]int{} }
    res, index := make([][]int, n + m - 1), n + m - 2
    res[index] = []int{ n - 1, m - 1 }  // 从终点开始倒推路径
    index--
    dirs := [][]int{{-1, 0}, {0, -1}} // 倒推方向为向左或向上
    for index >= 0 {
        lastCell := res[index + 1]
        for _, d := range dirs {
            x, y := lastCell[0] + d[0], lastCell[1] + d[1]
            if x >= 0 && y >= 0 && dp[x][y] {
                res[index] = []int{ x, y }
                break
            }
        }
        index--
    }
    return res
}

func pathWithObstacles1(obstacleGrid [][]int) [][]int {
    m, n := len(obstacleGrid), len(obstacleGrid[0])
    dp := make([][]bool,m)
    for i := 0; i < m; i++ {
        dp[i] = make([]bool,n)
        for j := 0;j < n;j++ {
            if obstacleGrid[i][j] == 1 { continue }
            if i == 0 && j == 0 { dp[i][j] = true }
            if i > 0 { dp[i][j] = dp[i][j] || dp[i - 1][j] }
            if j > 0 { dp[i][j] = dp[i][j] || dp[i][j - 1] }
        }
    }
    if !dp[m-1][n-1] { return [][]int{} }
    res, x, y := make([][]int,m + n - 1),m - 1,n - 1
    res[m + n - 2] = []int{ x, y }
    for i := m + n - 3; i >= 0; i-- {
        if x > 0 && dp[x - 1][y] {
            x--
        } else {
            y--
        }
        res[i] = []int{ x, y }
    }
    return res
}

func main() {
    // Example 1:
    // Input:
    // [
    //   [0,0,0],
    //   [0,1,0],
    //   [0,0,0]
    // ]
    // Output: [[0,0],[0,1],[0,2],[1,2],[2,2]]
    grid1 := [][]int{
        {0,0,0}, 
        {0,1,0}, 
        {0,0,0}, 
    }
    fmt.Println(pathWithObstacles(grid1)) // [[0,0],[0,1],[0,2],[1,2],[2,2]]

    grid2 := [][]int{
        {0,0,0,0}, 
        {1,1,1,0}, 
        {0,0,0,0}, 
    }
    fmt.Println(pathWithObstacles(grid2)) // [[0 0] [0 1] [0 2] [0 3] [1 3] [2 3]]

    fmt.Println(pathWithObstacles1(grid1)) // [[0,0],[0,1],[0,2],[1,2],[2,2]]
    fmt.Println(pathWithObstacles1(grid2)) // [[0 0] [0 1] [0 2] [0 3] [1 3] [2 3]]
}
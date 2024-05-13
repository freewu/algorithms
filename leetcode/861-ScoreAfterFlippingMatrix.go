package main

// 861. Score After Flipping Matrix
// You are given an m x n binary matrix grid.
// A move consists of choosing any row or column and toggling each value in that row or column (i.e., changing all 0's to 1's, and all 1's to 0's).
// Every row of the matrix is interpreted as a binary number, and the score of the matrix is the sum of these numbers.
// Return the highest possible score after making any number of moves (including zero moves).

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/07/23/lc-toogle1.jpg" />
// Input: grid = [[0,0,1,1],[1,0,1,0],[1,1,0,0]]
// Output: 39
// Explanation: 0b1111 + 0b1001 + 0b1111 = 15 + 9 + 15 = 39

// Example 2:
// Input: grid = [[0]]
// Output: 1
 
// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 20
//     grid[i][j] is either 0 or 1

import "fmt"

// 要想二元矩阵最终得分最大，必须首先确保第一列全为 1（可以通过行或列移动达到这一目的）
// 之后从低列到高列查看其它列 1 和 0 的个数，0 多则进行列移动
func matrixScore(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    res := m * (1 << (n-1)) // 第一列全 1 贡献的分数
    for j := 1; j < n; j++ {
        oneCnt := 0
        for i := 0; i < m; i++ {
            if grid[i][j] == grid[i][0] { // 和行首元素相同，行移动后值保持相同
                oneCnt++
            }
        }
        oneCnt = max(oneCnt, m - oneCnt)
        res += oneCnt * (1 << (n - j - 1)) // 第 j 列所有的 1 贡献的分数
    }
    return res 
}

func matrixScore1(grid [][]int) int {
    score := 0
    flipRow := func (grid [][]int, row int) {
        for i := 0; i < len(grid[row]); i++ {
            if grid[row][i] == 1 { grid[row][i] = 0; } else {; grid[row][i] = 1; }
        }
    }
    flipCol := func (grid [][]int, col int) {
        for i := 0; i < len(grid); i++ {
            if grid[i][col] == 1 {  grid[i][col] = 0 ; } else {; grid[i][col] = 1; }
        }
    }
    for i := 0; i < len(grid); i++ {
        if grid[i][0] == 0 {
            flipRow(grid, i)
        }
    }
    for i := 0; i < len(grid[0]); i++ {
        n0 := 0
        n1 := 0
        for j := 0; j < len(grid); j++ {
            if grid[j][i] == 1 {
                n1++
            } else {
                n0++
            }
        }
        if n0 > n1 {
            flipCol(grid, i)
        }
    }
    for _, v := range grid {
        n := 0
        for _, b := range v {
            n = n << 1 + b
        }
        score += n
    }
    return score
}

func main() {
    // Example 1:
    // Input: grid = [[0,0,1,1],[1,0,1,0],[1,1,0,0]]
    // Output: 39
    // Explanation: 0b1111 + 0b1001 + 0b1111 = 15 + 9 + 15 = 39
    fmt.Println(matrixScore([][]int{{0,0,1,1},{1,0,1,0},{1,1,0,0}})) // 39
    // Example 2:
    // Input: grid = [[0]]
    // Output: 1
    fmt.Println(matrixScore([][]int{{0}})) // 1

    fmt.Println(matrixScore1([][]int{{0,0,1,1},{1,0,1,0},{1,1,0,0}})) // 39
    fmt.Println(matrixScore1([][]int{{0}})) // 1
}

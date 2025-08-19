package main

// 1277. Count Square Submatrices with All Ones
// Given a m * n matrix of ones and zeros, return how many square submatrices have all ones.

// Example 1:
// Input: matrix =
// [
//   [0,1,1,1],
//   [1,1,1,1],
//   [0,1,1,1]
// ]
// Output: 15
// Explanation: 
// There are 10 squares of side 1.
// There are 4 squares of side 2.
// There is  1 square of side 3.
// Total number of squares = 10 + 4 + 1 = 15.

// Example 2:
// Input: matrix = 
// [
//   [1,0,1],
//   [1,1,0],
//   [1,1,0]
// ]
// Output: 7
// Explanation: 
// There are 6 squares of side 1.  
// There is 1 square of side 2. 
// Total number of squares = 6 + 1 = 7.

// Constraints:
//     1 <= arr.length <= 300
//     1 <= arr[0].length <= 300
//     0 <= arr[i][j] <= 1

import "fmt"

func countSquares(matrix [][]int) int {
    dp := make([][]int, len(matrix) + 1)
    for i := range dp {
        dp[i] = make([]int, len(matrix[0]) + 1)
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := range matrix {
        for j, v := range matrix[i] {
            if v == 0 { continue }
            t := min(dp[i][j+1], dp[i+1][j])
            if matrix[i-t][j-t] != 0 {
                dp[i+1][j+1] = t + 1
            } else {
                dp[i+1][j+1] = t
            }
        }
    }
    squares := 0
    for i := range dp {
        for _, v := range dp[i] {
            squares += v
        }
    }
    return squares
}

func countSquares1(matrix [][]int) int {
    // f(i,j): 以matrix[i][j]为右下角的正方向的最大边长， 
    // 同时f[i][j] = x 也表示以 (i, j) 为右下角的正方形的数目为 x（即边长为 1, 2, ..., x 的正方形各一个）。
    // 在计算出所有的 f[i][j] 后，我们将它们进行累加，就可以得到矩阵中正方形的数目。
    res, dp := 0, make([][]int, len(matrix))
    for row := 0; row < len(matrix); row++ {
        dp[row] = make([]int, len(matrix[row]))
        for col := 0; col < len(matrix[row]); col++ {
            // 顺便初始化了 row=0 或者 col = 0 的 dp = matrix[row][col]
            dp[row][col] = matrix[row][col]
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for row := 0; row < len(matrix); row++ {
        for col:=0; col < len(matrix[row]); col++{
            if row != 0 && col != 0 && dp[row][col] == 1 {
                dp[row][col] = min(dp[row-1][col-1], min(dp[row-1][col], dp[row][col-1])) + 1
            }
            res += dp[row][col]
        }
    }
    return res
}

func countSquares2(matrix [][]int) int {
    res, m, n := 0, len(matrix), len(matrix[0])
    f := make([][]int, m)
    for i := range f {
        f[i] = make([]int, n)
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i == 0 || j == 0 {
                f[i][j] = matrix[i][j]
            } else if matrix[i][j] == 0 {
                f[i][j] = 0
            } else {
                f[i][j] = min(min(f[i][j - 1], f[i - 1][j]), f[i - 1][j - 1]) + 1
            }
            res += f[i][j]
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: matrix =
    // [
    //   [0,1,1,1],
    //   [1,1,1,1],
    //   [0,1,1,1]
    // ]
    // Output: 15
    // Explanation: 
    // There are 10 squares of side 1.
    // There are 4 squares of side 2.
    // There is  1 square of side 3.
    // Total number of squares = 10 + 4 + 1 = 15.
    matrix1 := [][]int{
        {0,1,1,1},
        {1,1,1,1},
        {0,1,1,1},
    }
    fmt.Println(countSquares(matrix1)) // 15
    // Example 2:
    // Input: matrix = 
    // [
    //   [1,0,1],
    //   [1,1,0],
    //   [1,1,0]
    // ]
    // Output: 7
    // Explanation: 
    // There are 6 squares of side 1.  
    // There is 1 square of side 2. 
    // Total number of squares = 6 + 1 = 7.
    matrix2 := [][]int{
        {1,0,1},
        {1,1,0},
        {1,1,0},
    }
    fmt.Println(countSquares(matrix2)) // 7

    fmt.Println(countSquares1(matrix1)) // 15
    fmt.Println(countSquares1(matrix2)) // 7

    fmt.Println(countSquares2(matrix1)) // 15
    fmt.Println(countSquares2(matrix2)) // 7
}
package main

// 931. Minimum Falling Path Sum
// Given an n x n array of integers matrix, return the minimum sum of any falling path through matrix.
// A falling path starts at any element in the first row and chooses the element in the next row that is either directly below or diagonally left/right. 
// Specifically, the next element from position (row, col) will be (row + 1, col - 1), (row + 1, col), or (row + 1, col + 1).
 
// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/11/03/failing1-grid.jpg" />
// Input: matrix = [[2,1,3],[6,5,4],[7,8,9]]
// Output: 13
// Explanation: There are two falling paths with a minimum sum as shown.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/11/03/failing2-grid.jpg" />
// Input: matrix = [[-19,57],[-40,-5]]
// Output: -59
// Explanation: The falling path with a minimum sum is shown.
 
// Constraints:
//     n == matrix.length == matrix[i].length
//     1 <= n <= 100
//     -100 <= matrix[i][j] <= 100

import "fmt"
import "math"

// dp
func minFallingPathSum(matrix [][]int) int {
    n, flag := len(matrix), len(matrix) > 1
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := n - 2; i >= 0; i-- {
        if flag {
            // 在下一行选择的元素和当前行所选元素最多相隔一列（即位于正下方或者沿对角线向左或者向右的第一个元素）。
            // 具体来说，位置 (row, col) 的下一个元素应当是 (row + 1, col - 1)、(row + 1, col) 或者 (row + 1, col + 1)
            matrix[i][0] += min(matrix[i + 1][0], matrix[i + 1][1])
            matrix[i][n - 1] += min(matrix[i + 1][n - 1], matrix[i + 1][n - 2])
        }
        for k := 1; k < n - 1; k++ {
            matrix[i][k] += min( min( matrix[i + 1][k],  matrix[i + 1][k - 1] ), matrix[i + 1][k + 1])
        }
    }
    res := matrix[0][0]
    for k := 1; k < n; k++ {
        res = min(res, matrix[0][k])
    }
    return res
}

func minFallingPathSum1(matrix [][]int) int {
    n := len(matrix)
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n + 2)
        dp[i][0], dp[i][n + 1] = math.MaxInt, math.MaxInt
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := n - 1; i >= 0; i-- {
        for j := 1; j <= n; j++ {
            if i == n - 1 {
                dp[i][j] = matrix[i][j - 1]
            } else {
                dp[i][j] = min(dp[i + 1][j - 1], min(dp[i + 1][j], dp[i + 1][j + 1])) + matrix[i][j - 1]
            }
        }
    }
    res := math.MaxInt
    for _, num := range dp[0] {
        res = min(res, num)
    }
    return res
}

func main() {
    // 1 + 4 + 8  or  1 + 5 + 7
    fmt.Println(minFallingPathSum([][]int{{2,1,3},{6,5,4},{7,8,9}})) // 13
    //  -19 + -40
    fmt.Println(minFallingPathSum([][]int{{-19,57},{-40,-5}})) // -59

    fmt.Println(minFallingPathSum1([][]int{{2,1,3},{6,5,4},{7,8,9}})) // 13
    fmt.Println(minFallingPathSum1([][]int{{-19,57},{-40,-5}})) // -59
}

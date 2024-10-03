package main

// 2428. Maximum Sum of an Hourglass
// You are given an m x n integer matrix grid.

// We define an hourglass as a part of the matrix with the following form:
// <img src="https://assets.leetcode.com/uploads/2022/08/21/img.jpg" />

// Return the maximum sum of the elements of an hourglass.

// Note that an hourglass cannot be rotated and must be entirely contained within the matrix.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/08/21/1.jpg" />
// Input: grid = [[6,2,1,3],[4,2,1,5],[9,2,8,7],[4,1,2,9]]
// Output: 30
// Explanation: The cells shown above represent the hourglass with the maximum sum: 6 + 2 + 1 + 2 + 9 + 2 + 8 = 30.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/08/21/2.jpg" />
// Input: grid = [[1,2,3],[4,5,6],[7,8,9]]
// Output: 35
// Explanation: There is only one hourglass in the matrix, with the sum: 1 + 2 + 3 + 5 + 7 + 8 + 9 = 35.

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     3 <= m, n <= 150
//     0 <= grid[i][j] <= 10^6

import "fmt"

func maxSum(grid [][]int) int {
    res, m, n := 0, len(grid), len(grid[0])
    sum := func(i, j int) int {
        res := grid[i][j]
        for k := j - 1; k <= j + 1; k++ {
            res += (grid[i - 1][k] + grid[i + 1][k])
        }
        return res
    }
    for i := 1; i < m - 1; i++ {
        for j := 1; j < n - 1; j++ {
            s := sum(i,j)
            if s > res { 
                res = s 
            }
        }
    }
    return res
}

func maxSum1(grid [][]int) int {
    res, m, n := 0, len(grid), len(grid[0])
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := m - 3; i >= 0; i-- {
        sum := grid[i+1][n-2]
        for j := 1; j <= 3; j++ {
            sum += grid[i][n-j] + grid[i+2][n-j]
        }
        res = max(res, sum)
        for j := n - 4; j >= 0; j-- {
            sum += grid[i][j] + grid[i+1][j+1] + grid[i+2][j] - grid[i][j+3] - grid[i+1][j+2] - grid[i+2][j+3]
            res = max(res, sum)
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/08/21/1.jpg" />
    // Input: grid = [[6,2,1,3],[4,2,1,5],[9,2,8,7],[4,1,2,9]]
    // Output: 30
    // Explanation: The cells shown above represent the hourglass with the maximum sum: 6 + 2 + 1 + 2 + 9 + 2 + 8 = 30.
    fmt.Println(maxSum([][]int{{6,2,1,3},{4,2,1,5},{9,2,8,7},{4,1,2,9}})) // 30
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/08/21/2.jpg" />
    // Input: grid = [[1,2,3],[4,5,6],[7,8,9]]
    // Output: 35
    // Explanation: There is only one hourglass in the matrix, with the sum: 1 + 2 + 3 + 5 + 7 + 8 + 9 = 35.
    fmt.Println(maxSum([][]int{{1,2,3},{4,5,6},{7,8,9}})) // 35

    fmt.Println(maxSum1([][]int{{6,2,1,3},{4,2,1,5},{9,2,8,7},{4,1,2,9}})) // 30
    fmt.Println(maxSum1([][]int{{1,2,3},{4,5,6},{7,8,9}})) // 35
}
package main

// 2245. Maximum Trailing Zeros in a Cornered Path
// You are given a 2D integer array grid of size m x n, where each cell contains a positive integer.

// A cornered path is defined as a set of adjacent cells with at most one turn. 
// More specifically, the path should exclusively move either horizontally or vertically up to the turn (if there is one), without returning to a previously visited cell. 
// After the turn, the path will then move exclusively in the alternate direction: 
//     move vertically if it moved horizontally, and vice versa, also without returning to a previously visited cell.

// The product of a path is defined as the product of all the values in the path.

// Return the maximum number of trailing zeros in the product of a cornered path found in grid.

// Note:
//     Horizontal movement means moving in either the left or right direction.
//     Vertical movement means moving in either the up or down direction.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/03/23/ex1new2.jpg" />
// Input: grid = [[23,17,15,3,20],[8,1,20,27,11],[9,4,6,2,21],[40,9,1,10,6],[22,7,4,5,3]]
// Output: 3
// Explanation: The grid on the left shows a valid cornered path.
// It has a product of 15 * 20 * 6 * 1 * 10 = 18000 which has 3 trailing zeros.
// It can be shown that this is the maximum trailing zeros in the product of a cornered path.
// The grid in the middle is not a cornered path as it has more than one turn.
// The grid on the right is not a cornered path as it requires a return to a previously visited cell.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/03/25/ex2.jpg" />
// Input: grid = [[4,3,2],[7,6,1],[8,8,8]]
// Output: 0
// Explanation: The grid is shown in the figure above.
// There are no cornered paths in the grid that result in a product with a trailing zero.

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 10^5
//     1 <= m * n <= 10^5
//     1 <= grid[i][j] <= 1000

import "fmt"

func maxTrailingZeros(grid [][]int) int {
    res := 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    count := func(n, w int) int {
        res := 0
        for n % w == 0 {
            res++
            n /= w
        }
        return res
    }
    rotate := func(grid [][]int) [][]int {
        m, n := len(grid), len(grid[0])
        res := make([][]int, n)
        for i := range res {
            res[i] = make([]int, m)
        }
        for i := 0; i < m; i++ {
            for j := 0; j < n; j++ {
                res[j][m - 1 - i] = grid[i][j]
            }
        }
        return res
    }
    solve := func(grid [][]int) int {
        res, m, n := 0, len(grid), len(grid[0])
        fcol, tcol := make([][]int, m + 1), make([][]int, m + 1)
        frow, trow := make([][]int, n + 1), make([][]int, n + 1)
        for i := range fcol {
            fcol[i], tcol[i] = make([]int, n), make([]int, n)
        }
        for i := range frow {
            frow[i], trow[i] = make([]int, m), make([]int, m)
        }
        for i := 0; i < m; i++ {
            for j := 0; j < n; j++ {
                t, f := count(grid[i][j], 2), count(grid[i][j], 5)
                fcol[i + 1][j] += fcol[i][j] + f
                tcol[i + 1][j] += tcol[i][j] + t
                frow[j + 1][i] += frow[j][i] + f
                trow[j + 1][i] += trow[j][i] + t
            }
        }
        for i := 0; i < m; i++ {
            for j := 0; j < n; j++ {
                rowTwo, rowFive := trow[j + 1][i], frow[j + 1][i]
                colTwo, colFive := tcol[m][j] - tcol[i + 1][j], fcol[m][j] - fcol[i + 1][j];
                res = max(res, min(rowTwo + colTwo, rowFive + colFive))
            }
        }
        return res
    }
    for i := 0; i < 4; i++ {
        res = max(res, solve(grid))
        grid = rotate(grid)
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/03/23/ex1new2.jpg" />
    // Input: grid = [[23,17,15,3,20],[8,1,20,27,11],[9,4,6,2,21],[40,9,1,10,6],[22,7,4,5,3]]
    // Output: 3
    // Explanation: The grid on the left shows a valid cornered path.
    // It has a product of 15 * 20 * 6 * 1 * 10 = 18000 which has 3 trailing zeros.
    // It can be shown that this is the maximum trailing zeros in the product of a cornered path.
    // The grid in the middle is not a cornered path as it has more than one turn.
    // The grid on the right is not a cornered path as it requires a return to a previously visited cell.
    grid1 := [][]int{
        {23,17,15,3,20},
        {8,1,20,27,11},
        {9,4,6,2,21},
        {40,9,1,10,6},
        {22,7,4,5,3},
    }
    fmt.Println(maxTrailingZeros(grid1)) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/03/25/ex2.jpg" />
    // Input: grid = [[4,3,2],[7,6,1],[8,8,8]]
    // Output: 0
    // Explanation: The grid is shown in the figure above.
    // There are no cornered paths in the grid that result in a product with a trailing zero.
    grid2 := [][]int{
        {4,3,2},
        {7,6,1},
        {8,8,8},
    }
    fmt.Println(maxTrailingZeros(grid2)) // 0
}
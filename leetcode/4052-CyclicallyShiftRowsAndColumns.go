package main

// 4052. Cyclically Shift Rows and Columns
// You are given an integer n, a 2D integer array grid of size n x n, and two integer arrays rowShift and colShift, each of length n, where:
//     1. rowShift[i] represents the number of positions to cyclically shift the ith row of grid to the left.
//     2. colShift[j] represents the number of positions to cyclically shift the jth column of grid upward.

// First, cyclically shift each row according to rowShift, 
// then cyclically shift each column of the resulting grid according to colShift.

// Return the resulting grid after performing all the shifts.

// A cyclic left shift of a row by k positions moves the element at column j to column (j - k + n) % n. 
// All other rows remain unchanged.

// A cyclic upward shift of a column by k positions moves the element at row i to row (i - k + n) % n. 
// All other columns remain unchanged.

// Example 1:
// Input: n = 2, grid = [[1,2],[3,4]], rowShift = [1,0], colShift = [0,1]
// Output: [[2,4],[3,1]]
// Explanation:
// The grid changes as follows:
// <img src="https://assets.leetcode.com/uploads/2026/08/18/4743-1.png" />

// Example 2:
// Input: n = 3, grid = [[1,2,3],[4,5,6],[7,8,9]], rowShift = [1,2,0], colShift = [2,2,1]
// Output: [[7,8,5],[2,3,9],[6,4,1]]
// Explanation:
// The grid changes as follows:
// <img src="https://assets.leetcode.com/uploads/2026/08/18/4743-2.png" />

// Constraints:
//     1 <= n == grid.length == grid[i].length <= 10
//     1 <= grid[i][j] <= 100
//     rowShift.length == colShift.length == n
//     0 <= rowShift[i], colShift[i] < n

import "fmt"

func cyclicShift(n int, grid [][]int, rowShift, colShift []int) [][]int {
    for i, row := range grid {
        shift := rowShift[i]
        grid[i] = append(row[shift:], row[:shift]...)
    }
    column := make([]int, n)
    for j, shift := range colShift {
        // 收集列元素
        col := column[:0]
        for _, row := range grid[shift:] {
            col = append(col, row[j])
        }
        for _, row := range grid[:shift] {
            col = append(col, row[j])
        }
        // 填入列
        for i, row := range grid {
            row[j] = col[i]
        }
    }
    return grid
}

func cyclicShift1(n int, grid [][]int, rowShift []int, colShift []int) [][]int {
    it := make([][]int, n)
    for i := 0; i < n; i++ {
        it[i] = make([]int, n)
        for j := 0; j < n; j++ {
            it[i][j] = grid[i][(j + rowShift[i]) % n]
        }
    }
    res := make([][]int, n)
    for i := 0; i < n; i++ {
        res[i] = make([]int, n)
    }
    for j := 0; j < n; j++ {
        for i := 0; i < n; i++ {
            res[i][j] = it[(i + colShift[j]) % n][j]
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 2, grid = [[1,2],[3,4]], rowShift = [1,0], colShift = [0,1]
    // Output: [[2,4],[3,1]]
    // Explanation:
    // The grid changes as follows:
    // <img src="https://assets.leetcode.com/uploads/2026/08/18/4743-1.png" />
    fmt.Println(cyclicShift(2, [][]int{{1,2},{3,4}}, []int{1,0}, []int{0,1})) // [[2,4],[3,1]]
    // Example 2:
    // Input: n = 3, grid = [[1,2,3],[4,5,6],[7,8,9]], rowShift = [1,2,0], colShift = [2,2,1]
    // Output: [[7,8,5],[2,3,9],[6,4,1]]
    // Explanation:
    // The grid changes as follows:
    // <img src="https://assets.leetcode.com/uploads/2026/08/18/4743-2.png" />
    fmt.Println(cyclicShift(3, [][]int{{1,2,3},{4,5,6}, {7,8,9}}, []int{1,2, 0}, []int{2,2, 1})) // [[7,8,5],[2,3,9],[6,4,1]]

    fmt.Println(cyclicShift1(2, [][]int{{1,2},{3,4}}, []int{1,0}, []int{0,1})) // [[2,4],[3,1]]
    fmt.Println(cyclicShift1(3, [][]int{{1,2,3},{4,5,6}, {7,8,9}}, []int{1,2, 0}, []int{2,2, 1})) // [[7,8,5],[2,3,9],[6,4,1]]
}
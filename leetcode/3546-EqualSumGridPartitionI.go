package main

// 3546. Equal Sum Grid Partition I
// You are given an m x n matrix grid of positive integers. 
// Your task is to determine if it is possible to make either one horizontal or one vertical cut on the grid such that:
//     1. Each of the two resulting sections formed by the cut is non-empty.
//     2. The sum of the elements in both sections is equal.

// Return true if such a partition exists; otherwise return false.

// Example 1:
// Input: grid = [[1,4],[2,3]]
// Output: true
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/03/30/lc.jpeg" />
// A horizontal cut between row 0 and row 1 results in two non-empty sections, each with a sum of 5. Thus, the answer is true.

// Example 2:
// Input: grid = [[1,3],[2,4]]
// Output: false
// Explanation:
// No horizontal or vertical cut results in two non-empty sections with equal sums. Thus, the answer is false.

// Constraints:
//     1 <= m == grid.length <= 10^5
//     1 <= n == grid[i].length <= 10^5
//     2 <= m * n <= 10^5
//     1 <= grid[i][j] <= 10^5

import "fmt"

func canPartitionGrid(grid [][]int) bool {
    n, m := len(grid), len(grid[0])
    currRow, currCol, rows, cols := int64(0), int64(0), int64(0), int64(0)
    prefixRow, prefixCol := make([]int64, n), make([]int64, m)
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            prefixRow[i] += int64(grid[i][j])
            prefixCol[j] += int64(grid[i][j])
        }
    }
    for _, v := range prefixRow {
        rows += v
    }
    cols = rows
    for i := 0; i < n - 1; i++ {
        currRow += prefixRow[i]
        if currRow == (rows - currRow) { return true }
    }
    for i := 0; i < m - 1; i++ {
        currCol += prefixCol[i]
        if currCol == (cols - currCol) { return true }
    }
    return false
}

func canPartitionGrid1(grid [][]int) bool {
    n, m := len(grid), len(grid[0])
    rows, cols := make([]int, n), make([]int, m)
    for i, nums := range grid {
        for j, v := range nums {
            rows[i] += v
            cols[j] += v
        }
    }
    for i := 1; i < n; i++ {
        rows[i] += rows[i - 1]
    }
    for i := 1; i < m; i++ {
        cols[i] += cols[i - 1]
    }
    for i := 0; i < n - 1; i++ {
        if (rows[i] << 1) == rows[n - 1] { 
            return true
        }
    }
    for i := 0; i < m - 1; i++ {
        if (cols[i] << 1) == cols[m - 1] {
            return true
        }
    }
    return false
}

func canPartitionGrid2(grid [][]int) bool {
    sum, n, m := int64(0), len(grid), len(grid[0])
    rows, cols := make([]int64, n), make([]int64, m)
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            v := int64(grid[i][j])
            sum += v
            rows[i] += v
            cols[j] += v
        }
    }
    check := func(arr []int64) bool {
        for i := 0; i < len(arr)-1; i++ {
            if i > 0 { arr[i] += arr[i-1] }
            if arr[i] * 2 == sum { return true }
            if arr[i] * 2 > sum  { break }
        }
        return false
    }
    if sum % 2 != 0 { return false }
    return check(rows) || check(cols)
}

func main() {
    // Example 1:
    // Input: grid = [[1,4],[2,3]]
    // Output: true
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/03/30/lc.jpeg" />
    // A horizontal cut between row 0 and row 1 results in two non-empty sections, each with a sum of 5. Thus, the answer is true.
    fmt.Println(canPartitionGrid([][]int{{1,4},{2,3}})) // true
    // Example 2:
    // Input: grid = [[1,3],[2,4]]
    // Output: false
    // Explanation:
    // No horizontal or vertical cut results in two non-empty sections with equal sums. Thus, the answer is false.
    fmt.Println(canPartitionGrid([][]int{{1,3},{2,4}})) // false

    fmt.Println(canPartitionGrid1([][]int{{1,4},{2,3}})) // true
    fmt.Println(canPartitionGrid1([][]int{{1,3},{2,4}})) // false

    fmt.Println(canPartitionGrid2([][]int{{1,4},{2,3}})) // true
    fmt.Println(canPartitionGrid2([][]int{{1,3},{2,4}})) // false
}
package main

// 2397. Maximum Rows Covered by Columns
// You are given an m x n binary matrix matrix and an integer numSelect.

// Your goal is to select exactly numSelect distinct columns from matrix such that you cover as many rows as possible.

// A row is considered covered if all the 1's in that row are also part of a column that you have selected. 
// If a row does not have any 1s, it is also considered covered.

// More formally, let us consider selected = {c1, c2, ...., cnumSelect} as the set of columns selected by you. 
// A row i is covered by selected if:
//     For each cell where matrix[i][j] == 1, the column j is in selected.
//     Or, no cell in row i has a value of 1.

// Return the maximum number of rows that can be covered by a set of numSelect columns.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/07/14/rowscovered.png" />
// Input: matrix = [[0,0,0],[1,0,1],[0,1,1],[0,0,1]], numSelect = 2
// Output: 3
// Explanation:
// One possible way to cover 3 rows is shown in the diagram above.
// We choose s = {0, 2}.
// - Row 0 is covered because it has no occurrences of 1.
// - Row 1 is covered because the columns with value 1, i.e. 0 and 2 are present in s.
// - Row 2 is not covered because matrix[2][1] == 1 but 1 is not present in s.
// - Row 3 is covered because matrix[2][2] == 1 and 2 is present in s.
// Thus, we can cover three rows.
// Note that s = {1, 2} will also cover 3 rows, but it can be shown that no more than three rows can be covered.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/07/14/rowscovered2.png" />
// Input: matrix = [[1],[0]], numSelect = 1
// Output: 2
// Explanation:
// Selecting the only column will result in both rows being covered since the entire matrix is selected.

// Constraints:
//     m == matrix.length
//     n == matrix[i].length
//     1 <= m, n <= 12
//     matrix[i][j] is either 0 or 1.
//     1 <= numSelect <= n

import "fmt"

func maximumRows(mat [][]int, cols int) int {
    res, m, n := 0, len(mat), len(mat[0])
    rows := []int{}
    for i := 0; i < m; i++ {
        v := mat[i][0]
        for j := 1; j < n; j++ {
            v = (v << 1) + mat[i][j]
        }
        rows = append(rows, v)
    }
    countBits := func(n int) int {
        count := 0
        for n > 0 {
            count++
            n = n & (n - 1)
        }
        return count
    }
    for cm := 0; cm < 1 << n; cm++ {
        if countBits(cm) == cols {
            count := 0
            for _, row := range rows {
                c := cm
                y := true
                for i := 0; i < n; i++ {
                    if row & 1 == 1 && c & 1 != 1 {
                        y = false
                        break
                    }
                    row >>= 1
                    c >>= 1
                }
                if y {
                    count++
                }
            }
            if count > res {
                res = count
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/07/14/rowscovered.png" />
    // Input: matrix = [[0,0,0],[1,0,1],[0,1,1],[0,0,1]], numSelect = 2
    // Output: 3
    // Explanation:
    // One possible way to cover 3 rows is shown in the diagram above.
    // We choose s = {0, 2}.
    // - Row 0 is covered because it has no occurrences of 1.
    // - Row 1 is covered because the columns with value 1, i.e. 0 and 2 are present in s.
    // - Row 2 is not covered because matrix[2][1] == 1 but 1 is not present in s.
    // - Row 3 is covered because matrix[2][2] == 1 and 2 is present in s.
    // Thus, we can cover three rows.
    // Note that s = {1, 2} will also cover 3 rows, but it can be shown that no more than three rows can be covered.
    fmt.Println(maximumRows([][]int{{0,0,0},{1,0,1},{0,1,1},{0,0,1}}, 2)) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/07/14/rowscovered2.png" />
    // Input: matrix = [[1],[0]], numSelect = 1
    // Output: 2
    // Explanation:
    // Selecting the only column will result in both rows being covered since the entire matrix is selected.
    fmt.Println(maximumRows([][]int{{1},{0}}, 1)) // 2
}
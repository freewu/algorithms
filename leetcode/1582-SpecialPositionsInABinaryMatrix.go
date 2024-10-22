package main

// 1582. Special Positions in a Binary Matrix
// Given an m x n binary matrix mat, return the number of special positions in mat.

// A position (i, j) is called special if mat[i][j] == 1 
// and all other elements in row i and column j are 0 (rows and columns are 0-indexed).

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/12/23/special1.jpg" />
// Input: mat = [[1,0,0],[0,0,1],[1,0,0]]
// Output: 1
// Explanation: (1, 2) is a special position because mat[1][2] == 1 and all other elements in row 1 and column 2 are 0.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/12/24/special-grid.jpg" />
// Input: mat = [[1,0,0],[0,1,0],[0,0,1]]
// Output: 3
// Explanation: (0, 0), (1, 1) and (2, 2) are special positions.

// Constraints:
//     m == mat.length
//     n == mat[i].length
//     1 <= m, n <= 100
//     mat[i][j] is either 0 or 1.

import "fmt"

func numSpecial(mat [][]int) int {
    res, n, m := 0, len(mat), len(mat[0])
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if mat[i][j] != 1 { continue }
            isSpecial := true
            for k := 0; k < m; k++ {
                if k == j { continue }
                if mat[i][k] == 1 {
                    isSpecial = false // 有列又出现了 1 直接跳出
                    break
                }
            }
            for k := 0; k < n; k++ {
                if k == i { continue }
                if mat[k][j] == 1 { // 有行又出现了 1 直接跳出
                    isSpecial = false
                    break
                }
            }
            if isSpecial {
                res++
            }
        }
    }
    return res
}

func numSpecial1(mat [][]int) int {
    res, m, n := 0, len(mat), len(mat[0])
    row, col := make([]int, m), make([]int, n)
    for i := range mat {
        for j, v := range mat[i] {
            row[i] += v
            col[j] += v
        }
    }
    for i := range mat {
        for j, v := range mat[i] {
            if v == 1 && row[i] == 1 && col[j] == 1 {
                res++
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/12/23/special1.jpg" />
    // Input: mat = [[1,0,0],[0,0,1],[1,0,0]]
    // Output: 1
    // Explanation: (1, 2) is a special position because mat[1][2] == 1 and all other elements in row 1 and column 2 are 0.
    fmt.Println(numSpecial([][]int{{1,0,0},{0,0,1},{1,0,0}})) // 1
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/12/24/special-grid.jpg" />
    // Input: mat = [[1,0,0],[0,1,0],[0,0,1]]
    // Output: 3
    // Explanation: (0, 0), (1, 1) and (2, 2) are special positions.
    fmt.Println(numSpecial([][]int{{1,0,0},{0,1,0},{0,0,1}})) // 3

    fmt.Println(numSpecial1([][]int{{1,0,0},{0,0,1},{1,0,0}})) // 1
    fmt.Println(numSpecial1([][]int{{1,0,0},{0,1,0},{0,0,1}})) // 3
}
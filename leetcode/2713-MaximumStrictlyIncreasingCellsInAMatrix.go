package main

// 2713. Maximum Strictly Increasing Cells in a Matrix
// Given a 1-indexed m x n integer matrix mat, you can select any cell in the matrix as your starting cell.

// From the starting cell, you can move to any other cell in the same row or column, 
// but only if the value of the destination cell is strictly greater than the value of the current cell. 
// You can repeat this process as many times as possible, moving from cell to cell until you can no longer make any moves.

// Your task is to find the maximum number of cells that you can visit in the matrix by starting from some cell.
// Return an integer denoting the maximum number of cells that can be visited.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2023/04/23/diag1drawio.png" />
// Input: mat = [[3,1],[3,4]]
// Output: 2
// Explanation: The image shows how we can visit 2 cells starting from row 1, column 2. It can be shown that we cannot visit more than 2 cells no matter where we start from, so the answer is 2. 

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2023/04/23/diag3drawio.png" />
// Input: mat = [[1,1],[1,1]]
// Output: 1
// Explanation: Since the cells must be strictly increasing, we can only visit one cell in this example. 

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2023/04/23/diag4drawio.png" />
// Input: mat = [[3,1,6],[-9,5,7]]
// Output: 4
// Explanation: The image above shows how we can visit 4 cells starting from row 2, column 1. It can be shown that we cannot visit more than 4 cells no matter where we start from, so the answer is 4. 

// Constraints:
//     m == mat.length 
//     n == mat[i].length 
//     1 <= m, n <= 10^5
//     1 <= m * n <= 10^5
//     -10^5 <= mat[i][j] <= 10^5

import "fmt"
import "sort"

func maxIncreasingCells(mat [][]int) int {
    type X struct {
        r, c, v int
    }
    type Y struct {
        max, second, v int
    }
    res, m, n := 0, len(mat), len(mat[0])
    sorted, rows, columns := make([]*X, 0, m * n), make([]Y, m), make([]Y, n)
    for r, row := range mat {
        for c, v := range row {
            sorted = append(sorted, &X{r, c, v})
        }
    }
    sort.Slice(sorted, func(i, j int) bool {
        return sorted[i].v < sorted[j].v
    })
    for _, x := range sorted {
        nd := 0
        if rows[x.r].max == 0 || rows[x.r].v < x.v{
            nd = rows[x.r].max + 1
        } else {
            nd = rows[x.r].second + 1
        }
        if columns[x.c].max == 0 || columns[x.c].v < x.v {
            if columns[x.c].max + 1 > nd {
                nd = columns[x.c].max + 1
            }
        } else if columns[x.c].second + 1 > nd {
            nd = columns[x.c].second + 1
        }
        if rows[x.r].v == x.v && rows[x.r].max < nd {
            rows[x.r].max = nd
        } else if rows[x.r].max == 0 || rows[x.r].v < x.v {
            rows[x.r].second = rows[x.r].max 
            rows[x.r].max = nd
            rows[x.r].v = x.v
        }
        if columns[x.c].v == x.v && columns[x.c].max < nd {
            columns[x.c].max = nd
        } else if columns[x.c].max == 0 || columns[x.c].v < x.v {
            columns[x.c].second = columns[x.c].max
            columns[x.c].max = nd
            columns[x.c].v = x.v
        }
        if nd > res {
            res = nd
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2023/04/23/diag1drawio.png" />
    // Input: mat = [[3,1],[3,4]]
    // Output: 2
    // Explanation: The image shows how we can visit 2 cells starting from row 1, column 2. It can be shown that we cannot visit more than 2 cells no matter where we start from, so the answer is 2. 
    fmt.Println(maxIncreasingCells([][]int{{3,1},{3,4}})) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2023/04/23/diag3drawio.png" />
    // Input: mat = [[1,1],[1,1]]
    // Output: 1
    // Explanation: Since the cells must be strictly increasing, we can only visit one cell in this example. 
    fmt.Println(maxIncreasingCells([][]int{{1,1},{1,1}})) // 1
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2023/04/23/diag4drawio.png" />
    // Input: mat = [[3,1,6],[-9,5,7]]
    // Output: 4
    // Explanation: The image above shows how we can visit 4 cells starting from row 2, column 1. It can be shown that we cannot visit more than 4 cells no matter where we start from, so the answer is 4. 
    fmt.Println(maxIncreasingCells([][]int{{3,1,6},{-9,5,7}})) // 4
}
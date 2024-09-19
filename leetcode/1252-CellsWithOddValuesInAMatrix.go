package main

// 1252. Cells with Odd Values in a Matrix
// There is an m x n matrix that is initialized to all 0's. 
// There is also a 2D array indices where each indices[i] = [ri, ci] represents a 0-indexed location to perform some increment operations on the matrix.

// For each location indices[i], do both of the following:
//     Increment all the cells on row ri.
//     Increment all the cells on column ci.

// Given m, n, and indices, 
// return the number of odd-valued cells in the matrix after applying the increment to all locations in indices.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/10/30/e1.png" />
// Input: m = 2, n = 3, indices = [[0,1],[1,1]]
// Output: 6
// Explanation: Initial matrix = [[0,0,0],[0,0,0]].
// After applying first increment it becomes [[1,2,1],[0,1,0]].
// The final matrix is [[1,3,1],[1,3,1]], which contains 6 odd numbers.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2019/10/30/e2.png" />
// Input: m = 2, n = 2, indices = [[1,1],[0,0]]
// Output: 0
// Explanation: Final matrix = [[2,2],[2,2]]. There are no odd numbers in the final matrix.

// Constraints:
//     1 <= m, n <= 50
//     1 <= indices.length <= 100
//     0 <= ri < m
//     0 <= ci < n

// Follow up: Could you solve this in O(n + m + indices.length) time with only O(n + m) extra space?

import "fmt"

func oddCells(m int, n int, indices [][]int) int {
    res, matrix := 0, make([][]int, m)
    for i := range matrix { // make the matrix
        matrix[i] = make([]int,n)
    }
    for i := 0; i < len(indices); i++ { // action
        r, c := indices[i][0], indices[i][1]
        for j := 0; j < n; j++ {
            matrix[r][j]++
        }
        for j := 0; j < m; j++ {
            matrix[j][c]++
        }
    }
    for i := range matrix {
        for j := range matrix[0] {
            if matrix[i][j] % 2 != 0 { // stat the odd cell
                res++
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/10/30/e1.png" />
    // Input: m = 2, n = 3, indices = [[0,1],[1,1]]
    // Output: 6
    // Explanation: Initial matrix = [[0,0,0],[0,0,0]].
    // After applying first increment it becomes [[1,2,1],[0,1,0]].
    // The final matrix is [[1,3,1],[1,3,1]], which contains 6 odd numbers.
    fmt.Println(oddCells(2, 3, [][]int{{0,1},{1,1}})) // 6
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2019/10/30/e2.png" />
    // Input: m = 2, n = 2, indices = [[1,1],[0,0]]
    // Output: 0
    // Explanation: Final matrix = [[2,2],[2,2]]. There are no odd numbers in the final matrix.
    fmt.Println(oddCells(2, 2, [][]int{{1,1},{0,0}})) // 0
}
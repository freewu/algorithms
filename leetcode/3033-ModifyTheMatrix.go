package main

// 3033. Modify the Matrix
// Given a 0-indexed m x n integer matrix matrix, create a new 0-indexed matrix called answer. 
// Make answer equal to matrix, then replace each element with the value -1 with the maximum element in its respective column.
// Return the matrix answer.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2023/12/24/matrix1.png" />
// Input: matrix = [[1,2,-1],[4,-1,6],[7,8,9]]
// Output: [[1,2,9],[4,8,6],[7,8,9]]
// Explanation: The diagram above shows the elements that are changed (in blue).
// - We replace the value in the cell [1][1] with the maximum value in the column 1, that is 8.
// - We replace the value in the cell [0][2] with the maximum value in the column 2, that is 9.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2023/12/24/matrix2.png" />
// Input: matrix = [[3,-1],[5,2]]
// Output: [[3,2],[5,2]]
// Explanation: The diagram above shows the elements that are changed (in blue).

// Constraints:
//     m == matrix.length
//     n == matrix[i].length
//     2 <= m, n <= 50
//     -1 <= matrix[i][j] <= 100
//     The input is generated such that each column contains at least one non-negative integer.

import "fmt"

func modifiedMatrix(matrix [][]int) [][]int {
    n, m := len(matrix), len(matrix[0])
    for i := 0; i < m; i++ {
        mx, flag := 0, false // 列中最大值 & 是否有 -1
        for j := 0; j < n; j++ {
            if matrix[j][i] > mx {
                mx = matrix[j][i]
            }
            if matrix[j][i] == -1 {
                flag = true
            }
        }
        if flag {
            for j := 0; j < n; j++ {
                if matrix[j][i] == -1 { // 将 -1 的位置填充为本列最大的值
                    matrix[j][i] = mx
                }
            }
        }
    }
    return matrix
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2023/12/24/matrix1.png" />
    // Input: matrix = [[1,2,-1],[4,-1,6],[7,8,9]]
    // Output: [[1,2,9],[4,8,6],[7,8,9]]
    // Explanation: The diagram above shows the elements that are changed (in blue).
    // - We replace the value in the cell [1][1] with the maximum value in the column 1, that is 8.
    // - We replace the value in the cell [0][2] with the maximum value in the column 2, that is 9.
    matrix1 := [][]int{
        {1,2,-1},
        {4,-1,6},
        {7,8,9},
    }
    fmt.Println(modifiedMatrix(matrix1)) // [[1 2 9] [4 8 6] [7 8 9]]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2023/12/24/matrix2.png" />
    // Input: matrix = [[3,-1],[5,2]]
    // Output: [[3,2],[5,2]]
    // Explanation: The diagram above shows the elements that are changed (in blue).
    matrix2 := [][]int{
        {3,-1},
        {5,2},
    }
    fmt.Println(modifiedMatrix(matrix2)) // [[3 2] [5 2]]
}
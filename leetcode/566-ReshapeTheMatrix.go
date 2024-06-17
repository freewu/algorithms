package main

// 566. Reshape the Matrix
// In MATLAB, there is a handy function called reshape which can reshape an m x n matrix into a new one with a different size r x c keeping its original data.
// You are given an m x n matrix mat and two integers r and c representing the number of rows and the number of columns of the wanted reshaped matrix.
// The reshaped matrix should be filled with all the elements of the original matrix in the same row-traversing order as they were.
// If the reshape operation with given parameters is possible and legal, output the new reshaped matrix; Otherwise, output the original matrix.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/04/24/reshape1-grid.jpg" />
// Input: mat = [[1,2],[3,4]], r = 1, c = 4
// Output: [[1,2,3,4]]

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/04/24/reshape2-grid.jpg" />
// Input: mat = [[1,2],[3,4]], r = 2, c = 4
// Output: [[1,2],[3,4]]

// Constraints:
//     m == mat.length
//     n == mat[i].length
//     1 <= m, n <= 100
//     -1000 <= mat[i][j] <= 1000
//     1 <= r, c <= 300

import "fmt"

func matrixReshape(nums [][]int, r int, c int) [][]int {
    canReshape := func(nums [][]int, r, c int) bool {
        row, colume := len(nums), len(nums[0])
        if row * colume == r * c {
            return true
        }
        return false
    }
    reshape := func(nums [][]int, r, c int) [][]int {
        newShape := make([][]int, r)
        for index := range newShape {
            newShape[index] = make([]int, c)
        }
        rowIndex, colIndex := 0, 0
        for _, row := range nums {
            for _, col := range row {
                if colIndex == c {
                    colIndex = 0
                    rowIndex++
                }
                newShape[rowIndex][colIndex] = col
                colIndex++
            }
        }
        return newShape
    }
    if canReshape(nums, r, c) {
        return reshape(nums, r, c)
    }
    return nums
}

func matrixReshape1(mat [][]int, r int, c int) [][]int {
    m, n := len(mat), len(mat[0]) 
    if m * n != r * c {
        return mat
    }
    newMat := make([][]int,r)
    for i := range newMat {
        newMat[i] = make([]int,c)
    }
    for i := 0; i < r * c; i++ {
        newMat[i / c][i % c] = mat[i / n][i % n ]
    }
    return newMat
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/04/24/reshape1-grid.jpg" />
    // Input: mat = [[1,2],[3,4]], r = 1, c = 4
    // Output: [[1,2,3,4]]
    fmt.Println(matrixReshape([][]int{{1,2},{3,4}},1,4)) // [[1,2,3,4]]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/04/24/reshape2-grid.jpg" />
    // Input: mat = [[1,2],[3,4]], r = 2, c = 4
    // Output: [[1,2],[3,4]]
    fmt.Println(matrixReshape([][]int{{1,2},{3,4}},2,4)) // [[1,2],[3,4]]

    fmt.Println(matrixReshape1([][]int{{1,2},{3,4}},1,4)) // [[1,2,3,4]]
    fmt.Println(matrixReshape1([][]int{{1,2},{3,4}},2,4)) // [[1,2],[3,4]]
}
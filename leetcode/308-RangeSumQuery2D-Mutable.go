package main 

// 308. Range Sum Query 2D - Mutable
// Given a 2D matrix matrix, handle multiple queries of the following types:
//     Update the value of a cell in matrix.
//     Calculate the sum of the elements of matrix inside the rectangle defined by its upper left corner (row1, col1) and lower right corner (row2, col2).

// Implement the NumMatrix class:
//     NumMatrix(int[][] matrix) 
//         Initializes the object with the integer matrix matrix.
//     void update(int row, int col, int val) 
//         Updates the value of matrix[row][col] to be val.
//     int sumRegion(int row1, int col1, int row2, int col2) 
//         Returns the sum of the elements of matrix inside the rectangle defined by its upper left corner (row1, col1) 
//         and lower right corner (row2, col2).

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/03/14/summut-grid.jpg" />
// Input
// ["NumMatrix", "sumRegion", "update", "sumRegion"]
// [[[[3, 0, 1, 4, 2], [5, 6, 3, 2, 1], [1, 2, 0, 1, 5], [4, 1, 0, 1, 7], [1, 0, 3, 0, 5]]], [2, 1, 4, 3], [3, 2, 2], [2, 1, 4, 3]]
// Output
// [null, 8, null, 10]
// Explanation
// NumMatrix numMatrix = new NumMatrix([[3, 0, 1, 4, 2], [5, 6, 3, 2, 1], [1, 2, 0, 1, 5], [4, 1, 0, 1, 7], [1, 0, 3, 0, 5]]);
// numMatrix.sumRegion(2, 1, 4, 3); // return 8 (i.e. sum of the left red rectangle)
// numMatrix.update(3, 2, 2);       // matrix changes from left image to right image
// numMatrix.sumRegion(2, 1, 4, 3); // return 10 (i.e. sum of the right red rectangle)
 
// Constraints:
//     m == matrix.length
//     n == matrix[i].length
//     1 <= m, n <= 200
//     -1000 <= matrix[i][j] <= 1000
//     0 <= row < m
//     0 <= col < n
//     -1000 <= val <= 1000
//     0 <= row1 <= row2 < m
//     0 <= col1 <= col2 < n
//     At most 5000 calls will be made to sumRegion and update.

import "fmt"

type NumMatrix struct {
    matrix [][]int
    preArr [][]int
}

func Constructor(matrix [][]int) NumMatrix {
    preArr := make([][]int, len(matrix))
    for i := 0; i < len(preArr); i++ {
        preArr[i] = make([]int, len(matrix[i]))
        preArr[i][0] = matrix[i][0]
        if i != 0 {
            preArr[i][0] += preArr[i-1][0]
        }
    }
    for j := 0; j < len(preArr[0]); j++ {
        preArr[0][j] = matrix[0][j]
        if j != 0 {
            preArr[0][j] += preArr[0][j-1]
        }
    }
    for i := 1; i < len(preArr); i++ {
        for j := 1; j < len(preArr[i]); j++ {
            preArr[i][j] = matrix[i][j] + preArr[i-1][j] + preArr[i][j-1] - preArr[i-1][j-1]
        }
    }
	return NumMatrix{matrix: matrix, preArr: preArr}
}

func (this *NumMatrix) Update(row int, col int, val int) {
    now := this.matrix[row][col]
    this.matrix[row][col] = val
    for i := row; i < len(this.preArr); i++ {
        for j := col; j < len(this.preArr[i]); j++ {
            this.preArr[i][j] += val - now
        }
    }
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    res := this.preArr[row2][col2]
    if row1 != 0 {
        res -= this.preArr[row1-1][col2]
    }
    if col1 != 0 {
        res -= this.preArr[row2][col1-1]
    }
    if row1 != 0 && col1 != 0 {
        res += this.preArr[row1-1][col1-1]
    }
    return res
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * obj.Update(row,col,val);
 * param_2 := obj.SumRegion(row1,col1,row2,col2);
 */

func main() {
    // NumMatrix numMatrix = new NumMatrix([[3, 0, 1, 4, 2], [5, 6, 3, 2, 1], [1, 2, 0, 1, 5], [4, 1, 0, 1, 7], [1, 0, 3, 0, 5]]);
    obj := Constructor([][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}})
    fmt.Println(obj)
    // numMatrix.sumRegion(2, 1, 4, 3); // return 8 (i.e. sum of the left red rectangle)
    fmt.Println(obj.SumRegion(2, 1, 4, 3)) // 8
    // numMatrix.update(3, 2, 2);       // matrix changes from left image to right image
    obj.Update(3, 2, 2)
    fmt.Println(obj)
    // numMatrix.sumRegion(2, 1, 4, 3); // return 10 (i.e. sum of the right red rectangle)
    fmt.Println(obj.SumRegion(2, 1, 4, 3)) // 10
}
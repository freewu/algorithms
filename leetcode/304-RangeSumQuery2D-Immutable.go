package main

import "fmt"

// 304. Range Sum Query 2D - Immutable
// Given a 2D matrix matrix, handle multiple queries of the following type:

// Calculate the sum of the elements of matrix inside the rectangle defined by its upper left corner (row1, col1) and lower right corner (row2, col2).
// Implement the NumMatrix class:

// NumMatrix(int[][] matrix) Initializes the object with the integer matrix matrix.
// int sumRegion(int row1, int col1, int row2, int col2) Returns the sum of the elements of matrix inside the rectangle defined by its upper left corner (row1, col1) and lower right corner (row2, col2).
// You must design an algorithm where sumRegion works on O(1) time complexity.

 

// Example 1:


// Input
// ["NumMatrix", "sumRegion", "sumRegion", "sumRegion"]
// [
// 	[
// 		[
// 			[3, 0, 1, 4, 2], 
// 			[5, 6, 3, 2, 1], 
// 			[1, 2, 0, 1, 5], 
// 			[4, 1, 0, 1, 7], 
// 			[1, 0, 3, 0, 5]
// 		]
// 	], 
// 	[2, 1, 4, 3], 
// 	[1, 1, 2, 2], 
// 	[1, 2, 2, 4]
// ]
// Output
// [null, 8, 11, 12]

// Explanation
// NumMatrix numMatrix = new NumMatrix([[3, 0, 1, 4, 2], [5, 6, 3, 2, 1], [1, 2, 0, 1, 5], [4, 1, 0, 1, 7], [1, 0, 3, 0, 5]]);
// numMatrix.sumRegion(2, 1, 4, 3); // return 8 (i.e sum of the red rectangle)
// numMatrix.sumRegion(1, 1, 2, 2); // return 11 (i.e sum of the green rectangle)
// numMatrix.sumRegion(1, 2, 2, 4); // return 12 (i.e sum of the blue rectangle)

// Constraints:

// 	m == matrix.length
// 	n == matrix[i].length
// 	1 <= m, n <= 200
// 	-104 <= matrix[i][j] <= 104
// 	0 <= row1 <= row2 < m
// 	0 <= col1 <= col2 < n
// 	At most 104 calls will be made to sumRegion.

type NumMatrix struct {
	cumsum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 {
		return NumMatrix{nil}
	}
	// 生成一个二维数组保存结果
	cumsum := make([][]int, len(matrix)+1)
	cumsum[0] = make([]int, len(matrix[0])+1)
	for i := range matrix {
		cumsum[i+1] = make([]int, len(matrix[i])+1)
		for j := range matrix[i] {
			cumsum[i+1][j+1] = matrix[i][j] + cumsum[i][j+1] + cumsum[i+1][j] - cumsum[i][j]
		}
	}
	fmt.Println(cumsum)
	return NumMatrix{cumsum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	cumsum := this.cumsum
	return cumsum[row2+1][col2+1] - cumsum[row1][col2+1] - cumsum[row2+1][col1] + cumsum[row1][col1]
}

func main() {
	numMatrix := Constructor([][]int{
		{3, 0, 1, 4, 2}, 
		{5, 6, 3, 2, 1}, 
		{1, 2, 0, 1, 5}, 
		{4, 1, 0, 1, 7}, 
		{1, 0, 3, 0, 5},
	})
	fmt.Println(numMatrix.SumRegion(2, 1, 4, 3)); // return 8 (i.e sum of the red rectangle)
	fmt.Println(numMatrix.SumRegion(1, 1, 2, 2)); // return 11 (i.e sum of the green rectangle)
	fmt.Println(numMatrix.SumRegion(1, 2, 2, 4)); // return 12 (i.e sum of the blue rectangle)
}
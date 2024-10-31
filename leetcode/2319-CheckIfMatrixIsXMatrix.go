package main

// 2319. Check if Matrix Is X-Matrix
// A square matrix is said to be an X-Matrix if both of the following conditions hold:
//     All the elements in the diagonals of the matrix are non-zero.
//     All other elements are 0.
    
// Given a 2D integer array grid of size n x n representing a square matrix, return true if grid is an X-Matrix. 
// Otherwise, return false.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/05/03/ex1.jpg" />
// Input: grid = [[2,0,0,1],[0,3,1,0],[0,5,2,0],[4,0,0,2]]
// Output: true
// Explanation: Refer to the diagram above. 
// An X-Matrix should have the green elements (diagonals) be non-zero and the red elements be 0.
// Thus, grid is an X-Matrix.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/05/03/ex2.jpg" />
// Input: grid = [[5,7,0],[0,3,1],[0,5,0]]
// Output: false
// Explanation: Refer to the diagram above.
// An X-Matrix should have the green elements (diagonals) be non-zero and the red elements be 0.
// Thus, grid is not an X-Matrix.

// Constraints:
//     n == grid.length == grid[i].length
//     3 <= n <= 100
//     0 <= grid[i][j] <= 10^5

import "fmt"

func checkXMatrix(grid [][]int) bool {
    n := len(grid)
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            // 矩阵对角线上的所有元素都 不是 0   对角线时: true  = (v != 0)
            // 矩阵中所有其他元素都是 0        非对角线时: false = (v == 0)
            if (i == j || i + j == n - 1) == (grid[i][j] == 0) {
                return false
            }
        }
    }
    return true
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/05/03/ex1.jpg" />
    // Input: grid = [[2,0,0,1],[0,3,1,0],[0,5,2,0],[4,0,0,2]]
    // Output: true
    // Explanation: Refer to the diagram above. 
    // An X-Matrix should have the green elements (diagonals) be non-zero and the red elements be 0.
    // Thus, grid is an X-Matrix.
    fmt.Println(checkXMatrix([][]int{{2,0,0,1},{0,3,1,0},{0,5,2,0},{4,0,0,2}})) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/05/03/ex2.jpg" />
    // Input: grid = [[5,7,0],[0,3,1],[0,5,0]]
    // Output: false
    // Explanation: Refer to the diagram above.
    // An X-Matrix should have the green elements (diagonals) be non-zero and the red elements be 0.
    // Thus, grid is not an X-Matrix.
    fmt.Println(checkXMatrix([][]int{{5,7,0},{0,3,1},{0,5,0}})) // false
}
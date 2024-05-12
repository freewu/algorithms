package main

// 2373. Largest Local Values in a Matrix
// You are given an n x n integer matrix grid.
// Generate an integer matrix maxLocal of size (n - 2) x (n - 2) such that:
//     maxLocal[i][j] is equal to the largest value of the 3 x 3 matrix in grid centered around row i + 1 and column j + 1.

// In other words, we want to find the largest value in every contiguous 3 x 3 matrix in grid.
// Return the generated matrix.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/06/21/ex1.png" / >
// Input: grid = [[9,9,8,1],[5,6,2,6],[8,2,6,4],[6,2,2,2]]
// Output: [[9,9],[8,6]]
// Explanation: The diagram above shows the original matrix and the generated matrix.
// Notice that each value in the generated matrix corresponds to the largest value of a contiguous 3 x 3 matrix in grid.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/07/02/ex2new2.png" / >
// Input: grid = [[1,1,1,1,1],[1,1,1,1,1],[1,1,2,1,1],[1,1,1,1,1],[1,1,1,1,1]]
// Output: [[2,2,2],[2,2,2],[2,2,2]]
// Explanation: Notice that the 2 is contained within every contiguous 3 x 3 matrix in grid.
 
// Constraints:
//     n == grid.length == grid[i].length
//     3 <= n <= 100
//     1 <= grid[i][j] <= 100

import "fmt"

func largestLocal(grid [][]int) [][]int {
    n := len(grid)
    maxLocal := make([][]int, n - 2)

    max := func(args... int) int {
        res := -1 >> 32 - 1
        for _,v := range args {
            if v > res {
                res = v
            }
        }
        return res
    }

    for i := range maxLocal {
        maxLocal[i] = make([]int, n - 2)
        for j := range maxLocal[i] {
            maxLocal[i][j] = max(
                grid[i][j], grid[i][j+1], grid[i][j+2],
                grid[i+1][j], grid[i+1][j+1], grid[i+1][j+2],
                grid[i+2][j], grid[i+2][j+1], grid[i+2][j+2],
            )
        }
    }
    return maxLocal
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/06/21/ex1.png" / >
    // Input: grid = [[9,9,8,1],[5,6,2,6],[8,2,6,4],[6,2,2,2]]
    // Output: [[9,9],[8,6]]
    // Explanation: The diagram above shows the original matrix and the generated matrix.
    // Notice that each value in the generated matrix corresponds to the largest value of a contiguous 3 x 3 matrix in grid.
    fmt.Println(largestLocal([][]int{{1,1,1,1,1},{1,1,1,1,1},{1,1,2,1,1},{1,1,1,1,1},{1,1,1,1,1}})) // [[9,9],[8,6]]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/07/02/ex2new2.png" / >
    // Input: grid = [[1,1,1,1,1],[1,1,1,1,1],[1,1,2,1,1],[1,1,1,1,1],[1,1,1,1,1]]
    // Output: [[2,2,2],[2,2,2],[2,2,2]]
    // Explanation: Notice that the 2 is contained within every contiguous 3 x 3 matrix in grid.
    fmt.Println(largestLocal([][]int{{9,9,8,1},{5,6,2,6},{8,2,6,4},{6,2,2,2}})) // [[2,2,2],[2,2,2],[2,2,2]]
}
package main

// 840. Magic Squares In Grid
// A 3 x 3 magic square is a 3 x 3 grid filled with distinct numbers from 1 to 9 such that each row, column, and both diagonals all have the same sum.
// Given a row x col grid of integers, how many 3 x 3 contiguous magic square subgrids are there?
// Note: while a magic square can only contain numbers from 1 to 9, grid may contain numbers up to 15.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/09/11/magic_main.jpg" />
// Input: grid = [[4,3,8,4],[9,5,1,9],[2,7,6,2]]
// Output: 1
// Explanation: 
// The following subgrid is a 3 x 3 magic square:
// <img src="https://assets.leetcode.com/uploads/2020/09/11/magic_valid.jpg" />
// while this one is not:
// <img src="https://assets.leetcode.com/uploads/2020/09/11/magic_invalid.jpg" />
// In total, there is only one magic square inside the given grid.

// Example 2:
// Input: grid = [[8]]
// Output: 0

// Constraints:
//     row == grid.length
//     col == grid[i].length
//     1 <= row, col <= 10
//     0 <= grid[i][j] <= 15

import "fmt"

func numMagicSquaresInside(grid [][]int) int {
    res, col, prefix := 0, len(grid[0]), make([]int, len(grid))
    calc := func(grid [][]int, row, col int) bool {
        visited, sum := make(map[int]bool), make([]int, 3)
        for j := col; j > col-3; j -- {
            for i := row - 2; i <= row; i ++ {
                if visited[grid[i][j]] ||  grid[i][j]> 9 || grid[i][j] < 1 {
                    return false
                }
                visited[grid[i][j]] = true
                sum[col-j] += grid[i][j]
            } 
        }
        if sum[0] != sum[1] || sum[0] != sum[2] {
            return false
        } 
        if (grid[row][col] + grid[row-2][col-2]) != (grid[row][col-2] + grid[row-2][col]){
            return false
        }
        return true
    }
    for i := 0; i < col; i++ {
        for j := 0; j < len(grid); j ++ {
            prefix[j] += grid[j][i]
            if i >= 3 { prefix[j] -= grid[j][i-3] }
            if i >=2 && j >= 2 && prefix[j] == prefix[j-1] && prefix[j] == prefix[j-2] {
                if calc(grid, j, i) {
                    res ++
                }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/09/11/magic_main.jpg" />
    // Input: grid = [[4,3,8,4],[9,5,1,9],[2,7,6,2]]
    // Output: 1
    // Explanation: 
    // The following subgrid is a 3 x 3 magic square:
    // <img src="https://assets.leetcode.com/uploads/2020/09/11/magic_valid.jpg" />
    // while this one is not:
    // <img src="https://assets.leetcode.com/uploads/2020/09/11/magic_invalid.jpg" />
    // In total, there is only one magic square inside the given grid.
    fmt.Println(numMagicSquaresInside([][]int{{4,3,8,4}, {9,5,1,9},{2,7,6,2}})) // 1
    // Example 2:
    // Input: grid = [[8]]
    // Output: 0
    fmt.Println(numMagicSquaresInside([][]int{{8}})) // 0
}
package main

// 3127. Make a Square with the Same Color
// You are given a 2D matrix grid of size 3 x 3 consisting only of characters 'B' and 'W'. 
// Character 'W' represents the white color, and character 'B' represents the black color.

// Your task is to change the color of at most one cell so that the matrix has a 2 x 2 square where all cells are of the same color.
// Return true if it is possible to create a 2 x 2 square of the same color, otherwise, return false.

// Example 1:
// Input: grid = [["B","W","B"],["B","W","W"],["B","W","B"]]
// Output: true
// Explanation:
// It can be done by changing the color of the grid[0][2].

// Example 2:
// Input: grid = [["B","W","B"],["W","B","W"],["B","W","B"]]
// Output: false
// Explanation:
// It cannot be done by changing at most one cell.

// Example 3:
// Input: grid = [["B","W","B"],["B","W","W"],["B","W","W"]]
// Output: true
// Explanation:
// The grid already contains a 2 x 2 square of the same color.

// Constraints:
//     grid.length == 3
//     grid[i].length == 3
//     grid[i][j] is either 'W' or 'B'.

import "fmt"

func canMakeSquare(grid [][]byte) bool {
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if ((j - 1 >= 0 && grid[i][j] == grid[i][j-1]) || (j + 1 < 3  && grid[i][j] == grid[i][j+1])) &&
               ((i + 1 < 3  && grid[i][j] == grid[i+1][j]) || (i - 1 >= 0 && grid[i][j] == grid[i-1][j])) {
                return true
            }
        }
    }
    return false
}

func main() {
    // Example 1:
    // Input: grid = [["B","W","B"],["B","W","W"],["B","W","B"]]
    // Output: true
    // Explanation:
    // It can be done by changing the color of the grid[0][2].
    fmt.Println(canMakeSquare([][]byte{{'B','W','B'},{'B','W','W'},{'B','W','W'}})) // true
    // Example 2:
    // Input: grid = [["B","W","B"],["W","B","W"],["B","W","B"]]
    // Output: false
    // Explanation:
    // It cannot be done by changing at most one cell.
    fmt.Println(canMakeSquare([][]byte{{'B','W','B'},{'W','B','W'},{'B','W','B'}})) // false
    // Example 3:
    // Input: grid = [["B","W","B"],["B","W","W"],["B","W","W"]]
    // Output: true
    // Explanation:
    // The grid already contains a 2 x 2 square of the same color.
    fmt.Println(canMakeSquare([][]byte{{'B','W','B'},{'B','W','W'},{'B','W','B'}})) // true
}
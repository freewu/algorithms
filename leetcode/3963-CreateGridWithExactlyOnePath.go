package main

// 3963. Create Grid With Exactly One Path
// You are given two integers m and n, representing the number of rows and columns of a grid.

// Construct any m x n grid consisting only of the characters '.' and '#', where:
//     1. '.' represents a free cell.
//     2. '#'  represents an obstacle cell.

// A valid path is a sequence of free cells that:
//     1. Starts at the top-left cell (0, 0).
//     2. Ends at the bottom-right cell (m - 1, n - 1).

// Moves only:
//     1. Right, from (i, j) to (i, j + 1), or
//     2. Down, from (i, j) to (i + 1, j).

// Return any grid such that there is exactly one valid path from the top-left cell to the bottom-right cell.

// Example 1:
// Input: m = 2, n = 3
// Output: ["..#","#.."]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2026/05/26/screenshot-2026-05-26-at-61005pm.png" />
// The only valid path is: (0,0) → (0,1) → (1,1) → (1,2)

// Example 2:
// Input: m = 3, n = 3
// Output: ["..#","#..","##."]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2026/05/26/screenshot-2026-05-26-at-61129pm.png" />
// The only valid path is: (0,0) → (0,1) → (1,1) → (1,2) → (2,2)

// Example 3:
// Input: m = 1, n = 4
// Output: ["...."]
// Explanation:
// The only valid path is: (0,0) → (0,1) → (0,2) → (0,3)

// Constraints:
//     1 <= m, n <= 25

import "fmt"
import "strings"

func createGrid(m int, n int) []string {
    res := make([]string, m)
    res[0] = strings.Repeat(".", n)
    row := strings.Repeat("#", n - 1) + "." // 避免在循环中反复创建字符串
    for i := 1; i < m; i++ {
        res[i] = row
    }
    return res  
}

func main() {
    // Example 1:
    // Input: m = 2, n = 3
    // Output: ["..#","#.."]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2026/05/26/screenshot-2026-05-26-at-61005pm.png" />
    // The only valid path is: (0,0) → (0,1) → (1,1) → (1,2)
    fmt.Println(createGrid(2, 3)) // ["..#","#.."]
    // Example 2:
    // Input: m = 3, n = 3
    // Output: ["..#","#..","##."]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2026/05/26/screenshot-2026-05-26-at-61129pm.png" />
    // The only valid path is: (0,0) → (0,1) → (1,1) → (1,2) → (2,2)
    fmt.Println(createGrid(3, 3)) // ["..#","#..","##.""]
    // Example 3:
    // Input: m = 1, n = 4
    // Output: ["...."]
    // Explanation:
    // The only valid path is: (0,0) → (0,1) → (0,2) → (0,3)    
    fmt.Println(createGrid(1, 4)) // ["...."]

    fmt.Println(createGrid(1, 1)) // [.]
    fmt.Println(createGrid(1, 25)) // [.........................]
    fmt.Println(createGrid(25, 1)) // [. . . . . . . . . . . . . . . . . . . . . . . . .]
    fmt.Println(createGrid(25, 25)) // [......................... ########################. ########################. ########################. ########################. ########################. ########################. ########################. ########################. ########################. ########################. ########################. ########################. ########################. ########################. ########################. ########################. ########################. ########################. ########################. ########################. ########################. ########################. ########################. ########################.] 
}
package main

// 3988. Create Grid With Exactly K Paths I
// You are given three integers m, n, and k.
// Construct any m x n grid consisting only of the characters '.' and '#', where:
//     1. '.' represents a free cell.
//     2. '#'  represents an obstacle cell.

// A valid path is a sequence of free cells that:
//     1. Starts at the top-left cell (0, 0).
//     2. Ends at the bottom-right cell (m - 1, n - 1).
//     3. Moves only:
//         3.1. Right, from (i, j) to (i, j + 1), or
//         3.2. Down, from (i, j) to (i + 1, j).

// Return any grid such that there are exactly k valid paths from the top-left cell to the bottom-right cell. 
// If no such grid exists, return an empty array.

// Example 1:
// Input: m = 2, n = 3, k = 2
// Output: ["...","#.."]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2026/05/26/screenshot-2026-05-27-at-113554am.png" />
// There are exactly k = 2 valid paths from (0, 0) to (1, 2):
// (0, 0) → (0, 1) → (0, 2) → (1, 2)
// (0, 0) → (0, 1) → (1, 1) → (1, 2)

// Example 2:
// Input: m = 3, n = 3, k = 4
// Output: ["..#","...","#.."]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2026/05/26/screenshot-2026-05-27-at-113452am.png" />
// There are exactly k = 4 valid paths from (0, 0) to (2, 2):
// (0, 0) → (0, 1) → (1, 1) → (1, 2) → (2, 2)
// (0, 0) → (0, 1) → (1, 1) → (2, 1) → (2, 2)
// (0, 0) → (1, 0) → (1, 1) → (1, 2) → (2, 2)
// (0, 0) → (1, 0) → (1, 1) → (2, 1) → (2, 2)

// Example 3:
// Input: m = 1, n = 4, k = 2
// Output: []
// Explanation:​
// No grid exists with exactly k = 2 valid paths for a 1 x 4 grid, so the answer is an empty array.

// Constraints:
//     1 <= m, n <= 10
//     1 <= k <= 4

import "fmt"
import "bytes"
import "strings"

func createGrid(m, n, k int) []string {
    // 特判
    if k == 4 && m == 3 && n == 3 {
        return []string{"..#", "...", "#.."}
    }
    if m == 1 || n == 1 {
        // 一行或一列，只能有一种方案
        if k > 1 {
            return nil
        }
        // 全为 '.'
        res := make([]string, m)
        row := strings.Repeat(".", n)
        for i := range res {
            res[i] = row
        }
        return res
    }
    // 至少要有 k 行或 k 列（特殊情况上面已判断）
    if m < k && n < k {
        return nil
    }
    // 初始全为 '#'
    arr := make([][]byte, m)
    for i := range m - 1 {
        arr[i] = bytes.Repeat([]byte{'#'}, n)
        arr[i][0] = '.' // 第一列全为 '.'   
    }
    arr[m-1] = bytes.Repeat([]byte{'.'}, n) // 最后一行全为 '.'
    if n >= k { // 至少有 k 列 
        // 倒数第二行开头 k 个 '.'
        for j := 1; j < k; j++ {
            arr[m-2][j] = '.'
        }
    } else { // 至少有 k 行
        // 第二列末尾 k 个 '.'
        for _, row := range arr[m-k : m-1] {
            row[1] = '.'
        }
    }
    res := make([]string, m)
    for i, row := range arr {
        res[i] = string(row)
    }
    return res
}

func main() {
    // Example 1:
    // Input: m = 2, n = 3, k = 2
    // Output: ["...","#.."]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2026/05/26/screenshot-2026-05-27-at-113554am.png" />
    // There are exactly k = 2 valid paths from (0, 0) to (1, 2):
    // (0, 0) → (0, 1) → (0, 2) → (1, 2)
    // (0, 0) → (0, 1) → (1, 1) → (1, 2)
    fmt.Println(createGrid(2, 3, 2)) // ["...","#.."]
    // Example 2:
    // Input: m = 3, n = 3, k = 4
    // Output: ["..#","...","#.."]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2026/05/26/screenshot-2026-05-27-at-113452am.png" />
    // There are exactly k = 4 valid paths from (0, 0) to (2, 2):
    // (0, 0) → (0, 1) → (1, 1) → (1, 2) → (2, 2)
    // (0, 0) → (0, 1) → (1, 1) → (2, 1) → (2, 2)
    // (0, 0) → (1, 0) → (1, 1) → (1, 2) → (2, 2)
    // (0, 0) → (1, 0) → (1, 1) → (2, 1) → (2, 2)
    fmt.Println(createGrid(3, 3, 4)) // ["..#","...","#.."]
    // Example 3:
    // Input: m = 1, n = 4, k = 2
    // Output: []
    // Explanation:​
    // No grid exists with exactly k = 2 valid paths for a 1 x 4 grid, so the answer is an empty array.
    fmt.Println(createGrid(1, 4, 2)) // []

    fmt.Println(createGrid(1, 1, 1)) // ["."]
    fmt.Println(createGrid(1, 1, 4)) // []
    fmt.Println(createGrid(10, 1, 1)) // [".”,".”,".”,".”,".”,".”,".”,".”,".”,".”,"."]
    fmt.Println(createGrid(1, 10, 1)) // [".........."]
    fmt.Println(createGrid(10, 1, 4)) // []
    fmt.Println(createGrid(1, 10, 4)) // []
    fmt.Println(createGrid(10, 10, 1)) // [.######### .######### .######### .######### .######### .######### .######### .######### .######### ..........]
    fmt.Println(createGrid(10, 10, 4)) // [.######### .######### .######### .######### .######### .######### .######### .######### ....###### ..........]
}
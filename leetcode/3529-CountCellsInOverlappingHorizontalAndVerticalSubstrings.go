package main

// 3529. Count Cells in Overlapping Horizontal and Vertical Substrings
// You are given an m x n matrix grid consisting of characters and a string pattern.

// A horizontal substring is a contiguous sequence of characters read from left to right. 
// If the end of a row is reached before the substring is complete, it wraps to the first column of the next row and continues as needed. 
// You do not wrap from the bottom row back to the top.

// A vertical substring is a contiguous sequence of characters read from top to bottom. 
// If the bottom of a column is reached before the substring is complete, it wraps to the first row of the next column and continues as needed. 
// You do not wrap from the last column back to the first.

// Count the number of cells in the matrix that satisfy the following condition:
//     1. The cell must be part of at least one horizontal substring and at least one vertical substring, 
//        where both substrings are equal to the given pattern.

// Return the count of these cells.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2025/03/03/gridtwosubstringsdrawio.png" />
// Input: grid = [["a","a","c","c"],["b","b","b","c"],["a","a","b","a"],["c","a","a","c"],["a","a","c","c"]], pattern = "abaca"
// Output: 1
// Explanation:
// The pattern "abaca" appears once as a horizontal substring (colored blue) and once as a vertical substring (colored red), intersecting at one cell (colored purple).

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2025/03/03/gridexample2fixeddrawio.png" />
// Input: grid = [["c","a","a","a"],["a","a","b","a"],["b","b","a","a"],["a","a","b","a"]], pattern = "aba"
// Output: 4
// Explanation:
// The cells colored above are all part of at least one horizontal and one vertical substring matching the pattern "aba".

// Example 3:
// Input: grid = [["a"]], pattern = "a"
// Output: 1
 
// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 1000
//     1 <= m * n <= 10^5
//     1 <= pattern.length <= m * n
//     grid and pattern consist of only lowercase English letters.

import "fmt"
import "slices"

// kmp
func countCells(grid [][]byte, pattern string) int {
    calcPi := func(pattern string) []int {
        pi := make([]int, len(pattern))
        match := 0
        for i := 1; i < len(pi); i++ {
            b := pattern[i]
            for match > 0 && pattern[match] != b {
                match = pi[match-1]
            }
            if pattern[match] == b {
                match++
            }
            pi[i] = match
        }
        return pi
    }
    kmpSearch := func(text []byte, pattern string, pi []int) []int {
        n, match := len(text), 0
        diff := make([]int, n + 1)
        for i, b := range text {
            for match > 0 && pattern[match] != b {
                match = pi[match - 1]
            }
            if pattern[match] == b {
                match++
            }
            if match == len(pi) {
                diff[i - len(pi) + 1]++
                diff[i + 1]--
                match = pi[match - 1]
            }
        }
        for i := 1; i < n; i++ {
            diff[i] += diff[i - 1]
        }
        return diff[:n]
    }
    res, m, n := 0, len(grid), len(grid[0])
    hText, vText:= slices.Concat(grid...), make([]byte, 0, m * n)
    for i := 0; i < n; i++ {
        for _, row := range grid {
            vText = append(vText, row[i])
        }
    }
    pi := calcPi(pattern)
    inPatternH, inPatternV := kmpSearch(hText, pattern, pi), kmpSearch(vText, pattern, pi)
    for i, v := range inPatternH {
        if v > 0 && inPatternV[i % n * m + i / n] > 0 {
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2025/03/03/gridtwosubstringsdrawio.png" />
    // Input: grid = [["a","a","c","c"],["b","b","b","c"],["a","a","b","a"],["c","a","a","c"],["a","a","c","c"]], pattern = "abaca"
    // Output: 1
    // Explanation:
    // The pattern "abaca" appears once as a horizontal substring (colored blue) and once as a vertical substring (colored red), intersecting at one cell (colored purple).
    fmt.Println(countCells([][]byte{{'a','a','c','c'},{'b','b','b','c'},{'a','a','b','a'},{'c','a','a','c'},{'a','a','c','c'}}, "abaca")) // 1
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2025/03/03/gridexample2fixeddrawio.png" />
    // Input: grid = [["c","a","a","a"],["a","a","b","a"],["b","b","a","a"],["a","a","b","a"]], pattern = "aba"
    // Output: 4
    // Explanation:
    // The cells colored above are all part of at least one horizontal and one vertical substring matching the pattern "aba".
    fmt.Println(countCells([][]byte{{'c','a','a','a'},{'a','a','b','a'},{'b','b','a','a'},{'a','a','b','a'}}, "aba")) // 4
    // Example 3:
    // Input: grid = [["a"]], pattern = "a"
    // Output: 1
    fmt.Println(countCells([][]byte{{'a'}}, "a")) // 1
}
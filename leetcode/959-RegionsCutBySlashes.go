package main

// 959. Regions Cut By Slashes
// An n x n grid is composed of 1 x 1 squares where each 1 x 1 square consists of a '/', '\', or blank space ' '. 
// These characters divide the square into contiguous regions.

// Given the grid grid represented as a string array, return the number of regions.
// Note that backslash characters are escaped, so a '\' is represented as '\\'.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2018/12/15/1.png" />
// Input: grid = [" /","/ "]
// Output: 2

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2018/12/15/2.png" />
// Input: grid = [" /","  "]
// Output: 1

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2018/12/15/4.png" />
// Input: grid = ["/\\","\\/"]
// Output: 5
// Explanation: Recall that because \ characters are escaped, "\\/" refers to \/, and "/\\" refers to /\.
 
// Constraints:
//     n == grid.length == grid[i].length
//     1 <= n <= 30
//     grid[i][j] is either '/', '\', or ' '.

import "fmt"

// 并查集 + 欧拉公式
func regionsBySlashes(grid []string) int {
    res, n := 1, len(grid)
    // 并查集
    mat := make([]int, (n+1)*(n+1))
    var find func(x int) int
    find = func(x int) int {
        if mat[x]!=x {
            mat[x] = find(mat[x])
        }
        return mat[x]
    }
    var union func(x, y int) int
    union = func(x, y int) int {
        rx , ry := find(x), find(y)
        if rx == ry {
            return 1
        }
        mat[ry] = rx
        return 0
    }
    getIndex := func(n, r, c int) int { return r *(n+1) + c }
    for i := 0; i <= n; i++ {
        for j := 0; j <= n; j++ {
            if i == 0 || j == 0 || i == n || j == n {
                continue
            } else {
                mat[getIndex(n, i, j)] = getIndex(n, i, j)
            }
        }
    }
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == ' ' { continue }
            if grid[i][j] == '/' {
                res += union(mat[getIndex(n, i, j+1)], mat[getIndex(n, i+1, j)])
            } else {
                res += union(mat[getIndex(n, i, j)], mat[getIndex(n, i+1, j+1)])
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2018/12/15/1.png" />
    // Input: grid = [" /","/ "]
    // Output: 2
    fmt.Println(regionsBySlashes([]string{" /","/ "})) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2018/12/15/2.png" />
    // Input: grid = [" /","  "]
    // Output: 1
    fmt.Println(regionsBySlashes([]string{" /","  "})) // 1
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2018/12/15/4.png" />
    // Input: grid = ["/\\","\\/"]
    // Output: 5
    // Explanation: Recall that because \ characters are escaped, "\\/" refers to \/, and "/\\" refers to /\.
    fmt.Println(regionsBySlashes([]string{"/\\","\\/"})) // 5
}
package main

// 1559. Detect Cycles in 2D Grid
// Given a 2D array of characters grid of size m x n, 
// you need to find if there exists any cycle consisting of the same value in grid.

// A cycle is a path of length 4 or more in the grid that starts and ends at the same cell. 
// From a given cell, you can move to one of the cells adjacent to it - in one of the four directions (up, down, left, or right), if it has the same value of the current cell.

// Also, you cannot move to the cell that you visited in your last move. 
// For example, the cycle (1, 1) -> (1, 2) -> (1, 1) is invalid because from (1, 2) we visited (1, 1) which was the last visited cell.

// Return true if any cycle of the same value exists in grid, otherwise, return false.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/07/15/1.png" />
// Input: grid = [["a","a","a","a"],["a","b","b","a"],["a","b","b","a"],["a","a","a","a"]]
// Output: true
// Explanation: There are two valid cycles shown in different colors in the image below:
// <img src="https://assets.leetcode.com/uploads/2020/07/15/11.png" />

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/07/15/22.png" />
// Input: grid = [["c","c","c","a"],["c","d","c","c"],["c","c","e","c"],["f","c","c","c"]]
// Output: true
// Explanation: There is only one valid cycle highlighted in the image below:
// <img src="https://assets.leetcode.com/uploads/2020/07/15/2.png" />

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2020/07/15/3.png" />
// Input: grid = [["a","b","b"],["b","z","b"],["b","b","a"]]
// Output: false

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 500
//     grid consists only of lowercase English letters.

import "fmt"

func containsCycle(grid [][]byte) bool {
    m, n := len(grid), len(grid[0])
    set := make(map[[2]int]bool)
    directions := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

    var dfs func(grid [][]byte, i int, j int, set map[[2]int]bool, path map[[2]int]bool,lastI int, lastJ int) bool
    dfs = func(grid [][]byte, i int, j int, set map[[2]int]bool, path map[[2]int]bool,lastI int, lastJ int) bool {
        set[[2]int{i, j}], path[[2]int{i, j}] = true, true
        for _, dir := range directions {
            ii, jj := i + dir[0], j + dir[1]
            if ii == lastI && jj == lastJ { continue }
            if ii < 0 || jj < 0 || ii >= len(grid) || jj >= len(grid[0]) ||
            grid[ii][jj] != grid[i][j] { continue }
            if path[[2]int{ii, jj}] { return true }
            if dfs(grid, ii, jj, set, path, i, j) { return true }
        }
        return false
    } 
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if !set[[2]int{i, j}] {
                path := make(map[[2]int]bool)
                if dfs(grid, i, j, set, path, -1, -1) {
                    return true
                }
            }
        }
    }
    return false
}

// 并查集
func containsCycle1(grid [][]byte) bool {
    m, n := len(grid), len(grid[0])
    pa, size := make([]int, m*n), make([]int, m*n)
    for i := range pa {
        pa[i], size[i] = i, 1
    }
    var find func(x int) int
    find = func(x int) int {
        if pa[x] != x {
            pa[x] = find(pa[x])
        }
        return pa[x]
    }
    union := func(x, y int) bool {
        x, y = find(x), find(y)
        if x == y {
            return false
        }
        if x < y {
            x, y = y, x
        }
        pa[y] = x
        size[x] += size[y]
        return true
    }
    for i, g := range grid {
        for j, v := range g {
            if i > 0 && grid[i-1][j] == v {
                if !union(i*n+j, (i-1)*n+j) {
                    return true
                }
            }
            if j > 0 && grid[i][j-1] == v {
                if !union(i*n+j, i*n+j-1) {
                    return true
                }
            }
        }
    }
    return false
}


func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/07/15/1.png" />
    // Input: grid = [["a","a","a","a"],["a","b","b","a"],["a","b","b","a"],["a","a","a","a"]]
    // Output: true
    // Explanation: There are two valid cycles shown in different colors in the image below:
    // <img src="https://assets.leetcode.com/uploads/2020/07/15/11.png" />
    fmt.Println(containsCycle([][]byte{{'a','a','a','a'},{'a','b','b','a'},{'a','b','b','a'},{'a','a','a','a'}})) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/07/15/22.png" />
    // Input: grid = [["c","c","c","a"],["c","d","c","c"],["c","c","e","c"],["f","c","c","c"]]
    // Output: true
    // Explanation: There is only one valid cycle highlighted in the image below:
    // <img src="https://assets.leetcode.com/uploads/2020/07/15/2.png" />
    fmt.Println(containsCycle([][]byte{{'c','c','c','a'},{'c','d','c','c'},{'c','c','e','c'},{'f','c','c','c'}})) // true
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2020/07/15/3.png" />
    // Input: grid = [["a","b","b"],["b","z","b"],["b","b","a"]]
    // Output: false
    fmt.Println(containsCycle([][]byte{{'a','b','b'},{'b','z','b'},{'b','b','a'}})) // false

    fmt.Println(containsCycle1([][]byte{{'a','a','a','a'},{'a','b','b','a'},{'a','b','b','a'},{'a','a','a','a'}})) // true
    fmt.Println(containsCycle1([][]byte{{'c','c','c','a'},{'c','d','c','c'},{'c','c','e','c'},{'f','c','c','c'}})) // true
    fmt.Println(containsCycle1([][]byte{{'a','b','b'},{'b','z','b'},{'b','b','a'}})) // false
}
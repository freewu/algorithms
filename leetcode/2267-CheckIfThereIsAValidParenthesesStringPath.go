package main

// 2267. Check if There Is a Valid Parentheses String Path
// A parentheses string is a non-empty string consisting only of '(' and ')'. 
// It is valid if any of the following conditions is true:
//     1. It is ().
//     2. It can be written as AB (A concatenated with B), where A and B are valid parentheses strings.
//     3. It can be written as (A), where A is a valid parentheses string.

// You are given an m x n matrix of parentheses grid. 
// A valid parentheses string path in the grid is a path satisfying all of the following conditions:
//     1. The path starts from the upper left cell (0, 0).
//     2. The path ends at the bottom-right cell (m - 1, n - 1).
//     3. The path only ever moves down or right.
//     4. The resulting parentheses string formed by the path is valid.

// Return true if there exists a valid parentheses string path in the grid. Otherwise, return false.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/03/15/example1drawio.png" />
// Input: grid = [["(","(","("],[")","(",")"],["(","(",")"],["(","(",")"]]
// Output: true
// Explanation: The above diagram shows two possible paths that form valid parentheses strings.
// The first path shown results in the valid parentheses string "()(())".
// The second path shown results in the valid parentheses string "((()))".
// Note that there may be other valid parentheses string paths.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/03/15/example2drawio.png" />
// Input: grid = [[")",")"],["(","("]]
// Output: false
// Explanation: The two possible paths form the parentheses strings "))(" and ")((". Since neither of them are valid parentheses strings, we return false.

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 100
//     grid[i][j] is either '(' or ')'.

import "fmt"

// dp
func hasValidPath(grid [][]byte) bool {
    m, n := len(grid), len(grid[0])
    dp := make([][][]bool, m+1)
    for i := 0; i <= m; i++ {
        dp[i] = make([][]bool, n + 1)
        for j := 0; j <= n; j++ {
            dp[i][j] = make([]bool, m + n + 1)
        }
    }
    dp[0][1][0] = true
    for i, x := 0, 1; i < m; i, x = i+1, x+1 {
        for j, y := 0, 1; j < n; j, y = j+1, y+1 {
            for k := 0; k <= m + n; k++ {
                dp[x][y][k] = dp[i][y][k] || dp[x][j][k]
            }
            if grid[i][j] == '(' {
                for k := m + n; k >= 1; k-- {
                    dp[x][y][k] = dp[x][y][k-1]
                }
                dp[x][y][0] = false
            } else {
                for k := 0; k < m + n; k++ {
                    dp[x][y][k] = dp[x][y][k+1]
                }
            }
        }
    }
    return dp[m][n][0]
}

// dfs
func hasValidPath1(grid [][]byte) bool {
    m, n := len(grid), len(grid[0])
    if (m + n) % 2 == 0 || grid[0][0] == ')' || grid[m - 1][n - 1] == '(' { return false } // 剪枝
    visited := make([][][]bool, m)
    for i := range visited {
        visited[i] = make([][]bool, n)
        for j := range visited[i] {
            visited[i][j] = make([]bool, (m + n + 1)/2)
        }
    }
    var dfs func(x, y, c int) bool
    dfs = func(x, y, c int) bool {
        if c > m - x + n - y - 1 { return false } // 剪枝：即使后面都是 ')' 也不能将 c 减为 0
        if x == m - 1 && y == n - 1  { return c == 1} // 终点, 终点一定是 ')'
        if visited[x][y][c] { return false }// 重复访问
        visited[x][y][c] = true
        if grid[x][y] == '(' {
            c++
        } else if c--; c < 0 { // 非法括号字符串
            return false
        }
        return x < m - 1 && dfs(x + 1, y, c) || y < n - 1 && dfs(x, y+1, c) // 往下或者往右
    }
    return dfs(0, 0, 0) // 起点
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/03/15/example1drawio.png" />
    // Input: grid = [["(","(","("],[")","(",")"],["(","(",")"],["(","(",")"]]
    // Output: true
    // Explanation: The above diagram shows two possible paths that form valid parentheses strings.
    // The first path shown results in the valid parentheses string "()(())".
    // The second path shown results in the valid parentheses string "((()))".
    // Note that there may be other valid parentheses string paths.
    fmt.Println(hasValidPath([][]byte{{'(','(','('},{')','(',')'},{'(','(',')'},{'(','(',')'}})) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/03/15/example2drawio.png" />
    // Input: grid = [[")",")"],["(","("]]
    // Output: false
    // Explanation: The two possible paths form the parentheses strings "))(" and ")((". Since neither of them are valid parentheses strings, we return false.
    fmt.Println(hasValidPath([][]byte{{')',')'},{'(','('}})) // false

    fmt.Println(hasValidPath1([][]byte{{'(','(','('},{')','(',')'},{'(','(',')'},{'(','(',')'}})) // true
    fmt.Println(hasValidPath1([][]byte{{')',')'},{'(','('}})) // false
}
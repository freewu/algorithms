package main

// 1905. Count Sub Islands
// You are given two m x n binary matrices grid1 and grid2 containing only 0's (representing water) and 1's (representing land). 
// An island is a group of 1's connected 4-directionally (horizontal or vertical). 
// Any cells outside of the grid are considered water cells.

// An island in grid2 is considered a sub-island if there is an island in grid1 that contains all the cells that make up this island in grid2.
// Return the number of islands in grid2 that are considered sub-islands.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/06/10/test1.png" />
// Input: grid1 = [[1,1,1,0,0],[0,1,1,1,1],[0,0,0,0,0],[1,0,0,0,0],[1,1,0,1,1]], grid2 = [[1,1,1,0,0],[0,0,1,1,1],[0,1,0,0,0],[1,0,1,1,0],[0,1,0,1,0]]
// Output: 3
// Explanation: In the picture above, the grid on the left is grid1 and the grid on the right is grid2.
// The 1s colored red in grid2 are those considered to be part of a sub-island. There are three sub-islands.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/06/03/testcasex2.png" />
// Input: grid1 = [[1,0,1,0,1],[1,1,1,1,1],[0,0,0,0,0],[1,1,1,1,1],[1,0,1,0,1]], grid2 = [[0,0,0,0,0],[1,1,1,1,1],[0,1,0,1,0],[0,1,0,1,0],[1,0,0,0,1]]
// Output: 2 
// Explanation: In the picture above, the grid on the left is grid1 and the grid on the right is grid2.
// The 1s colored red in grid2 are those considered to be part of a sub-island. There are two sub-islands.

// Constraints:
//     m == grid1.length == grid2.length
//     n == grid1[i].length == grid2[i].length
//     1 <= m, n <= 500
//     grid1[i][j] and grid2[i][j] are either 0 or 1.

import "fmt"

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
    m, n := len(grid1), len(grid1[0])
    var dfs func(i int, j int, visited map[[2]int]bool) bool
    dfs = func(i, j int, visited map[[2]int]bool) bool {
        if i < 0 || j < 0 || i == m || j == n || grid2[i][j] == 0 || visited[[2]int{i, j}] {
            return true
        }
        visited[[2]int{i, j}] = true
        res := grid1[i][j] == 1
        res = dfs(i+1, j, visited) && res
        res = dfs(i-1, j, visited) && res
        res = dfs(i, j+1, visited) && res
        res = dfs(i, j-1, visited) && res
        return res
    }
    res, visited := 0, map[[2]int]bool{}
    for i := 0; i < len(grid2); i++ {
        for j := 0; j < len(grid2[i]); j++ {
            if grid1[i][j] == 1 && grid2[i][j] == 1 && !visited[[2]int{i, j}] && dfs(i, j, visited) {
                res += 1
            }
        }
    }
    return res
}

func countSubIslands1(grid1 [][]int, grid2 [][]int) int {
    SEA, LAND := 0, 1
    m, n := len(grid1), len(grid1[0])
    var traverse func(r, c int)
    traverse = func(r, c int) {
        if !(r >= 0 && r < m && c >= 0 && c < n) { return } // 边界
        if grid2[r][c] == SEA { return }
        grid2[r][c] = SEA
        traverse(r+1, c)
        traverse(r-1, c)
        traverse(r, c+1)
        traverse(r, c-1)
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid1[i][j] == SEA && grid2[i][j] == LAND {
                traverse(i, j)
            }
        }
    }
    res := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid2[i][j] == LAND {
                res += 1
                traverse(i, j)
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/06/10/test1.png" />
    // Input: grid1 = [[1,1,1,0,0],[0,1,1,1,1],[0,0,0,0,0],[1,0,0,0,0],[1,1,0,1,1]], grid2 = [[1,1,1,0,0],[0,0,1,1,1],[0,1,0,0,0],[1,0,1,1,0],[0,1,0,1,0]]
    // Output: 3
    // Explanation: In the picture above, the grid on the left is grid1 and the grid on the right is grid2.
    // The 1s colored red in grid2 are those considered to be part of a sub-island. There are three sub-islands.
    grid11 := [][]int{
        {1,1,1,0,0},
        {0,1,1,1,1},
        {0,0,0,0,0},
        {1,0,0,0,0},
        {1,1,0,1,1},
    }
    grid12 := [][]int{
        {1,1,1,0,0},
        {0,0,1,1,1},
        {0,1,0,0,0},
        {1,0,1,1,0},
        {0,1,0,1,0},
    }
    fmt.Println(countSubIslands(grid11,grid12)) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/06/03/testcasex2.png" />
    // Input: grid1 = [[1,0,1,0,1],[1,1,1,1,1],[0,0,0,0,0],[1,1,1,1,1],[1,0,1,0,1]], grid2 = [[0,0,0,0,0],[1,1,1,1,1],[0,1,0,1,0],[0,1,0,1,0],[1,0,0,0,1]]
    // Output: 2 
    // Explanation: In the picture above, the grid on the left is grid1 and the grid on the right is grid2.
    // The 1s colored red in grid2 are those considered to be part of a sub-island. There are two sub-islands.
    grid21 := [][]int{
        {1,0,1,0,1},
        {1,1,1,1,1},
        {0,0,0,0,0},
        {1,1,1,1,1},
        {1,0,1,0,1},
    }
    grid22 := [][]int{
        {0,0,0,0,0},
        {1,1,1,1,1},
        {0,1,0,1,0},
        {0,1,0,1,0},
        {1,0,0,0,1},
    }
    fmt.Println(countSubIslands(grid21,grid22)) // 2

    grid111 := [][]int{
        {1,1,1,0,0},
        {0,1,1,1,1},
        {0,0,0,0,0},
        {1,0,0,0,0},
        {1,1,0,1,1},
    }
    grid112 := [][]int{
        {1,1,1,0,0},
        {0,0,1,1,1},
        {0,1,0,0,0},
        {1,0,1,1,0},
        {0,1,0,1,0},
    }
    fmt.Println(countSubIslands1(grid111,grid112)) // 3
    grid121 := [][]int{
        {1,0,1,0,1},
        {1,1,1,1,1},
        {0,0,0,0,0},
        {1,1,1,1,1},
        {1,0,1,0,1},
    }
    grid122 := [][]int{
        {0,0,0,0,0},
        {1,1,1,1,1},
        {0,1,0,1,0},
        {0,1,0,1,0},
        {1,0,0,0,1},
    }
    fmt.Println(countSubIslands1(grid121,grid122)) // 2
}
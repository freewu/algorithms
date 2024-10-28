package main

// 2123. Minimum Operations to Remove Adjacent Ones in Matrix
// You are given a 0-indexed binary matrix grid. 
// In one operation, you can flip any 1 in grid to be 0.

// A binary matrix is well-isolated if there is no 1 in the matrix that is 4-directionally connected (i.e., horizontal and vertical) to another 1.

// Return the minimum number of operations to make grid well-isolated.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/12/23/image-20211223181501-1.png" />
// Input: grid = [[1,1,0],[0,1,1],[1,1,1]]
// Output: 3
// Explanation: Use 3 operations to change grid[0][1], grid[1][2], and grid[2][1] to 0.
// After, no more 1's are 4-directionally connected and grid is well-isolated.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/12/23/image-20211223181518-2.png" />
// Input: grid = [[0,0,0],[0,0,0],[0,0,0]]
// Output: 0
// Explanation: There are no 1's in grid and it is well-isolated.
// No operations were done so return 0.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/12/23/image-20211223181817-3.png" />
// Input: grid = [[0,1],[1,0]]
// Output: 0
// Explanation: None of the 1's are 4-directionally connected and grid is well-isolated.
// No operations were done so return 0.

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 300
//     grid[i][j] is either 0 or 1.

import "fmt"

// 转化为二分图，边只能存在于删除的节点和保留的节点之间
func minimumOperations(grid [][]int) int {
    diresctions := [][]int{{-1, 0}, {0, 1}, {0, -1}, {1, 0}}
    res, m, n := 0, len(grid), len(grid[0])
    pre, edges, visited := make(map[int]int), make(map[int][]int), make(map[int]bool)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                for _, dir := range diresctions {
                    ii, jj := i + dir[0], j + dir[1]
                    if ii >= 0 && ii < m && jj >= 0 && jj < n && grid[ii][jj] == 1 {
                        edges[i*n+j] = append(edges[i * n + j], ii * n + jj)
                    }
                }
            }
        }
    }
    for i := range edges {
        pre[i]=-1
    }
    var dfs func(i int) bool
    dfs = func(i int) bool {
        for _, j := range edges[i] {
            if visited[j] { continue }
            visited[j] = true
            if pre[j] == -1 || dfs(pre[j]) {
                pre[j], pre[i] = i, j
                return true
            }
        }
        return false
    }
    for i := range edges {
        visited = make(map[int]bool)
        if pre[i]==-1 && dfs(i) {
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/12/23/image-20211223181501-1.png" />
    // Input: grid = [[1,1,0],[0,1,1],[1,1,1]]
    // Output: 3
    // Explanation: Use 3 operations to change grid[0][1], grid[1][2], and grid[2][1] to 0.
    // After, no more 1's are 4-directionally connected and grid is well-isolated.
    fmt.Println(minimumOperations([][]int{{1,1,0},{0,1,1},{1,1,1}})) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/12/23/image-20211223181518-2.png" />
    // Input: grid = [[0,0,0],[0,0,0],[0,0,0]]
    // Output: 0
    // Explanation: There are no 1's in grid and it is well-isolated.
    // No operations were done so return 0.
    fmt.Println(minimumOperations([][]int{{0,0,0},{0,0,0},{0,0,0}})) // 0
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/12/23/image-20211223181817-3.png" />
    // Input: grid = [[0,1],[1,0]]
    // Output: 0
    // Explanation: None of the 1's are 4-directionally connected and grid is well-isolated.
    // No operations were done so return 0.
    fmt.Println(minimumOperations([][]int{{0,1},{1,0}})) // 0
}
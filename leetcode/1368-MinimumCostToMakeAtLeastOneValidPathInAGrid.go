package main

// 1368. Minimum Cost to Make at Least One Valid Path in a Grid
// Given an m x n grid. Each cell of the grid has a sign pointing to the next cell you should visit if you are currently in this cell. 
// The sign of grid[i][j] can be:
//     1 which means go to the cell to the right. (i.e go from grid[i][j] to grid[i][j + 1])
//     2 which means go to the cell to the left. (i.e go from grid[i][j] to grid[i][j - 1])
//     3 which means go to the lower cell. (i.e go from grid[i][j] to grid[i + 1][j])
//     4 which means go to the upper cell. (i.e go from grid[i][j] to grid[i - 1][j])

// Notice that there could be some signs on the cells of the grid that point outside the grid.

// You will initially start at the upper left cell (0, 0). 
// A valid path in the grid is a path that starts from the upper left cell (0, 0) 
// and ends at the bottom-right cell (m - 1, n - 1) following the signs on the grid. 
// The valid path does not have to be the shortest.

// You can modify the sign on a cell with cost = 1. 
// You can modify the sign on a cell one time only.

// Return the minimum cost to make the grid have at least one valid path.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/02/13/grid1.png" />
// Input: grid = [[1,1,1,1],[2,2,2,2],[1,1,1,1],[2,2,2,2]]
// Output: 3
// Explanation: You will start at point (0, 0).
// The path to (3, 3) is as follows. (0, 0) --> (0, 1) --> (0, 2) --> (0, 3) change the arrow to down with cost = 1 --> (1, 3) --> (1, 2) --> (1, 1) --> (1, 0) change the arrow to down with cost = 1 --> (2, 0) --> (2, 1) --> (2, 2) --> (2, 3) change the arrow to down with cost = 1 --> (3, 3)
// The total cost = 3.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/02/13/grid2.png" />
// Input: grid = [[1,1,3],[3,2,2],[1,1,4]]
// Output: 0
// Explanation: You can follow the path from (0, 0) to (2, 2).

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2020/02/13/grid3.png" />
// Input: grid = [[1,2],[4,3]]
// Output: 1
 
// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 100
//     1 <= grid[i][j] <= 4

import "fmt"

// bfs
func minCost(grid [][]int) int {
    directions := [][]int{{0,1},{0,-1},{1,0},{-1,0}}
    k, m, n := 0, len(grid), len(grid[0])
    dp := make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
        for j := 0; j < n; j++ {
            dp[i][j] = -1
        }
    }
    queue := [][]int{}
    var dfs func(row int, col int, k int)
    dfs = func(row int, col int, k int) {
        if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) || dp[row][col] != -1 { // 边界检测
            return
        }
        dp[row][col] = k
        queue = append(queue, []int{row, col})
        curr := grid[row][col] - 1
        nr, nc := row + directions[curr][0], col + directions[curr][1]
        dfs(nr, nc, k)
    }
    dfs( 0, 0, k)
    for len(queue) != 0 {
        k++
        n := len(queue)
        for i := 0; i < n; i++ {
            curr := queue[0] // pop
            queue = queue[1:]
            for _, dir := range directions {
                nr, nc := dir[0] + curr[0],  dir[1] + curr[1]
                dfs(nr, nc, k)
            }
        }
    }
    return dp[m - 1][n - 1]
}

// dfs
func minCost1(grid [][]int) int {
    dist, m, n := 0, len(grid), len(grid[0])
    queue, visited := [][]int{}, make([][]bool, m)
    for i, _ := range visited {
        visited[i] = make([]bool, n)
    }
    directions := [][]int{ {0, 1}, {0, -1}, {1, 0}, {-1, 0} }
    var dfs func(i, j int)
    dfs = func(i, j int) {
        if i < 0 || i >= m { return }
        if j < 0 || j >= n { return }
        if visited[i][j]   { return }
        visited[i][j] = true
        queue = append(queue, []int{ i, j, dist } )
        dir := directions[grid[i][j] - 1]
        dfs(i + dir[0], j + dir[1])
    }
    dfs(0, 0)
    for len(queue) > 0 {
        top := queue[0]
        queue = queue[1: ]
        if top[0] == m - 1 && top[1] == n - 1 { return top[2] }
        dist = top[2] + 1
        for _, dir := range directions {
            dfs(top[0] + dir[0], top[1] + dir[1])
        }
    }
    return 0
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/02/13/grid1.png" />
    // Input: grid = [[1,1,1,1],[2,2,2,2],[1,1,1,1],[2,2,2,2]]
    // Output: 3
    // Explanation: You will start at point (0, 0).
    // The path to (3, 3) is as follows. (0, 0) --> (0, 1) --> (0, 2) --> (0, 3) change the arrow to down with cost = 1 --> (1, 3) --> (1, 2) --> (1, 1) --> (1, 0) change the arrow to down with cost = 1 --> (2, 0) --> (2, 1) --> (2, 2) --> (2, 3) change the arrow to down with cost = 1 --> (3, 3)
    // The total cost = 3.
    grid1 := [][]int{
        {1,1,1,1},
        {2,2,2,2},
        {1,1,1,1},
        {2,2,2,2},
    }
    fmt.Println(minCost(grid1)) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/02/13/grid2.png" />
    // Input: grid = [[1,1,3],[3,2,2],[1,1,4]]
    // Output: 0
    // Explanation: You can follow the path from (0, 0) to (2, 2).
    grid2 := [][]int{
        {1,1,3},
        {3,2,2},
        {1,1,4},
    }
    fmt.Println(minCost(grid2)) // 0
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2020/02/13/grid3.png" />
    // Input: grid = [[1,2],[4,3]]
    // Output: 1
    grid3 := [][]int{
        {1,2},
        {4,3},
    }
    fmt.Println(minCost(grid3)) // 1

    fmt.Println(minCost1(grid1)) // 3
    fmt.Println(minCost1(grid2)) // 0
    fmt.Println(minCost1(grid3)) // 1
}
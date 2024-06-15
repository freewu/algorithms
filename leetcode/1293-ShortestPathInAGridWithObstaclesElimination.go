package main

// 1293. Shortest Path in a Grid with Obstacles Elimination
// You are given an m x n integer matrix grid where each cell is either 0 (empty) or 1 (obstacle). 
// You can move up, down, left, or right from and to an empty cell in one step.

// Return the minimum number of steps to walk from the upper left corner (0, 0) to the lower right corner (m - 1, n - 1) 
// given that you can eliminate at most k obstacles. 
// If it is not possible to find such walk return -1.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/09/30/short1-grid.jpg" />
// Input: grid = [[0,0,0],[1,1,0],[0,0,0],[0,1,1],[0,0,0]], k = 1
// Output: 6
// Explanation: 
// The shortest path without eliminating any obstacle is 10.
// The shortest path with one obstacle elimination at position (3,2) is 6. Such path is (0,0) -> (0,1) -> (0,2) -> (1,2) -> (2,2) -> (3,2) -> (4,2).

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/09/30/short2-grid.jpg" />
// Input: grid = [[0,1,1],[1,1,1],[1,0,0]], k = 1
// Output: -1
// Explanation: We need to eliminate at least two obstacles to find such a walk.
 
// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 40
//     1 <= k <= m * n
//     grid[i][j] is either 0 or 1.
//     grid[0][0] == grid[m - 1][n - 1] == 0

import "fmt"

// bfs
func shortestPath(grid [][]int, k int) int {
    dirs := [][]int{[]int{1, 0}, []int{-1, 0}, []int{0, 1}, []int{0, -1}}
    m, n := len(grid), len(grid[0])
    if k >= m + n - 3 {
        return m + n - 2
    }
    type tuple struct {
        x, y, z int
    }
    q, vis := []tuple{tuple{0, 0, k}}, make([][]bool, m * n)
    for i := range vis {
        vis[i] = make([]bool, k + 1)
    }
    d := 0
    for len(q) > 0 {
        for t := len(q); t > 0; t-- {
            i, j, z := q[0].x, q[0].y, q[0].z
            q = q[1:]
            if z < 0 || vis[i * n + j][z] {
                continue
            }
            vis[i * n + j][z] = true
            if i == m - 1 && j == n - 1 {
                return d
            }
            for _, dir := range dirs {
                x, y := i + dir[0], j + dir[1]
                if x < 0 || x >= m || y < 0 || y >= n {
                    continue
                } 
                if grid[x][y] == 1 && z > 0 {
                    q = append(q, tuple{x, y, z - 1})
                } else if grid[x][y] == 0 {
                    q = append(q, tuple{x, y, z})
                }
            }
        }
        d++
    }
    return -1
}

// dfs
func shortestPath1(grid [][]int, k int) int {
    m, n, inf := len(grid), len(grid[0]), 1 << 32 - 1
    if k >= m + n - 3 { // 假设网格中没有障碍物，只向右或向下走，将会是最短路径，共（m-1）+（n-1）步。 如果k>=m+n-3，不需要搜索。
        return m + n - 2
    }
    vis := make([][][]int, m) // 三维记忆数组， [x][y][cnt1] 分别表示坐标和当前墙的数量，存储的值为当前状态下遍历过的最短路径长度
    for i := range vis {
        vis[i] = make([][]int, n)
        for j := range vis[i] {
            vis[i][j] = make([]int, k+1)
            for l := range vis[i][j] {
                vis[i][j][l] = inf
            }
        }
    }
    dirs := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
    res := inf
    var dfs func(int, int, int, int) 
    dfs = func(x, y, cnt, wallCnt int) {
        if grid[x][y] == 1 {
            wallCnt++   // 墙数量+1
        }
        cnt++ // 步数+1
        if wallCnt > k { // 超过允许消除的障碍数
            return
        }
        if cnt >= res { // 步数 >= 之前的最短答案，就没必要了
            return
        }
        if cnt >= vis[x][y][wallCnt] { // 步数 >= 记忆化的那个
            return
        }
        if x == m-1 && y == n-1 { // 到达，更新答案
            res = cnt
            return
        }
        vis[x][y][wallCnt] = cnt // 记忆化
        for i := 0; i < 4; i++ {
            dx, dy := x + dirs[i][0], y + dirs[i][1]
            if dx >= 0 && dx < m && dy >= 0 && dy < n {
                dfs(dx, dy, cnt, wallCnt)
            }
        }
    }
    dfs(0, 0, -1, 0)
    if res == inf {
        return -1
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/09/30/short1-grid.jpg" />
    // Input: grid = [[0,0,0],[1,1,0],[0,0,0],[0,1,1],[0,0,0]], k = 1
    // Output: 6
    // Explanation: 
    // The shortest path without eliminating any obstacle is 10.
    // The shortest path with one obstacle elimination at position (3,2) is 6. Such path is (0,0) -> (0,1) -> (0,2) -> (1,2) -> (2,2) -> (3,2) -> (4,2).
    grid1 := [][]int{
        {0,0,0},
        {1,1,0},
        {0,0,0},
        {0,1,1},
        {0,0,0},
    }
    fmt.Println(shortestPath(grid1, 1)) // 6
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/09/30/short2-grid.jpg" />
    // Input: grid = [[0,1,1],[1,1,1],[1,0,0]], k = 1
    // Output: -1
    // Explanation: We need to eliminate at least two obstacles to find such a walk.
    grid2 := [][]int{
        {0,1,1},
        {1,1,1},
        {1,0,0},
    }
    fmt.Println(shortestPath(grid2, 1)) // -1

    fmt.Println(shortestPath1(grid1, 1)) // 6
    fmt.Println(shortestPath1(grid2, 1)) // -1

}
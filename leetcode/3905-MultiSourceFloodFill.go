package main

// 3905. Multi Source Flood Fill
// You are given two integers n and m representing the number of rows and columns of a grid, respectively.

// You are also given a 2D integer array sources, where sources[i] = [ri, ci, color​​​​​​​i] indicates that the cell (ri, ci) is initially colored with colori. 
// All other cells are initially uncolored and represented as 0.

// At each time step, every currently colored cell spreads its color to all adjacent uncolored cells in the four directions: up, down, left, and right. 
// All spreads happen simultaneously.

// If multiple colors reach the same uncolored cell at the same time step, the cell takes the color with the maximum value.

// The process continues until no more cells can be colored.

// Return a 2D integer array representing the final state of the grid, where each cell contains its final color.

// Example 1:
// Input: n = 3, m = 3, sources = [[0,0,1],[2,2,2]]
// Output: [[1,1,2],[1,2,2],[2,2,2]]
// Explanation:
// The grid at each time step is as follows:
// ​​​​​​​​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/03/29/g50new.png" />
// At time step 2, cells (0, 2), (1, 1), and (2, 0) are reached by both colors, so they are assigned color 2 as it has the maximum value among them.

// Example 2:
// Input: n = 3, m = 3, sources = [[0,1,3],[1,1,5]]
// Output: [[3,3,3],[5,5,5],[5,5,5]]
// Explanation:
// The grid at each time step is as follows:
// ​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/03/29/g51new.png" />

// Example 3:
// Input: n = 2, m = 2, sources = [[1,1,5]]
// Output: [[5,5],[5,5]]
// Explanation:
// The grid at each time step is as follows:
// ​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/03/29/g52new.png" />
// Since there is only one source, all cells are assigned the same color.

// Constraints:
//     1 <= n, m <= 10^5
//     1 <= n * m <= 10^5
//     1 <= sources.length <= n * m
//     sources[i] = [ri, ci, colori]
//     0 <= ri <= n - 1
//     0 <= ci <= m - 1
//     1 <= colori <= 10^6​​​​​​​
//     All (ri, ci​​​​​​​) in sources are distinct.

import "fmt"
import "slices"

func colorGrid(n, m int, sources [][]int) [][]int {
    dirs := []struct{ x, y int }{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} // 左右上下
    slices.SortFunc(sources, func(a, b []int) int {
            return b[2] - a[2] 
    })
    res := make([][]int, n)
    for i := range res {
        res[i] = make([]int, m)
    }
    for _, p := range sources {
        res[p[0]][p[1]] = p[2] // 初始颜色
    }
    queue := sources
    for len(queue) > 0 {
        p := queue[0]
        queue = queue[1:]
        x, y, c := p[0], p[1], p[2]
        for _, d := range dirs { // 向四个方向扩散
            i, j := x+d.x, y+d.y
            if 0 <= i && i < n && 0 <= j && j < m && res[i][j] == 0 { // (i, j) 未着色
                res[i][j] = c // 着色
                queue = append(queue, []int{i, j, c}) // 继续扩散
            }
        }
    }
    return res
}

func colorGrid1(n int, m int, sources [][]int) [][]int {
    res, dist := make([][]int, n), make([][]int, n)
    for i := 0; i < n; i++ {
        res[i],dist[i] = make([]int, m), make([]int, m)
        for j := 0; j < m; j++ {
            dist[i][j] = -1
        }
    }
    type Cell struct{ r, c int }
    list := make([]Cell, 0, n*m)
    for _, s := range sources {
        r, c, color := s[0], s[1], s[2]
        if dist[r][c] == -1 {
            dist[r][c] = 0
            res[r][c] = color
            list = append(list, Cell{r, c})
        } else {
            if color > res[r][c] {
                res[r][c] = color
            }
        }
    }
    dirs := []struct{ dr, dc int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    for len(list) > 0 {
        curr := list[0]
        list = list[1:]
        r, c := curr.r, curr.c
        for _, d := range dirs {
            nr, nc := r+d.dr, c+d.dc
            if nr >= 0 && nr < n && nc >= 0 && nc < m {
                if dist[nr][nc] == -1 {
                    dist[nr][nc] = dist[r][c] + 1
                    res[nr][nc] = res[r][c]
                    list = append(list, Cell{nr, nc})
                } else if dist[nr][nc] == dist[r][c]+1 {
                    if res[r][c] > res[nr][nc] {
                        res[nr][nc] = res[r][c]
                    }
                }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 3, m = 3, sources = [[0,0,1],[2,2,2]]
    // Output: [[1,1,2],[1,2,2],[2,2,2]]
    // Explanation:
    // The grid at each time step is as follows:
    // ​​​​​​​​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/03/29/g50new.png" />
    // At time step 2, cells (0, 2), (1, 1), and (2, 0) are reached by both colors, so they are assigned color 2 as it has the maximum value among them.
    fmt.Println(colorGrid(3, 3, [][]int{{0,0,1},{2,2,2}})) // [[1,1,2],[1,2,2],[2,2,2]]
    // Example 2:
    // Input: n = 3, m = 3, sources = [[0,1,3],[1,1,5]]
    // Output: [[3,3,3],[5,5,5],[5,5,5]]
    // Explanation:
    // The grid at each time step is as follows:
    // ​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/03/29/g51new.png" />
    fmt.Println(colorGrid(3, 3, [][]int{{0,1,3},{1,1,5}})) // [[3,3,3],[5,5,5],[5,5,5]] 
    // Example 3:
    // Input: n = 2, m = 2, sources = [[1,1,5]]
    // Output: [[5,5],[5,5]]
    // Explanation:
    // The grid at each time step is as follows:
    // ​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/03/29/g52new.png" />
    // Since there is only one source, all cells are assigned the same color.
    fmt.Println(colorGrid(2, 2, [][]int{{1,1,5}})) // [[5,5],[5,5]]

    fmt.Println(colorGrid1(3, 3, [][]int{{0,0,1},{2,2,2}})) // [[1,1,2],[1,2,2],[2,2,2]]
    fmt.Println(colorGrid1(3, 3, [][]int{{0,1,3},{1,1,5}})) // [[3,3,3],[5,5,5],[5,5,5]] 
    fmt.Println(colorGrid1(2, 2, [][]int{{1,1,5}})) // [[5,5],[5,5]]
}
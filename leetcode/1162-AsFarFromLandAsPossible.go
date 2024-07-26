package main

// 1162. As Far from Land as Possible
// Given an n x n grid containing only values 0 and 1, 
// where 0 represents water and 1 represents land, find a water cell such that its distance to the nearest land cell is maximized, 
// and return the distance. If no land or water exists in the grid, return -1.

// The distance used in this problem is the Manhattan distance: 
// the distance between two cells (x0, y0) and (x1, y1) is |x0 - x1| + |y0 - y1|.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/05/03/1336_ex1.JPG" />
// Input: grid = [[1,0,1],[0,0,0],[1,0,1]]
// Output: 2
// Explanation: The cell (1, 1) is as far as possible from all the land with distance 2.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2019/05/03/1336_ex2.JPG" />
// Input: grid = [[1,0,0],[0,0,0],[0,0,0]]
// Output: 4
// Explanation: The cell (2, 2) is as far as possible from all the land with distance 4.

// Constraints:
//     n == grid.length
//     n == grid[i].length
//     1 <= n <= 100
//     grid[i][j] is 0 or 1

import "fmt"

func maxDistance(grid [][]int) int {
    n, queue := len(grid), [][2]int{}
    directions := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {      
                queue = append(queue, [2]int{i, j})
            }  
        }
    }
    if len(queue) == n * n {
        return -1
    }
    bfs := func(grid [][]int, queue [][2]int) int {
        n, res := len(grid), -1
        visited := make(map[[2]int]bool)
        for len(queue) > 0 {
            l := len(queue)
            res++
            for i := 0; i < l; i++ {
                for _, dir := range directions {
                    ii, jj := queue[i][0] + dir[0], queue[i][1] + dir[1]
                    if ii >= 0 && jj >= 0 && ii < n && jj < n &&
                    grid[ii][jj] == 0 && !visited[[2]int{ii, jj}] {
                        visited[[2]int{ii, jj}] = true
                        queue = append(queue, [2]int{ii, jj})
                    }
                }
            }
            queue = queue[l:]
        }
        return res
    }
    return bfs(grid, queue)
}

func maxDistance1(grid [][]int) int {
    queue := [101 * 101][2]int{}
    visited := [101][101]bool{}
    move := []int{-1, 0, 1, 0, -1} // 0:上，1:右，2:下，3:左
    l, r, seas := 0, 0, 0
    n, m := len(grid), len(grid[0])
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if grid[i][j] == 1 {
                visited[i][j] = true
                queue[r][0] = i
                queue[r][1] = j
                r++
            } else {
                visited[i][j] = false
                seas++
            }
        }
    }
    if seas == 0 || seas == n * m {
        return -1
    }
    level := 0
    for l < r {
        level++
        size := r - l
        x, y, nx, ny := 0, 0, 0, 0
        for k := 0; k < size; k++ {
            x = queue[l][0]
            y = queue[l][1]
            l++
            for i := 0; i < 4; i++ {
                nx = x + move[i] // 上、右、下、左
                ny = y + move[i+1]
                if nx >= 0 && nx < n && ny >= 0 && ny < m && !visited[nx][ny] {
                    visited[nx][ny] = true
                    queue[r][0] = nx
                    queue[r][1] = ny
                    r++
                }
            }
        }
    }
    return level - 1
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/05/03/1336_ex1.JPG" />
    // Input: grid = [[1,0,1],[0,0,0],[1,0,1]]
    // Output: 2
    // Explanation: The cell (1, 1) is as far as possible from all the land with distance 2.
    fmt.Println(maxDistance([][]int{{1,0,1},{0,0,0},{1,0,1}})) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2019/05/03/1336_ex2.JPG" />
    // Input: grid = [[1,0,0],[0,0,0],[0,0,0]]
    // Output: 4
    // Explanation: The cell (2, 2) is as far as possible from all the land with distance 4.
    fmt.Println(maxDistance([][]int{{1,0,0},{0,0,0},{0,0,0}})) // 4

    fmt.Println(maxDistance1([][]int{{1,0,1},{0,0,0},{1,0,1}})) // 2
    fmt.Println(maxDistance1([][]int{{1,0,0},{0,0,0},{0,0,0}})) // 4
}
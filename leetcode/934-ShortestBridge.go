package main

// 934. Shortest Bridge
// You are given an n x n binary matrix grid where 1 represents land and 0 represents water.

// An island is a 4-directionally connected group of 1's not connected to any other 1's. 
// There are exactly two islands in grid.

// You may change 0's to 1's to connect the two islands to form one island.
// Return the smallest number of 0's you must flip to connect the two islands.

// Example 1:
// Input: grid = [[0,1],[1,0]]
// Output: 1

// Example 2:
// Input: grid = [[0,1,0],[0,0,0],[0,0,1]]
// Output: 2

// Example 3:
// Input: grid = [[1,1,1,1,1],[1,0,0,0,1],[1,0,1,0,1],[1,0,0,0,1],[1,1,1,1,1]]
// Output: 1

// Constraints:
//     n == grid.length == grid[i].length
//     2 <= n <= 100
//     grid[i][j] is either 0 or 1.
//     There are exactly two islands in grid.

import "fmt"

func shortestBridge(grid [][]int) int {
   land, label := 1,2
    res, r, c := 0, -1, -1 // row, column
    for i, row := range grid {
        for j := range row {
            if grid[i][j] == land {
                r, c = i, j
                break
            }
        }
    }
    queue := [][2]int{} // BFS queue with 2 (x,y) coordinates per cell
    var dfs func(x, y int) // find an island, label it and add it to the BFS queue
    dfs = func(x, y int) {
        if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] != land { // 边界检测
            return
        }
        grid[x][y] = label
        queue = append(queue, [2]int{x, y})
        dfs(x + 1, y) // 右
        dfs(x - 1, y) // 左
        dfs(x, y + 1) // 下
        dfs(x, y - 1) // 上
    }
    dfs(r, c)
    dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} // directions
    // expand countour of the queued island till it touches another one, return
    // the number of expansions as result
    for len(queue) > 0 {
        ln := len(queue)
        for i := 0; i < ln; i++ {
            cur := queue[0] // pop a cell from the head and
            queue = queue[1:]   // truncate the queue's head
            for _, dir := range dirs {
                x := cur[0] + dir[0]
                y := cur[1] + dir[1]
                if x < 0 || y < 0 ||
                    x >= len(grid) || y >= len(grid) || // 'grid' is square!
                    grid[x][y] == label {
                    continue
                }
                if grid[x][y] == land {
                    return res
                }
                queue = append(queue, [2]int{x, y})
                grid[x][y] = label
            }
        }
        res++
    }
    return res
}

func shortestBridge1(grid [][]int) int {
    type pair struct { x, y int }
    m, n := len(grid), len(grid[0])
    delta := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    for i, row := range grid {
        for j, v := range row {
            if v != 1 {
                continue
            }
            island := []pair{}
            grid[i][j] = -1
            q := []pair{{i, j}}
            for len(q) > 0 {
                p := q[0]
                q = q[1:]
                island = append(island, p)
                for _, d := range delta {
                    x, y := p.x + d.x, p.y + d.y
                    if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
                        grid[x][y] = -1
                        q = append(q, pair{x, y})
                    }
                }
            }
            q = island
            depth := 0
            for len(q) > 0 {
                aux := []pair{}
                for _, p := range q {
                    for _, d := range delta {
                        x, y := p.x + d.x, p.y + d.y
                        if x >= 0 && x < m && y >= 0 && y < n {
                            if grid[x][y] == 1 {
                                return depth
                            }
                            if grid[x][y] == 0 {
                                grid[x][y] = -1
                                aux = append(aux, pair{x, y})
                            }
                        }
                    }
                }
                q = aux
                depth++
            }
        }
    }
    return -1
}

func main() {
    // Example 1:
    // Input: grid = [[0,1],[1,0]]
    // Output: 1
    grid1 := [][]int{
        {0, 1},
        {1, 0},
    }
    fmt.Println(shortestBridge(grid1)) // 1
    // Example 2:
    // Input: grid = [[0,1,0],[0,0,0],[0,0,1]]
    // Output: 2
    grid2 := [][]int{
        {0,1,0},
        {0,0,0},
        {0,0,1},
    }
    fmt.Println(shortestBridge(grid2)) // 2
    // Example 3:
    // Input: grid = [[1,1,1,1,1],[1,0,0,0,1],[1,0,1,0,1],[1,0,0,0,1],[1,1,1,1,1]]
    // Output: 1
    grid3 := [][]int{
        {1,1,1,1,1},
        {1,0,0,0,1},
        {1,0,1,0,1},
        {1,0,0,0,1},
        {1,1,1,1,1},
    }
    fmt.Println(shortestBridge(grid3)) // 1

    grid11 := [][]int{
        {0, 1},
        {1, 0},
    }
    fmt.Println(shortestBridge1(grid11)) // 1
    grid12 := [][]int{
        {0,1,0},
        {0,0,0},
        {0,0,1},
    }
    fmt.Println(shortestBridge1(grid12)) // 2
    grid13 := [][]int{
        {1,1,1,1,1},
        {1,0,0,0,1},
        {1,0,1,0,1},
        {1,0,0,0,1},
        {1,1,1,1,1},
    }
    fmt.Println(shortestBridge1(grid13)) // 1
}
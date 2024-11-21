package main

// 2257. Count Unguarded Cells in the Grid
// You are given two integers m and n representing a 0-indexed m x n grid. 
// You are also given two 2D integer arrays guards and walls where guards[i] = [rowi, coli] and walls[j] = [rowj, colj] represent the positions of the ith guard and jth wall respectively.

// A guard can see every cell in the four cardinal directions (north, east, south, or west) starting from their position unless obstructed by a wall or another guard. 
// A cell is guarded if there is at least one guard that can see it.

// Return the number of unoccupied cells that are not guarded.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/03/10/example1drawio2.png" />
// Input: m = 4, n = 6, guards = [[0,0],[1,1],[2,3]], walls = [[0,1],[2,2],[1,4]]
// Output: 7
// Explanation: The guarded and unguarded cells are shown in red and green respectively in the above diagram.
// There are a total of 7 unguarded cells, so we return 7.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/03/10/example2drawio.png" />
// Input: m = 3, n = 3, guards = [[1,1]], walls = [[0,1],[1,0],[2,1],[1,2]]
// Output: 4
// Explanation: The unguarded cells are shown in green in the above diagram.
// There are a total of 4 unguarded cells, so we return 4.

// Constraints:
//     1 <= m, n <= 10^5
//     2 <= m * n <= 10^5
//     1 <= guards.length, walls.length <= 5 * 10^4
//     2 <= guards.length + walls.length <= m * n
//     guards[i].length == walls[j].length == 2
//     0 <= rowi, rowj < m
//     0 <= coli, colj < n
//     All the positions in guards and walls are unique.

import "fmt"

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
    res, grid := m * n, make([][]byte, m)
    for i := 0; i < m; i++ {
        grid[i] = make([]byte, n)
    }
    markGuarded := func (grid [][]byte, m, n, i, j int, dx, dy int) int {
        res, out := 0, false
        for !out {
            i += dx
            j += dy
            if i < 0 || i >=m || j < 0 || j >= n {
                break
            }
            switch grid[i][j] {
                case 'W': // wall
                    fallthrough
                case 'G': // guard
                    out = true
                case 0:
                    grid[i][j] = 'C'
                    res++
            }
        }
        return res
    }
    for i := range guards {
        grid[guards[i][0]][guards[i][1]] = 'G'
        res--
    }
    for i := range walls {
        grid[walls[i][0]][walls[i][1]] = 'W'
        res--
    }
    for i := range grid {
        for j := range grid[0] {
            if grid[i][j] != 'G' {
                continue
            }
            res -= markGuarded(grid, m, n, i, j, -1, 0)
            res -= markGuarded(grid, m, n, i, j, 1, 0)
            res -= markGuarded(grid, m, n, i, j, 0, 1)
            res -= markGuarded(grid, m, n, i, j, 0, -1)
        }
    }
    return res
}

func countUnguarded1(m int, n int, guards [][]int, walls [][]int) int {
    res, grid := 0, make([][]int, m)
    for i := 0; i < m; i++ {
        grid[i] = make([]int, n)
    }
    for _, point := range guards {
        grid[point[0]][point[1]] = 1
    }
    for _, point := range walls {
        grid[point[0]][point[1]] = -1
    }
    for row := 0; row < m; row++ {
        for col := 1; col < n; col++ {
            if grid[row][col] == 0 && (grid[row][col - 1] == 1 || grid[row][col - 1] == 2) {
                grid[row][col] = 2
            }
        }
    }
    for row := 0; row < m; row++ {
        for col := n - 2; col >= 0; col-- {
            if grid[row][col] == 0 && (grid[row][col + 1] == 1 || grid[row][col + 1] == 2) {
                grid[row][col] = 2
            }
        }
    }
    for col := 0; col < n; col++ {
        for row := 1; row < m; row++ {
            if grid[row][col] != -1 && grid[row - 1][col] == 1 {
                grid[row][col] = 1
            }
        }
    }
    for col := 0; col < n; col++ {
        for row := m - 2; row >= 0; row-- {
            if grid[row][col] != -1 && grid[row + 1][col] == 1 {
                grid[row][col] = 1
            }
        }
    }
    for _, row := range grid {
        for _, v := range row {
            if v == 0 {
                res++
            }
        }
    }
    return res
}

func countUnguarded2(m int, n int, guards [][]int, walls [][]int) int {
    res := 0
    arr := make([]byte, m * n)
    for _, w := range walls { // 墙的位置
        arr[w[0] * n + w[1]] = 1
    }
    for _, g := range guards { // 守卫的位置
        arr[g[0]*n + g[1]] = 1
    }
    for _, g := range guards { // 守卫能观察到的位置
        r, c := g[0], g[1]
        for r := r - 1; r >= 0 && arr[r * n + c] != 1; r-- { arr[r * n + c] = 2 }
        for r := r + 1; r <  m && arr[r * n + c] != 1; r++ { arr[r * n + c] = 2 }
        for c := c - 1; c >= 0 && arr[r * n + c] != 1; c-- { arr[r * n + c] = 2 }
        for c := c + 1; c <  n && arr[r * n + c] != 1; c++ { arr[r * n + c] = 2 }
    }
    for i := 0; i < m * n; i++ {
        if arr[i] == 0 {
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/03/10/example1drawio2.png" />
    // Input: m = 4, n = 6, guards = [[0,0],[1,1],[2,3]], walls = [[0,1],[2,2],[1,4]]
    // Output: 7
    // Explanation: The guarded and unguarded cells are shown in red and green respectively in the above diagram.
    // There are a total of 7 unguarded cells, so we return 7.
    fmt.Println(countUnguarded(4,6,[][]int{{0,0},{1,1},{2,3}},[][]int{{0,1},{2,2},{1,4}})) // 7
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/03/10/example2drawio.png" />
    // Input: m = 3, n = 3, guards = [[1,1]], walls = [[0,1],[1,0],[2,1],[1,2]]
    // Output: 4
    // Explanation: The unguarded cells are shown in green in the above diagram.
    // There are a total of 4 unguarded cells, so we return 4.
    fmt.Println(countUnguarded(3,3,[][]int{{1,1}},[][]int{{0,1},{1,0},{2,1},{1,2}})) // 4

    fmt.Println(countUnguarded1(4,6,[][]int{{0,0},{1,1},{2,3}},[][]int{{0,1},{2,2},{1,4}})) // 7
    fmt.Println(countUnguarded1(3,3,[][]int{{1,1}},[][]int{{0,1},{1,0},{2,1},{1,2}})) // 4

    fmt.Println(countUnguarded2(4,6,[][]int{{0,0},{1,1},{2,3}},[][]int{{0,1},{2,2},{1,4}})) // 7
    fmt.Println(countUnguarded2(3,3,[][]int{{1,1}},[][]int{{0,1},{1,0},{2,1},{1,2}})) // 4
}
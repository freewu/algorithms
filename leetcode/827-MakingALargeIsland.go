package main

// 827. Making A Large Island
// You are given an n x n binary matrix grid. You are allowed to change at most one 0 to be 1.
// Return the size of the largest island in grid after applying this operation.
// An island is a 4-directionally connected group of 1s.

// Example 1:
// Input: grid = [[1,0],[0,1]]
// Output: 3
// Explanation: Change one 0 to 1 and connect two 1s, then we get an island with area = 3.

// Example 2:
// Input: grid = [[1,1],[1,0]]
// Output: 4
// Explanation: Change the 0 to 1 and make the island bigger, only one island with area = 4.

// Example 3:
// Input: grid = [[1,1],[1,1]]
// Output: 4
// Explanation: Can't change any 0 to 1, only one island with area = 4.

// Constraints:
//     n == grid.length
//     n == grid[i].length
//     1 <= n <= 500
//     grid[i][j] is either 0 or 1.

import "fmt"

func largestIsland(grid [][]int) int {
    color, zeroFlag, sizeMap, res := 2, false, make(map[int]int), -1 << 31
    m, n := len(grid), len(grid[0])
    islandPainter := func (row, col, color int) int {
        stack, visited, size := [][]int{{row, col}}, make(map[[2]int]bool), 0
        rows, cols := []int{0, 1, 0, -1}, []int{1, 0, -1, 0}
        for len(stack) > 0 {
            node := stack[len(stack)-1] // pop
            stack = stack[:len(stack)-1]
            r, c := node[0], node[1]
            if !visited[[2]int{r, c}] {
                visited[[2]int{r, c}] = true
                grid[r][c] = color
                size++
                for i := 0; i < 4; i++ {
                    newRow, newCol := r + rows[i], c + cols[i]
                    if newRow < 0 || newCol < 0 || newRow > m - 1 || newCol > n - 1 || grid[newRow][newCol] != 1 {
                        continue
                    }
                    stack = append(stack, []int{newRow, newCol})
                }
            }
        }
        return size
    }
    for r, row := range grid {
        for c, num := range row {
            if num == 1 {
                size := islandPainter(r, c, color)
                sizeMap[color] = size
                color++
            }
            if num == 0 {
                zeroFlag = true
            }
        }
    }
    if !zeroFlag { // 全部为 1
        return len(grid) * len(grid[0])
    }
    islandMaker := func (row, col int) int {
        rows, cols := []int{0, 1, 0, -1}, []int{1, 0, -1, 0}
        size, visited := 1, make(map[int]bool)
        for i := 0; i < 4; i++ {
            newRow, newCol := row + rows[i], col + cols[i]
            if newRow < 0 || newCol < 0 || newRow > m - 1 || newCol > n - 1 || grid[newRow][newCol] == 0 || visited[grid[newRow][newCol]] {
                continue
            }
            size += sizeMap[grid[newRow][newCol]]
            visited[grid[newRow][newCol]] = true
        }
        return size
    }
    for r, row := range grid {
        for c, num := range row {
            if num == 0 {
                posSize := islandMaker(r, c)
                if posSize > res {
                    res = posSize
                }
            }
        }
    }
    return res
}

func largestIsland1(grid [][]int) int {
    n, m, id := len(grid), len(grid[0]), 2
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var infect func(grid [][]int, i, j, v, n, m int)
    infect = func(grid [][]int, i, j, v, n, m int) {
        if i < 0 || i == n || j < 0 || j == m || grid[i][j] != 1 { return } // 边界检测
        grid[i][j] = v
        infect(grid, i-1, j, v, n, m)
        infect(grid, i+1, j, v, n, m)
        infect(grid, i, j-1, v, n, m)
        infect(grid, i, j+1, v, n, m)
    }
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if grid[i][j] == 1 {
                infect(grid, i, j, id, n, m)
                id++
            }
        }
    }
    res, sizes := 0, make([]int, id)
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if grid[i][j] != 0 {
                sizes[grid[i][j]]++
                res = max(res, sizes[grid[i][j]])
            }
        }
    }
    visited := make([]bool, id)
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if grid[i][j] == 0 {
                up, down, left, right := 0, 0, 0, 0
                if i-1 >= 0 { up = grid[i-1][j] }
                if i+1 < n  { down = grid[i+1][j] }
                if j-1 >= 0 { left = grid[i][j-1] }
                if j+1 < m  { right = grid[i][j+1] }
                merge := 1 + sizes[up]
                visited[up] = true
                if !visited[down] {
                    merge += sizes[down]
                    visited[down] = true
                }
                if !visited[left] {
                    merge += sizes[left]
                    visited[left] = true
                }
                if !visited[right] {
                    merge += sizes[right]
                    visited[right] = true
                }
                res = max(res, merge)
                visited[up] = false
                visited[down] = false
                visited[left] = false
                visited[right] = false
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: grid = [[1,0],[0,1]]
    // Output: 3
    // Explanation: Change one 0 to 1 and connect two 1s, then we get an island with area = 3.
    fmt.Println(largestIsland([][]int{{1,0},{0,1}})) // 3
    // Example 2:
    // Input: grid = [[1,1],[1,0]]
    // Output: 4
    // Explanation: Change the 0 to 1 and make the island bigger, only one island with area = 4.
    fmt.Println(largestIsland([][]int{{1,1},{0,1}})) // 4
    // Example 3:
    // Input: grid = [[1,1],[1,1]]
    // Output: 4
    // Explanation: Can't change any 0 to 1, only one island with area = 4.
    fmt.Println(largestIsland([][]int{{1,1},{1,1}})) // 4

    fmt.Println(largestIsland1([][]int{{1,0},{0,1}})) // 3
    fmt.Println(largestIsland1([][]int{{1,1},{0,1}})) // 4
    fmt.Println(largestIsland1([][]int{{1,1},{1,1}})) // 4
}
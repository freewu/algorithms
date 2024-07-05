package main

// 1091. Shortest Path in Binary Matrix
// Given an n x n binary matrix grid, return the length of the shortest clear path in the matrix. 
// If there is no clear path, return -1.

// A clear path in a binary matrix is a path from the top-left cell (i.e., (0, 0)) to the bottom-right cell (i.e., (n - 1, n - 1)) such that:
//     All the visited cells of the path are 0.
//     All the adjacent cells of the path are 8-directionally connected (i.e., they are different and they share an edge or a corner).

// The length of a clear path is the number of visited cells of this path.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/02/18/example1_1.png" />
// Input: grid = [[0,1],[1,0]]
// Output: 2

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/02/18/example2_1.png" />
// Input: grid = [[0,0,0],[1,1,0],[1,1,0]]
// Output: 4

// Example 3:
// Input: grid = [[1,0,0],[1,1,0],[1,1,0]]
// Output: -1

// Constraints:
//     n == grid.length
//     n == grid[i].length
//     1 <= n <= 100
//     grid[i][j] is 0 or 1

import "fmt"

// bfs
func shortestPathBinaryMatrix(grid [][]int) int {
    if grid[0][0] == 1 { // 起始点不为 0
        return -1
    }
    n := len(grid)
    if n == 1 && grid[0][0] == 0 {
        return 1
    }
    directions := [][]int{ {-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}, }
    type Node struct {
        col, row, distance int
    }
    var queue []Node
    queue = append(queue, Node{0, 0, 1})
    for len(queue) != 0 {
        node := queue[0]
        queue = queue[1:]
        grid[node.row][node.col] = 1
        for _, dir := range directions {
            row := node.row + dir[0]
            col := node.col + dir[1]
            if row < 0 || row >= n || col < 0 || col >= n { continue } // 越界
            if grid[row][col] != 0 { continue } // 障碍物
            if row == n - 1 && col == n - 1 { // 到达目标位置
                return node.distance + 1
            }
            queue = append(queue, Node{col, row, node.distance + 1})
            grid[row][col] = 1
        }
    }
    return -1
}

func shortestPathBinaryMatrix1(grid [][]int) int {
    res, n := 0, len(grid)
    if n == 0 || grid[0][0] != 0 {
        return -1
    }
    queue, visited := []int{}, make([]bool, n*n)
    queue = append(queue, 0)
    visited[0] = true
    directions := [][]int{{1, 0}, {1, 1}, {1, -1}, {0, 1}, {0, -1}, {-1, 0}, {-1, 1}, {-1, -1}}
    for len(queue) > 0 {
        res++
        newQueue := []int{}
        for _, cur := range queue {
            x, y := cur / n, cur % n
            if x == n - 1 && y == n - 1 {
                return res
            }
            for _, dir := range directions {
                r, c := x + dir[0], y + dir[1]
                next := r * n + c
                if r < 0 || r == n || c < 0 || c == n || grid[r][c] != 0 || visited[next] {
                    continue
                }
                visited[next] = true
                newQueue = append(newQueue, next)
            }
        }
        queue = newQueue
    }
    return -1
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/02/18/example1_1.png" />
    // Input: grid = [[0,1],[1,0]]
    // Output: 2
    grid1 := [][]int{
        {0, 1},
        {1, 0},
    }
    fmt.Println(shortestPathBinaryMatrix(grid1)) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/02/18/example2_1.png" />
    // Input: grid = [[0,0,0],[1,1,0],[1,1,0]]
    // Output: 4
    grid2 := [][]int{
        {0,0,0},
        {1,1,0},
        {1,1,0},
    }
    fmt.Println(shortestPathBinaryMatrix(grid2)) // 4
    // Example 3:
    // Input: grid = [[1,0,0],[1,1,0],[1,1,0]]
    // Output: -1
    grid3:= [][]int{
        {1,0,0},
        {1,1,0},
        {1,1,0},
    }
    fmt.Println(shortestPathBinaryMatrix(grid3)) // -1

    grid11 := [][]int{
        {0, 1},
        {1, 0},
    }
    fmt.Println(shortestPathBinaryMatrix1(grid11)) // 2
    grid12 := [][]int{
        {0,0,0},
        {1,1,0},
        {1,1,0},
    }
    fmt.Println(shortestPathBinaryMatrix1(grid12)) // 4
    grid13:= [][]int{
        {1,0,0},
        {1,1,0},
        {1,1,0},
    }
    fmt.Println(shortestPathBinaryMatrix1(grid13)) // -1
}
package main

// 1765. Map of Highest Peak
// You are given an integer matrix isWater of size m x n that represents a map of land and water cells.
//     1. If isWater[i][j] == 0, cell (i, j) is a land cell.
//     2. If isWater[i][j] == 1, cell (i, j) is a water cell.

// You must assign each cell a height in a way that follows these rules:
//     1. The height of each cell must be non-negative.
//     2. If the cell is a water cell, its height must be 0.
//     3. Any two adjacent cells must have an absolute height difference of at most 1. 
//        A cell is adjacent to another cell if the former is directly north, east, south, or west of the latter 
//        (i.e., their sides are touching).

// Find an assignment of heights such that the maximum height in the matrix is maximized.

// Return an integer matrix height of size m x n where height[i][j] is cell (i, j)'s height. 
// If there are multiple solutions, return any of them.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/01/10/screenshot-2021-01-11-at-82045-am.png" /> 
// Input: isWater = [[0,1],[0,0]]
// Output: [[1,0],[2,1]]
// Explanation: The image shows the assigned heights of each cell.
// The blue cell is the water cell, and the green cells are the land cells.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/01/10/screenshot-2021-01-11-at-82050-am.png" /> 
// Input: isWater = [[0,0,1],[1,0,0],[0,0,0]]
// Output: [[1,1,0],[0,1,1],[1,2,2]]
// Explanation: A height of 2 is the maximum possible height of any assignment.
// Any height assignment that has a maximum height of 2 while still meeting the rules will also be accepted.

// Constraints:
//     m == isWater.length
//     n == isWater[i].length
//     1 <= m, n <= 1000
//     isWater[i][j] is 0 or 1.
//     There is at least one water cell.

import "fmt"

// bfs
func highestPeak(isWater [][]int) [][]int {
    directions := [][]int{{0,1},{0,-1},{1,0},{-1,0}}
    n, m, queue := len(isWater), len(isWater[0]), [][]int{}
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if isWater[i][j] == 1 {
                isWater[i][j] = 0
                queue = append(queue, []int{i, j})
            } else {
                isWater[i][j] = -1
            }
        }
    }
    for len(queue) != 0 {
        size := len(queue)
        for i := 0; i < size; i++ {
            cell := queue[0]
            queue = queue[1:] // pop
            for _, dir := range directions {
                nr, nc := cell[0] + dir[0], cell[1] + dir[1]
                if nr >= 0 && nc >= 0 && nr < n && nc < m && isWater[nr][nc] == -1 { // 边界检测
                    isWater[nr][nc] = isWater[cell[0]][cell[1]] + 1   
                    queue = append(queue, []int{ nr, nc })
                }
            }
        }
    }
    return isWater
}

var queue [1024*1024]int
var visited [1024][1024]bool
// bfs
func highestPeak1(isWater [][]int) [][]int {
    rows, cols, queue:= len(isWater), len(isWater[0]), queue[:0]
    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            visited[r][c] = false
            if isWater[r][c] == 1 { 
                queue, visited[r][c] = append(queue, r * 1024 + c), true 
            }
            isWater[r][c] ^= 1
        }
    }
    enq := func(r, c int) {
        if r < 0 || r == rows || c < 0 || c == cols || visited[r][c] { return }
        queue, visited[r][c] = append(queue, r * 1024 + c), true
    }
    for i := 0; len(queue) > 0; i++ {
        for _, v := range queue {
            r, c := v / 1024, v % 1024
            isWater[r][c], queue = i, queue[1:]
            enq(r - 1, c)
            enq(r + 1, c)
            enq(r, c - 1)
            enq(r, c + 1)
        }
    }
    return isWater
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/01/10/screenshot-2021-01-11-at-82045-am.png" /> 
    // Input: isWater = [[0,1],[0,0]]
    // Output: [[1,0],[2,1]]
    // Explanation: The image shows the assigned heights of each cell.
    // The blue cell is the water cell, and the green cells are the land cells.
    fmt.Println(highestPeak([][]int{{0,1},{0,0}})) // [[1,0],[2,1]]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/01/10/screenshot-2021-01-11-at-82050-am.png" /> 
    // Input: isWater = [[0,0,1],[1,0,0],[0,0,0]]
    // Output: [[1,1,0],[0,1,1],[1,2,2]]
    // Explanation: A height of 2 is the maximum possible height of any assignment.
    // Any height assignment that has a maximum height of 2 while still meeting the rules will also be accepted.
    fmt.Println(highestPeak([][]int{{0,0,1},{1,0,0},{0,0,0}})) // [[1,1,0],[0,1,1],[1,2,2]]

    fmt.Println(highestPeak1([][]int{{0,1},{0,0}})) // [[1,0],[2,1]]
    fmt.Println(highestPeak1([][]int{{0,0,1},{1,0,0},{0,0,0}})) // [[1,1,0],[0,1,1],[1,2,2]]
}
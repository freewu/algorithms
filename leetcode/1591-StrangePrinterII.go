package main

// 1591. Strange Printer II
// There is a strange printer with the following two special requirements:
//     1. On each turn, the printer will print a solid rectangular pattern of a single color on the grid. 
//        This will cover up the existing colors in the rectangle.
//     2. Once the printer has used a color for the above operation, the same color cannot be used again.

// You are given a m x n matrix targetGrid, where targetGrid[row][col] is the color in the position (row, col) of the grid.

// Return true if it is possible to print the matrix targetGrid, otherwise, return false.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/12/23/print1.jpg" />
// Input: targetGrid = [[1,1,1,1],[1,2,2,1],[1,2,2,1],[1,1,1,1]]
// Output: true

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/12/23/print2.jpg" />
// Input: targetGrid = [[1,1,1,1],[1,1,3,3],[1,1,3,4],[5,5,1,4]]
// Output: true

// Example 3:
// Input: targetGrid = [[1,2,1],[2,1,2],[1,2,1]]
// Output: false
// Explanation: It is impossible to form targetGrid because it is not allowed to print the same color in different turns.

// Constraints:
//     m == targetGrid.length
//     n == targetGrid[i].length
//     1 <= m, n <= 60
//     1 <= targetGrid[row][col] <= 60

import "fmt"

func isPrintable(targetGrid [][]int) bool {
    colors := make(map[int][]int) // Let's find the top-left and bottom-right coordinates for each color
    m, n, inf  := len(targetGrid), len(targetGrid[0]), 1 << 31
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            c := targetGrid[i][j]
            if _, ok := colors[c]; !ok {
                colors[c] = []int{ inf, inf, -inf, -inf }
            }
            colors[c][0] = min(colors[c][0], i)
            colors[c][1] = min(colors[c][1], j)
            colors[c][2] = max(colors[c][2], i)
            colors[c][3] = max(colors[c][3], j)
        }
    }
    indegrees, overlaps:= make(map[int]int),make(map[int]map[int]struct{}) // Now we're going to build a graph on colors based on overlaps.
    for c, coords := range colors {
        overlaps[c] = make(map[int]struct{})
        for i := coords[0]; i <= coords[2]; i++ {
            for j := coords[1]; j <= coords[3]; j++ {
                c2 := targetGrid[i][j]
                if c != c2 {
                    if _, ok := overlaps[c][c2]; !ok {
                        indegrees[c2]++
                        overlaps[c][c2] = struct{}{}
                    }
                }
            }
        }
    }
    // We'll do a topological sorting based on the overlaps. We start by processing
    // all nodes that aren't overlapping with anything else. 
    handled, queue := 0, []int{}
    for c := range colors {
        if indegrees[c] == 0 {
            queue = append(queue, c)
        }
    }
    for len(queue) > 0 {
        newQueue := []int{}
        for _, c := range queue {
            handled++
            for c2 := range overlaps[c] {
                indegrees[c2]--
                if indegrees[c2] == 0 {
                    newQueue = append(newQueue, c2)
                }
            }
        }
        queue = newQueue
    }
    return handled == len(colors) // Return whether we can paint all colors
}

func isPrintable1(targetGrid [][]int) bool {
    pos := make(map[int][]int)
    m, n := len(targetGrid), len(targetGrid[0])
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            color := targetGrid[i][j]
            _, ok := pos[color]
            if (!ok) {
                pos[color] = []int{ m - 1, n - 1, 0, 0}
            }
            coord := pos[color]
            if (i < coord[0]) { coord[0] = i }
            if (j < coord[1]) { coord[1] = j }
            if (i > coord[2]) { coord[2] = i }
            if (j > coord[3]) { coord[3] = j }
        }
    }
    colors := make(map[int]bool)
    for color := range pos {
        colors[color] = true
    }
    erase := func(targetGrid [][]int, coord []int, color int) bool {
        for i := coord[0]; i <= coord[2]; i++ {
            for j := coord[1]; j <= coord[3]; j++ {
                if (targetGrid[i][j] != 0 && targetGrid[i][j] != color) { return false }
            }
        }
        for i := coord[0]; i <= coord[2]; i++ {
            for j := coord[1]; j <= coord[3]; j++ {
                targetGrid[i][j] = 0
            }
        }
        return true
    }
    for len(colors) > 0 {
        next := make(map[int]bool)
        for color := range colors {
            if (!erase(targetGrid, pos[color], color)) {
                next[color] = true
            }
        }
        if len(colors) == len(next) {
            return false
        }
        colors = next
    }
    return true
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/12/23/print1.jpg" />
    // Input: targetGrid = [[1,1,1,1],[1,2,2,1],[1,2,2,1],[1,1,1,1]]
    // Output: true
    fmt.Println(isPrintable([][]int{{1,1,1,1},{1,2,2,1},{1,2,2,1},{1,1,1,1}})) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/12/23/print2.jpg" />
    // Input: targetGrid = [[1,1,1,1],[1,1,3,3],[1,1,3,4],[5,5,1,4]]
    // Output: true
    fmt.Println(isPrintable([][]int{{1,1,1,1},{1,1,3,3},{1,1,3,4},{5,5,1,4}})) // true
    // Example 3:
    // Input: targetGrid = [[1,2,1],[2,1,2],[1,2,1]]
    // Output: false
    // Explanation: It is impossible to form targetGrid because it is not allowed to print the same color in different turns.
    fmt.Println(isPrintable([][]int{{1,2,1},{2,1,2},{1,2,1}})) // false

    fmt.Println(isPrintable1([][]int{{1,1,1,1},{1,2,2,1},{1,2,2,1},{1,1,1,1}})) // true
    fmt.Println(isPrintable1([][]int{{1,1,1,1},{1,1,3,3},{1,1,3,4},{5,5,1,4}})) // true
    fmt.Println(isPrintable1([][]int{{1,2,1},{2,1,2},{1,2,1}})) // false
}
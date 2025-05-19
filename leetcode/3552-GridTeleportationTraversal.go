package main

// 3552. Grid Teleportation Traversal
// You are given a 2D character grid matrix of size m x n, 
// represented as an array of strings, 
// where matrix[i][j] represents the cell at the intersection of the ith row and jth column. 
// Each cell is one of the following:
//    1.  '.' representing an empty cell.
//    2.  '#' representing an obstacle.
//    3.  An uppercase letter ('A'-'Z') representing a teleportation portal.

// You start at the top-left cell (0, 0), and your goal is to reach the bottom-right cell (m - 1, n - 1). 
// You can move from the current cell to any adjacent cell (up, down, left, right) as long as the destination cell is within the grid bounds and is not an obstacle.

// If you step on a cell containing a portal letter and you haven't used that portal letter before, you may instantly teleport to any other cell in the grid with the same letter. 
// This teleportation does not count as a move, but each portal letter can be used at most once during your journey.

// Return the minimum number of moves required to reach the bottom-right cell.
// If it is not possible to reach the destination, return -1.

// Example 1:
// Input: matrix = ["A..",".A.","..."]
// Output: 2
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/03/15/example04140.png" />
// Before the first move, teleport from (0, 0) to (1, 1).
// In the first move, move from (1, 1) to (1, 2).
// In the second move, move from (1, 2) to (2, 2).

// Example 2:
// Input: matrix = [".#...",".#.#.",".#.#.","...#."]
// Output: 13
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/03/15/ezgifcom-animated-gif-maker.gif" />

// Constraints:
//     1 <= m == matrix.length <= 10^3
//     1 <= n == matrix[i].length <= 10^3
//     matrix[i][j] is either '#', '.', or an uppercase English letter.
//     matrix[0][0] is not an obstacle.

import "fmt"

func minMoves(matrix []string) int {
    m, n := len(matrix), len(matrix[0])
    if matrix[m-1][n-1] == '#' {
        return -1
    }
    isUpper := func(c rune) bool { return c >= 'A' && c <= 'Z' }
    type Pair struct{ x, y int }
    pos := ['Z' + 1][]Pair{}
    for i, row := range matrix {
        for j, c := range row {
            if isUpper(c) {
                pos[c] = append(pos[c], Pair{i, j})
            }
        }
    }
    dirs := []Pair{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
    dis := make([][]int, m)
    for i := range dis {
        dis[i] = make([]int, n)
        for j := range dis[i] {
            dis[i][j] = 1 << 31
        }
    }
    dis[0][0] = 0
    // 两个 slice 头对头，模拟 deque
    q0, q1 := []Pair{{}}, []Pair{}
    for len(q0) > 0 || len(q1) > 0 {
        // 弹出队首
        var p Pair
        if len(q0) > 0 {
            p, q0 = q0[len(q0) - 1], q0[:len(q0) - 1]
        } else {
            p, q1 = q1[0], q1[1:]
        }
        d := dis[p.x][p.y]
        if p.x == m-1 && p.y == n-1 {
            return d
        }
        if c := matrix[p.x][p.y]; c != '.' {
            // 使用所有传送门
            for _, q := range pos[c] {
                x, y := q.x, q.y
                if d < dis[x][y] {
                    dis[x][y] = d
                    q0 = append(q0, Pair{x, y}) // 加到队首
                }
            }
            pos[c] = nil // 避免重复使用传送门
        }
        // 下面代码和普通 BFS 是一样的
        for _, dir := range dirs {
            x, y := p.x+dir.x, p.y+dir.y
            if 0 <= x && x < m && 0 <= y && y < n && matrix[x][y] != '#' && d+1 < dis[x][y] {
                dis[x][y] = d + 1
                q1 = append(q1, Pair{x, y}) // 加到队尾
            }
        }
    }
    return -1
}

func main() {
    // Example 1:
    // Input: matrix = ["A..",".A.","..."]
    // Output: 2
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/03/15/example04140.png" />
    // Before the first move, teleport from (0, 0) to (1, 1).
    // In the first move, move from (1, 1) to (1, 2).
    // In the second move, move from (1, 2) to (2, 2).
    fmt.Println(minMoves([]string{"A..",".A.","..."})) // 2
    // Example 2:
    // Input: matrix = [".#...",".#.#.",".#.#.","...#."]
    // Output: 13
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/03/15/ezgifcom-animated-gif-maker.gif" />
    fmt.Println(minMoves([]string{".#...",".#.#.",".#.#.","...#."})) // 13
}
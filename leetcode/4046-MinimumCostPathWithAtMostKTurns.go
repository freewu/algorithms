package main

// 4046. Minimum Cost Path With At Most K Turns
// You are given a 2D integer array grid of size m x n, where grid[i][j] represents the cost of visiting cell (i, j), and an integer k.

// You start at the top-left cell (0, 0) and want to reach the bottom-right cell (m - 1, n - 1).

// From each cell, you may move one step in any of the four directions: up, down, left, or right.

// The cost of a path is the sum of the values of all visited cells, including the starting and ending cells. 
// If a cell is visited more than once, its value is included each time it is visited.

// Return the minimum possible path cost to reach (m - 1, n - 1) using at most k turns. 
// If no such path exists, return -1.

// A turn occurs when the direction changes between two consecutive moves. 
// For example, moving right and then down counts as one turn, while moving right and then right does not.

// Example 1:
// Input: grid = [[2,7,3],[1,4,5]], k = 1
// Output: 12
// Explanation:
// An optimal path is (0, 0) → (1, 0) → (1, 1) → (1, 2). The moves are down, right, right.
// The direction changes from down to right once, so the path uses exactly k = 1 turn.
// The total path cost is 2 + 1 + 4 + 5 = 12.

// Example 2:
// Input: grid = [[4,1,9],[3,2,5],[4,8,6]], k = 2
// Output: 20
// Explanation:​​​​​​​
// An optimal path is (0, 0) → (1, 0) → (1, 1) → (1, 2) → (2, 2). The moves are down, right, right, down.
// The direction changes from down to right and from right to down, so the path uses exactly k = 2 turns.
// The total path cost is 4 + 3 + 2 + 5 + 6 = 20.

// Example 3:
// Input: grid = [[1,9],[3,4]], k = 0
// Output: -1
// Explanation:
// It is impossible to reach (1, 1) using k = 0 turns. Thus, the answer is -1.

// Constraints:
//     1 <= m == grid.length <= 75
//     1 <= n == grid[i].length <= 75
//     0 <= grid[i][j] <= 1000
//     0 <= k < min(m, n)

import "fmt"

func minCost(grid [][]int, k int) int {
    dirs := []struct{ x, y int }{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} // 左右上下
    m, n := len(grid), len(grid[0])
    memo := make([][][][4]int, k + 1)
    for i := range memo {
        memo[i] = make([][][4]int, m)
        for j := range memo[i] {
            memo[i][j] = make([][4]int, n)
            for p := range memo[i][j] {
                for q := range memo[i][j][p] {
                    memo[i][j][p][q] = -1 // -1 表示该状态没有计算过
                }
            }
        }
    }
    var dfs func(int, int, int, int) int
    dfs = func(k, i, j, index int) int {
        if i == 0 && j == 0 {
            return grid[0][0]
        }
        p := &memo[k][i][j][index]
        if *p != -1 { // 之前计算过
            return *p
        }
        res := 1 << 61
        for newIndex, dir := range dirs {
            x, y := i+dir.x, j+dir.y
            if 0 <= x && x < m && 0 <= y && y < n {
                newK := k
                if newIndex != index {
                    if k == 0 {
                        continue
                    }
                    newK--
                }
                res = min(res, dfs(newK, x, y, newIndex))
            }
        }
        res += grid[i][j]
        *p = res // 记忆化
        return res
    }
    res := min(dfs(k, m-1, n-1, 0), dfs(k, m-1, n-1, 2))
    if res < 1 << 61 {
        return res
    }
    return -1
}

func main() {
    // Example 1:
    // Input: grid = [[2,7,3],[1,4,5]], k = 1
    // Output: 12
    // Explanation:
    // An optimal path is (0, 0) → (1, 0) → (1, 1) → (1, 2). The moves are down, right, right.
    // The direction changes from down to right once, so the path uses exactly k = 1 turn.
    // The total path cost is 2 + 1 + 4 + 5 = 12.
    fmt.Println(minCost([][]int{{2,7,3},{1,4,5}}, 1)) // 12
    // Example 2:
    // Input: grid = [[4,1,9],[3,2,5],[4,8,6]], k = 2
    // Output: 20
    // Explanation:​​​​​​​
    // An optimal path is (0, 0) → (1, 0) → (1, 1) → (1, 2) → (2, 2). The moves are down, right, right, down.
    // The direction changes from down to right and from right to down, so the path uses exactly k = 2 turns.
    // The total path cost is 4 + 3 + 2 + 5 + 6 = 20.
    fmt.Println(minCost([][]int{{4,1,9},{3,2,5},{4,8,6}}, 2)) // 20
    // Example 3:
    // Input: grid = [[1,9],[3,4]], k = 0
    // Output: -1
    // Explanation:
    // It is impossible to reach (1, 1) using k = 0 turns. Thus, the answer is -1.
    fmt.Println(minCost([][]int{{1,9},{3,4}}, 0)) // -1
}
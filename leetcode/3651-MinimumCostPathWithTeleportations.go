package main

// 3651. Minimum Cost Path with Teleportations
// You are given a m x n 2D integer array grid and an integer k. 
// You start at the top-left cell (0, 0) and your goal is to reach the bottom‐right cell (m - 1, n - 1).

// There are two types of moves available:
//     1. Normal move: You can move right or down from your current cell (i, j), 
//        i.e. you can move to (i, j + 1) (right) or (i + 1, j) (down). 
//        The cost is the value of the destination cell.
//     2. Teleportation: You can teleport from any cell (i, j), 
//        to any cell (x, y) such that grid[x][y] <= grid[i][j]; the cost of this move is 0. 
//        You may teleport at most k times.

// Return the minimum total cost to reach cell (m - 1, n - 1) from (0, 0).

// Example 1:
// Input: grid = [[1,3,3],[2,5,4],[4,3,5]], k = 2
// Output: 7
// Explanation:
// Initially we are at (0, 0) and cost is 0.
// Current Position	Move	New Position	Total Cost
// (0, 0)	Move Down	(1, 0)	0 + 2 = 2
// (1, 0)	Move Right	(1, 1)	2 + 5 = 7
// (1, 1)	Teleport to (2, 2)	(2, 2)	7 + 0 = 7
// The minimum cost to reach bottom-right cell is 7.

// Example 2:
// Input: grid = [[1,2],[2,3],[3,4]], k = 1
// Output: 9
// Explanation:
// Initially we are at (0, 0) and cost is 0.
// Current Position	Move	New Position	Total Cost
// (0, 0)	Move Down	(1, 0)	0 + 2 = 2
// (1, 0)	Move Right	(1, 1)	2 + 3 = 5
// (1, 1)	Move Down	(2, 1)	5 + 4 = 9
// The minimum cost to reach bottom-right cell is 9.

// Constraints:
//     2 <= m, n <= 80
//     m == grid.length
//     n == grid[i].length
//     0 <= grid[i][j] <= 10^4
//     0 <= k <= 10

import "fmt"
import "slices"

func minCost(grid [][]int, k int) int {
    m, n := len(grid), len(grid[0])
    if k > 0 && grid[0][0] > grid[m-1][n-1] {
        return 0
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    mx := 0
    for _, row := range grid {
        mx = max(mx, slices.Max(row))
    }
    suff := make([]int, mx + 2)
    for i := range suff {
        suff[i] = 1 << 61
    }
    minF := make([]int, mx + 1)
    f := make([]int, n+1)
    for range k + 1 {
        for i := range minF {
            minF[i] = 1 << 61
        }
        //最小路径和（空间优化写法）
        for i := range f {
            f[i] = 1 << 61 / 2
        }
        f[1] = -grid[0][0] // 起点的成本不算
        for _, row := range grid {
            for j, x := range row {
                f[j+1] = min(f[j]+x, f[j+1]+x, suff[x])
                minF[x] = min(minF[x], f[j+1])
            }
        }
        done := true
        // 计算 minF 的后缀最小值
        for i := mx; i >= 0; i-- {
            mn := min(suff[i+1], minF[i])
            if mn < suff[i] {
                suff[i] = mn
                done = false
            }
        }
        if done {
            // 收敛了：传送不改变 sufMinF，那么无论再传送多少次都不会改变 sufMinF
            break
        }
    }
    return f[n]
}

func main() {
    // Example 1:
    // Input: grid = [[1,3,3],[2,5,4],[4,3,5]], k = 2
    // Output: 7
    // Explanation:
    // Initially we are at (0, 0) and cost is 0.
    // Current Position	Move	New Position	Total Cost
    // (0, 0)	Move Down	(1, 0)	0 + 2 = 2
    // (1, 0)	Move Right	(1, 1)	2 + 5 = 7
    // (1, 1)	Teleport to (2, 2)	(2, 2)	7 + 0 = 7
    // The minimum cost to reach bottom-right cell is 7.
    fmt.Println(minCost([][]int{{1,3,3},{2,5,4},{4,3,5}}, 2)) // 7
    // Example 2:
    // Input: grid = [[1,2],[2,3],[3,4]], k = 1
    // Output: 9
    // Explanation:
    // Initially we are at (0, 0) and cost is 0.
    // Current Position	Move	New Position	Total Cost
    // (0, 0)	Move Down	(1, 0)	0 + 2 = 2
    // (1, 0)	Move Right	(1, 1)	2 + 3 = 5
    // (1, 1)	Move Down	(2, 1)	5 + 4 = 9
    // The minimum cost to reach bottom-right cell is 9.
    fmt.Println(minCost([][]int{{1,2},{2,3},{3,4}}, 1)) // 9
}
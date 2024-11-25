package main

// 2017. Grid Game
// You are given a 0-indexed 2D array grid of size 2 x n, 
// where grid[r][c] represents the number of points at position (r, c) on the matrix. 
// Two robots are playing a game on this matrix.

// Both robots initially start at (0, 0) and want to reach (1, n-1). 
// Each robot may only move to the right ((r, c) to (r, c + 1)) or down ((r, c) to (r + 1, c)).

// At the start of the game, the first robot moves from (0, 0) to (1, n-1), 
// collecting all the points from the cells on its path. For all cells (r, c) traversed on the path, grid[r][c] is set to 0. 
// Then, the second robot moves from (0, 0) to (1, n-1), collecting the points on its path. 
// Note that their paths may intersect with one another.

// The first robot wants to minimize the number of points collected by the second robot. 
// In contrast, the second robot wants to maximize the number of points it collects. 
// If both robots play optimally, return the number of points collected by the second robot.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/09/08/a1.png" />
// Input: grid = [[2,5,4],[1,5,1]]
// Output: 4
// Explanation: The optimal path taken by the first robot is shown in red, and the optimal path taken by the second robot is shown in blue.
// The cells visited by the first robot are set to 0.
// The second robot will collect 0 + 0 + 4 + 0 = 4 points.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/09/08/a2.png" />
// Input: grid = [[3,3,1],[8,5,2]]
// Output: 4
// Explanation: The optimal path taken by the first robot is shown in red, and the optimal path taken by the second robot is shown in blue.
// The cells visited by the first robot are set to 0.
// The second robot will collect 0 + 3 + 1 + 0 = 4 points.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/09/08/a3.png" />
// Input: grid = [[1,3,1,15],[1,3,3,1]]
// Output: 7
// Explanation: The optimal path taken by the first robot is shown in red, and the optimal path taken by the second robot is shown in blue.
// The cells visited by the first robot are set to 0.
// The second robot will collect 0 + 1 + 3 + 3 + 0 = 7 points.

// Constraints:
//     grid.length == 2
//     n == grid[r].length
//     1 <= n <= 5 * 10^4
//     1 <= grid[r][c] <= 10^5

import "fmt"
import "math"

func gridGame(grid [][]int) int64 {
    res, n := int64(math.MaxInt64), len(grid[0])
    prefix, postfix := make([]int64, n), make([]int64, n) // setup postfix and prefix arrays
    for i := 1; i < n; i++ {
        j := n - 1 - i
        prefix[i], postfix[j] = prefix[i-1] + int64(grid[1][i-1]), postfix[j+1] + int64(grid[0][j+1])
    }
    min := func (x, y int64) int64 { if x < y { return x; }; return y; }
    max := func (x, y int64) int64 { if x > y { return x; }; return y; }
    // actual logic, by checking every time the min points if robot2 plays optimally
    for i := 0; i < n; i++ {
        res = min(res, max(postfix[i], prefix[i]))
    }
    return res
}

func gridGame1(grid [][]int) int64 {
    res, sum1, sum2 := math.MaxInt64, 0, 0
    for _, v := range grid[0] {
        sum1 += v
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for j, v := range grid[0] {
        sum1 -= v
        res = min(res, max(sum1, sum2))
        sum2 += grid[1][j]
    }
    return int64(res)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/09/08/a1.png" />
    // Input: grid = [[2,5,4],[1,5,1]]
    // Output: 4
    // Explanation: The optimal path taken by the first robot is shown in red, and the optimal path taken by the second robot is shown in blue.
    // The cells visited by the first robot are set to 0.
    // The second robot will collect 0 + 0 + 4 + 0 = 4 points.
    fmt.Println(gridGame([][]int{{2,5,4},{1,5,1}})) // 4
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/09/08/a2.png" />
    // Input: grid = [[3,3,1],[8,5,2]]
    // Output: 4
    // Explanation: The optimal path taken by the first robot is shown in red, and the optimal path taken by the second robot is shown in blue.
    // The cells visited by the first robot are set to 0.
    // The second robot will collect 0 + 3 + 1 + 0 = 4 points.
    fmt.Println(gridGame([][]int{{3,3,1},{8,5,2}})) // 4
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/09/08/a3.png" />
    // Input: grid = [[1,3,1,15],[1,3,3,1]]
    // Output: 7
    // Explanation: The optimal path taken by the first robot is shown in red, and the optimal path taken by the second robot is shown in blue.
    // The cells visited by the first robot are set to 0.
    // The second robot will collect 0 + 1 + 3 + 3 + 0 = 7 points.
    fmt.Println(gridGame([][]int{{1,3,1,15},{1,3,3,1}})) // 7

    fmt.Println(gridGame1([][]int{{2,5,4},{1,5,1}})) // 4
    fmt.Println(gridGame1([][]int{{3,3,1},{8,5,2}})) // 4
    fmt.Println(gridGame1([][]int{{1,3,1,15},{1,3,3,1}})) // 7
}
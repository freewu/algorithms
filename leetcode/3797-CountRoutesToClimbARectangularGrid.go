package main

// 3797. Count Routes to Climb a Rectangular Grid
// You are given a string array grid of size n, where each string grid[i] has length m. 
// The character grid[i][j] is one of the following symbols:
//     1. '.': The cell is available.
//     2. '#': The cell is blocked.

// You want to count the number of different routes to climb grid. 
// Each route must start from any cell in the bottom row (row n - 1) and end in the top row (row 0).

// However, there are some constraints on the route.
//     1. You can only move from one available cell to another available cell.
//     2. The Euclidean distance of each move is at most d, where d is an integer parameter given to you. 
//        The Euclidean distance between two cells (r1, c1), (r2, c2) is sqrt((r1 - r2)2 + (c1 - c2)2).
//     3. Each move either stays on the same row or moves to the row directly above (from row r to r - 1).
//     4. You cannot stay on the same row for two consecutive turns. 
//        If you stay on the same row in a move (and this move is not the last move), your next move must go to the row above.

// Return an integer denoting the number of such routes. Since the answer may be very large, return it modulo 10^9 + 7.

// Example 1:
// Input: grid = ["..","#."], d = 1
// Output: 2
// Explanation:
// We label the cells we visit in the routes sequentially, starting from 1. The two routes are:
// .2
// #1
// 32
// #1
// We can move from the cell (1, 1) to the cell (0, 1) because the Euclidean distance is sqrt((1 - 0)^2 + (1 - 1)^2) = sqrt(1) <= d.
// However, we cannot move from the cell (1, 1) to the cell (0, 0) because the Euclidean distance is sqrt((1 - 0)^2 + (1 - 0)^2) = sqrt(2) > d.

// Example 2:
// Input: grid = ["..","#."], d = 2
// Output: 4
// Explanation:
// Two of the routes are given in example 1. The other two routes are:
// 2.
// #1
// 23
// #1
// Note that we can move from (1, 1) to (0, 0) because the Euclidean distance is sqrt(2) <= d.

// Example 3:
// Input: grid = ["#"], d = 750
// Output: 0
// Explanation:
// We cannot choose any cell as the starting cell. Therefore, there are no routes.

// Example 4:
// Input: grid = [".."], d = 1
// Output: 4
// Explanation:
// The possible routes are:
// .1
// 1.
// 12
// 21

// Constraints:
//     1 <= n == grid.length <= 750
//     1 <= m == grid[i].length <= 750
//     grid[i][j] is '.' or '#'.
//     1 <= d <= 750

import "fmt"

func numberOfRoutes(grid []string, d int) int {
    n, mod := len(grid[0]), 1_000_000_007
    sum, sumF := make([]int, n + 1), make([]int, n + 1)
    for i, row := range grid {
        // f 的前缀和
        for j, ch := range row {
            if ch == '#' {
                sumF[j+1] = sumF[j]
            } else if i == 0 { // 第一行（起点）
                sumF[j+1] = sumF[j] + 1 // DP 初始值
            } else {
                sumF[j+1] = (sumF[j] + sum[min(j+d, n)] - sum[max(j-d+1, 0)]) % mod
            }
        }
        // f[j] + g[j] 的前缀和
        for j, ch := range row {
            if ch == '#' {
                sum[j+1] = sum[j]
            } else {
                // -f[j] 和 +f[j] 抵消了
                sum[j+1] = (sum[j] + sumF[min(j+d+1, n)] - sumF[max(j-d, 0)]) % mod
            }
        }
    }
    return (sum[n] + mod) % mod // 保证结果非负 
}

func main() {
    // Example 1:
    // Input: grid = ["..","#."], d = 1
    // Output: 2
    // Explanation:
    // We label the cells we visit in the routes sequentially, starting from 1. The two routes are:
    // .2
    // #1
    // 32
    // #1
    // We can move from the cell (1, 1) to the cell (0, 1) because the Euclidean distance is sqrt((1 - 0)^2 + (1 - 1)^2) = sqrt(1) <= d.
    // However, we cannot move from the cell (1, 1) to the cell (0, 0) because the Euclidean distance is sqrt((1 - 0)^2 + (1 - 0)^2) = sqrt(2) > d.
    fmt.Println(numberOfRoutes([]string{"..", "#."}, 1)) // 2
    // Example 2:
    // Input: grid = ["..","#."], d = 2
    // Output: 4
    // Explanation:
    // Two of the routes are given in example 1. The other two routes are:
    // 2.
    // #1
    // 23
    // #1
    // Note that we can move from (1, 1) to (0, 0) because the Euclidean distance is sqrt(2) <= d.
    fmt.Println(numberOfRoutes([]string{"..", "#."}, 2)) // 4    
    // Example 3:
    // Input: grid = ["#"], d = 750
    // Output: 0
    // Explanation:
    // We cannot choose any cell as the starting cell. Therefore, there are no routes.
    fmt.Println(numberOfRoutes([]string{"#"}, 750)) // 0
    // Example 4:
    // Input: grid = [".."], d = 1
    // Output: 4
    // Explanation:
    // The possible routes are:
    // .1
    // 1.
    // 12
    // 21
    fmt.Println(numberOfRoutes([]string{".."}, 1)) // 4
}
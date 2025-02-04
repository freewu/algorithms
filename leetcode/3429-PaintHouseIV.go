package main

// 3429. Paint House IV
// You are given an even integer n representing the number of houses arranged in a straight line, 
// and a 2D array cost of size n x 3, where cost[i][j] represents the cost of painting house i with color j + 1.

// The houses will look beautiful if they satisfy the following conditions:
//     1. No two adjacent houses are painted the same color.
//     2. Houses equidistant from the ends of the row are not painted the same color. 
//        For example, if n = 6, houses at positions (0, 5), (1, 4), and (2, 3) are considered equidistant.

// Return the minimum cost to paint the houses such that they look beautiful.

// Example 1:
// Input: n = 4, cost = [[3,5,7],[6,2,9],[4,8,1],[7,3,5]]
// Output: 9
// Explanation:
// The optimal painting sequence is [1, 2, 3, 2] with corresponding costs [3, 2, 1, 3]. 
// This satisfies the following conditions:
//     1. No adjacent houses have the same color.
//     2. Houses at positions 0 and 3 (equidistant from the ends) are not painted the same color (1 != 2).
//     3. Houses at positions 1 and 2 (equidistant from the ends) are not painted the same color (2 != 3).
// The minimum cost to paint the houses so that they look beautiful is 3 + 2 + 1 + 3 = 9.

// Example 2:
// Input: n = 6, cost = [[2,4,6],[5,3,8],[7,1,9],[4,6,2],[3,5,7],[8,2,4]]
// Output: 18
// Explanation:
// The optimal painting sequence is [1, 3, 2, 3, 1, 2] with corresponding costs [2, 8, 1, 2, 3, 2]. 
// This satisfies the following conditions:
//     1. No adjacent houses have the same color.
//     2. Houses at positions 0 and 5 (equidistant from the ends) are not painted the same color (1 != 2).
//     3. Houses at positions 1 and 4 (equidistant from the ends) are not painted the same color (3 != 1).
//     4. Houses at positions 2 and 3 (equidistant from the ends) are not painted the same color (2 != 3).
// The minimum cost to paint the houses so that they look beautiful is 2 + 8 + 1 + 2 + 3 + 2 = 18.

// Constraints:
//     2 <= n <= 10^5
//     n is even.
//     cost.length == n
//     cost[i].length == 3
//     0 <= cost[i][j] <= 10^5

import "fmt"

func minCost(n int, cost [][]int) int64 {
    dp := make([][4][4]int64, n)
    for i := range dp {
        for j := range dp[i] {
            for k := range dp[i][j] {
                dp[i][j][k] = -1
            }
        }
    }
    min := func (x, y int64) int64 { if x < y { return x; }; return y; }
    var dfs func(i, left, right int) int64
    dfs = func(i, left, right int) int64 {
        if i >= n / 2 { return 0 }
        if dp[i][left][right] != -1 { return dp[i][left][right] }
        res := int64(1 << 61)
        for col1 := 0; col1 < 3; col1++ {
            if col1 != left {
                for col2 := 0; col2 < 3; col2++ {
                    if col2 != right && col2 != col1 {
                        res = min(res, int64(cost[i][col1] + cost[n - 1 - i][col2]) + dfs(i + 1, col1, col2))
                    }
                }
            }
        }
        dp[i][left][right] = res
        return res
    }
    return dfs(0, 3, 3)
}

func minCost1(n int, cost [][]int) int64 {
    x0y1, x0y2, x1y0, x1y2, x2y0, x2y1 := 0, 0, 0, 0, 0, 0
    for i := n/2-1; i >= 0; i-- {
        x0y1, x0y2, x1y0, x1y2, x2y0, x2y1 =
            cost[i][0] + cost[n-1-i][1] + min(x1y0, x1y2, x2y0), cost[i][0]+cost[n-1-i][2] + min(x1y0, x2y0, x2y1),
            cost[i][1] + cost[n-1-i][0] + min(x0y1, x0y2, x2y1), cost[i][1]+cost[n-1-i][2] + min(x0y1, x2y0, x2y1),
            cost[i][2] + cost[n-1-i][0] + min(x0y1, x0y2, x1y2), cost[i][2]+cost[n-1-i][1] + min(x0y2, x1y0, x1y2)
    }
    return int64(min(x0y1, x0y2, x1y0, x1y2, x2y0, x2y1))
}

func main() {
    // Example 1:
    // Input: n = 4, cost = [[3,5,7],[6,2,9],[4,8,1],[7,3,5]]
    // Output: 9
    // Explanation:
    // The optimal painting sequence is [1, 2, 3, 2] with corresponding costs [3, 2, 1, 3]. 
    // This satisfies the following conditions:
    //     1. No adjacent houses have the same color.
    //     2. Houses at positions 0 and 3 (equidistant from the ends) are not painted the same color (1 != 2).
    //     3. Houses at positions 1 and 2 (equidistant from the ends) are not painted the same color (2 != 3).
    // The minimum cost to paint the houses so that they look beautiful is 3 + 2 + 1 + 3 = 9.
    fmt.Println(minCost(4, [][]int{{3,5,7},{6,2,9},{4,8,1},{7,3,5}})) // 9
    // Example 2:
    // Input: n = 6, cost = [[2,4,6],[5,3,8],[7,1,9],[4,6,2],[3,5,7],[8,2,4]]
    // Output: 18
    // Explanation:
    // The optimal painting sequence is [1, 3, 2, 3, 1, 2] with corresponding costs [2, 8, 1, 2, 3, 2]. 
    // This satisfies the following conditions:
    //     1. No adjacent houses have the same color.
    //     2. Houses at positions 0 and 5 (equidistant from the ends) are not painted the same color (1 != 2).
    //     3. Houses at positions 1 and 4 (equidistant from the ends) are not painted the same color (3 != 1).
    //     4. Houses at positions 2 and 3 (equidistant from the ends) are not painted the same color (2 != 3).
    // The minimum cost to paint the houses so that they look beautiful is 2 + 8 + 1 + 2 + 3 + 2 = 18.
    fmt.Println(minCost(6, [][]int{{2,4,6},{5,3,8},{7,1,9},{4,6,2},{3,5,7},{8,2,4}})) // 18

    fmt.Println(minCost1(4, [][]int{{3,5,7},{6,2,9},{4,8,1},{7,3,5}})) // 9
    fmt.Println(minCost1(6, [][]int{{2,4,6},{5,3,8},{7,1,9},{4,6,2},{3,5,7},{8,2,4}})) // 18
}
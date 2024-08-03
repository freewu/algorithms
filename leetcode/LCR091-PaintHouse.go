package main

// LCR 091. 粉刷房子
// 假如有一排房子，共 n 个，每个房子可以被粉刷成红色、蓝色或者绿色这三种颜色中的一种，你需要粉刷所有的房子并且使其相邻的两个房子颜色不能相同。

// 当然，因为市场上不同颜色油漆的价格不同，所以房子粉刷成不同颜色的花费成本也是不同的。
// 每个房子粉刷成不同颜色的花费是以一个 n x 3 的正整数矩阵 costs 来表示的。

// 例如，costs[0][0] 表示第 0 号房子粉刷成红色的成本花费；costs[1][2] 表示第 1 号房子粉刷成绿色的花费，以此类推。
// 请计算出粉刷完所有房子最少的花费成本。

// 示例 1：
// 输入: costs = [[17,2,17],[16,16,5],[14,3,19]]
// 输出: 10
// 解释: 将 0 号房子粉刷成蓝色，1 号房子粉刷成绿色，2 号房子粉刷成蓝色。
//      最少花费: 2 + 5 + 3 = 10。

// 示例 2：
// 输入: costs = [[7,6,2]]
// 输出: 2

// 提示:
//     costs.length == n
//     costs[i].length == 3
//     1 <= n <= 100
//     1 <= costs[i][j] <= 20

import "fmt"

func minCost(costs [][]int) int {
    r, b, g := 0, 0, 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < len(costs) ; i++ {
        // 需要粉刷所有的房子并且使其相邻的两个房子颜色不能相同 所以 r 只能取 b, g
        r, b, g = costs[i][0] + min(b, g), costs[i][1] + min(r, g), costs[i][2] + min(r, b)
    }
    return min(r, min(b, g))
}

func minCost1(costs [][]int) int {
    n, inf := len(costs), 1 << 32 - 1
    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, 3)
        for j := 0; j < 3; j++ {
            dp[i][j] = inf
        }
    }
    dp[0][0], dp[0][1], dp[0][2] = costs[0][0], costs[0][1], costs[0][2]
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i < n; i++ {
        for j := 0; j < 3; j++ {
            for k := 0; k < 3; k++ {
                if k == j {
                    continue
                }
                dp[i][j] = min(dp[i-1][k], dp[i][j])
            }
            dp[i][j] += costs[i][j]
        }
        // // 颜色 0 只能挨着颜色 1 或颜色 2
        // dp[i][0] = min(dp[i-1][1], dp[i-1][2]) + costs[i][0]
        // // 颜色 1 只能挨着颜色 0 或颜色 2
        // dp[i][1] = min(dp[i-1][0], dp[i-1][2]) + costs[i][1]
        // // 颜色 2 只能挨着颜色 0 或颜色 1
        // dp[i][2] = min(dp[i-1][0], dp[i-1][1]) + costs[i][2]
    }
    return min(dp[n-1][0], min(dp[n-1][1], dp[n-1][2]))
}

func main() {
    // Example 1:
    // Input: costs = [[17,2,17],[16,16,5],[14,3,19]]
    // Output: 10
    // Explanation: Paint house 0 into blue, paint house 1 into green, paint house 2 into blue.
    // Minimum cost: 2 + 5 + 3 = 10.
    fmt.Println(minCost([][]int{{17,2,17},{16,16,5},{14,3,19}})) // 10
    // Example 2:
    // Input: costs = [[7,6,2]]
    // Output: 2
    fmt.Println(minCost([][]int{{7,6,2}})) // 2

    fmt.Println(minCost1([][]int{{17,2,17},{16,16,5},{14,3,19}})) // 10
    fmt.Println(minCost1([][]int{{7,6,2}})) // 2
}
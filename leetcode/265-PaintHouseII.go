package main

// 265. Paint House II
// There are a row of n houses, each house can be painted with one of the k colors. 
// The cost of painting each house with a certain color is different. 
// You have to paint all the houses such that no two adjacent houses have the same color.

// The cost of painting each house with a certain color is represented by an n x k cost matrix costs.
//     For example, costs[0][0] is the cost of painting house 0 with color 0; 
//     costs[1][2] is the cost of painting house 1 with color 2, and so on...

// Return the minimum cost to paint all houses.

// Example 1:
// Input: costs = [[1,5,3],[2,9,4]]
// Output: 5
// Explanation:
// Paint house 0 into color 0, paint house 1 into color 2. Minimum cost: 1 + 4 = 5; 
// Or paint house 0 into color 2, paint house 1 into color 0. Minimum cost: 3 + 2 = 5.

// Example 2:
// Input: costs = [[1,3],[2,4]]
// Output: 5
 
// Constraints:
//     costs.length == n
//     costs[i].length == k
//     1 <= n <= 100
//     2 <= k <= 20
//     1 <= costs[i][j] <= 20

// Follow up: Could you solve it in O(nk) runtime?

import "fmt"

func minCostII(costs [][]int) int {
    n, k := len(costs), len(costs[0])
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, k)
    }
    dp[0] = costs[0]
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i < n; i ++ {
        for j := 0; j < k; j ++ {
            dp[i][j] = 2000
            for m := 0; m < k; m ++ {
                if m == j {
                    continue
                }
                dp[i][j] = min(dp[i][j], dp[i - 1][m] + costs[i][j])
            }
        }
    }
    res := dp[n - 1][0]
    for i := range dp[n - 1] {
        res = min(res, dp[n - 1][i])
    }
    return res
}

func minCostII1(costs [][]int) int {
    n, k := len(costs), len(costs[0])
    if n == 0 {
        return 0
    }
    // 找到当前房子的涂色方案中，不是 nextColor 方案的最小代价
    minColor := func(colors []int, nextColor int) int {
        res := 1 << 32 - 1 // 找最小，初始值为最大
        for i, cost := range colors {
            if i != nextColor { // 遍历 [0,k) 种方案，跳过 nextColor 方案
                if res > cost {
                    res = cost
                }
            }
        }
        return res
    }
    for i := 1; i < n; i++ {
        for j := 0; j < k; j++ {
            // 选择第 i-1 号房子的方案中，涂色方案不为 j 的花费最小的那个
            // 用 minColor(costs[i-1], j) 更新当前 第 i 号房子选择 j 涂色方案的最小总花费
            costs[i][j] = minColor(costs[i-1], j) + costs[i][j]
        }
    }
    // 返回第 n-1 号房子最小的花费, k+1 是一个遍历时取不到方案
    return minColor(costs[n-1], k+1)
}


func minCostII2(costs [][]int) int {
    m, n := len(costs), len(costs[0])
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    dp[0] = costs[0]
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i < m; i++ {
        for j := 0; j < n; j++ {
            mc := 1 << 32 - 1 
            for k := 0; k < n; k++ {
                if k != j {
                    mc = min(mc, dp[i-1][k]+costs[i][j])
                }
            }
            dp[i][j] = mc
        }
    }
    var res = 1 << 32 - 1 
    for _, e := range dp[m-1] {
        res = min(res, e)
    }
    return res
}

func main() {
    // Paint house 0 into color 0, paint house 1 into color 2. Minimum cost: 1 + 4 = 5; 
    // Or paint house 0 into color 2, paint house 1 into color 0. Minimum cost: 3 + 2 = 5.
    fmt.Println(minCostII([][]int{{1,5,3},{2,9,4}})) // 5
    // 1 + 4 or 3 +2
    fmt.Println(minCostII([][]int{{1,3},{2,4}})) // 5

    fmt.Println(minCostII1([][]int{{1,5,3},{2,9,4}})) // 5
    fmt.Println(minCostII1([][]int{{1,3},{2,4}})) // 5

    fmt.Println(minCostII2([][]int{{1,5,3},{2,9,4}})) // 5
    fmt.Println(minCostII2([][]int{{1,3},{2,4}})) // 5
}
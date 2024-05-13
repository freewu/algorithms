package main

// 1473. Paint House III
// There is a row of m houses in a small city, each house must be painted with one of the n colors (labeled from 1 to n), 
// some houses that have been painted last summer should not be painted again.

// A neighborhood is a maximal group of continuous houses that are painted with the same color.
//     For example: houses = [1,2,2,3,3,2,1,1] contains 5 neighborhoods [{1}, {2,2}, {3,3}, {2}, {1,1}].

// Given an array houses, an m x n matrix cost and an integer target where:
//     houses[i]: is the color of the house i, and 0 if the house is not painted yet.
//     cost[i][j]: is the cost of paint the house i with the color j + 1.

// Return the minimum cost of painting all the remaining houses in such a way that there are exactly target neighborhoods. 
// If it is not possible, return -1.

// Example 1:
// Input: houses = [0,0,0,0,0], cost = [[1,10],[10,1],[10,1],[1,10],[5,1]], m = 5, n = 2, target = 3
// Output: 9
// Explanation: Paint houses of this way [1,2,2,1,1]
// This array contains target = 3 neighborhoods, [{1}, {2,2}, {1,1}].
// Cost of paint all houses (1 + 1 + 1 + 1 + 5) = 9.

// Example 2:
// Input: houses = [0,2,1,2,0], cost = [[1,10],[10,1],[10,1],[1,10],[5,1]], m = 5, n = 2, target = 3
// Output: 11
// Explanation: Some houses are already painted, Paint the houses of this way [2,2,1,2,2]
// This array contains target = 3 neighborhoods, [{2,2}, {1}, {2,2}]. 
// Cost of paint the first and last house (10 + 1) = 11.

// Example 3:
// Input: houses = [3,1,2,3], cost = [[1,1,1],[1,1,1],[1,1,1],[1,1,1]], m = 4, n = 3, target = 3
// Output: -1
// Explanation: Houses are already painted with a total of 4 neighborhoods [{3},{1},{2},{3}] different of target = 3.
 
// Constraints:
//     m == houses.length == cost.length
//     n == cost[i].length
//     1 <= m <= 100
//     1 <= n <= 20
//     1 <= target <= m
//     0 <= houses[i] <= n
//     1 <= cost[i][j] <= 10^4

import "fmt"
import "math"

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
    dp, res := make([][][]float64,m), math.Inf(1)
    for i,_ := range dp {
        dp[i] = make([][]float64, target+1)
        for j,_ := range(dp[i]){
            dp[i][j] = make([]float64, n)
        }
    }
    for i:=0; i<m;i++{
        for j:=0; j<target+1;j++{
            for k:=0; k<n; k++{
                dp[i][j][k] = math.Inf(1)
            }
        }
    }
    for c:=1;c<n+1;c++{
        if houses[0] == c{
            dp[0][1][c - 1] = 0
        }else if houses[0] == 0{
            dp[0][1][c - 1] = float64(cost[0][c - 1])
        }
    }
    for i:=1; i<m;i++{
        x := int(math.Min(float64(target),float64(i+1)))
        for k:=1; k<x+1; k++{
            for c:=1; c<n+1; c++{
                if houses[i] != 0 && c != houses[i]{
                    continue
                }
                same_neighbor_cost := dp[i - 1][k][c - 1]
                diff_neighbor_cost := math.Inf(1)
                for c_:=0; c_<n;c_++{
                    if c_ == c-1{
                        continue    
                    }
                    diff_neighbor_cost = math.Min(diff_neighbor_cost,dp[i - 1][k - 1][c_])   
                }
                var multiplier int
                if houses[i] == 0{
                    multiplier = 1
                }else{
                    multiplier = 0
                }
                paint_cost := float64(cost[i][c - 1] * multiplier)
                dp[i][k][c - 1] = math.Min(same_neighbor_cost, diff_neighbor_cost) + paint_cost
            }
        }
    }
    for _, value:= range(dp[m-1][target]){
        res = math.Min(res,value)
    }
    if res == math.Inf(1){
        return -1
    }
    return int(res)
}

func minCost1(houses []int, cost [][]int, m int, n int, target int) int {
    res, dp := 0x3f3f3f3f, [101][21][101]int{}
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < m+1; i++ { // 所有的都先设置为不可能状态
        for j := 0; j < n+1; j++ {
            for t := 0; t < target+1; t++ {
                if i == 0 && (t == 0 || t == 1) {
                    // 0 个房子的时候只能形成 0 个街区
                    // 0 个 没有颜色 只是为了判断1号房子和0号房子颜色异同使用
                    dp[i][j][t] = 0
                } else {
                    dp[i][j][t] = 0x3f3f3f3f
                }
            }
        }
    }
    // 第i个房子在颜色j状态下形成 k个街区的状态转移
    // 颜色和上一个房子相同
    // [i][j][t] 依赖 [i-1][j][t]
    // 颜色和上一个房子不同
    // [i][j][t] 依赖 [i-1][!j][t-1]
    // 从这两种方案中选最小的
    for i := 1; i < m+1; i++ {
        for j := 1; j < n+1; j++ {
            for t := 1; t < target+1 && t<=i; t++ {
                if houses[i-1] == 0 { // 没有涂色随便选
                    
                    dp[i][j][t] = dp[i-1][j][t] + cost[i-1][j-1] // 相同的方案
                    for c:= 1; c<=n; c++ { // 不同的方案
                        if c != j {
                            dp[i][j][t] = min(dp[i][j][t], dp[i-1][c][t-1]+cost[i-1][j-1])
                        }
                    }
                } else {
                    if j != houses[i-1] { // 已经染色别的颜色没有意义
                        continue
                    }
                    dp[i][j][t] = dp[i-1][j][t] // 相同的方案
                    for c:= 1; c<=n; c++ { // 不同的方案
                        if c != j {
                            dp[i][j][t] = min(dp[i][j][t], dp[i-1][c][t-1])
                        }
                    }
                }
            }
        }
    }
    for j := 1; j <= n; j++ {
        res = min(res, dp[m][j][target])
    }
    if res == 0x3f3f3f3f {
        return -1
    }
    return res
}

func main() {
    // Example 1:
    // Input: houses = [0,0,0,0,0], cost = [[1,10],[10,1],[10,1],[1,10],[5,1]], m = 5, n = 2, target = 3
    // Output: 9
    // Explanation: Paint houses of this way [1,2,2,1,1]
    // This array contains target = 3 neighborhoods, [{1}, {2,2}, {1,1}].
    // Cost of paint all houses (1 + 1 + 1 + 1 + 5) = 9.
    fmt.Println(minCost([]int{0,0,0,0,0}, [][]int{{1,10},{10,1},{10,1},{1,10},{5,1}}, 5, 2, 3)) // 9
    // Example 2:
    // Input: houses = [0,2,1,2,0], cost = [[1,10],[10,1],[10,1],[1,10],[5,1]], m = 5, n = 2, target = 3
    // Output: 11
    // Explanation: Some houses are already painted, Paint the houses of this way [2,2,1,2,2]
    // This array contains target = 3 neighborhoods, [{2,2}, {1}, {2,2}]. 
    // Cost of paint the first and last house (10 + 1) = 11.
    fmt.Println(minCost([]int{0,2,1,2,0}, [][]int{{1,10},{10,1},{10,1},{1,10},{5,1}}, 5, 2, 3)) // 11
    // Example 3:
    // Input: houses = [3,1,2,3], cost = [[1,1,1],[1,1,1],[1,1,1],[1,1,1]], m = 4, n = 3, target = 3
    // Output: -1
    // Explanation: Houses are already painted with a total of 4 neighborhoods [{3},{1},{2},{3}] different of target = 3.
    fmt.Println(minCost([]int{3,1,2,3}, [][]int{{1,1,1},{1,1,1},{1,1,1},{1,1,1}}, 4, 3, 3)) // -1

    fmt.Println(minCost1([]int{0,0,0,0,0}, [][]int{{1,10},{10,1},{10,1},{1,10},{5,1}}, 5, 2, 3)) // 9
    fmt.Println(minCost1([]int{0,2,1,2,0}, [][]int{{1,10},{10,1},{10,1},{1,10},{5,1}}, 5, 2, 3)) // 11
    fmt.Println(minCost1([]int{3,1,2,3}, [][]int{{1,1,1},{1,1,1},{1,1,1},{1,1,1}}, 4, 3, 3)) // -1
}
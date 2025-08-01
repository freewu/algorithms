package main

// 3562. Maximum Profit from Trading Stocks with Discounts
// You are given an integer n, representing the number of employees in a company.
// Each employee is assigned a unique ID from 1 to n, and employee 1 is the CEO. 
// You are given two 1-based integer arrays, present and future, each of length n, where:
//     1. present[i] represents the current price at which the ith employee can buy a stock today.
//     2. future[i] represents the expected price at which the ith employee can sell the stock tomorrow.

// The company's hierarchy is represented by a 2D integer array hierarchy, where hierarchy[i] = [ui, vi] means that employee ui is the direct boss of employee vi.

// Additionally, you have an integer budget representing the total funds available for investment.

// However, the company has a discount policy: if an employee's direct boss purchases their own stock, then the employee can buy their stock at half the original price (floor(present[v] / 2)).

// Return the maximum profit that can be achieved without exceeding the given budget.

// Note:
//     1. You may buy each stock at most once.
//     2. You cannot use any profit earned from future stock prices to fund additional investments and must buy only from budget.

// Example 1:
// Input: n = 2, present = [1,2], future = [4,3], hierarchy = [[1,2]], budget = 3
// Output: 5
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/04/09/screenshot-2025-04-10-at-053641.png" />
// Employee 1 buys the stock at price 1 and earns a profit of 4 - 1 = 3.
// Since Employee 1 is the direct boss of Employee 2, Employee 2 gets a discounted price of floor(2 / 2) = 1.
// Employee 2 buys the stock at price 1 and earns a profit of 3 - 1 = 2.
// The total buying cost is 1 + 1 = 2 <= budget. Thus, the maximum total profit achieved is 3 + 2 = 5.

// Example 2:
// Input: n = 2, present = [3,4], future = [5,8], hierarchy = [[1,2]], budget = 4
// Output: 4
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/04/09/screenshot-2025-04-10-at-053641.png" />
// Employee 2 buys the stock at price 4 and earns a profit of 8 - 4 = 4.
// Since both employees cannot buy together, the maximum profit is 4.

// Example 3:
// Input: n = 3, present = [4,6,8], future = [7,9,11], hierarchy = [[1,2],[1,3]], budget = 10
// Output: 10
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/04/09/image.png" />
// Employee 1 buys the stock at price 4 and earns a profit of 7 - 4 = 3.
// Employee 3 would get a discounted price of floor(8 / 2) = 4 and earns a profit of 11 - 4 = 7.
// Employee 1 and Employee 3 buy their stocks at a total cost of 4 + 4 = 8 <= budget. Thus, the maximum total profit achieved is 3 + 7 = 10.

// Example 4:
// Input: n = 3, present = [5,2,3], future = [8,5,6], hierarchy = [[1,2],[2,3]], budget = 7
// Output: 12
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/04/09/screenshot-2025-04-10-at-054114.png" />
// Employee 1 buys the stock at price 5 and earns a profit of 8 - 5 = 3.
// Employee 2 would get a discounted price of floor(2 / 2) = 1 and earns a profit of 5 - 1 = 4.
// Employee 3 would get a discounted price of floor(3 / 2) = 1 and earns a profit of 6 - 1 = 5.
// The total cost becomes 5 + 1 + 1 = 7 <= budget. Thus, the maximum total profit achieved is 3 + 4 + 5 = 12.

// Constraints:
//     1 <= n <= 160
//     present.length, future.length == n
//     1 <= present[i], future[i] <= 50
//     hierarchy.length == n - 1
//     hierarchy[i] == [ui, vi]
//     1 <= ui, vi <= n
//     ui != vi
//     1 <= budget <= 160
//     There are no duplicate edges.
//     Employee 1 is the direct or indirect boss of every employee.
//     The input graph hierarchy is guaranteed to have no cycles.

import "fmt"
import "slices"

func maxProfit(n int, present []int, future []int, hierarchy [][]int, budget int) int {
    g := make([][]int, n)
    for _, e := range hierarchy {
        x, y := e[0] - 1, e[1] - 1
        g[x] = append(g[x], y)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(x int) [][2]int
    dfs = func(x int) [][2]int {
        // 计算从 x 的所有儿子子树 y 中，能得到的最大利润之和
        subF := make([][2]int, budget+1)
        for _, y := range g[x] {
            fy := dfs(y)
            for j := budget; j >= 0; j-- {
                // 枚举子树 y 的预算为 jy
                // 当作一个体积为 jy，价值为 resY=fy[jy][k] 的物品
                for jy, p := range fy[:j+1] {
                    for k, resY := range p {
                        subF[j][k] = max(subF[j][k], subF[j-jy][k]+resY)
                    }
                }
            }
        }
        f := make([][2]int, budget + 1)
        for j, p := range subF {
            for k := 0; k < 2; k++ {
                cost := present[x] / (k + 1)
                if j >= cost {
                    // 不买 x，转移来源是 subF[j][0]
                    // 买 x，转移来源为 subF[j-cost][1]，因为对于子树来说，父节点一定买
                    f[j][k] = max(p[0], subF[j-cost][1]+future[x]-cost)
                } else { // 只能不买 x
                    f[j][k] = p[0]
                }
            }
        }
        return f
    }
    return dfs(0)[budget][0]
}

func maxProfit1(n int, present []int, future []int, hierarchy [][]int, budget int) int {
    g, inf := make([][]int, n), 1 << 31
    for _, e := range hierarchy {
        x, y := e[0]-1, e[1]-1
        g[x] = append(g[x], y)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(x int) [2][]int
    dfs = func(x int) [2][]int {
        // 计算从 x 的所有儿子子树 y 中，能得到的最大利润之和
        subF := [2][]int{make([]int, budget+1), make([]int, budget+1)}
        for i := 1; i <= budget; i++ {
            subF[0][i] = -inf // 表示不存在对应的花费总和
            subF[1][i] = -inf
        }
        for _, y := range g[x] {
            fy := dfs(y)
            for k, fyk := range fy {
                nf := make([]int, budget+1)
                for i := 1; i <= budget; i++ {
                    nf[i] = -inf
                }
                for jy, resY := range fyk {
                    if resY < 0 { // 重要优化：物品价值为负数，一定不选
                        continue
                    }
                    for j := jy; j <= budget; j++ {
                        nf[j] = max(nf[j], subF[k][j-jy] + resY)
                    }
                }
                subF[k] = nf
            }
        }
        f := [2][]int{}
        for k := 0; k < 2; k++ {
            // 不买 x，转移来源为 subF[0]，因为对于子树来说，父节点一定不买
            f[k] = slices.Clone(subF[0])
            cost := present[x] / (k + 1)
            // 买 x，转移来源为 subF[1]，因为对于子树来说，父节点一定买
            for j := cost; j <= budget; j++ {
                f[k][j] = max(f[k][j], subF[1][j-cost]+future[x]-cost)
            }
        }
        return f
    }
    return slices.Max(dfs(0)[0])
}

func main() {
    // Example 1:
    // Input: n = 2, present = [1,2], future = [4,3], hierarchy = [[1,2]], budget = 3
    // Output: 5
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/04/09/screenshot-2025-04-10-at-053641.png" />
    // Employee 1 buys the stock at price 1 and earns a profit of 4 - 1 = 3.
    // Since Employee 1 is the direct boss of Employee 2, Employee 2 gets a discounted price of floor(2 / 2) = 1.
    // Employee 2 buys the stock at price 1 and earns a profit of 3 - 1 = 2.
    // The total buying cost is 1 + 1 = 2 <= budget. Thus, the maximum total profit achieved is 3 + 2 = 5.
    fmt.Println(maxProfit(2, []int{1,2}, []int{4,3}, [][]int{{1,2}}, 3)) // 5
    // Example 2:
    // Input: n = 2, present = [3,4], future = [5,8], hierarchy = [[1,2]], budget = 4
    // Output: 4
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/04/09/screenshot-2025-04-10-at-053641.png" />
    // Employee 2 buys the stock at price 4 and earns a profit of 8 - 4 = 4.
    // Since both employees cannot buy together, the maximum profit is 4.
    fmt.Println(maxProfit(2, []int{3,4}, []int{5,8}, [][]int{{1,2}}, 4)) // 4
    // Example 3:
    // Input: n = 3, present = [4,6,8], future = [7,9,11], hierarchy = [[1,2],[1,3]], budget = 10
    // Output: 10
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/04/09/image.png" />
    // Employee 1 buys the stock at price 4 and earns a profit of 7 - 4 = 3.
    // Employee 3 would get a discounted price of floor(8 / 2) = 4 and earns a profit of 11 - 4 = 7.
    // Employee 1 and Employee 3 buy their stocks at a total cost of 4 + 4 = 8 <= budget. Thus, the maximum total profit achieved is 3 + 7 = 10.
    fmt.Println(maxProfit(3, []int{4,6,8}, []int{7,9,11}, [][]int{{1,2},{1,3}}, 10)) // 10
    // Example 4:
    // Input: n = 3, present = [5,2,3], future = [8,5,6], hierarchy = [[1,2],[2,3]], budget = 7
    // Output: 12
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/04/09/screenshot-2025-04-10-at-054114.png" />
    // Employee 1 buys the stock at price 5 and earns a profit of 8 - 5 = 3.
    // Employee 2 would get a discounted price of floor(2 / 2) = 1 and earns a profit of 5 - 1 = 4.
    // Employee 3 would get a discounted price of floor(3 / 2) = 1 and earns a profit of 6 - 1 = 5.
    // The total cost becomes 5 + 1 + 1 = 7 <= budget. Thus, the maximum total profit achieved is 3 + 4 + 5 = 12.
    fmt.Println(maxProfit(3, []int{5,2,3}, []int{8,5,6}, [][]int{{1,2},{2,3}}, 7)) // 12

    fmt.Println(maxProfit1(2, []int{1,2}, []int{4,3}, [][]int{{1,2}}, 3)) // 5
    fmt.Println(maxProfit1(2, []int{3,4}, []int{5,8}, [][]int{{1,2}}, 4)) // 4
    fmt.Println(maxProfit1(3, []int{4,6,8}, []int{7,9,11}, [][]int{{1,2},{1,3}}, 10)) // 10
    fmt.Println(maxProfit1(3, []int{5,2,3}, []int{8,5,6}, [][]int{{1,2},{2,3}}, 7)) // 12
}
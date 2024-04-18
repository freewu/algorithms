package main

// 309. Best Time to Buy and Sell Stock with Cooldown
// You are given an array prices where prices[i] is the price of a given stock on the ith day.

// Find the maximum profit you can achieve. 
// You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times) with the following restrictions:
//     After you sell your stock, you cannot buy stock on the next day (i.e., cooldown one day).

// Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).

// Example 1:
// Input: prices = [1,2,3,0,2]
// Output: 3
// Explanation: transactions = [buy, sell, cooldown, buy, sell]

// Example 2:
// Input: prices = [1]
// Output: 0
 
// Constraints:
//     1 <= prices.length <= 5000
//     0 <= prices[i] <= 1000

import "fmt"

// dp
func maxProfit(prices []int) int {
    l := len(prices)
    if l <= 1 {
        return 0
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    buy := []int{ -prices[0], max(-prices[0], -prices[1]), 1 << 32 -1 }
    sell := []int{0, max(0, -prices[0] + prices[1]), 0}
    for i := 2; i < l; i++ {
        // 第 i 天如果是 sell，那么这天能获得的最大收益是 buy[i - 1] + price[i - 1]，因为只有 buy 了才能 sell。
        //      如果这一天是 cooldown，那么这天能获得的最大收益还是 sell[i - 1]。
        //      所以 sell[i] 的状态转移方程 sell[i] = max(buy[i - 1] + price[i - 1], sell[i - 1])。
        //      sell[0] = 0 代表第一天就卖了，由于第一天不持有股票，所以 sell[0] = 0。
        //      sell[1] = max(sell[0], buy[0]+prices[1]) 代表第一天卖了，和第一天不卖，第二天卖做对比，钱多的保存至 sell[1]
        // 第 i 天如果是 buy，那么这天能获得的最大收益是 sell[i - 2] - price[i - 1]，
        //      因为 i - 1 天是 cooldown。如果这一天是 cooldown，那么这天能获得的最大收益还是 buy[i - 1]。
        //      所以 buy[i] 的状态转移方程 buy[i] = max(sell[i - 2] - price[i - 1], buy[i - 1])。
        //      buy[0] = -prices[0] 代表第一天就买入，所以金钱变成了负的。
        //      buy[1] = max(buy[0], -prices[1]) 代表第一天不买入，第二天再买入
        sell[i % 3] = max(sell[(i-1) % 3], buy[(i-1) % 3] + prices[i])
        buy[i % 3] = max(buy[(i-1) % 3], sell[(i-2) % 3] - prices[i])
    }
    return sell[(l-1) % 3]
}

func maxProfit1(prices []int) int {
    // 持有0只股票，手上现金最大值: dp[i][0] = maxNum(dp[i-1][0], dp[i-1][1] + prices[i])
    // 持有0只股票，冷冻期，手上现金最大值: dp[i][1] = dp[i][0]
    // 持有1只股票，手上现金最大值: dp[i][2] = maxNum(dp[i-1][1], dp[i-1][0] - prices[i])
    if len(prices) == 0 {
        return 0
    }
    profit0, profit1, profit2 := 0, 0, 0 - prices[0]
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < len(prices); i++ {
        tmp0 := profit0
        tmp1 := profit1
        profit0 = max(profit0, profit2 + prices[i])
        profit1 = tmp0
        profit2 = max(profit2, tmp1 - prices[i])
    }
    return profit0
}

func main() {
    // Explanation: transactions = [buy, sell, cooldown, buy, sell]
    fmt.Println(maxProfit([]int{1,2,3,0,2})) // 3
    fmt.Println(maxProfit([]int{1})) // 0

    fmt.Println(maxProfit1([]int{1,2,3,0,2})) // 3
    fmt.Println(maxProfit1([]int{1})) // 0
}
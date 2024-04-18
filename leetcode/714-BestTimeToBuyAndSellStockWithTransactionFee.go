package main

// 714. Best Time to Buy and Sell Stock with Transaction Fee
// You are given an array prices where prices[i] is the price of a given stock on the ith day, and an integer fee representing a transaction fee.

// Find the maximum profit you can achieve. 
// You may complete as many transactions as you like, but you need to pay the transaction fee for each transaction.

// Note:
//     You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).
//     The transaction fee is only charged once for each stock purchase and sale.
 
// Example 1:
// Input: prices = [1,3,2,8,4,9], fee = 2
// Output: 8
// Explanation: The maximum profit can be achieved by:
// - Buying at prices[0] = 1
// - Selling at prices[3] = 8
// - Buying at prices[4] = 4
// - Selling at prices[5] = 9
// The total profit is ((8 - 1) - 2) + ((9 - 4) - 2) = 8.

// Example 2:
// Input: prices = [1,3,7,5,10,3], fee = 3
// Output: 6
 
// Constraints:
//     1 <= prices.length <= 5 * 10^4
//     1 <= prices[i] < 5 * 10^4
//     0 <= fee < 5 * 10^4

import "fmt"

// DP
// 需要维护买和卖的两种状态:
//      buy[i] 代表第 i 天买入的最大收益，  buy[i] = max(buy[i-1], sell[i-1]-prices[i])
//      sell[i] 代表第 i 天卖出的最大收益，sell[i] = max(sell[i-1], buy[i-1]+prices[i]-fee)
func maxProfit(prices []int, fee int) int {
    l, inf := len(prices), 1 << 32 -1
    if l <= 1 {
        return 0
    }
    buy, sell := make([]int, l), make([]int, l)
    for i := range buy {
        buy[i] = inf
    }
    buy[0] = -prices[0]
    for i := 1; i < l; i++ {
        buy[i] = max(buy[i-1], sell[i-1] - prices[i])
        sell[i] = max(sell[i-1], buy[i-1] + prices[i] - fee)
    }
    return sell[l - 1]
}

func maxProfit1(prices []int, fee int) int {
    sell, buy := 0, -prices[0]
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < len(prices); i++ {
        sell = max(sell, buy + prices[i] - fee)
        buy = max(buy, sell - prices[i])
    }
    return sell
}

func maxProfit2(prices []int, fee int) int {
    buy, sell := -0x3f3f3f, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := range prices {
        buy = max(sell - prices[i], buy)
        sell = max(buy + prices[i] - fee, sell)
    }
    return sell
}

func main() {

    // Explanation: The maximum profit can be achieved by:
    // - Buying at prices[0] = 1
    // - Selling at prices[3] = 8
    // - Buying at prices[4] = 4
    // - Selling at prices[5] = 9
    // The total profit is ((8 - 1) - 2) + ((9 - 4) - 2) = 8.
    fmt.Println(maxProfit([]int{1,3,2,8,4,9}, 2)) // 8
    fmt.Println(maxProfit([]int{1,3,7,5,10,3}, 3)) // 6

    fmt.Println(maxProfit1([]int{1,3,2,8,4,9}, 2)) // 8
    fmt.Println(maxProfit1([]int{1,3,7,5,10,3}, 3)) // 6

    fmt.Println(maxProfit2([]int{1,3,2,8,4,9}, 2)) // 8
    fmt.Println(maxProfit2([]int{1,3,7,5,10,3}, 3)) // 6
}
package main

// 123. Best Time to Buy and Sell Stock III
// You are given an array prices where prices[i] is the price of a given stock on the ith day.
// Find the maximum profit you can achieve. You may complete at most two transactions.
// Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).

// Example 1:
// Input: prices = [3,3,5,0,0,3,1,4]
// Output: 6
// Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
// Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 = 3.

// Example 2:
// Input: prices = [1,2,3,4,5]
// Output: 4
// Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
// Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are engaging multiple transactions at the same time. You must sell before buying again.

// Example 3:
// Input: prices = [7,6,4,3,1]
// Output: 0
// Explanation: In this case, no transaction is done, i.e. max profit = 0.
 
// Constraints:
//     1 <= prices.length <= 10^5
//     0 <= prices[i] <= 10^5


import "fmt"

func maxProfit(prices []int) int {
    // The idea is to find the two profits separatly
    mp1,mp2, p1, p2  := 1 << 32 -1, 1 << 32 -1, 0, 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    // Here we are using a trick to calculate max using min function
    // pass the negative values of the items and make the result negative
    // in this way we can use min function to calculate the max as well
    for _, price := range prices {
        mp1 = min(mp1, price)
        p1 = -min(-p1, -(price - mp1))
        mp2 = min(mp2, price - p1)
        p2 = -min(-p2, -(price - mp2))
    }
    return p2
}

func maxProfit1(prices []int) int {
    buy1, buy2, sell1, sell2 := -prices[0], -prices[0], 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := range prices {
        buy1 = max(buy1, -prices[i])
        sell1 = max(sell1, prices[i] + buy1)
        buy2 = max(buy2, sell1 - prices[i])
        sell2 = max(sell2, prices[i] + buy2)
    }
    return sell2
}

func main() {
    // Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
    // Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 = 3.
    fmt.Println(maxProfit([]int{3,3,5,0,0,3,1,4})) // 6
    // Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
    // Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are engaging multiple transactions at the same time. You must sell before buying again.
    fmt.Println(maxProfit([]int{1,2,3,4,5})) // 4
    // Explanation: In this case, no transaction is done, i.e. max profit = 0.
    fmt.Println(maxProfit([]int{7,6,4,3,1})) // 0

    fmt.Println(maxProfit1([]int{3,3,5,0,0,3,1,4})) // 6
    fmt.Println(maxProfit1([]int{1,2,3,4,5})) // 4
    fmt.Println(maxProfit1([]int{7,6,4,3,1})) // 0
}
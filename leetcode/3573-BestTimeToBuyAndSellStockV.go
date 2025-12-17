package main

// 3573. Best Time to Buy and Sell Stock V
// You are given an integer array prices where prices[i] is the price of a stock in dollars on the ith day, and an integer k.

// You are allowed to make at most k transactions, where each transaction can be either of the following:

// Normal transaction: Buy on day i, then sell on a later day j where i < j. You profit prices[j] - prices[i].

// Short selling transaction: Sell on day i, then buy back on a later day j where i < j. You profit prices[i] - prices[j].

// Note that you must complete each transaction before starting another. 
// Additionally, you can't buy or sell on the same day you are selling or buying back as part of a previous transaction.

// Return the maximum total profit you can earn by making at most k transactions.

// Example 1:
// Input: prices = [1,7,9,8,2], k = 2
// Output: 14
// Explanation:
// We can make $14 of profit through 2 transactions:
// A normal transaction: buy the stock on day 0 for $1 then sell it on day 2 for $9.
// A short selling transaction: sell the stock on day 3 for $8 then buy back on day 4 for $2.

// Example 2:
// Input: prices = [12,16,19,19,8,1,19,13,9], k = 3
// Output: 36
// Explanation:
// We can make $36 of profit through 3 transactions:
// A normal transaction: buy the stock on day 0 for $12 then sell it on day 2 for $19.
// A short selling transaction: sell the stock on day 3 for $19 then buy back on day 4 for $8.
// A normal transaction: buy the stock on day 5 for $1 then sell it on day 6 for $19.

// Constraints:
//     2 <= prices.length <= 10^3
//     1 <= prices[i] <= 10^9
//     1 <= k <= prices.length / 2

import "fmt"

func maximumProfit(prices []int, k int) int64 {
    const FREE, HOLD, SHORT_SELLING = 0, 1, 2
    n := len(prices)
    dp := make([][]int64, k + 1)
    for j := 0; j <= k; j++ {
        dp[j] = make([]int64, 3)
    }
    for j := 1; j <= k; j++ {
        dp[j][HOLD] = int64(-prices[0])
        dp[j][SHORT_SELLING] = int64(prices[0])
    }
    max := func (x, y int64) int64 { if x > y { return x; }; return y; }
    for i := 1; i < n; i++ {
        for j := k; j > 0; j-- {
            dp[j][FREE] = max(dp[j][FREE], max(dp[j][HOLD] + int64(prices[i]), dp[j][SHORT_SELLING] - int64(prices[i])))
            dp[j][HOLD] = max(dp[j][HOLD], dp[j - 1][FREE] - int64(prices[i]))
            dp[j][SHORT_SELLING] = max(dp[j][SHORT_SELLING], dp[j - 1][FREE] + int64(prices[i]))
        }
    }
    return dp[k][FREE]
}

func maximumProfit1(prices []int, k int) int64 {
    res, dp := 0, make([]int, len(prices) + 1)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for t := 1; t <= k; t++ {
        buy, sell, prev := -prices[0], prices[0], dp[t-1]
        for i := t; i < len(prices); i++ {
            tmp := dp[i]
            dp[i] = max(dp[i-1], max(buy + prices[i], sell - prices[i]))
            res = max(res, dp[i])
            buy = max(buy, prev - prices[i])
            sell = max(sell, prev + prices[i])
            prev = tmp
        }
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: prices = [1,7,9,8,2], k = 2
    // Output: 14
    // Explanation:
    // We can make $14 of profit through 2 transactions:
    // A normal transaction: buy the stock on day 0 for $1 then sell it on day 2 for $9.
    // A short selling transaction: sell the stock on day 3 for $8 then buy back on day 4 for $2.
    fmt.Println(maximumProfit([]int{1,7,9,8,2}, 2)) // 14
    // Example 2:
    // Input: prices = [12,16,19,19,8,1,19,13,9], k = 3
    // Output: 36
    // Explanation:
    // We can make $36 of profit through 3 transactions:
    // A normal transaction: buy the stock on day 0 for $12 then sell it on day 2 for $19.
    // A short selling transaction: sell the stock on day 3 for $19 then buy back on day 4 for $8.
    // A normal transaction: buy the stock on day 5 for $1 then sell it on day 6 for $19.
    fmt.Println(maximumProfit([]int{12,16,19,19,8,1,19,13,9}, 3)) // 36

    fmt.Println(maximumProfit([]int{1,2,3,4,5,6,7,8,9}, 3)) // 8
    fmt.Println(maximumProfit([]int{9,8,7,6,5,4,3,2,1}, 3)) // 8

    fmt.Println(maximumProfit1([]int{1,7,9,8,2}, 2)) // 14
    fmt.Println(maximumProfit1([]int{12,16,19,19,8,1,19,13,9}, 3)) // 36
    fmt.Println(maximumProfit1([]int{1,2,3,4,5,6,7,8,9}, 3)) // 8
    fmt.Println(maximumProfit1([]int{9,8,7,6,5,4,3,2,1}, 3)) // 8
}
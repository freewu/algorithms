package main

// 4069. Best Time to Buy and Sell Stock with Cooldown II
// You are given an integer array prices of length n, where prices[i] is the price of a stock on day i.

// You are also given an integer cooldown and an integer array costs of length n. 
// The value costs[k] is the total holding fee for a transaction in which the stock is held for exactly k days.

// You may perform any number of transactions, including zero, subject to the following rules:
//     1. You may hold at most one share at a time. 
//        You must sell your current share before buying another.
//     2. If you sell on day j, the earliest day you may buy again is j + cooldown + 1.
//     3. If you buy on day i and sell on day j, the holding duration is j - i days. 
//        You pay the fee costs[j - i] once for that transaction.

// The profit from buying on day i and selling on day j is prices[j] - prices[i] - costs[j - i].

// Return the maximum total profit you can achieve. 
// If no profitable transactions are possible, return 0.

// Example 1:
// Input: prices = [1,5,3], cooldown = 1, costs = [2,0,1]
// Output: 4
// Explanation:
// The optimal strategy is:
// Buy on day 0 at prices[0] = 1.
// Sell on day 1 at prices[1] = 5. The holding duration is 1 day, so the holding cost is costs[1] = 0.
// Therefore, the profit from this transaction is 5 - 1 - 0 = 4.
// After selling, cooldown = 1 prevents buying any stock on day 2, so no further transactions are possible.
// Maximum profit achievable is 4.

// Example 2:
// Input: prices = [1,2,5], cooldown = 0, costs = [0,2,1]
// Output: 3
// Explanation:
// The optimal strategy is:
// Buy on day 0 at prices[0] = 1.
// Sell on day 2 at prices[2] = 5. The holding duration is 2 days, so the holding cost is costs[2] = 1.
// Therefore, the profit from this transaction is 5 - 1 - 1 = 3.
// No further transactions improve the profit. Maximum profit achievable is 3.

// Example 3:
// Input: prices = [4,1,7], cooldown = 2, costs = [3,0,1]
// Output: 6
// Explanation:
// The optimal strategy is:
// Buy on day 1 at prices[1] = 1.
// Sell on day 2 at prices[2] = 7. The holding duration is 1 day, so the holding cost is costs[1] = 0.
// Profit from this transaction = 7 - 1 - 0 = 6.
// No further transactions are possible. Maximum profit achievable = 6.

// Example 4:
// Input: prices = [3,1,4], cooldown = 0, costs = [2,1,0]
// Output: 2
// Explanation:
// The optimal strategy is:
// Buy on day 1 at prices[1] = 1.
// Sell on day 2 at prices[2] = 4. The holding duration is 1 day, so the holding cost is costs[1] = 1.
// Therefore, the profit from this transaction is 4 - 1 - 1 = 2.
// No further transactions improve the profit. Maximum profit achievable is 2.

// Constraints:
//     1 <= n == prices.length <= 1500
//     1 <= prices[i] <= 10^5
//     0 <= cooldown <= n - 1
//     costs.length == n
//     0 <= costs[i] <= 10^5

import "fmt"

func maxProfit(prices []int, cooldown int, costs []int) int {
    res, n := 0, len(prices)
    dpSell, dpFree := make([]int, n), make([]int, n)
    for i := range dpSell {
        dpSell[i] = -1e18
    }
    // dpFree initialized to 0 by default in go
    for j := 0; j < n; j++ {
        // sell on day j, buy on i < j
        for i := 0; i < j; i++ {
            if dpFree[i] == -1e18 {
                continue
            }
            holdDays := j - i
            profit := dpFree[i] + prices[j] - prices[i] - costs[holdDays]
            if profit > dpSell[j] {
                dpSell[j] = profit
            }
        }
        // carry free state forward to next day
        if j + 1 < n {
            if dpFree[j] > dpFree[j+1] {
                dpFree[j+1] = dpFree[j]
            }
        }
        // after sell j, cooldown ends at freeDay = j + cooldown +1
        freeDay := j + cooldown + 1
        if freeDay < n {
            if dpSell[j] > dpFree[freeDay] {
                dpFree[freeDay] = dpSell[j]
            }
        }
    }
    for _, v := range dpSell {
        if v > res {
            res = v
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: prices = [1,5,3], cooldown = 1, costs = [2,0,1]
    // Output: 4
    // Explanation:
    // The optimal strategy is:
    // Buy on day 0 at prices[0] = 1.
    // Sell on day 1 at prices[1] = 5. The holding duration is 1 day, so the holding cost is costs[1] = 0.
    // Therefore, the profit from this transaction is 5 - 1 - 0 = 4.
    // After selling, cooldown = 1 prevents buying any stock on day 2, so no further transactions are possible.
    // Maximum profit achievable is 4.
    fmt.Println(maxProfit([]int{1,5,3}, 1, []int{2,0,1})) // 4 
    // Example 2:
    // Input: prices = [1,2,5], cooldown = 0, costs = [0,2,1]
    // Output: 3
    // Explanation:
    // The optimal strategy is:
    // Buy on day 0 at prices[0] = 1.
    // Sell on day 2 at prices[2] = 5. The holding duration is 2 days, so the holding cost is costs[2] = 1.
    // Therefore, the profit from this transaction is 5 - 1 - 1 = 3.
    // No further transactions improve the profit. Maximum profit achievable is 3.
    fmt.Println(maxProfit([]int{1,2,5}, 0, []int{0,2,1})) // 3
    // Example 3:
    // Input: prices = [4,1,7], cooldown = 2, costs = [3,0,1]
    // Output: 6
    // Explanation:
    // The optimal strategy is:
    // Buy on day 1 at prices[1] = 1.
    // Sell on day 2 at prices[2] = 7. The holding duration is 1 day, so the holding cost is costs[1] = 0.
    // Profit from this transaction = 7 - 1 - 0 = 6.
    // No further transactions are possible. Maximum profit achievable = 6.
    fmt.Println(maxProfit([]int{4,1,7}, 2, []int{3,0,1})) // 6
    // Example 4:
    // Input: prices = [3,1,4], cooldown = 0, costs = [2,1,0]
    // Output: 2
    // Explanation:
    // The optimal strategy is:
    // Buy on day 1 at prices[1] = 1.
    // Sell on day 2 at prices[2] = 4. The holding duration is 1 day, so the holding cost is costs[1] = 1.
    // Therefore, the profit from this transaction is 4 - 1 - 1 = 2.
    // No further transactions improve the profit. Maximum profit achievable is 2.
    fmt.Println(maxProfit([]int{3,1,4}, 0, []int{2,1,0})) // 2

    fmt.Println(maxProfit([]int{1,2,3,4,5,6,7,8,9}, 1, []int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(maxProfit([]int{1,2,3,4,5,6,7,8,9}, 1, []int{9,8,7,6,5,4,3,2,1})) // 7
    fmt.Println(maxProfit([]int{9,8,7,6,5,4,3,2,1}, 1, []int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(maxProfit([]int{9,8,7,6,5,4,3,2,1}, 1, []int{9,8,7,6,5,4,3,2,1})) // 0
}
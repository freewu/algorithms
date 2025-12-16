package main

// 3652. Best Time to Buy and Sell Stock using Strategy
// You are given two integer arrays prices and strategy, where:
//     1. prices[i] is the price of a given stock on the ith day.
//     2. strategy[i] represents a trading action on the ith day, where:
//         2.1 -1 indicates buying one unit of the stock.
//         2.2 0 indicates holding the stock.
//         2.3 1 indicates selling one unit of the stock.

// You are also given an even integer k, and may perform at most one modification to strategy. 
// A modification consists of:
//     1. Selecting exactly k consecutive elements in strategy.
//     2. Set the first k / 2 elements to 0 (hold).
//     3. Set the last k / 2 elements to 1 (sell).

// The profit is defined as the sum of strategy[i] * prices[i] across all days.

// Return the maximum possible profit you can achieve.

// Note: There are no constraints on budget or stock ownership, so all buy and sell operations are feasible regardless of past actions.

// Example 1:
// Input: prices = [4,2,8], strategy = [-1,0,1], k = 2
// Output: 10
// Explanation:
// Modification	Strategy	Profit Calculation	Profit
// Original	[-1, 0, 1]	(-1 × 4) + (0 × 2) + (1 × 8) = -4 + 0 + 8	4
// Modify [0, 1]	[0, 1, 1]	(0 × 4) + (1 × 2) + (1 × 8) = 0 + 2 + 8	10
// Modify [1, 2]	[-1, 0, 1]	(-1 × 4) + (0 × 2) + (1 × 8) = -4 + 0 + 8	4
// Thus, the maximum possible profit is 10, which is achieved by modifying the subarray [0, 1]​​​​​​​.

// Example 2:
// Input: prices = [5,4,3], strategy = [1,1,0], k = 2
// Output: 9
// Explanation:
// Modification	Strategy	Profit Calculation	Profit
// Original	[1, 1, 0]	(1 × 5) + (1 × 4) + (0 × 3) = 5 + 4 + 0	9
// Modify [0, 1]	[0, 1, 0]	(0 × 5) + (1 × 4) + (0 × 3) = 0 + 4 + 0	4
// Modify [1, 2]	[1, 0, 1]	(1 × 5) + (0 × 4) + (1 × 3) = 5 + 0 + 3	8
// Thus, the maximum possible profit is 9, which is achieved without any modification.

// Constraints:
//     2 <= prices.length == strategy.length <= 10^5
//     1 <= prices[i] <= 10^5
//     -1 <= strategy[i] <= 1
//     2 <= k <= prices.length
//     k is even

import "fmt"

func maxProfit(prices []int, strategy []int, k int) int64 {
    n := len(prices)
    sum, sell := make([]int, n + 1), make([]int, n + 1)
    for i, v := range prices {
        sum[i+1] = sum[i] + v * strategy[i]
        sell[i+1] = sell[i] + v
    }
    res := sum[n]
    for i := k; i <= n; i++ {
        val := sum[i - k] + sum[n] - sum[i] + sell[i] - sell[i - k / 2]
        res = max(res, val)
    }
    return int64(res)
}

func maxProfit1(prices, strategy []int, k int) int64 {
    total, sum := 0, 0
    // 计算第一个窗口
    for i := 0; i < k / 2; i++ {
        p, s := prices[i], strategy[i]
        total += p * s
        sum -= p * s
    }
    for i := k / 2; i < k; i++ {
        p, s := prices[i], strategy[i]
        total += p * s
        sum += p * (1 - s)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    mx := max(sum, 0)
    for i := k; i < len(prices); i++ {
        p, s := prices[i], strategy[i]
        total += p * s
        sum += p * (1-s) - prices[i-k/2] + prices[i-k] * strategy[i-k]
        mx = max(mx, sum)
    }
    return int64(total + mx)
}

func maxProfit2(prices []int, strategy []int, k int) int64 {
    profit, mx, n := 0, 0, len(prices)
    for i := 0; i < n; i++ {
        profit += (strategy[i] * prices[i])
    }
    sum, curr, half := 0, 0, k / 2
    for i := 0; i < n; i++ {
        sum += (prices[i] * strategy[i])
        // Calculate modification effect for current sliding window end
        if i >= half {
            curr += prices[i]
        }
        // Enter the window setup
        if i >= k-1 {
            val := curr - sum
            if val > mx {
                mx = val
            }
            // Slide the window, so adjust effect
            if i - k + 1 >= 0 {
                sum -= (prices[i-k+1] * strategy[i-k+1])
                if i - k + 1 + half >= 0 {
                    curr -= (prices[i-k+1+half])
                }
            }
        }
    }
    return int64(profit + mx)
}

func main() {
    // Example 1:
    // Input: prices = [4,2,8], strategy = [-1,0,1], k = 2
    // Output: 10
    // Explanation:
    // Modification	Strategy	Profit Calculation	Profit
    // Original	[-1, 0, 1]	(-1 × 4) + (0 × 2) + (1 × 8) = -4 + 0 + 8	4
    // Modify [0, 1]	[0, 1, 1]	(0 × 4) + (1 × 2) + (1 × 8) = 0 + 2 + 8	10
    // Modify [1, 2]	[-1, 0, 1]	(-1 × 4) + (0 × 2) + (1 × 8) = -4 + 0 + 8	4
    // Thus, the maximum possible profit is 10, which is achieved by modifying the subarray [0, 1]​​​​​​​.
    fmt.Println(maxProfit([]int{4,2,8}, []int{-1,0,1}, 2)) // 10
    // Example 2:
    // Input: prices = [5,4,3], strategy = [1,1,0], k = 2
    // Output: 9
    // Explanation:
    // Modification	Strategy	Profit Calculation	Profit
    // Original	[1, 1, 0]	(1 × 5) + (1 × 4) + (0 × 3) = 5 + 4 + 0	9
    // Modify [0, 1]	[0, 1, 0]	(0 × 5) + (1 × 4) + (0 × 3) = 0 + 4 + 0	4
    // Modify [1, 2]	[1, 0, 1]	(1 × 5) + (0 × 4) + (1 × 3) = 5 + 0 + 3	8
    // Thus, the maximum possible profit is 9, which is achieved without any modification.
    fmt.Println(maxProfit([]int{5,4,3}, []int{1,1,0}, 2)) // 9

    fmt.Println(maxProfit([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1}, 2)) // 165
    fmt.Println(maxProfit([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9}, 2)) // 285
    fmt.Println(maxProfit([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1}, 2)) // 285
    fmt.Println(maxProfit([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9}, 2)) // 165

    fmt.Println(maxProfit1([]int{4,2,8}, []int{-1,0,1}, 2)) // 10
    fmt.Println(maxProfit1([]int{5,4,3}, []int{1,1,0}, 2)) // 9
    fmt.Println(maxProfit1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1}, 2)) // 165
    fmt.Println(maxProfit1([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9}, 2)) // 285
    fmt.Println(maxProfit1([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1}, 2)) // 285
    fmt.Println(maxProfit1([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9}, 2)) // 165

    fmt.Println(maxProfit2([]int{4,2,8}, []int{-1,0,1}, 2)) // 10
    fmt.Println(maxProfit2([]int{5,4,3}, []int{1,1,0}, 2)) // 9
    fmt.Println(maxProfit2([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1}, 2)) // 165
    fmt.Println(maxProfit2([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9}, 2)) // 285
    fmt.Println(maxProfit2([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1}, 2)) // 285
    fmt.Println(maxProfit2([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9}, 2)) // 165
}
package main

// 2291. Maximum Profit From Trading Stocks
// You are given two 0-indexed integer arrays of the same length present 
// and future where present[i] is the current price of the ith stock 
// and future[i] is the price of the ith stock a year in the future. 
// You may buy each stock at most once. 
// You are also given an integer budget representing the amount of money you currently have.

// Return the maximum amount of profit you can make.

// Example 1:
// Input: present = [5,4,6,2,3], future = [8,5,4,3,5], budget = 10
// Output: 6
// Explanation: One possible way to maximize your profit is to:
// Buy the 0th, 3rd, and 4th stocks for a total of 5 + 2 + 3 = 10.
// Next year, sell all three stocks for a total of 8 + 3 + 5 = 16.
// The profit you made is 16 - 10 = 6.
// It can be shown that the maximum profit you can make is 6.

// Example 2:
// Input: present = [2,2,5], future = [3,4,10], budget = 6
// Output: 5
// Explanation: The only possible way to maximize your profit is to:
// Buy the 2nd stock, and make a profit of 10 - 5 = 5.
// It can be shown that the maximum profit you can make is 5.

// Example 3:
// Input: present = [3,3,12], future = [0,3,15], budget = 10
// Output: 0
// Explanation: One possible way to maximize your profit is to:
// Buy the 1st stock, and make a profit of 3 - 3 = 0.
// It can be shown that the maximum profit you can make is 0.

// Constraints:
//     n == present.length == future.length
//     1 <= n <= 1000
//     0 <= present[i], future[i] <= 100
//     0 <= budget <= 1000

import "fmt"

func maximumProfit(present []int, future []int, budget int) int {
    res, n, dp := 0, len(present), make([]int, budget + 1)
    for i := 1; i <= budget; i++ {
        dp[i] = -1
    }
    for i := 0; i < n; i++ {
        for j := budget; j >=0 ; j-- {
            if dp[j] < 0 || j+ present[i] > budget { continue }
            if dp[j + present[i]] < dp[j] + future[i] - present[i] {
                dp[j + present[i]] = dp[j] + future[i] - present[i]
            }
            if dp[j + present[i]] > res {
                res = dp[j + present[i]]
            }
        }
    }
    return res
}

func maximumProfit1(present []int, future []int, budget int) int {
    n := len(present)
    dp := make([]int,  budget + 1)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n; i++ {
        cost := present[i]
        profit := future[i] - cost
        if profit <= 0 { continue } // no point here
        // Update the DP array in reverse to avoid reusing a stock
        for j := budget; j >= cost; j-- {
            dp[j] = max(dp[j], dp[j-cost] + profit)
        }
    }
    return dp[budget]
}

func main() {
    // Example 1:
    // Input: present = [5,4,6,2,3], future = [8,5,4,3,5], budget = 10
    // Output: 6
    // Explanation: One possible way to maximize your profit is to:
    // Buy the 0th, 3rd, and 4th stocks for a total of 5 + 2 + 3 = 10.
    // Next year, sell all three stocks for a total of 8 + 3 + 5 = 16.
    // The profit you made is 16 - 10 = 6.
    // It can be shown that the maximum profit you can make is 6.
    fmt.Println(maximumProfit([]int{5,4,6,2,3}, []int{8,5,4,3,5}, 10)) // 6
    // Example 2:
    // Input: present = [2,2,5], future = [3,4,10], budget = 6
    // Output: 5
    // Explanation: The only possible way to maximize your profit is to:
    // Buy the 2nd stock, and make a profit of 10 - 5 = 5.
    // It can be shown that the maximum profit you can make is 5.
    fmt.Println(maximumProfit([]int{2,2,5}, []int{3,4,10}, 6)) // 5
    // Example 3:
    // Input: present = [3,3,12], future = [0,3,15], budget = 10
    // Output: 0
    // Explanation: One possible way to maximize your profit is to:
    // Buy the 1st stock, and make a profit of 3 - 3 = 0.
    // It can be shown that the maximum profit you can make is 0.
    fmt.Println(maximumProfit([]int{3,3,12}, []int{0,3,15}, 10)) // 0

    fmt.Println(maximumProfit([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9}, 10)) // 0
    fmt.Println(maximumProfit([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1}, 10)) // 20
    fmt.Println(maximumProfit([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1}, 10)) // 0
    fmt.Println(maximumProfit([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9}, 10)) // 20

    fmt.Println(maximumProfit1([]int{5,4,6,2,3}, []int{8,5,4,3,5}, 10)) // 6
    fmt.Println(maximumProfit1([]int{2,2,5}, []int{3,4,10}, 6)) // 5
    fmt.Println(maximumProfit1([]int{3,3,12}, []int{0,3,15}, 10)) // 0
    fmt.Println(maximumProfit1([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9}, 10)) // 0
    fmt.Println(maximumProfit1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1}, 10)) // 20
    fmt.Println(maximumProfit1([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1}, 10)) // 0
    fmt.Println(maximumProfit1([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9}, 10)) // 20
}
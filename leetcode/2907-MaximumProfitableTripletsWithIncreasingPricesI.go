package main

// 2907. Maximum Profitable Triplets With Increasing Prices I
// Given the 0-indexed arrays prices and profits of length n. 
// There are n items in an store where the ith item has a price of prices[i] and a profit of profits[i].

// We have to pick three items with the following condition:
//     prices[i] < prices[j] < prices[k] where i < j < k.

// If we pick items with indices i, j and k satisfying the above condition, 
// the profit would be profits[i] + profits[j] + profits[k].

// Return the maximum profit we can get, and -1 if it's not possible to pick three items with the given condition.

// Example 1:
// Input: prices = [10,2,3,4], profits = [100,2,7,10]
// Output: 19
// Explanation: We can't pick the item with index i=0 since there are no indices j and k such that the condition holds.
// So the only triplet we can pick, are the items with indices 1, 2 and 3 and it's a valid pick since prices[1] < prices[2] < prices[3].
// The answer would be sum of their profits which is 2 + 7 + 10 = 19.

// Example 2:
// Input: prices = [1,2,3,4,5], profits = [1,5,3,4,6]
// Output: 15
// Explanation: We can select any triplet of items since for each triplet of indices i, j and k such that i < j < k, the condition holds.
// Therefore the maximum profit we can get would be the 3 most profitable items which are indices 1, 3 and 4.
// The answer would be sum of their profits which is 5 + 4 + 6 = 15.

// Example 3:
// Input: prices = [4,3,2,1], profits = [33,20,19,87]
// Output: -1
// Explanation: We can't select any triplet of indices such that the condition holds, so we return -1.

// Constraints:
//     3 <= prices.length == profits.length <= 2000
//     1 <= prices[i] <= 10^6
//     1 <= profits[i] <= 10^6

import "fmt"

func maxProfit(prices []int, profits []int) int {
    res, dp1, dp2 := -1, make([]int, len(prices)), make([]int, len(prices))
    for i := 0; i < len(dp1); i++ {
        dp1[i], dp2[i] = -2, -2
    }
    var dfs func(nowCount, m, n, index int) int 
    dfs = func(nowCount, m, n, index int) int {
        if n == 3 { return nowCount }
        if index >= len(prices) { return -1 }
        if n == 1 {
            if v := dp1[index]; v != -2 {
                if v == -1 { return v }
                return nowCount + v
            }
        }
        if n == 2 {
            if v := dp2[index]; v != -2 {
                if v == -1 {  return v }
                return nowCount + v
            }
        }
        res := -1
        for i := index; i < len(prices); i++ {
            if prices[i] > m {
                v := dfs(nowCount + profits[i], prices[i], n + 1, i + 1)
                if v > res {
                    res = v
                }
            }
        }
        if res == -1 {
            if n == 1 { dp1[index] = -1 }
            if n == 2 { dp2[index] = -1 }
            return -1
        }
        if n == 1 { dp1[index] = res - nowCount }
        if n == 2 { dp2[index] = res - nowCount }
        return res
    }
    for i := 0; i < len(prices) - 2; i++ {
        v := dfs(profits[i], prices[i], 1, i + 1)
        if v > res {
            res = v
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: prices = [10,2,3,4], profits = [100,2,7,10]
    // Output: 19
    // Explanation: We can't pick the item with index i=0 since there are no indices j and k such that the condition holds.
    // So the only triplet we can pick, are the items with indices 1, 2 and 3 and it's a valid pick since prices[1] < prices[2] < prices[3].
    // The answer would be sum of their profits which is 2 + 7 + 10 = 19.
    fmt.Println(maxProfit([]int{10,2,3,4}, []int{100,2,7,10})) // 19
    // Example 2:
    // Input: prices = [1,2,3,4,5], profits = [1,5,3,4,6]
    // Output: 15
    // Explanation: We can select any triplet of items since for each triplet of indices i, j and k such that i < j < k, the condition holds.
    // Therefore the maximum profit we can get would be the 3 most profitable items which are indices 1, 3 and 4.
    // The answer would be sum of their profits which is 5 + 4 + 6 = 15.
    fmt.Println(maxProfit([]int{1,2,3,4,5}, []int{1,5,3,4,6})) // 15
    // Example 3:
    // Input: prices = [4,3,2,1], profits = [33,20,19,87]
    // Output: -1
    // Explanation: We can't select any triplet of indices such that the condition holds, so we return -1.
    fmt.Println(maxProfit([]int{4,3,2,1}, []int{33,20,19,87})) // -1
}
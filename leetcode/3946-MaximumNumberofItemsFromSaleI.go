package main

// 3946. Maximum Number of Items From Sale I
// You are given a 2D integer array items, where items[i] = [factori, pricei] represents the ith item. 
// You are also given an integer budget.

// There are unlimited copies of each item available for purchase.
// You may buy any number of copies of any items such that the total cost of the purchased copies is at most budget.

// After buying items, you may receive free copies according to the following rules:
//     1. For each item i that you bought at least one copy of, you receive one free copy of every item j such that j != i and factori divides factorj.
//     2. Buying multiple copies of the same item i does not give additional free copies through item i.
//     3. The same item j can be received multiple times for free if it is received from purchases of different item types.

// Return the maximum total number of item copies you can obtain, including both purchased copies and free copies, while spending at most budget on purchased items.

// Example 1:
// Input: items = [[6,2],[2,6],[3,4]], budget = 9
// Output: 4
// Explanation:
// You can buy 2 copies of item 0 and 1 copy of item 2 for a total cost of 2 * 2 + 4 = 8, which is not greater than budget = 9.
// Buying item 2 gives 1 free copy of item 0, because factor2 = 3 divides factor0 = 6.
// You leave with 3 purchased copies and 1 free copy, for a total of 4 item copies.

// Example 2:
// Input: items = [[2,4],[3,2],[4,1],[6,4],[12,4]], budget = 8
// Output: 10
// Explanation:
// You can buy 1 copy of item 0, 1 copy of item 1, and 2 copies of item 2 for a total cost of 4 + 2 + 2 * 1 = 8.
// Buying item 0 gives 1 free copy of items 2, 3, and 4.
// Buying item 1 gives 1 free copy of items 3 and 4.
// Buying item 2 gives 1 free copy of item 4.
// Thus, you receive 6 free copies. You leave with 4 purchased copies and 6 free copies, for a total of 10 item copies.

// Constraints:
//     1 <= items.length <= 1000
//     items[i] = [factori, pricei]
//     1 <= factori, pricei <= 1500
//     1 <= budget <= 1500

import "fmt"

func maximumSaleItems(items [][]int, budget int) int {
    freq := make([]int, budget + 1)
    res, minPrice := 0, 1 << 31
    for _, p := range items {
        factor, price := p[0], p[1]
        minPrice = min(minPrice, price)
        count := 0 // 统计 factor 的倍数（包括 factor）
        for _, q := range items {
            if q[0] % factor == 0 {
                count++
            }
        }
        // 视作一个体积为 price，价值为 count 的物品
        for i := budget; i >= price; i-- {
            freq[i] = max(freq[i], freq[i - price] + count)
        }
    }
    for i, count := range freq {
        res = max(res, count + (budget - i) / minPrice)
    }
    return res
}

func maximumSaleItems1(items [][]int, budget int) int {
    n := len(items)
    count := make([]int, n)
    for i, vi := range items {
        for j, vj := range items {
            if i != j && vj[0] % vi[0] == 0 {
                count[i]++
            }
        }
    }
    dp := make([]int, budget + 1)
    for i, v := range items {
        for c := budget; c >= v[1]; c-- { // 01背包一次
            dp[c] = max(dp[c], dp[c - v[1]] + 1 + count[i])
        }
        for c := v[1]; c <= budget; c++ { // 再计算完全背包
            dp[c] = max(dp[c], dp[c - v[1]] + 1)
        }
    }
    return dp[budget]
}

func maximumSaleItems2(items [][]int, budget int) int {
    factors := make(map[int]int)
    res, maxFactor, minPrice := 0, 0, items[0][1]
    for _, item := range items {
        factor, price := item[0], item[1]
        factors[factor]++
        maxFactor, minPrice = max(maxFactor, factor), min(minPrice, price)
    }
    hasFactor := make(map[int]int)
    for factor := 1; factor <= maxFactor; factor++ {
        if factors[factor] == 0 {
            continue
        }
        hasFactor[factor] += factors[factor] - 1
        for i := factor * 2; i <= maxFactor; i += factor {
            hasFactor[factor] += factors[i]
        }
    }
    dp := make([]int, budget + 1)
    for _, item := range items {
        factor, price := item[0], item[1]
        gain := hasFactor[factor] + 1
        for b := budget; b >= price; b-- {
            dp[b] = max(dp[b], dp[b-price]+gain)
        }
    }
    for i := 0; i <= budget; i++ {
        res = max(res, dp[i]+(budget - i) / minPrice)
    }
    return res
}

func main() {
    // Example 1:
    // Input: items = [[6,2],[2,6],[3,4]], budget = 9
    // Output: 4
    // Explanation:
    // You can buy 2 copies of item 0 and 1 copy of item 2 for a total cost of 2 * 2 + 4 = 8, which is not greater than budget = 9.
    // Buying item 2 gives 1 free copy of item 0, because factor2 = 3 divides factor0 = 6.
    // You leave with 3 purchased copies and 1 free copy, for a total of 4 item copies.
    fmt.Println(maximumSaleItems([][]int{{6,2},{2,6},{3,4}}, 9)) // 4
    // Example 2:
    // Input: items = [[2,4],[3,2],[4,1],[6,4],[12,4]], budget = 8
    // Output: 10
    // Explanation:
    // You can buy 1 copy of item 0, 1 copy of item 1, and 2 copies of item 2 for a total cost of 4 + 2 + 2 * 1 = 8.
    // Buying item 0 gives 1 free copy of items 2, 3, and 4.
    // Buying item 1 gives 1 free copy of items 3 and 4.
    // Buying item 2 gives 1 free copy of item 4.
    // Thus, you receive 6 free copies. You leave with 4 purchased copies and 6 free copies, for a total of 10 item copies.
    fmt.Println(maximumSaleItems([][]int{{2,4},{3,2},{4,1},{6,4},{12,4}}, 8)) // 10

    fmt.Println(maximumSaleItems1([][]int{{6,2},{2,6},{3,4}}, 9)) // 4
    fmt.Println(maximumSaleItems1([][]int{{2,4},{3,2},{4,1},{6,4},{12,4}}, 8)) // 10

    fmt.Println(maximumSaleItems2([][]int{{6,2},{2,6},{3,4}}, 9)) // 4
    fmt.Println(maximumSaleItems2([][]int{{2,4},{3,2},{4,1},{6,4},{12,4}}, 8)) // 10
}
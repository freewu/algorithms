package main

// 322. Coin Change
// You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.
// Return the fewest number of coins that you need to make up that amount. 
// If that amount of money cannot be made up by any combination of the coins, return -1.
// You may assume that you have an infinite number of each kind of coin.

// Example 1:
// Input: coins = [1,2,5], amount = 11
// Output: 3
// Explanation: 11 = 5 + 5 + 1

// Example 2:
// Input: coins = [2], amount = 3
// Output: -1

// Example 3:
// Input: coins = [1], amount = 0
// Output: 0
 
// Constraints:
//     1 <= coins.length <= 12
//     1 <= coins[i] <= 2^31 - 1
//     0 <= amount <= 10^4

import "fmt"

// db
func coinChange(coins []int, amount int) int {
    if amount == 0 { return 0 }
    min := func (x, y int) int { if x > y { return y; }; return x; }
    dp, inf := make([]int, amount + 1), 1 << 31
    // 遍历 1 - amount 所有最小组合
    for i := 1; i <= amount; i++ {
        minCoin := inf
        for _, coin := range coins {
            if i - coin >= 0 && dp[i-coin] != -1 {
                minCoin = min(minCoin, dp[i-coin] + 1) 
            }
        }
        if minCoin == inf {
            dp[i] = -1 // 无解
        } else {
            dp[i] = minCoin
        }
    }
    return dp[amount]
}

func coinChange1(coins []int, amount int) int {
    dp, inf := make([]int, amount+1), 1 << 31
    dp[0] = 0
    min := func (x, y int) int { if x > y { return y; }; return x; }
    for i := 1; i < len(dp); i++ {
        mn := inf
        for _, c := range coins {
            if i-c >= 0 {
                mn = min(mn, dp[i-c] + 1)
            }
        }
        dp[i] = mn
    }
    if dp[amount] == inf {
        return -1
    }
    return dp[amount]
}

func main() {
    fmt.Println(coinChange([]int{1,2,5},11)) // 3  11 = 5 + 5 + 1
    fmt.Println(coinChange([]int{2},3)) // -1
    fmt.Println(coinChange([]int{1},0)) // 0
    
    fmt.Println(coinChange1([]int{1,2,5},11)) // 3  11 = 5 + 5 + 1
    fmt.Println(coinChange1([]int{2},3)) // -1
    fmt.Println(coinChange1([]int{1},0)) // 0
}
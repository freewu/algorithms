package main

// LCR 103. 零钱兑换
// 给定不同面额的硬币 coins 和一个总金额 amount。
// 编写一个函数来计算可以凑成总金额所需的最少的硬币个数。
// 如果没有任何一种硬币组合能组成总金额，返回 -1。

// 你可以认为每种硬币的数量是无限的。

// 示例 1：
// 输入：coins = [1, 2, 5], amount = 11
// 输出：3 
// 解释：11 = 5 + 5 + 1

// 示例 2：
// 输入：coins = [2], amount = 3
// 输出：-1

// 示例 3：
// 输入：coins = [1], amount = 0
// 输出：0

// 示例 4：
// 输入：coins = [1], amount = 1
// 输出：1

// 示例 5：
// 输入：coins = [1], amount = 2
// 输出：2

// 提示：
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
package main

// 1230. Toss Strange Coins
// You have some coins.  The i-th coin has a probability prob[i] of facing heads when tossed.
// Return the probability that the number of coins facing heads equals target if you toss every coin exactly once.

// Example 1:
// Input: prob = [0.4], target = 1
// Output: 0.40000

// Example 2:
// Input: prob = [0.5,0.5,0.5,0.5,0.5], target = 0
// Output: 0.03125

// Constraints:
//     1 <= prob.length <= 1000
//     0 <= prob[i] <= 1
//     0 <= target <= prob.length
//     Answers will be accepted as correct if they are within 10^-5 of the correct answer.

import "fmt"

// dp
func probabilityOfHeads(prob []float64, target int) float64 {
    n := len(prob)
    dp := make([][]float64, n + 1) //dp[i][j]：前i个硬币，正面朝上j次的概率
    for i:=0; i<=n; i++ {
        dp[i] = make([]float64, target + 1)
    }
    dp[0][0] = 1 //前 0 个硬币，正面朝上0次是一定出现的，所以概率为 1
    for i := 1; i <= n; i++ {
        for j := 0; j <= target; j++ {
            dp[i][j] = dp[i-1][j]*(1-prob[i-1]) // 前i-1个硬币正面朝上j次的概率 * 第i次背面朝上的概率
            if j > 0 {
                dp[i][j] += dp[i-1][j-1]*prob[i-1] // 加上 前i-1个硬币正面朝上j-1次的概率 * 第i次正面朝上的概率
            }
        }
    }
    return dp[n][target]
}

func main() {
    // Example 1:
    // Input: prob = [0.4], target = 1
    // Output: 0.40000
    fmt.Println(probabilityOfHeads([]float64{0.4}, 1)) //  0.40000
    // Example 2:
    // Input: prob = [0.5,0.5,0.5,0.5,0.5], target = 0
    // Output: 0.03125
    fmt.Println(probabilityOfHeads([]float64{0.5,0.5,0.5,0.5,0.5}, 0)) // 0.03125
}
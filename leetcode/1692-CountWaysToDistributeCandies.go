package main

// 1692. Count Ways to Distribute Candies
// There are n unique candies (labeled 1 through n) and k bags. 
// You are asked to distribute all the candies into the bags such that every bag has at least one candy.

// There can be multiple ways to distribute the candies. 
// Two ways are considered different if the candies in one bag in the first way are not all in the same bag in the second way. 
// The order of the bags and the order of the candies within each bag do not matter.

// For example, (1), (2,3) and (2), (1,3) are considered different because candies 2 and 3 in the bag (2,3) in the first way are not in the same bag in the second way (they are split between the bags (2) and (1,3)). 
// However, (1), (2,3) and (3,2), (1) are considered the same because the candies in each bag are all in the same bags in both ways.

// Given two integers, n and k, return the number of different ways to distribute the candies. 
// As the answer may be too large, return it modulo 10^9 + 7.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/12/16/candies-1.png" />
// Input: n = 3, k = 2
// Output: 3
// Explanation: You can distribute 3 candies into 2 bags in 3 ways:
// (1), (2,3)
// (1,2), (3)
// (1,3), (2)

// Example 2:
// Input: n = 4, k = 2
// Output: 7
// Explanation: You can distribute 4 candies into 2 bags in 7 ways:
// (1), (2,3,4)
// (1,2), (3,4)
// (1,3), (2,4)
// (1,4), (2,3)
// (1,2,3), (4)
// (1,2,4), (3)
// (1,3,4), (2)

// Example 3:
// Input: n = 20, k = 5
// Output: 206085257
// Explanation: You can distribute 20 candies into 5 bags in 1881780996 ways. 1881780996 modulo 10^9 + 7 = 206085257.
 
// Constraints:
//     1 <= k <= n <= 1000

import "fmt"

func waysToDistribute(n int, k int) int {
    dp, dp1 := make([]int, n + 1), make([]int, n+1)
    for j := 1; j <= k; j++ {
        dp1[j] = 1
        for i := j + 1; i <= n; i++ {
            dp1[i] = (dp1[i-1] * j + dp[i-1]) % 1_000_000_007
        }
        dp, dp1 = dp1, dp
    }
    return dp[n]
}

func waysToDistribute1(n int, k int) int {
    dp := make([][]int, n + 1)
    for i := range dp {
        dp[i] = make([]int, k + 1)
    }
    dp[0][0] = 1 // dps[x][0] = 0 当 x != 0 时
    for i := 1; i <= n; i++ {
        for j := 1; j <= k; j++ {
            // 递推式: n个人进k个房间, 第n个可以单独建一个房间,剩下的人建k-1个, 或者进入n-1个人建好的k个房间中的任意一个
            dp[i][j] = (dp[i-1][j-1] + j * dp[i-1][j]) % 1_000_000_007
        }
    }
    return dp[n][k]
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/12/16/candies-1.png" />
    // Input: n = 3, k = 2
    // Output: 3
    // Explanation: You can distribute 3 candies into 2 bags in 3 ways:
    // (1), (2,3)
    // (1,2), (3)
    // (1,3), (2)
    fmt.Println(waysToDistribute(3, 2)) // 3
    // Example 2:
    // Input: n = 4, k = 2
    // Output: 7
    // Explanation: You can distribute 4 candies into 2 bags in 7 ways:
    // (1), (2,3,4)
    // (1,2), (3,4)
    // (1,3), (2,4)
    // (1,4), (2,3)
    // (1,2,3), (4)
    // (1,2,4), (3)
    // (1,3,4), (2)
    fmt.Println(waysToDistribute(4, 2)) // 7
    // Example 3:
    // Input: n = 20, k = 5
    // Output: 206085257
    // Explanation: You can distribute 20 candies into 5 bags in 1881780996 ways. 1881780996 modulo 10^9 + 7 = 206085257.
    fmt.Println(waysToDistribute(20, 5)) // 206085257

    fmt.Println(waysToDistribute1(3, 2)) // 3
    fmt.Println(waysToDistribute1(4, 2)) // 7
    fmt.Println(waysToDistribute1(20, 5)) // 206085257
}
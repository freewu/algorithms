package main

// 790. Domino and Tromino Tiling
// You have two types of tiles: a 2 x 1 domino shape and a tromino shape. You may rotate these shapes.
// <img src="https://assets.leetcode.com/uploads/2021/07/15/lc-domino.jpg" />

// Given an integer n, return the number of ways to tile an 2 x n board. 
// Since the answer may be very large, return it modulo 10^9 + 7.

// In a tiling, every square must be covered by a tile. 
// Two tilings are different if and only if there are two 4-directionally adjacent cells on the board such that exactly one of the tilings has both squares occupied by a tile.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/07/15/lc-domino1.jpg" />
// Input: n = 3
// Output: 5
// Explanation: The five different ways are show above.

// Example 2:
// Input: n = 1
// Output: 1
 
// Constraints:
//     1 <= n <= 1000

import "fmt"

func numTilings(n int) int {
    dp := [4]int{}
    dp[0] = 1
    const MOD = 1e9 + 7
    for i := 1; i <= n; i++ {
        nextCombo := [4]int{}
        // When the last column is filled (:)
        nextCombo[0] = (dp[0] + dp[1] + dp[2] + dp[3]) % MOD
        // When the upper cell of the last column is empty (.)
        nextCombo[1] = (dp[2] + dp[3]) % MOD
        // When the lower cell of the last column is empty (`)
        nextCombo[2] = (dp[1] + dp[3]) % MOD
        // When the last cell is empty
        nextCombo[3] = dp[0] % MOD
        dp = nextCombo
    }
    return dp[0]
}

func main() {
    fmt.Println(numTilings(3)) // 5
    fmt.Println(numTilings(1)) // 1
}
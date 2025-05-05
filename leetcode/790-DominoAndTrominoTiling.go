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
    dp, mod := make([]int, 4), 1_000_000_007
    dp[0] = 1
    for i := 1; i <= n; i++ {
        nextCombo := make([]int, 4)
        // When the last column is filled (:)
        nextCombo[0] = (dp[0] + dp[1] + dp[2] + dp[3]) % mod
        // When the upper cell of the last column is empty (.)
        nextCombo[1] = (dp[2] + dp[3]) % mod
        // When the lower cell of the last column is empty (`)
        nextCombo[2] = (dp[1] + dp[3]) % mod
        // When the last cell is empty
        nextCombo[3] = dp[0] % mod
        dp = nextCombo
    }
    return dp[0]
}

func numTilings1(n int) int {
    if n == 1 { return 1 }
    dp := [4]int{2, 1, 1, 1}
    for i := 0;  i < (n-2); i++ {
        x := (dp[0] + dp[1] + dp[2] + dp[3]) % 1_000_000_007
        dp[0], dp[1], dp[2], dp[3] = x, dp[2] + dp[3], dp[1] + dp[3], dp[0]
    }
    return dp[0]
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/07/15/lc-domino1.jpg" />
    // Input: n = 3
    // Output: 5
    // Explanation: The five different ways are show above.
    fmt.Println(numTilings(3)) // 5
    // Example 2:
    // Input: n = 1
    // Output: 1
    fmt.Println(numTilings(1)) // 1

    fmt.Println(numTilings(2)) // 2
    fmt.Println(numTilings(999)) // 326038248
    fmt.Println(numTilings(1000)) // 979232805

    fmt.Println(numTilings1(3)) // 5
    fmt.Println(numTilings1(1)) // 1
    fmt.Println(numTilings1(2)) // 2
    fmt.Println(numTilings1(999)) // 326038248
    fmt.Println(numTilings1(1000)) // 979232805
}
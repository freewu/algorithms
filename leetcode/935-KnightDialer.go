package main

// 935. Knight Dialer
// The chess knight has a unique movement, it may move two squares vertically and one square horizontally, 
// or two squares horizontally and one square vertically (with both forming the shape of an L). 
// The possible movements of chess knight are shown in this diagram:
// <img src="https://assets.leetcode.com/uploads/2020/08/18/chess.jpg" />

// A chess knight can move as indicated in the chess diagram below:
// <img src="https://assets.leetcode.com/uploads/2020/08/18/phone.jpg" />

// We have a chess knight and a phone pad as shown below, the knight can only stand on a numeric cell (i.e. blue cell).

// Given an integer n, return how many distinct phone numbers of length n we can dial.

// You are allowed to place the knight on any numeric cell initially and then you should perform n - 1 jumps to dial a number of length n. 
// All jumps should be valid knight jumps.

// As the answer may be very large, return the answer modulo 10^9 + 7.

// Example 1:
// Input: n = 1
// Output: 10
// Explanation: We need to dial a number of length 1, so placing the knight over any numeric cell of the 10 cells is sufficient.

// Example 2:
// Input: n = 2
// Output: 20
// Explanation: All the valid number we can dial are [04, 06, 16, 18, 27, 29, 34, 38, 40, 43, 49, 60, 61, 67, 72, 76, 81, 83, 92, 94]

// Example 3:
// Input: n = 3131
// Output: 136006598
// Explanation: Please take care of the mod.

// Constraints:
//     1 <= n <= 5000

import "fmt"

// dfs + memo
func knightDialer(n int) int {
    // All the different places knight can jump from given location on keypad
    // From 0 can jump to {4,6} , from 1 can jump to {8,6} ....
    graph := [][]int{{4, 6}, {6, 8}, {7, 9}, {4, 8}, {3, 9, 0}, {}, {1, 7, 0}, {2, 6}, {1, 3}, {2, 4}}
    res, dp, mod := 0,  make([][]int, n + 1), 1_000_000_007
    for i := range dp {
        dp[i] = make([]int, 10)
        for j:= range dp[i] {
            dp[i][j] = -1
        }
    }
    var dfs func(n, cur int) int 
    dfs = func(n, cur int) int {
        if n == 0 { return 1 }
        if dp[n][cur] != -1 {
            return dp[n][cur]
        }
        res := 0
        for _, v := range graph[cur] {
            res = (res + dfs(n-1, v)) % mod
        }
        dp[n][cur] = res
        return res
    }
    for i := 0; i <= 9; i++ {
        res = (res + dfs(n - 1, i)) % mod
    }
    return res
}

func knightDialer1(n int) int {
    mod := 1_000_000_007
    if n == 1 {
        return 10
    }
    a, b, c, d := 4, 2, 2, 1
    for i := 2; i <= n; i++ {
        an, bn, cn, dn := (b*2 + c*2) % mod, a, (a + d*2) % mod, c
        a, b, c, d = an, bn, cn, dn
    }
    return (a + b + c + d) % mod
}

func main() {
    // Example 1:
    // Input: n = 1
    // Output: 10
    // Explanation: We need to dial a number of length 1, so placing the knight over any numeric cell of the 10 cells is sufficient.
    fmt.Println(knightDialer(1)) // 10
    // Example 2:
    // Input: n = 2
    // Output: 20
    // Explanation: All the valid number we can dial are [04, 06, 16, 18, 27, 29, 34, 38, 40, 43, 49, 60, 61, 67, 72, 76, 81, 83, 92, 94]
    fmt.Println(knightDialer(2)) // 20
    // Example 3:
    // Input: n = 3131
    // Output: 136006598
    // Explanation: Please take care of the mod.
    fmt.Println(knightDialer(3131)) // 136006598
    fmt.Println(knightDialer(5000)) // 406880451

    fmt.Println(knightDialer1(1)) // 10
    fmt.Println(knightDialer1(2)) // 20
    fmt.Println(knightDialer1(3131)) // 136006598
    fmt.Println(knightDialer1(5000)) // 406880451
}
package main

// 1269. Number of Ways to Stay in the Same Place After Some Steps
// You have a pointer at index 0 in an array of size arrLen. 
// At each step, you can move 1 position to the left, 1 position to the right in the array, 
// or stay in the same place (The pointer should not be placed outside the array at any time).

// Given two integers steps and arrLen, 
// return the number of ways such that your pointer is still at index 0 after exactly steps steps. 
// Since the answer may be too large, return it modulo 10^9 + 7.

// Example 1:
// Input: steps = 3, arrLen = 2
// Output: 4
// Explanation: There are 4 differents ways to stay at index 0 after 3 steps.
// Right, Left, Stay
// Stay, Right, Left
// Right, Stay, Left
// Stay, Stay, Stay

// Example 2:
// Input: steps = 2, arrLen = 4
// Output: 2
// Explanation: There are 2 differents ways to stay at index 0 after 2 steps
// Right, Left
// Stay, Stay

// Example 3:
// Input: steps = 4, arrLen = 2
// Output: 8

// Constraints:
//     1 <= steps <= 500
//     1 <= arrLen <= 10^6

import "fmt"

func numWays(steps int, arrLen int) int {
    // you can only go as right as much you have steps left or the arr size
    min := func (x, y int) int { if x < y { return x; }; return y; }
    maxRange, mod := min(steps, arrLen) - 1, 1_000_000_007
    var count func(curr int, steps int, dp [][]int) int
    count = func(curr int, steps int, dp [][]int) int {
        if steps == 0 {
            if curr == 0 { return 1 }
            return 0
        }
        if dp[curr][steps] != -1 { return dp[curr][steps] }
        res := count(curr, steps - 1, dp) // staying at the same position
        if curr > 0 { // go left
            res = (res + count(curr - 1, steps - 1, dp)) % mod
        }
        if curr < maxRange { // go right
            res = (res + count(curr + 1, steps - 1, dp)) % mod
        }
        dp[curr][steps] = res
        return res
    }
    dp := make([][]int, maxRange + 10)
    for i := range dp {
        dp[i] = make([]int, steps+1)
        for j := range dp[i] {
            dp[i][j] = -1
        }
    }
    return count(0, steps, dp)
}

func numWays1(steps int, arrLen int) int {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    mx, mod := min( arrLen - 1, steps / 2), 1_000_000_007
    dp := make([][]int, steps + 1)
    for i := range dp {
        dp[i] = make([]int, mx + 1)
    }
    dp[steps][0] = 1
    for i := steps - 1; i >= 0; i-- {
        for j := mx; j >= 0; j-- {
            dp[i][j] = dp[i+1][j]
            if j < mx {
                dp[i][j] = (dp[i][j] + dp[i+1][j+1]) % mod
            }
            if j > 0 {
                dp[i][j] = (dp[i][j] + dp[i+1][j-1]) % mod
            }
        }
    }
    return dp[0][0]
}

func main() {
    // Example 1:
    // Input: steps = 3, arrLen = 2
    // Output: 4
    // Explanation: There are 4 differents ways to stay at index 0 after 3 steps.
    // Right, Left, Stay
    // Stay, Right, Left
    // Right, Stay, Left
    // Stay, Stay, Stay
    fmt.Println(numWays(3,2)) // 4
    // Example 2:
    // Input: steps = 2, arrLen = 4
    // Output: 2
    // Explanation: There are 2 differents ways to stay at index 0 after 2 steps
    // Right, Left
    // Stay, Stay
    fmt.Println(numWays(2,4)) // 2
    // Example 3:
    // Input: steps = 4, arrLen = 2
    // Output: 8
    fmt.Println(numWays(4,2)) // 8

    fmt.Println(numWays1(3,2)) // 4
    fmt.Println(numWays1(2,4)) // 2
    fmt.Println(numWays1(4,2)) // 8
}
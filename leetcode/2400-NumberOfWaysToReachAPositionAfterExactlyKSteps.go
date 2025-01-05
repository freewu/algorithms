package main

// 2400. Number of Ways to Reach a Position After Exactly k Steps
// You are given two positive integers startPos and endPos. 
// Initially, you are standing at position startPos on an infinite number line.
// With one step, you can move either one position to the left, or one position to the right.

// Given a positive integer k, return the number of different ways to reach the position endPos starting from startPos, such that you perform exactly k steps. 
// Since the answer may be very large, return it modulo 10^9 + 7.

// Two ways are considered different if the order of the steps made is not exactly the same.

// Note that the number line includes negative integers.

// Example 1:
// Input: startPos = 1, endPos = 2, k = 3
// Output: 3
// Explanation: We can reach position 2 from 1 in exactly 3 steps in three ways:
// - 1 -> 2 -> 3 -> 2.
// - 1 -> 2 -> 1 -> 2.
// - 1 -> 0 -> 1 -> 2.
// It can be proven that no other way is possible, so we return 3.

// Example 2:
// Input: startPos = 2, endPos = 5, k = 10
// Output: 0
// Explanation: It is impossible to reach position 5 from position 2 in exactly 10 steps.
 
// Constraints:
//     1 <= startPos, endPos, k <= 1000

import "fmt"

func numberOfWays(startPos int, endPos int, k int) int {
    // Bottom-up DP with clever indexing
    offset, mod := 500, 1_000_000_007
    curr, next := make([]int, 2002), make([]int, 2002)
    startPos += offset
    endPos += offset
    curr[startPos] = 1
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i <= k; i++ {
        l, r := max(startPos - i, endPos - (k - i)), min(startPos + i, endPos + (k - i))
        for j := l; j <= r; j++ {
            next[j] = (curr[j-1] + curr[j+1]) % mod
        }
        curr, next = next, curr
    }
    return curr[endPos]
}

func numberOfWays1(startPos int, endPos int, k int) int {
    mod := 1_000_000_007
    facts := make([][]int, k + 1)
    for i := range facts {
        facts[i] = make([]int, k + 1)
        for j := range facts[i] {
            facts[i][j] = -1
        }
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    var dfs func(i, j int) int
    dfs = func(i, j int) int {
        if i > j || j < 0 { return 0 }
        if j == 0 {
            if i == 0 { return 1 }
            return 0
        }
        if facts[i][j] != -1 { return facts[i][j] }
        facts[i][j] = (dfs(i + 1, j - 1) + dfs(abs(i - 1), j - 1)) % mod
        return facts[i][j]
    }
    return dfs(abs(startPos - endPos), k)
}

func main() {
    // Example 1:
    // Input: startPos = 1, endPos = 2, k = 3
    // Output: 3
    // Explanation: We can reach position 2 from 1 in exactly 3 steps in three ways:
    // - 1 -> 2 -> 3 -> 2.
    // - 1 -> 2 -> 1 -> 2.
    // - 1 -> 0 -> 1 -> 2.
    // It can be proven that no other way is possible, so we return 3.
    fmt.Println(numberOfWays(1,2,3)) // 3
    // Example 2:
    // Input: startPos = 2, endPos = 5, k = 10
    // Output: 0
    // Explanation: It is impossible to reach position 5 from position 2 in exactly 10 steps.
    fmt.Println(numberOfWays(2,5,10)) // 0

    fmt.Println(numberOfWays(1, 1, 1)) // 0
    fmt.Println(numberOfWays(1, 1000, 1000)) // 0
    fmt.Println(numberOfWays(1000, 1, 1000)) // 0
    fmt.Println(numberOfWays(1000, 1000, 1)) // 0
    fmt.Println(numberOfWays(1, 1, 1000)) // 159835829
    fmt.Println(numberOfWays(1000, 1, 1)) // 0
    fmt.Println(numberOfWays(1000, 1000, 1000)) // 159835829

    fmt.Println(numberOfWays1(1,2,3)) // 3
    fmt.Println(numberOfWays1(2,5,10)) // 0
    fmt.Println(numberOfWays1(1, 1, 1)) // 0
    fmt.Println(numberOfWays1(1, 1000, 1000)) // 0
    fmt.Println(numberOfWays1(1000, 1, 1000)) // 0
    fmt.Println(numberOfWays1(1000, 1000, 1)) // 0
    fmt.Println(numberOfWays1(1, 1, 1000)) // 159835829
    fmt.Println(numberOfWays1(1000, 1, 1)) // 0
    fmt.Println(numberOfWays1(1000, 1000, 1000)) // 159835829
}
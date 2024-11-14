package main

// 1866. Number of Ways to Rearrange Sticks With K Sticks Visible
// There are n uniquely-sized sticks whose lengths are integers from 1 to n. 
// You want to arrange the sticks such that exactly k sticks are visible from the left. 
// A stick is visible from the left if there are no longer sticks to the left of it.
//     For example, if the sticks are arranged [1,3,2,5,4], 
//     then the sticks with lengths 1, 3, and 5 are visible from the left.

// Given n and k, return the number of such arrangements. 
// Since the answer may be large, return it modulo 10^9 + 7.

// Example 1:
// Input: n = 3, k = 2
// Output: 3
// Explanation: [1,3,2], [2,3,1], and [2,1,3] are the only arrangements such that exactly 2 sticks are visible.
// The visible sticks are underlined.

// Example 2:
// Input: n = 5, k = 5
// Output: 1
// Explanation: [1,2,3,4,5] is the only arrangement such that all 5 sticks are visible.
// The visible sticks are underlined.

// Example 3:
// Input: n = 20, k = 11
// Output: 647427950
// Explanation: There are 647427950 (mod 10^9 + 7) ways to rearrange the sticks such that exactly 11 sticks are visible.

// Constraints:
//     1 <= n <= 1000
//     1 <= k <= n

import "fmt"

func rearrangeSticks(n int, k int) int {
    if n == k { return 1 }
    if k > n  { return 0 }
    dp, mod := make([][]int, n + 1), 1_000_000_007
    for i := n; i >= 0; i-- {
        dp[i] = make([]int, k + 1)
    }
    dp[1][1] = 1
    for i := 2; i <= n; i++ {
        for j := 1; j <= k; j++ {
            dp[i][j] = (dp[i-1][j-1] + (i-1) * dp[i-1][j]) % mod
        }
    }
    return dp[n][k] % mod
}

func main() {
    // Example 1:
    // Input: n = 3, k = 2
    // Output: 3
    // Explanation: [1,3,2], [2,3,1], and [2,1,3] are the only arrangements such that exactly 2 sticks are visible.
    // The visible sticks are underlined.
    fmt.Println(rearrangeSticks(3, 2)) // 3
    // Example 2:
    // Input: n = 5, k = 5
    // Output: 1
    // Explanation: [1,2,3,4,5] is the only arrangement such that all 5 sticks are visible.
    // The visible sticks are underlined.
    fmt.Println(rearrangeSticks(5, 5)) // 1
    // Example 3:
    // Input: n = 20, k = 11
    // Output: 647427950
    // Explanation: There are 647427950 (mod 10^9 + 7) ways to rearrange the sticks such that exactly 11 sticks are visible.
    fmt.Println(rearrangeSticks(20, 11)) // 647427950

    fmt.Println(rearrangeSticks(1, 1)) // 1
    fmt.Println(rearrangeSticks(1000, 1000)) // 1
    fmt.Println(rearrangeSticks(1, 1000)) // 0
    fmt.Println(rearrangeSticks(1000, 1)) // 756641425
}
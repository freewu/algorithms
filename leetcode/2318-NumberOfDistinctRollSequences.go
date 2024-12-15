package main

// 2318. Number of Distinct Roll Sequences
// You are given an integer n. You roll a fair 6-sided dice n times. 
// Determine the total number of distinct sequences of rolls possible such that the following conditions are satisfied:
//     1. The greatest common divisor of any adjacent values in the sequence is equal to 1.
//     2. There is at least a gap of 2 rolls between equal valued rolls. 
//        More formally, if the value of the ith roll is equal to the value of the jth roll, then abs(i - j) > 2.

// Return the total number of distinct sequences possible. 
// Since the answer may be very large, return it modulo 10^9 + 7.

// Two sequences are considered distinct if at least one element is different.

// Example 1:
// Input: n = 4
// Output: 184
// Explanation: Some of the possible sequences are (1, 2, 3, 4), (6, 1, 2, 3), (1, 2, 3, 1), etc.
// Some invalid sequences are (1, 2, 1, 3), (1, 2, 3, 6).
// (1, 2, 1, 3) is invalid since the first and third roll have an equal value and abs(1 - 3) = 2 (i and j are 1-indexed).
// (1, 2, 3, 6) is invalid since the greatest common divisor of 3 and 6 = 3.
// There are a total of 184 distinct sequences possible, so we return 184.

// Example 2:
// Input: n = 2
// Output: 22
// Explanation: Some of the possible sequences are (1, 2), (2, 1), (3, 2).
// Some invalid sequences are (3, 6), (2, 4) since the greatest common divisor is not equal to 1.
// There are a total of 22 distinct sequences possible, so we return 22.

// Constraints:
//     1 <= n <= 10^4

import "fmt"

func distinctSequences(n int) int {
    if n == 1 { return 6 }
    res, mod := 0, 1_000_000_007
    dp := make([][][]int, n + 1)
    for k := range dp {
        dp[k] = make([][]int, 6)
        for i := range dp[k] {
            dp[k][i] = make([]int, 6)
        }
    }
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    for i := 0; i < 6; i++ {
        for j := 0; j < 6; j++ {
            if gcd(i+1, j+1) == 1 && i != j {
                dp[2][i][j] = 1
            }
        }
    }
    for k := 3; k <= n; k++ {
        for i := 0; i < 6; i++ {
            for j := 0; j < 6; j++ {
                if gcd(i+1, j+1) == 1 && i != j {
                    for h := 0; h < 6; h++ {
                        if gcd(h+1, i+1) == 1 && h != i && h != j {
                            dp[k][i][j] = (dp[k][i][j] + dp[k-1][h][i]) % mod
                        }
                    }
                }
            }
        }
    }
    for i := 0; i < 6; i++ {
        for j := 0; j < 6; j++ {
            res = (res + dp[n][i][j]) % mod
        }
    }
    return res
}

func distinctSequences1(n int) int {
    if n == 1 { return 6 }
    memo := make([][][]int, n+1)
    for i := 0; i < len(memo); i++ {
        memo[i] = make([][]int, 7)
        for j := 0; j < len(memo[i]); j++ {
            memo[i][j] = make([]int, 7)
            for k := 0; k < len(memo[i][j]); k++ {
                memo[i][j][k] = -1
            }
        }
    }
    var dfs func(idx int, pre1 int, pre2 int) int
    dfs = func(idx int, pre1 int, pre2 int) int {
        if idx == n { return 1 }
        if idx > 0 && memo[idx][pre1][pre2] != -1 { return memo[idx][pre1][pre2] }
        res := 0
        for i := 1; i <= 6; i++ {
            if idx > 0 && (i == pre1 || i == pre2) { continue }
            if idx > 0 && ((i%2 == 0 && pre1%2 == 0) || (i == 3 && pre1 == 6) || (i == 6 && pre1 == 3)) { continue }
            res = (res + dfs(idx+1, i, pre1)) % 1_000_000_007
        }
        memo[idx][pre1][pre2] = res
        return res
    }
    return dfs(0, 0, 0)
}

func main() {
    // Example 1:
    // Input: n = 4
    // Output: 184
    // Explanation: Some of the possible sequences are (1, 2, 3, 4), (6, 1, 2, 3), (1, 2, 3, 1), etc.
    // Some invalid sequences are (1, 2, 1, 3), (1, 2, 3, 6).
    // (1, 2, 1, 3) is invalid since the first and third roll have an equal value and abs(1 - 3) = 2 (i and j are 1-indexed).
    // (1, 2, 3, 6) is invalid since the greatest common divisor of 3 and 6 = 3.
    // There are a total of 184 distinct sequences possible, so we return 184.
    fmt.Println(distinctSequences(4)) // 184
    // Example 2:
    // Input: n = 2
    // Output: 22
    // Explanation: Some of the possible sequences are (1, 2), (2, 1), (3, 2).
    // Some invalid sequences are (3, 6), (2, 4) since the greatest common divisor is not equal to 1.
    // There are a total of 22 distinct sequences possible, so we return 22.
    fmt.Println(distinctSequences(2)) // 22

    fmt.Println(distinctSequences(1)) // 6
    fmt.Println(distinctSequences(8)) // 11672
    fmt.Println(distinctSequences(1024)) // 17407727
    fmt.Println(distinctSequences(9999)) // 455330915
    fmt.Println(distinctSequences(100000)) // 507965010

    fmt.Println(distinctSequences1(4)) // 184
    fmt.Println(distinctSequences1(2)) // 22
    fmt.Println(distinctSequences1(1)) // 6
    fmt.Println(distinctSequences1(8)) // 11672
    fmt.Println(distinctSequences1(1024)) // 17407727
    fmt.Println(distinctSequences1(9999)) // 455330915
    fmt.Println(distinctSequences1(100000)) // 507965010
}
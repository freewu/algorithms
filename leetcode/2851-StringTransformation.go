package main

// 2851. String Transformation
// You are given two strings s and t of equal length n. 
// You can perform the following operation on the string s:
//     Remove a suffix of s of length l where 0 < l < n and append it at the start of s.
//     For example, let s = 'abcd' then in one operation you can remove the suffix 'cd' and append it in front of s making s = 'cdab'.

// You are also given an integer k. 
// Return the number of ways in which s can be transformed into t in exactly k operations.

// Since the answer can be large, return it modulo 10^9 + 7.

// Example 1:
// Input: s = "abcd", t = "cdab", k = 2
// Output: 2
// Explanation: 
// First way:
// In first operation, choose suffix from index = 3, so resulting s = "dabc".
// In second operation, choose suffix from index = 3, so resulting s = "cdab".
// Second way:
// In first operation, choose suffix from index = 1, so resulting s = "bcda".
// In second operation, choose suffix from index = 1, so resulting s = "cdab".

// Example 2:
// Input: s = "ababab", t = "ababab", k = 1
// Output: 2
// Explanation: 
// First way:
// Choose suffix from index = 2, so resulting s = "ababab".
// Second way:
// Choose suffix from index = 4, so resulting s = "ababab".

// Constraints:
//     2 <= s.length <= 5 * 10^5
//     1 <= k <= 10^15
//     s.length == t.length
//     s and t consist of only lowercase English alphabets.

import "fmt"

func numberOfWays(s string, t string, k int64) int {
    n := len(t)
    x := s + s[:n - 1]
    kmp_table := func(W string) []int {
        T := make([]int, len(W)+1)
        pos, cnd := 1, 0
        T[0] = -1
        for pos < len(W) {
            if W[pos] == W[cnd] {
                T[pos] = T[cnd]
            } else {
                T[pos] = cnd
                for cnd >= 0 && W[pos] != W[cnd] {
                    cnd = T[cnd]
                }
            }
            pos++
            cnd++
        }
        T[pos] = cnd
        return T
    }
    kmp_search := func(S string, W string) ([]int, int) {
        j, k, T := 0, 0, kmp_table(W)
        T[0] = -1
        np, p := 0, []int{}
        for j < len(S) {
            if W[k] == S[j] {
                j++
                k++
                if k == len(W) {
                    p = append(p, j-k)
                    np++
                    k = T[k]
                }
    
            } else {
                k = T[k]
                if k < 0 {
                    j++
                    k++
                }
            }
        }
        return p, np
    }    
    _, good := kmp_search(x,t)
    bad := n - good
    if good == 0 { return 0 }
    same := 0
    if s == t {
        same++
    }
    mul := func(A [][]int, B [][]int) [][]int {
        n, l, m, mod := len(A), len(B), len(B[0]), 1_000_000_007
        C := make([][]int, n)
        for i := range C {
            C[i] = make([]int, m)
        }
        for i := 0; i < n; i++ {
            for j := 0; j < m; j++ {
                for k := 0; k < l; k++ {
                    C[i][j] = (C[i][j] + A[i][k] * B[k][j]) % mod
                }
            }
        }
        return C
    }
    exp_power := func(matrix [][]int, n int64) [][]int {
        res := [][]int{{1,0},{0,1}}
        for n > 0 {
            if n % 2 == 1 {
                res = mul(res, matrix)
            }
            n >>= 1
            matrix = mul(matrix, matrix)
        }
        return res
    }
    dp := [][]int{{ same }, {1 - same}}
    mat := exp_power([][]int{{good - 1, good},{ bad, bad - 1}}, k)
    dp = mul(mat, dp)
    return dp[0][0]
}

func main() {
    // Example 1:
    // Input: s = "abcd", t = "cdab", k = 2
    // Output: 2
    // Explanation: 
    // First way:
    // In first operation, choose suffix from index = 3, so resulting s = "dabc".
    // In second operation, choose suffix from index = 3, so resulting s = "cdab".
    // Second way:
    // In first operation, choose suffix from index = 1, so resulting s = "bcda".
    // In second operation, choose suffix from index = 1, so resulting s = "cdab".
    fmt.Println(numberOfWays("abcd", "cdab", 2)) // 2
    // Example 2:
    // Input: s = "ababab", t = "ababab", k = 1
    // Output: 2
    // Explanation: 
    // First way:
    // Choose suffix from index = 2, so resulting s = "ababab".
    // Second way:
    // Choose suffix from index = 4, so resulting s = "ababab".
    fmt.Println(numberOfWays("ababab", "ababab", 1)) // 2

    fmt.Println(numberOfWays("bluefrog", "leetcode", 1)) // 0
}
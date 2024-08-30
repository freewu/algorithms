package main

// 903. Valid Permutations for DI Sequence
// You are given a string s of length n where s[i] is either:
//     'D' means decreasing, or
//     'I' means increasing.

// A permutation perm of n + 1 integers of all the integers in the range [0, n] is called a valid permutation if for all valid i:
//     If s[i] == 'D', then perm[i] > perm[i + 1], and
//     If s[i] == 'I', then perm[i] < perm[i + 1].

// Return the number of valid permutations perm. Since the answer may be large, return it modulo 10^9 + 7.

// Example 1:
// Input: s = "DID"
// Output: 5
// Explanation: The 5 valid permutations of (0, 1, 2, 3) are:
// (1, 0, 3, 2)
// (2, 0, 3, 1)
// (2, 1, 3, 0)
// (3, 0, 2, 1)
// (3, 1, 2, 0)

// Example 2:
// Input: s = "D"
// Output: 1

// Constraints:
//     n == s.length
//     1 <= n <= 200
//     s[i] is either 'I' or 'D'.

import "fmt"

// Time Limit Exceeded
// func numPermsDISequence(s string) int {
//     res, n := 0, len(s)
//     seen := make([]int, n + 1)

//     var dfs func(j, p int) int 
//     dfs = func(j, p int) int {
//         if j == n {
//             return 1
//         }
//         res := 0
//         if s[j] == 'D' {
//             for i := p - 1; i >= 0; i-- {
//                 if seen[i] == 1 { continue }
//                 seen[i] = 1
//                 res += dfs(j + 1, i)
//                 seen[i] = 0
//             }
//         } else {
//             for i := p + 1; i <= n; i++ {
//                 if seen[i] == 1 { continue }
//                 seen[i] = 1;
//                 res += dfs(j+1, i)
//                 seen[i] = 0
//             }
//         }
//         return res
//     }
//     for i := 0; i <= n; i++ {
//         seen[i] = 1
//         res += dfs(0, i)
//         seen[i] = 0
//     }
//     return res
// }

func numPermsDISequence(s string) int {
    res, n, mod := 0, len(s), 1_000_000_007
    dp := make([][]int, 2)
    for i := range dp {
        dp[i] = make([]int, n + 1)
    }
    dp[0][0] = 1
    for i := 1; i <= n; i++ {
        for j := 0; j <= i; j++ {
            l, r := 0, j // s[i-1] == 'I'
            if s[i-1] == 'D' {
                l, r = j, i
            }
            dp[1][j] = 0
            for k := l; k < r; k++ {
                dp[1][j] += dp[0][k]
                dp[1][j] %= mod
            }
        }
        dp[0], dp[1] = dp[1], dp[0]
    }
    for j := range dp[0] {
        res += dp[0][j]
        res %= mod
    }
    return res
}

func numPermsDISequence1(s string) int {
    res, n, mod := 0, len(s), 1_000_000_007
    dp := make([][]int, n + 1)
    for i := range dp {
        dp[i] = make([]int, n + 1)
    }
    dp[0][0] = 1
    // dp[i][j]   = dp[i-1][j]   + ... + dp[i-1][i-1]
    // dp[i][j+1] = dp[i-1][j+1] + ... + dp[i-1][i-1]
    // dp[i][j]   = dp[i][j+1] + dp[i-1][j]
    for i := 1; i <= n; i++ {
        if s[i-1] == 'D' {
            for j := i - 1; j >= 0; j-- {
                dp[i][j] = (dp[i][j+1] + dp[i-1][j]) % mod
            }
        } else { // s[i-1] == 'I'
            for j := 1; j <= i; j++ {
                dp[i][j] = (dp[i][j-1] + dp[i-1][j-1]) % mod
            }
        }
    }
    for i := 0; i <= n; i++ {
        res = (res + dp[n][i]) % mod
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "DID"
    // Output: 5
    // Explanation: The 5 valid permutations of (0, 1, 2, 3) are:
    // (1, 0, 3, 2)
    // (2, 0, 3, 1)
    // (2, 1, 3, 0)
    // (3, 0, 2, 1)
    // (3, 1, 2, 0)
    fmt.Println(numPermsDISequence("DID")) // 5
    // Example 2:
    // Input: s = "D"
    // Output: 1
    fmt.Println(numPermsDISequence("D")) // 1
 
    fmt.Println(numPermsDISequence("IDDDIIDIIIIIIIIDIDID")) // 853197538

    fmt.Println(numPermsDISequence1("DID")) // 5
    fmt.Println(numPermsDISequence1("D")) // 1
    fmt.Println(numPermsDISequence1("IDDDIIDIIIIIIIIDIDID")) // 853197538
}
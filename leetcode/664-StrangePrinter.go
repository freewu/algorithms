package main

// 664. Strange Printer
// There is a strange printer with the following two special properties:
//     The printer can only print a sequence of the same character each time.
//     At each turn, the printer can print new characters starting from and ending at any place and will cover the original existing characters.

// Given a string s, return the minimum number of turns the printer needed to print it.

// Example 1:
// Input: s = "aaabbb"
// Output: 2
// Explanation: Print "aaa" first and then print "bbb".

// Example 2:
// Input: s = "aba"
// Output: 2
// Explanation: Print "aaa" first and then print "b" from the second place of the string, which will cover the existing character 'a'.
 
// Constraints:
//     1 <= s.length <= 100
//     s consists of lowercase English letters.

import "fmt"

func strangePrinter(s string) int {
    n := len(s)
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := n - 1; i > -1; i-- {
        dp[i][i] = 1
        for j := i + 1; j < n; j++ {
            dp[i][j] = dp[i][j-1] + 1
            for k := i; k < j; k++ {
                if s[k] == s[j] {
                    if k+1 <= j-1 {
                        dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j-1])
                    } else {
                        dp[i][j] = min(dp[i][j], dp[i][k])
                    }
                }
            }
        }
    }
    return dp[0][n-1]
}

func strangePrinter1(s string) int {
    n, inf := len(s), 1 << 32 - 1
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := n-1; i >= 0; i-- {
        dp[i][i]=1
        for j:= i+1; j < n; j++ {
            if s[i] == s[j] {
                dp[i][j] = dp[i][j-1]
            } else {
                dp[i][j] = inf
                for k := i; k < j; k++ {
                    dp[i][j] = min(dp[i][j], dp[i][k] + dp[k+1][j])
                }
            }
        }
    }
    return dp[0][n-1]
}

func main() {
    // Example 1:
    // Input: s = "aaabbb"
    // Output: 2
    // Explanation: Print "aaa" first and then print "bbb".
    fmt.Println(strangePrinter("aaabbb")) // 2
    // Example 2:
    // Input: s = "aba"
    // Output: 2
    // Explanation: Print "aaa" first and then print "b" from the second place of the string, which will cover the existing character 'a'.
    fmt.Println(strangePrinter("aba")) // 2

    fmt.Println(strangePrinter1("aaabbb")) // 2
    fmt.Println(strangePrinter1("aba")) // 2
}
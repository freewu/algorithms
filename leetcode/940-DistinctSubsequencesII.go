package main

// 940. Distinct Subsequences II
// Given a string s, return the number of distinct non-empty subsequences of s. 
// Since the answer may be very large, return it modulo 10^9 + 7.

// A subsequence of a string is a new string that is formed 
// from the original string by deleting some (can be none) of the characters 
// without disturbing the relative positions of the remaining characters. 
// (i.e., "ace" is a subsequence of "abcde" while "aec" is not.)

// Example 1:
// Input: s = "abc"
// Output: 7
// Explanation: The 7 distinct subsequences are "a", "b", "c", "ab", "ac", "bc", and "abc".

// Example 2:
// Input: s = "aba"
// Output: 6
// Explanation: The 6 distinct subsequences are "a", "b", "ab", "aa", "ba", and "aba".

// Example 3:
// Input: s = "aaa"
// Output: 3
// Explanation: The 3 distinct subsequences are "a", "aa" and "aaa".

// Constraints:
//     1 <= s.length <= 2000
//     s consists of lowercase English letters.

import "fmt"

func distinctSubseqII(s string) int {
    n, mod, dp, res := len(s), 1_000_000_007, make([]int,len(s)), 0
    for x,_ := range dp {
        dp[x] = 1
    }
    for i := 0; i < n; i++ {
        for j := 0; j < i; j++ {
            if s[i] != s[j] {
                dp[i] = (dp[i] + dp[j]) % mod
            }
        }
        res = (res + dp[i]) % mod
    }
    return res
}

func distinctSubseqII1(s string) int {
    all, mod, dp := 1, 1_000_000_007, make([]int,26)
    for i := range s {
        add := (all - dp[s[i]-'a']) % mod
        all = (all + add) % mod
        dp[s[i]-'a'] = (dp[s[i]-'a'] + add) % mod
    }
    return (all - 1 + mod) % mod
}

func main() {
    // Example 1:
    // Input: s = "abc"
    // Output: 7
    // Explanation: The 7 distinct subsequences are "a", "b", "c", "ab", "ac", "bc", and "abc".
    fmt.Println(distinctSubseqII("abc")) // 7
    // Example 2:
    // Input: s = "aba"
    // Output: 6
    // Explanation: The 6 distinct subsequences are "a", "b", "ab", "aa", "ba", and "aba".
    fmt.Println(distinctSubseqII("aba")) // 6
    // Example 3:
    // Input: s = "aaa"
    // Output: 3
    // Explanation: The 3 distinct subsequences are "a", "aa" and "aaa".
    fmt.Println(distinctSubseqII("aaa")) // 3
    fmt.Println(distinctSubseqII("blljuffdyfrkqtwfyfztpdiyktrhftgtabxxoibcclbjvirnqyynkyaqlxgyybkgyzvcahmytjdqqtctirnxfjpktxmjkojlvvrr")) // 589192369


    fmt.Println(distinctSubseqII1("abc")) // 7
    fmt.Println(distinctSubseqII1("aba")) // 6
    fmt.Println(distinctSubseqII1("aaa")) // 3

    fmt.Println(distinctSubseqII1("blljuffdyfrkqtwfyfztpdiyktrhftgtabxxoibcclbjvirnqyynkyaqlxgyybkgyzvcahmytjdqqtctirnxfjpktxmjkojlvvrr")) // 589192369
}
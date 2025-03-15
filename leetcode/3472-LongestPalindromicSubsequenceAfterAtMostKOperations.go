package main

// 3472. Longest Palindromic Subsequence After at Most K Operations
// You are given a string s and an integer k.

// In one operation, you can replace the character at any position with the next or previous letter in the alphabet (wrapping around so that 'a' is after 'z'). 
// For example, replacing 'a' with the next letter results in 'b', and replacing 'a' with the previous letter results in 'z'. 
// Similarly, replacing 'z' with the next letter results in 'a', and replacing 'z' with the previous letter results in 'y'.

// Return the length of the longest palindromic subsequence of s that can be obtained after performing at most k operations.

// Example 1:
// Input: s = "abced", k = 2
// Output: 3
// Explanation:
// Replace s[1] with the next letter, and s becomes "acced".
// Replace s[4] with the previous letter, and s becomes "accec".
// The subsequence "ccc" forms a palindrome of length 3, which is the maximum.

// Example 2:
// Input: s = "aaazzz", k = 4
// Output: 6
// Explanation:
// Replace s[0] with the previous letter, and s becomes "zaazzz".
// Replace s[4] with the next letter, and s becomes "zaazaz".
// Replace s[3] with the next letter, and s becomes "zaaaaz".
// The entire string forms a palindrome of length 6.

// Constraints:
//     1 <= s.length <= 200
//     1 <= k <= 200
//     s consists of only lowercase English letters.

import "fmt"

func longestPalindromicSubsequence(s string, k int) int {
    memo := make([][][]int, len(s) + 1)
    for i := range memo {
        memo[i] = make([][]int, len(s) + 1)
        for j := range memo[i] {
            memo[i][j] = make([]int, k + 1)
            for l := range memo[i][j] {
                memo[i][j][l] = -1
            }
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    var dfs func(start, end, k int) int
    dfs = func(start, end, k int) int {
        if memo[start][end][k] != -1 { return memo[start][end][k] }
        if start == end { return 1 }
        if start > end { return 0 }
        res := max(dfs(start + 1, end, k), dfs(start, end-1, k))
        cost := min(abs(int(s[start]) - int(s[end])), 26 - abs(int(s[start]) - int(s[end])))
        if cost <= k {
            res = max(res, 2 + dfs(start + 1, end - 1, k - cost))
        }
        memo[start][end][k] = res
        return res
    }
    return dfs(0, len(s)-1, k)
}

func longestPalindromicSubsequence1(s string, k int) int {
    n, count := len(s), 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i := 0; i < n / 2; i++ {
        d := abs(int(s[i]) - int(s[n-1-i]))
        count += min(d, 26-d)
    }
    if count <= k {
        return n
    }
    f := make([][][]int, k + 1)
    for a := range f {
        f[a] = make([][]int, n)
        for i := n - 1; i >= 0; i-- {
            f[a][i] = make([]int, n)
            f[a][i][i] = 1
            for j := i + 1; j < n; j++ {
                res := max(f[a][i+1][j], f[a][i][j-1])
                d := abs(int(s[i]) - int(s[j]))
                op := min(d, 26 - d)
                if op <= a {
                    res = max(res, f[a - op][i+1][j-1]+2)
                }
                f[a][i][j] = res
            }
        }
    }
    return f[k][0][n - 1]
}

func main() {
    // Example 1:
    // Input: s = "abced", k = 2
    // Output: 3
    // Explanation:
    // Replace s[1] with the next letter, and s becomes "acced".
    // Replace s[4] with the previous letter, and s becomes "accec".
    // The subsequence "ccc" forms a palindrome of length 3, which is the maximum.
    fmt.Println(longestPalindromicSubsequence("abced", 2)) // 3
    // Example 2:
    // Input: s = "aaazzz", k = 4
    // Output: 6
    // Explanation:
    // Replace s[0] with the previous letter, and s becomes "zaazzz".
    // Replace s[4] with the next letter, and s becomes "zaazaz".
    // Replace s[3] with the next letter, and s becomes "zaaaaz".
    // The entire string forms a palindrome of length 6.
    fmt.Println(longestPalindromicSubsequence("aaazzz", 4)) // 6

    fmt.Println(longestPalindromicSubsequence("bluefrog", 4)) // 4
    fmt.Println(longestPalindromicSubsequence("leetcode", 4)) // 5

    fmt.Println(longestPalindromicSubsequence1("abced", 2)) // 3
    fmt.Println(longestPalindromicSubsequence1("aaazzz", 4)) // 6
    fmt.Println(longestPalindromicSubsequence1("bluefrog", 4)) // 4
    fmt.Println(longestPalindromicSubsequence1("leetcode", 4)) // 5
}
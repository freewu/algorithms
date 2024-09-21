package main

// 1278. Palindrome Partitioning III
// You are given a string s containing lowercase letters and an integer k. You need to :
//     First, change some characters of s to other lowercase English letters.
//     Then divide s into k non-empty disjoint substrings such that each substring is a palindrome.

// Return the minimal number of characters that you need to change to divide the string.

// Example 1:
// Input: s = "abc", k = 2
// Output: 1
// Explanation: You can split the string into "ab" and "c", and change 1 character in "ab" to make it palindrome.

// Example 2:
// Input: s = "aabbc", k = 3
// Output: 0
// Explanation: You can split the string into "aa", "bb" and "c", all of them are palindrome.

// Example 3:
// Input: s = "leetcode", k = 8
// Output: 0

// Constraints:
//     1 <= k <= s.length <= 100.
//     s only contains lowercase English letters.

import "fmt"

// two dp
func palindromePartition(s string, k int) int {
    n := len(s)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    var dp1 [105][105]int
    for i := n-1; i >= 0; i-- {
        for j := i+1; j < n; j++ {
            if j - i == 1 {
                if s[i] != s[j] { dp1[i][j] = 1 }
            } else {
                if s[i] == s[j] {
                    dp1[i][j] = dp1[i+1][j-1] 
                } else {
                    dp1[i][j] = dp1[i+1][j-1] + 1
                }
            }
        }
    }
    var dp2[105][105]int
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            dp2[i][j] = int(1e9)
        }
    }
    for i := 0; i < n; i++ {
        dp2[i][1] = dp1[0][i]
    }
    for i := 0; i < n; i++ {
        for j := 2; j < n; j++ {
            for l := 0; l < i; l++ {
                dp2[i][j] = min(dp2[i][j], dp2[l][j-1] + dp1[l+1][i])
            }
        }
    }
    return dp2[n-1][k]
}

// dfs
func palindromePartition1(s string, k int) int {
    n, inf := len(s), 1 << 31
    f := make([][]int, n)
    for i := range f {
        f[i] = make([]int, k+1)
        for j := range f[i] {
            f[i][j] = -1
        }
    }
    t := make([][]int, n)
    for i := range t {
        t[i] = make([]int, n)
    }
    for i := 0; i < 2*n-1; i++ {
        l, r := i/2, i/2+i%2
        for ; l >= 0 && r < n; {
            a := 0
            if s[l] != s[r] {
                a = 1
            }
            b := 0
            if l < r {
                b = t[l+1][r-1]
            }
            t[l][r] = a+b
            l--
            r++
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    var dfs func(i, k int) int
    dfs = func(i, k int) int {
        if i < 0 && k == 0 {
            return 0
        }
        if i < 0 || k == 0 {
            return inf/2
        }
        if f[i][k] != -1 {
            return f[i][k]
        }
        res := inf
        for j := i; j >= 0; j-- {
            res = min(res, t[j][i]+dfs(j-1, k-1))
        }
        f[i][k] = res
        return res
    }
    return dfs(n-1, k)
}

func main() {
    // Example 1:
    // Input: s = "abc", k = 2
    // Output: 1
    // Explanation: You can split the string into "ab" and "c", and change 1 character in "ab" to make it palindrome.
    fmt.Println(palindromePartition("abc", 2)) // 1
    // Example 2:
    // Input: s = "aabbc", k = 3
    // Output: 0
    // Explanation: You can split the string into "aa", "bb" and "c", all of them are palindrome.
    fmt.Println(palindromePartition("aabbc", 3)) // 0
    // Example 3:
    // Input: s = "leetcode", k = 8
    // Output: 0
    fmt.Println(palindromePartition("leetcode", 8)) // 0

    fmt.Println(palindromePartition1("abc", 2)) // 1
    fmt.Println(palindromePartition1("aabbc", 3)) // 0
    fmt.Println(palindromePartition1("leetcode", 8)) // 0
}
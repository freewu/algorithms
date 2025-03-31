package main

// 3504. Longest Palindrome After Substring Concatenation II
// You are given two strings, s and t.

// You can create a new string by selecting a substring from s (possibly empty) and a substring from t (possibly empty), then concatenating them in order.

// Return the length of the longest palindrome that can be formed this way.

// Example 1:
// Input: s = "a", t = "a"
// Output: 2
// Explanation:
// Concatenating "a" from s and "a" from t results in "aa", which is a palindrome of length 2.

// Example 2:
// Input: s = "abc", t = "def"
// Output: 1
// Explanation:
// Since all characters are different, the longest palindrome is any single character, so the answer is 1.

// Example 3:
// Input: s = "b", t = "aaaa"
// Output: 4
// Explanation:
// Selecting "aaaa" from t is the longest palindrome, so the answer is 4.

// Example 4:
// Input: s = "abcde", t = "ecdba"
// Output: 5
// Explanation:
// Concatenating "abc" from s and "ba" from t results in "abcba", which is a palindrome of length 5.

// Constraints:
//     1 <= s.length, t.length <= 1000
//     s and t consist of lowercase English letters.


import "fmt"
import "slices"

func longestPalindrome(s, t string) int {
    m, n := len(s), len(t)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    expand := func(s string, g []int, l, r int) {
        for l >= 0 && r < len(s) && s[l] == s[r] {
            g[l] = max(g[l], r-l+1)
            l, r = l-1, r+1
        }
    }
    calc := func(s string) []int {
        n, g := len(s), make([]int, len(s))
        for i := 0; i < n; i++ {
            expand(s, g, i, i)
            expand(s, g, i, i+1)
        }
        return g
    }
    reverse := func(s string) string {
        r := []rune(s)
        slices.Reverse(r)
        return string(r)
    }
    t = reverse(t)
    g1, g2 := calc(s), calc(t)
    res := max(slices.Max(g1), slices.Max(g2))
    f := make([][]int, m + 1)
    for i := range f {
        f[i] = make([]int, n + 1)
    }
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if s[i-1] == t[j-1] {
                f[i][j] = f[i-1][j-1] + 1
                a, b := 0, 0
                if i < m {
                    a = g1[i]
                }
                if j < n {
                    b = g2[j]
                }
                res = max(res, max(f[i][j] * 2 + a, f[i][j] * 2 + b))
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "a", t = "a"
    // Output: 2
    // Explanation:
    // Concatenating "a" from s and "a" from t results in "aa", which is a palindrome of length 2.
    fmt.Println(longestPalindrome("a", "a")) // 2
    // Example 2:
    // Input: s = "abc", t = "def"
    // Output: 1
    // Explanation:
    // Since all characters are different, the longest palindrome is any single character, so the answer is 1.
    fmt.Println(longestPalindrome("abc", "def")) // 1
    // Example 3:
    // Input: s = "b", t = "aaaa"
    // Output: 4
    // Explanation:
    // Selecting "aaaa" from t is the longest palindrome, so the answer is 4.
    fmt.Println(longestPalindrome("b", "aaaa")) // 1
    // Example 4:
    // Input: s = "abcde", t = "ecdba"
    // Output: 5
    // Explanation:
    // Concatenating "abc" from s and "ba" from t results in "abcba", which is a palindrome of length 5.
    fmt.Println(longestPalindrome("abcde", "ecdba")) // 5

    fmt.Println(longestPalindrome("bluefrog", "leetcode")) // 3
    fmt.Println(longestPalindrome("leetcode", "bluefrog")) // 4
}
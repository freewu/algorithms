package main

// 2911. Minimum Changes to Make K Semi-palindromes
// Given a string s and an integer k, partition s into k substrings such that the letter changes needed to make each substring a semi-palindrome are minimized.

// Return the minimum number of letter changes required.

// A semi-palindrome is a special type of string that can be divided into palindromes based on a repeating pattern. 
// To check if a string is a semi-palindrome:​
//     1. Choose a positive divisor d of the string's length. 
//        d can range from 1 up to, but not including, the string's length. 
//        For a string of length 1, it does not have a valid divisor as per this definition, since the only divisor is its length, which is not allowed.
//     2. For a given divisor d, divide the string into groups where each group contains characters from the string that follow a repeating pattern of length d. 
//        Specifically, the first group consists of characters at positions 1, 1 + d, 1 + 2d, and so on; the second group includes characters at positions 2, 2 + d, 2 + 2d, etc.
//     3. The string is considered a semi-palindrome if each of these groups forms a palindrome.

// Consider the string "abcabc":
//     1. The length of "abcabc" is 6. Valid divisors are 1, 2, and 3.
//     2. For d = 1: The entire string "abcabc" forms one group. Not a palindrome.
//     3. For d = 2:
//         3.1 Group 1 (positions 1, 3, 5): "acb"
//         3.2 Group 2 (positions 2, 4, 6): "bac"
//         3.3 Neither group forms a palindrome.
//     4. For d = 3:
//         4.1 Group 1 (positions 1, 4): "aa"
//         4.2 Group 2 (positions 2, 5): "bb"
//         4.3 Group 3 (positions 3, 6): "cc"
//         4.4 All groups form palindromes. Therefore, "abcabc" is a semi-palindrome.

// Example 1:
// Input: s = "abcac", k = 2
// Output: 1
// Explanation: Divide s into "ab" and "cac". "cac" is already semi-palindrome. Change "ab" to "aa", it becomes semi-palindrome with d = 1.

// Example 2:
// Input: s = "abcdef", k = 2
// Output: 2
// Explanation: Divide s into substrings "abc" and "def". Each needs one change to become semi-palindrome.

// Example 3:
// Input: s = "aabbaa", k = 3
// Output: 0
// Explanation: Divide s into substrings "aa", "bb" and "aa". All are already semi-palindromes.

// Constraints:
//     2 <= s.length <= 200
//     1 <= k <= s.length / 2
//     s contains only lowercase English letters.

import "fmt"

func minimumChanges(s string, k int) int {
    n, inf := len(s), 1 << 31
    g, f := make([][]int, n + 1), make([][]int, n + 1)
    for i := range g {
        g[i], f[i] = make([]int, n + 1),  make([]int, k + 1)
        for j := range g[i] {
            g[i][j] = inf
        }
        for j := range f[i] {
            f[i][j] = inf
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    f[0][0] = 0
    for i := 1; i <= n; i++ {
        for j := i; j <= n; j++ {
            m := j - i + 1
            for d := 1; d < m; d++ {
                if m % d == 0 {
                    count := 0
                    for l := 0; l < m; l++ {
                        r := (m / d - 1 - l / d) * d + l % d
                        if l >= r { break }
                        if s[i-1+l] != s[i-1+r] {
                            count++
                        }
                    }
                    g[i][j] = min(g[i][j], count)
                }
            }
        }
    }
    for i := 1; i <= n; i++ {
        for j := 1; j <= k; j++ {
            for h := 0; h < i-1; h++ {
                f[i][j] = min(f[i][j], f[h][j-1] + g[h+1][i])
            }
        }
    }
    return f[n][k]
}

func main() {
    // Example 1:
    // Input: s = "abcac", k = 2
    // Output: 1
    // Explanation: Divide s into "ab" and "cac". "cac" is already semi-palindrome. Change "ab" to "aa", it becomes semi-palindrome with d = 1.
    fmt.Println(minimumChanges("abcac", 2)) // 1
    // Example 2:
    // Input: s = "abcdef", k = 2
    // Output: 2
    // Explanation: Divide s into substrings "abc" and "def". Each needs one change to become semi-palindrome.
    fmt.Println(minimumChanges("abcdef", 2)) // 2
    // Example 3:
    // Input: s = "aabbaa", k = 3
    // Output: 0
    // Explanation: Divide s into substrings "aa", "bb" and "aa". All are already semi-palindromes.
    fmt.Println(minimumChanges("aabbaa", 3)) // 0

    fmt.Println(minimumChanges("bluefrog", 2)) // 3
    fmt.Println(minimumChanges("leetcode", 2)) // 3
}
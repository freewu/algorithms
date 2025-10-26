package main 

// 3722. Lexicographically Smallest String After Reverse
// You are given a string s of length n consisting of lowercase English letters.

// You must perform exactly one operation by choosing any integer k such that 1 <= k <= n and either:
//     reverse the first k characters of s, or
//     reverse the last k characters of s.

// Return the lexicographically smallest string that can be obtained after exactly one such operation.

// A string a is lexicographically smaller than a string b if, 
// at the first position where they differ, 
// a has a letter that appears earlier in the alphabet than the corresponding letter in b. 
// If the first min(a.length, b.length) characters are the same, 
// then the shorter string is considered lexicographically smaller.

// Example 1:
// Input: s = "dcab"
// Output: "acdb"
// Explanation:
// Choose k = 3, reverse the first 3 characters.
// Reverse "dca" to "acd", resulting string s = "acdb", which is the lexicographically smallest string achievable.

// Example 2:
// Input: s = "abba"
// Output: "aabb"
// Explanation:
// Choose k = 3, reverse the last 3 characters.
// Reverse "bba" to "abb", so the resulting string is "aabb", which is the lexicographically smallest string achievable.

// Example 3:
// Input: s = "zxy"
// Output: "xzy"
// Explanation:
// Choose k = 2, reverse the first 2 characters.
// Reverse "zx" to "xz", so the resulting string is "xzy", which is the lexicographically smallest string achievable.
 
// Constraints:
//     1 <= n == s.length <= 1000
//     s consists of lowercase English letters.

import "fmt"
import "slices"
import "strings"

func lexSmallest(s string) string {
    res, n := s, len(s)
    for k := 2; k <= n; k++ { // k = 1 时，操作不改变 s
        t := []byte(s[:k])
        slices.Reverse(t)
        res = min(res, string(t) + s[k:])
        t = []byte(s[n - k:])
        slices.Reverse(t)
        res = min(res, s[:n - k] + string(t))
    }
    return res
}

func lexSmallest1(s string) string {
    res, left, right, n := "", "", "", len(s)
    for i := 1; i <= n; i++ {
        if i == 1 {
            left, right = s[:i], s[n - i:]
        } else {
            left = s[i - 1 : i] + left
            right = right + s[n - i :n - i + 1]
        }
        t, s1, s2 := "", left + s[i:], s[:n - i] + right
        if strings.Compare(s1, s2) <= 0 {
            t = s1
        } else {
            t = s2
        }
        if len(res) == 0 || strings.Compare(t, res) < 0 {
            res = t
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "dcab"
    // Output: "acdb"
    // Explanation:
    // Choose k = 3, reverse the first 3 characters.
    // Reverse "dca" to "acd", resulting string s = "acdb", which is the lexicographically smallest string achievable.
    fmt.Println(lexSmallest("dcab")) // acdb
    // Example 2:
    // Input: s = "abba"
    // Output: "aabb"
    // Explanation:
    // Choose k = 3, reverse the last 3 characters.
    // Reverse "bba" to "abb", so the resulting string is "aabb", which is the lexicographically smallest string achievable.
    fmt.Println(lexSmallest("abba")) // aabb
    // Example 3:
    // Input: s = "zxy"
    // Output: "xzy"
    // Explanation:
    // Choose k = 2, reverse the first 2 characters.
    // Reverse "zx" to "xz", so the resulting string is "xzy", which is the lexicographically smallest string achievable.
    fmt.Println(lexSmallest("zxy")) // xzy

    fmt.Println(lexSmallest("leetcode")) // cteelode
    fmt.Println(lexSmallest("bluefrog")) // bgorfeul

    fmt.Println(lexSmallest1("dcab")) // acdb
    fmt.Println(lexSmallest1("abba")) // aabb
    fmt.Println(lexSmallest1("zxy")) // xzy
    fmt.Println(lexSmallest1("leetcode")) // cteelode
    fmt.Println(lexSmallest1("bluefrog")) // bgorfeul
}
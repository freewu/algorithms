package main

// 2014. Longest Subsequence Repeated k Times
// You are given a string s of length n, and an integer k. 
// You are tasked to find the longest subsequence repeated k times in string s.

// A subsequence is a string that can be derived from another string by deleting some 
// or no characters without changing the order of the remaining characters.

// A subsequence seq is repeated k times in the string s if seq * k is a subsequence of s, 
// where seq * k represents a string constructed by concatenating seq k times.
//     For example, "bba" is repeated 2 times in the string "bababcba", 
//     because the string "bbabba", constructed by concatenating "bba" 2 times, is a subsequence of the string "bababcba".

// Return the longest subsequence repeated k times in string s. 
// If multiple such subsequences are found, return the lexicographically largest one. 
// If there is no such subsequence, return an empty string.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/08/30/longest-subsequence-repeat-k-times.png" />
// Input: s = "letsleetcode", k = 2
// Output: "let"
// Explanation: There are two longest subsequences repeated 2 times: "let" and "ete".
// "let" is the lexicographically largest one.

// Example 2:
// Input: s = "bb", k = 2
// Output: "b"
// Explanation: The longest subsequence repeated 2 times is "b".

// Example 3:
// Input: s = "ab", k = 2
// Output: ""
// Explanation: There is no subsequence repeated 2 times. Empty string is returned.

// Constraints:
//     n == s.length
//     2 <= n, k <= 2000
//     2 <= n < k * 8
//     s consists of lowercase English letters.

import "fmt"

func longestSubsequenceRepeatedK(s string, k int) string {
    res, arr := "", []byte(s)
    queue, count, bit := []string{""}, make([]int, 26), make([]bool, 26)
    for _, v := range arr {
        count[v - 'a']++
        if count[v - 'a'] >= k {
            bit[v - 'a'] = true
        }
    }
    check := func(k int, subsequence string) bool {
        count, n := 0, len(subsequence)
        for i := range arr {
            if byte(subsequence[count % n]) == arr[i] {
                count++
                if count >= k * n {
                    return true
                }
            }
        }
        return false
    }
    for len(queue) > 0 {
        sb := queue[0]
        queue = queue[1:] // pop
        for i := 0; i < 26; i++ {
            if bit[i] {
                v := sb + string('a' + i)
                if check(k, v) {
                    res = v
                    queue = append(queue, v)
                }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/08/30/longest-subsequence-repeat-k-times.png" />
    // Input: s = "letsleetcode", k = 2
    // Output: "let"
    // Explanation: There are two longest subsequences repeated 2 times: "let" and "ete".
    // "let" is the lexicographically largest one.
    fmt.Println(longestSubsequenceRepeatedK("letsleetcode", 2)) // "let"
    // Example 2:
    // Input: s = "bb", k = 2
    // Output: "b"
    // Explanation: The longest subsequence repeated 2 times is "b".
    fmt.Println(longestSubsequenceRepeatedK("bb", 2)) // "b"
    // Example 3:
    // Input: s = "ab", k = 2
    // Output: ""
    // Explanation: There is no subsequence repeated 2 times. Empty string is returned.
    fmt.Println(longestSubsequenceRepeatedK("ab", 2)) // ""
}
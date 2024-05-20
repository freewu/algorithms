package main

// 1163. Last Substring in Lexicographical Order
// Given a string s, return the last substring of s in lexicographical order.

// Example 1:
// Input: s = "abab"
// Output: "bab"
// Explanation: The substrings are ["a", "ab", "aba", "abab", "b", "ba", "bab"]. The lexicographically maximum substring is "bab".

// Example 2:
// Input: s = "leetcode"
// Output: "tcode"
 
// Constraints:
//     1 <= s.length <= 4 * 10^5
//     s contains only lowercase English letters.

import "fmt"

func lastSubstring(s string) string {
    maxIndex := len(s) - 1
    for currIndex := len(s) - 1; currIndex >= 0; currIndex-- {
        if s[currIndex] > s[maxIndex] {
            maxIndex = currIndex
        } else if s[currIndex] == s[maxIndex] {
            i, j := currIndex + 1, maxIndex + 1
            for i < maxIndex && j < len(s) && s[i] == s[j] {
                i++
                j++
            }
            if i == maxIndex || j == len(s) || s[i] > s[j] {
                maxIndex = currIndex
            }
        }
    }
    return s[maxIndex:]
}

func lastSubstring1(s string) string {
    idx1, idx2, sameLen := 0, 1, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for idx2 + sameLen < len(s) {
        c1 := s[idx1 + sameLen]
        c2 := s[idx2 + sameLen]
        if c1 == c2 {
            sameLen++
            continue
        }
        if c1 > c2 {
            idx2 = idx2 + sameLen + 1
        } else {
            idx1 = max(idx1 + sameLen + 1, idx2)
            idx2 = idx1 + 1
        }
        sameLen = 0
    }
    return s[idx1:]
}

func main() {
    // Example 1:
    // Input: s = "abab"
    // Output: "bab"
    // Explanation: The substrings are ["a", "ab", "aba", "abab", "b", "ba", "bab"]. The lexicographically maximum substring is "bab".
    fmt.Println(lastSubstring("abab")) // "bab"
    // Example 2:
    // Input: s = "leetcode"
    // Output: "tcode"
    fmt.Println(lastSubstring("leetcode")) // "tcode"

    fmt.Println(lastSubstring1("abab")) // "bab"
    fmt.Println(lastSubstring1("leetcode")) // "tcode"
}
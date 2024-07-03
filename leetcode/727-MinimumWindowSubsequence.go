package main

// 727. Minimum Window Subsequence
// Given strings s1 and s2, return the minimum contiguous substring part of s1, so that s2 is a subsequence of the part.
// If there is no such window in s1 that covers all characters in s2, return the empty string "". 
// If there are multiple such minimum-length windows, return the one with the left-most starting index.

// Example 1:
// Input: s1 = "abcdebdde", s2 = "bde"
// Output: "bcde"
// Explanation: 
// "bcde" is the answer because it occurs before "bdde" which has the same length.
// "deb" is not a smaller window because the elements of s2 in the window must occur in order.

// Example 2:
// Input: s1 = "jmeqksfrsdcmsiwvaovztaqenprpvnbstl", s2 = "u"
// Output: ""

// Constraints:
//     1 <= s1.length <= 2 * 10^4
//     1 <= s2.length <= 100
//     s1 and s2 consist of lowercase English letters.

import "fmt"

// 滑动窗口
func minWindow(s1 string, s2 string) string {
    if s1 == "" || s2 == "" || len(s1) < len(s2) {
        return ""
    }
    inf := 1 << 32 - 1
    rer, rel := inf, 0 // 返回值的坐标
    i, j := 0, 0 // 分别作为s1，s2的坐标
    for i < len(s1) {
        if s1[i] == s2[j] {
            j++
        }
        if j == len(s2) {
            right := i
            j--
            for j >= 0 {
                if s2[j] == s1[i] {
                    j--
                }
                i--
            }
            i++
            if right - i < rer - rel {
                rer = right
                rel = i
            }
            j = 0
        }
        i++
    }
    if rer == inf {
        return ""
    }
    return s1[rel:rer+1]
}

func main() {
    // Example 1:
    // Input: s1 = "abcdebdde", s2 = "bde"
    // Output: "bcde"
    // Explanation: 
    // "bcde" is the answer because it occurs before "bdde" which has the same length.
    // "deb" is not a smaller window because the elements of s2 in the window must occur in order.
    fmt.Println(minWindow("abcdebdde","bde")) // "bcde"
    // Example 2:
    // Input: s1 = "jmeqksfrsdcmsiwvaovztaqenprpvnbstl", s2 = "u"
    // Output: ""
    fmt.Println(minWindow("jmeqksfrsdcmsiwvaovztaqenprpvnbstl","u")) // ""
}
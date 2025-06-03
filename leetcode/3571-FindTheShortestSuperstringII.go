package main

// 3571. Find the Shortest Superstring II
// You are given two strings, s1 and s2. 
// Return the shortest possible string that contains both s1 and s2 as substrings. 
// If there are multiple valid answers, return any one of them.

// A substring is a contiguous sequence of characters within a string.

// Example 1:
// Input: s1 = "aba", s2 = "bab"
// Output: "abab"
// Explanation:
// "abab" is the shortest string that contains both "aba" and "bab" as substrings.

// Example 2:
// Input: s1 = "aa", s2 = "aaa"
// Output: "aaa"
// Explanation:
// "aa" is already contained within "aaa", so the shortest superstring is "aaa".

// Constraints:
//     1 <= s1.length <= 100
//     1 <= s2.length <= 100
//     s1 and s2 consist of lowercase English letters only.

import "fmt"
import "strings"

func shortestSuperstring(s1 string, s2 string) string {
    isSubstring := func(s1, s2 string) bool { // 检查s1是否是s2的子串
        //return len(s1) <= len(s2) && (s2 == s1 || len(s2) >= len(s1) && s2[len(s2)-len(s1):] == s1)
        return len(s1) <= len(s2) && strings.Contains(s2, s1)
    }
    maxOverlap := func(s1, s2 string) int { // 计算s1和s2的最大重叠部分
        n := len(s1)
        if len(s2) < n {
            n = len(s2)
        }
        for i := n; i > 0; i-- {
            if s1[len(s1) - i:] == s2[:i] {
                return i
            }
        }
        return 0
    }
    if isSubstring(s1, s2) { return s2 } // 检查 s1 是否是 s2 的子串
    if isSubstring(s2, s1) { return s1 } // 检查 s2 是否是s1的子串
    // 计算s1和s2的最大重叠部分
    overlap1, overlap2 := maxOverlap(s1, s2), maxOverlap(s2, s1)
    // 根据重叠部分选择最优的拼接方式
    if overlap1 >= overlap2 {
        return s1 + s2[overlap1:]
    }
    return s2 + s1[overlap2:]
}

func main() {
    // Example 1:
    // Input: s1 = "aba", s2 = "bab"
    // Output: "abab"
    // Explanation:
    // "abab" is the shortest string that contains both "aba" and "bab" as substrings.
    fmt.Println(shortestSuperstring("aba","bab")) // "abab"
    // Example 2:
    // Input: s1 = "aa", s2 = "aaa"
    // Output: "aaa"
    // Explanation:
    // "aa" is already contained within "aaa", so the shortest superstring is "aaa".
    fmt.Println(shortestSuperstring("aa","aaa")) // "aaa"

    fmt.Println(shortestSuperstring("m","azmvzfh")) // "azmvzfh"

    fmt.Println(shortestSuperstring("bluefrog","leetcode")) // "bluefrogleetcode"
    fmt.Println(shortestSuperstring("leetcode","bluefrog")) // "leetcodebluefrog"
}
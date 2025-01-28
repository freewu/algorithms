package main

// 2730. Find the Longest Semi-Repetitive Substring
// You are given a digit string s that consists of digits from 0 to 9.

// A string is called semi-repetitive if there is at most one adjacent pair of the same digit. 
// For example, "0010", "002020", "0123", "2002", and "54944" are semi-repetitive while the following are not: "00101022" (adjacent same digit pairs are 00 and 22), and "1101234883" (adjacent same digit pairs are 11 and 88).

// Return the length of the longest semi-repetitive substring of s.

// Example 1:
// Input: s = "52233"
// Output: 4
// Explanation:
// The longest semi-repetitive substring is "5223". Picking the whole string "52233" has two adjacent same digit pairs 22 and 33, but at most one is allowed.

// Example 2:
// Input: s = "5494"
// Output: 4
// Explanation:
// s is a semi-repetitive string.

// Example 3:
// Input: s = "1111111"
// Output: 2
// Explanation:
// The longest semi-repetitive substring is "11". Picking the substring "111" has two adjacent same digit pairs, but at most one is allowed.

// Constraints:
//     1 <= s.length <= 50
//     '0' <= s[i] <= '9'

import "fmt"

func longestSemiRepetitiveSubstring(s string) int {
    res, n := 0, len(s)
    for i := 1; i < n; i++ {
        if s[i] == s[i - 1]{
            left := i - 2
            for left >= 0 && s[left] != s[left + 1]{
                left--
            }
            right := i + 1
            for right < n && s[right] != s[right - 1]{
                right++
            }
            res = max(res,right - left - 1)
        }
    }
    if res > 0 {
        return res
    }
    return n
}

func longestSemiRepetitiveSubstring1(s string) int {
    n := len(s)
    if n == 1 { return 1 }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res, left, last := 0, 0, -1 // 用于记录最近一次重复字符的位置
    for right := 1; right < n; right++ {
        if s[right] == s[right - 1] { // 如果当前字符和前一个字符重复
            if last != -1 { // 如果已经有一个重复字符，调整左边界
                left = last + 1
            }
            last = right - 1 // 更新最近一次重复字符的位置
        }
        res = max(res, right - left + 1) // 更新最长子串的长度
    }
    return res
}

func longestSemiRepetitiveSubstring2(s string) int {
    res, left, same := 1, 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for right := 1; right < len(s); right++ {
        if s[right] == s[right-1] {
            same++
        }
        if same > 1 { // same == 2
            left++
            for s[left] != s[left-1] {
                left++
            }
            same = 1
        }
        res = max(res, right - left + 1)
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "52233"
    // Output: 4
    // Explanation:
    // The longest semi-repetitive substring is "5223". Picking the whole string "52233" has two adjacent same digit pairs 22 and 33, but at most one is allowed.
    fmt.Println(longestSemiRepetitiveSubstring("52233")) // 4
    // Example 2:
    // Input: s = "5494"
    // Output: 4
    // Explanation:
    // s is a semi-repetitive string.
    fmt.Println(longestSemiRepetitiveSubstring("5494")) // 4
    // Example 3:
    // Input: s = "1111111"
    // Output: 2
    // Explanation:
    // The longest semi-repetitive substring is "11". Picking the substring "111" has two adjacent same digit pairs, but at most one is allowed.
    fmt.Println(longestSemiRepetitiveSubstring("1111111")) // 2

    fmt.Println(longestSemiRepetitiveSubstring("123456789")) // 9
    fmt.Println(longestSemiRepetitiveSubstring("987654321")) // 9

    fmt.Println(longestSemiRepetitiveSubstring1("52233")) // 4
    fmt.Println(longestSemiRepetitiveSubstring1("5494")) // 4
    fmt.Println(longestSemiRepetitiveSubstring1("1111111")) // 2
    fmt.Println(longestSemiRepetitiveSubstring1("123456789")) // 9
    fmt.Println(longestSemiRepetitiveSubstring1("987654321")) // 9

    fmt.Println(longestSemiRepetitiveSubstring2("52233")) // 4
    fmt.Println(longestSemiRepetitiveSubstring2("5494")) // 4
    fmt.Println(longestSemiRepetitiveSubstring2("1111111")) // 2
    fmt.Println(longestSemiRepetitiveSubstring2("123456789")) // 9
    fmt.Println(longestSemiRepetitiveSubstring2("987654321")) // 9
}
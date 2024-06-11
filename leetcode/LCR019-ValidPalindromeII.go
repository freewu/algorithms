package main

// LCR 019. 验证回文串 II
// 给定一个非空字符串 s，请判断如果 最多 从字符串中删除一个字符能否得到一个回文字符串。

// 示例 1:
// 输入: s = "aba"
// 输出: true

// 示例 2:
// 输入: s = "abca"
// 输出: true
// 解释: 可以删除 "c" 字符 或者 "b" 字符

// 示例 3:
// 输入: s = "abc"
// 输出: false
 
// 提示:
// 1 <= s.length <= 10^5
// s 由小写英文字母组成

import "fmt"

func validPalindrome(s string) bool {
    i, j := 0, len(s) - 1
    validSubstring := func (s string) bool {
        i,j := 0, len(s) - 1
        for i < j {
            if s[i] == s[j] {
                i++
                j--
            } else {
                return false
            }
        }
        return true
    }
    for i < j {
        if s[i] == s[j] {
            i++
            j--
        } else {
            return validSubstring(s[i+1:j+1]) || validSubstring(s[i:j-1+1])
        }
    }
    return true
}

func main() {
    // Example 1:
    // Input: s = "aba"
    // Output: true
    fmt.Println(validPalindrome("aba")) // true
    // Example 2:
    // Input: s = "abca"
    // Output: true
    // Explanation: You could delete the character 'c'.
    fmt.Println(validPalindrome("abca")) // true
    // Example 3:
    // Input: s = "abc"
    // Output: false
    fmt.Println(validPalindrome("abc")) // false

    fmt.Println(validPalindrome("tebbem")) // true
}
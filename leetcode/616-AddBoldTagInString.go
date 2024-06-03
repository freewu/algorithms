package main

// 616. Add Bold Tag in String
// You are given a string s and an array of strings words.
// You should add a closed pair of bold tag <b> and </b> to wrap the substrings in s that exist in words.
//     If two such substrings overlap, you should wrap them together with only one pair of closed bold-tag.
//     If two substrings wrapped by bold tags are consecutive, you should combine them.

// Return s after adding the bold tags.

// Example 1:
// Input: s = "abcxyz123", words = ["abc","123"]
// Output: "<b>abc</b>xyz<b>123</b>"
// Explanation: The two strings of words are substrings of s as following: "abcxyz123".
// We add <b> before each substring and </b> after each substring.

// Example 2:
// Input: s = "aaabbb", words = ["aa","b"]
// Output: "<b>aaabbb</b>"
// Explanation: 
// "aa" appears as a substring two times: "aaabbb" and "aaabbb".
// "b" appears as a substring three times: "aaabbb", "aaabbb", and "aaabbb".
// We add <b> before each substring and </b> after each substring: "<b>a<b>a</b>a</b><b>b</b><b>b</b><b>b</b>".
// Since the first two <b>'s overlap, we merge them: "<b>aaa</b><b>b</b><b>b</b><b>b</b>".
// Since now the four <b>'s are consecutive, we merge them: "<b>aaabbb</b>".

// Constraints:
//     1 <= s.length <= 1000
//     0 <= words.length <= 100
//     1 <= words[i].length <= 1000
//     s and words[i] consist of English letters and digits.
//     All the values of words are unique.

import "fmt"
import "strings"

func addBoldTag(s string, words []string) string {
    // 寻找从index开始的substring，返回substring的末字符字符的索引值
    // 如果找得到，返回其中最长的substring的末字符的索引值
    // 找不到就返回-1
    findSubstring := func(s string, words []string, index int) int {
        sLen, i, j, res := len(s), 0, 0, -1
        for _, word := range words {
            if word[0] == s[index] {
                wordLen := len(word)
                for i, j = index + 1, 1; i < sLen && j < wordLen; i, j = i + 1, j + 1 {
                    if s[i] != word[j] {
                        break
                    }
                }
                if j == wordLen && i - 1 > res {
                    res = i - 1
                }
            }
        }
        return res
    }
    sLen := len(s)
    marked := make([]bool, sLen) // 用于记录该位置的字符是否加粗
    for index := 0; index < sLen; index++ {
        subStringEnd := findSubstring(s, words, index)
        if subStringEnd != -1 {
            // subStringEnd != -1 说明
            // 从index到substringEnd的字符串在words中
            // 这些位置需要加粗，标记为true
            for i := index; i <= subStringEnd; i++ {
                marked[i] = true
            }
        }
    }
    isFirst := true // 用于记录是否是第一个被加粗的字符
    var builder strings.Builder
    for index := 0; index < sLen; index++ {
        if marked[index] {
            if isFirst {
                isFirst = false
                builder.WriteString("<b>")
            } 
            builder.WriteByte(s[index])
            // 如果是s中最后一个字符，且加粗
            // 需要手动加上</b>
            if index == sLen - 1 {
                builder.WriteString("</b>")
            }
        } else {
            // isFirst为false说明前面是加粗的
            // 且这回碰上的字符不用加粗
            // 那就得加上</b>
            if !isFirst {
                builder.WriteString("</b>")
                isFirst = true
            }
            builder.WriteByte(s[index])
        }
    }
    return builder.String()
}

func main() {
    // Example 1:
    // Input: s = "abcxyz123", words = ["abc","123"]
    // Output: "<b>abc</b>xyz<b>123</b>"
    // Explanation: The two strings of words are substrings of s as following: "abcxyz123".
    // We add <b> before each substring and </b> after each substring.
    fmt.Println(addBoldTag("abcxyz123",[]string{"abc","123"})) // "<b>abc</b>xyz<b>123</b>"
    // Example 2:
    // Input: s = "aaabbb", words = ["aa","b"]
    // Output: "<b>aaabbb</b>"
    // Explanation: 
    // "aa" appears as a substring two times: "aaabbb" and "aaabbb".
    // "b" appears as a substring three times: "aaabbb", "aaabbb", and "aaabbb".
    // We add <b> before each substring and </b> after each substring: "<b>a<b>a</b>a</b><b>b</b><b>b</b><b>b</b>".
    // Since the first two <b>'s overlap, we merge them: "<b>aaa</b><b>b</b><b>b</b><b>b</b>".
    // Since now the four <b>'s are consecutive, we merge them: "<b>aaabbb</b>".
    fmt.Println(addBoldTag("aaabbb",[]string{"aa","b"})) // "<b>aaabbb</b>"
}
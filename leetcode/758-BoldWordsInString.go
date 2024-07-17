package main

// 758. Bold Words in String
// Given an array of keywords words and a string s, make all appearances of all keywords words[i] in s bold. 
// Any letters between <b> and </b> tags become bold.

// Return s after adding the bold tags. 
// The returned string should use the least number of tags possible, and the tags should form a valid combination.

// Example 1:
// Input: words = ["ab","bc"], s = "aabcd"
// Output: "a<b>abc</b>d"
// Explanation: Note that returning "a<b>a<b>b</b>c</b>d" would use more tags, so it is incorrect.

// Example 2:
// Input: words = ["ab","cb"], s = "aabcd"
// Output: "a<b>ab</b>cd"

// Constraints:
//     1 <= s.length <= 500
//     0 <= words.length <= 50
//     1 <= words[i].length <= 10
//     s and words[i] consist of lowercase English letters.

import "fmt"
import "strings"

func boldWords(words []string, s string) string {
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
    // Input: words = ["ab","bc"], s = "aabcd"
    // Output: "a<b>abc</b>d"
    // Explanation: Note that returning "a<b>a<b>b</b>c</b>d" would use more tags, so it is incorrect.
    fmt.Println(boldWords([]string{"ab","bc"},"aabcd")) // "a<b>abc</b>d"
    // Example 2:
    // Input: words = ["ab","cb"], s = "aabcd"
    // Output: "a<b>ab</b>cd"
    fmt.Println(boldWords([]string{"ab","cb"},"aabcd")) // "a<b>ab</b>cd"
}
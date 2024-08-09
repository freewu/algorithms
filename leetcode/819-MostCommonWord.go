package main

// 819. Most Common Word
// Given a string paragraph and a string array of the banned words banned, 
// return the most frequent word that is not banned. 
// It is guaranteed there is at least one word that is not banned, and that the answer is unique.

// The words in paragraph are case-insensitive and the answer should be returned in lowercase.

// Example 1:
// Input: paragraph = "Bob hit a ball, the hit BALL flew far after it was hit.", banned = ["hit"]
// Output: "ball"
// Explanation: 
// "hit" occurs 3 times, but it is a banned word.
// "ball" occurs twice (and no other word does), so it is the most frequent non-banned word in the paragraph. 
// Note that words in the paragraph are not case sensitive,
// that punctuation is ignored (even if adjacent to words, such as "ball,"), 
// and that "hit" isn't the answer even though it occurs more because it is banned.

// Example 2:
// Input: paragraph = "a.", banned = []
// Output: "a"
 
// Constraints:
//     1 <= paragraph.length <= 1000
//     paragraph consists of English letters, space ' ', or one of the symbols: "!?',;.".
//     0 <= banned.length <= 100
//     1 <= banned[i].length <= 10
//     banned[i] consists of only lowercase English letters.

import "fmt"
import "strings"

func mostCommonWord(paragraph string, banned []string) string {
    ban, dic := make(map[string]bool), make(map[string]int)
    word, freq := "", 0
    for _, v := range banned {
        ban[v] = true
    }
    para, check := strings.ToLower(paragraph) + " ", "" // 转成 小写
    for _, c := range para {
        a := int(c)
        if a > 96 && a < 123 { // 取出连续的字符
            check += string(c)
        } else {
            if check != "" {
                if ban[check] { // 如果是ban的字符直接处理下个
                    check = ""
                    continue
                } else {
                    dic[check]++ // 累加出现次数
                    if dic[check] > freq { // 出现最多的字符串
                        freq = dic[check]
                        word = check
                    }
                    check = ""
                }
            }
        }
    }
    return word
}

func main() {
    // Example 1:
    // Input: paragraph = "Bob hit a ball, the hit BALL flew far after it was hit.", banned = ["hit"]
    // Output: "ball"
    // Explanation: 
    // "hit" occurs 3 times, but it is a banned word.
    // "ball" occurs twice (and no other word does), so it is the most frequent non-banned word in the paragraph. 
    // Note that words in the paragraph are not case sensitive,
    // that punctuation is ignored (even if adjacent to words, such as "ball,"), 
    // and that "hit" isn't the answer even though it occurs more because it is banned.
    fmt.Println(mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"})) // "ball"
    // Example 2:
    // Input: paragraph = "a.", banned = []
    // Output: "a"
    fmt.Println(mostCommonWord("a.", []string{})) // "a"
}
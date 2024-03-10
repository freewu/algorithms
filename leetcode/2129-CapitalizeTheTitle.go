package main

// 2129. Capitalize the Title
// You are given a string title consisting of one or more words separated by a single space, where each word consists of English letters. 
// Capitalize the string by changing the capitalization of each word such that:
//         If the length of the word is 1 or 2 letters, change all letters to lowercase.
//         Otherwise, change the first letter to uppercase and the remaining letters to lowercase.

// Return the capitalized title.

// Example 1:
// Input: title = "capiTalIze tHe titLe"
// Output: "Capitalize The Title"
// Explanation:
// Since all the words have a length of at least 3, the first letter of each word is uppercase, and the remaining letters are lowercase.

// Example 2:
// Input: title = "First leTTeR of EACH Word"
// Output: "First Letter of Each Word"
// Explanation:
// The word "of" has length 2, so it is all lowercase.
// The remaining words have a length of at least 3, so the first letter of each remaining word is uppercase, and the remaining letters are lowercase.

// Example 3:
// Input: title = "i lOve leetcode"
// Output: "i Love Leetcode"
// Explanation:
// The word "i" has length 1, so it is lowercase.
// The remaining words have a length of at least 3, so the first letter of each remaining word is uppercase, and the remaining letters are lowercase.

// Constraints:
//         1 <= title.length <= 100
//         title consists of words separated by a single space without any leading or trailing spaces.
//         Each word consists of uppercase and lowercase English letters and is non-empty.

import "fmt"
import "strings"

// func capitalizeTitle1(title string) string {
//     l := len(title)
//     res := make([]byte, l)
//     flag := true // 开头需要大写
//     for i := 0; i < l; i++ {
//         res[i] = title[i]
//         // 遇到了空格下一个就要大写
//         if title[i] == ' ' {
//             // 判段只有1-2个字符的单词开头不需要转成大写
//             if (i + 2 < l &&  title[i + 2] == ' ') || (i + 3 < l &&  title[i + 3] == ' ') {
//             } else {
//                 flag = true
//                 continue
//             }
//         }
//         if flag && (title[i] >= 'a' && title[i] <= 'z') { // 判断单词开头是否小写
//             res[i] = title[i] - 32
//         } else if !flag && title[i] >= 'A' && title[i] <= 'Z' { // 判断单词其它部份(除了开头)是否大写
//             res[i] = title[i] + 32
//         }
//         flag = false
//     }
//     return string(res)
// }

func capitalizeTitle(title string) string {
    title = strings.ToLower(title)
    v := strings.Split(title, " ")
    for i := 0; i < len(v); i++ {
        // 如果单词的长度为 1 或者 2 ，所有字母变成小写。
        if len(v[i]) < 3 {
            v[i] = strings.ToLower(v[i])
        } else { // 否则，将单词首字母大写，剩余字母变成小写。
            v[i] = strings.Title(v[i])
        }
    }
    title = strings.Join(v, " ")
    return title
}

func main() {
    fmt.Println(capitalizeTitle("capiTalIze tHe titLe")) // capiTalIze tHe titLe
    fmt.Println(capitalizeTitle("First leTTeR of EACH Word")) // First Letter of Each Word
    fmt.Println(capitalizeTitle("i lOve leetcode")) // i Love Leetcode

    // fmt.Println(capitalizeTitle1("capiTalIze tHe titLe")) // capiTalIze tHe titLe
    // fmt.Println(capitalizeTitle1("First leTTeR of EACH Word")) // First Letter of Each Word
    // fmt.Println(capitalizeTitle1("i lOve leetcode")) // i Love Leetcode
}
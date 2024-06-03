package main

// 408. Valid Word Abbreviation
// A string can be abbreviated by replacing any number of non-adjacent, non-empty substrings with their lengths. 
// The lengths should not have leading zeros.

// For example, a string such as "substitution" could be abbreviated as (but not limited to):
//     "s10n" ("s ubstitutio n")
//     "sub4u4" ("sub stit u tion")
//     "12" ("substitution")
//     "su3i1u2on" ("su bst i t u ti on")
//     "substitution" (no substrings replaced)

// The following are not valid abbreviations:
//     "s55n" ("s ubsti tutio n", the replaced substrings are adjacent)
//     "s010n" (has leading zeros)
//     "s0ubstitution" (replaces an empty substring)

// Given a string word and an abbreviation abbr, return whether the string matches the given abbreviation.
// A substring is a contiguous non-empty sequence of characters within a string.

// Example 1:
// Input: word = "internationalization", abbr = "i12iz4n"
// Output: true
// Explanation: The word "internationalization" can be abbreviated as "i12iz4n" ("i nternational iz atio n").

// Example 2:
// Input: word = "apple", abbr = "a2e"
// Output: false
// Explanation: The word "apple" cannot be abbreviated as "a2e".
 
// Constraints:
//     1 <= word.length <= 20
//     word consists of only lowercase English letters.
//     1 <= abbr.length <= 10
//     abbr consists of lowercase English letters and digits.
//     All the integers in abbr will fit in a 32-bit integer.

import "fmt"

// func validWordAbbreviation(word string, abbr string) bool {
//     index, num := 0, 0
//     for i := 0; i < len(abbr); i++ {
//         if abbr[i] >= '0' && abbr[i] <= '9' { // 处理数字 "s010n" (has leading zeros) 的情况
//             n := int(abbr[i] - '0')
//             if num == 0 && n == 0 { // 处理
//                 return false
//             }
//             num = num * 10 + n
//         } else {
//             if num != 0 {
//                 index += num
//                 num = 0 // 重新置 0
//             }
//             fmt.Printf("word[%v] = %v, abbr[%v]= %v\n",index, word[index], i, abbr[i])
//             if word[index] != abbr[i] {
//                 return false
//             }
//             index++
//         }
//     }
//     return true
// }

func validWordAbbreviation(word string, abbr string) bool {
    m, n :=len(word), len(abbr)
    if m==0 && n == 0{
        return true
    }
    if m==0 || n == 0 {
        return false
    }
    if isAlpha(abbr[0]) {
        if abbr[0] != word[0] {
            return false
        }
        return validWordAbbreviation(word[1:], abbr[1:])
    }
    if abbr[0] == '0' {
        return false
    }
    num, i := 0, 0
    for ; i < n && !isAlpha(abbr[i]); i++ {
        num = num * 10 + int(abbr[i] - '0')
    }
    if num > m {
        return false
    }
    return validWordAbbreviation(word[num:],abbr[i:])
}

func isAlpha(ch byte)bool{
    return 'a' <= ch && ch <= 'z'
}

func main() {
    // Example 1:
    // Input: word = "internationalization", abbr = "i12iz4n"
    // Output: true
    // Explanation: The word "internationalization" can be abbreviated as "i12iz4n" ("i nternational iz atio n").
    fmt.Println(validWordAbbreviation("internationalization","i12iz4n")) // true
    // Example 2:
    // Input: word = "apple", abbr = "a2e"
    // Output: false
    // Explanation: The word "apple" cannot be abbreviated as "a2e".
    fmt.Println(validWordAbbreviation("apple","a2e")) // false
}

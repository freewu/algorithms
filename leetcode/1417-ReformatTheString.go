package main

// 1417. Reformat The String
// You are given an alphanumeric string s. 
// (Alphanumeric string is a string consisting of lowercase English letters and digits).

// You have to find a permutation of the string where no letter is followed by another letter and no digit is followed by another digit. 
// That is, no two adjacent characters have the same type.

// Return the reformatted string or return an empty string if it is impossible to reformat the string.

// Example 1:
// Input: s = "a0b1c2"
// Output: "0a1b2c"
// Explanation: No two adjacent characters have the same type in "0a1b2c". "a0b1c2", "0a1b2c", "0c2a1b" are also valid permutations.

// Example 2:
// Input: s = "leetcode"
// Output: ""
// Explanation: "leetcode" has only characters so we cannot separate them by digits.

// Example 3:
// Input: s = "1229857369"
// Output: ""
// Explanation: "1229857369" has only digits so we cannot separate them by characters.

// Constraints:
//     1 <= s.length <= 500
//     s consists of only lowercase English letters and/or digits.

import "fmt"

// func reformat(s string) string {
//     res := []byte{}
//     isLetter := func (c byte) bool { return 'a' <= c && c <= 'z' }
//     isNumber := func (c byte) bool { return '0' <= c && c <= '9' }
//     for i := 0; i < len(s) - 1; i += 2 {
//         if isLetter(s[i]) && isNumber(s[i + 1]) {
//             res = append(res, s[i + 1])
//             res = append(res, s[i])
//         }
//     }
//     return string(res)
// }

func reformat(s string) string {
    charCount, digitCount := 0, 0
    chars, digits, res := []byte{}, []byte{}, []byte{}
    for _, v := range []byte(s) {
        if '0' <= v && v <= '9' {
            digitCount++
            digits = append(digits, v)
        } else {
            charCount++
            chars = append(chars, v)
        }
    }
    if charCount - digitCount > 1 || charCount - digitCount < -1 { return "" }
    if len(chars) == len(digits) {
        for i, v := range chars {
            res = append(res, []byte{v, digits[i]}...)
        }
    } else if len(chars) > len(digits) {
        for i, v := range chars {
            res = append(res, v)
            if i < len(digits) {
                res = append(res, digits[i])
            }
        }
    } else {
        for i, v := range digits {
            res = append(res, v)
            if i < len(chars) {
                res = append(res, chars[i])
            }
        }
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: s = "a0b1c2"
    // Output: "0a1b2c"
    // Explanation: No two adjacent characters have the same type in "0a1b2c". "a0b1c2", "0a1b2c", "0c2a1b" are also valid permutations.
    fmt.Println(reformat("a0b1c2")) // "0a1b2c"
    // Example 2:
    // Input: s = "leetcode"
    // Output: ""
    // Explanation: "leetcode" has only characters so we cannot separate them by digits.
    fmt.Println(reformat("leetcode")) // ""
    // Example 3:
    // Input: s = "1229857369"
    // Output: ""
    // Explanation: "1229857369" has only digits so we cannot separate them by characters.
    fmt.Println(reformat("1229857369")) // ""
    fmt.Println(reformat("covid2019")) // "c2o0v1i9d"
}
package main

// 2060. Check if an Original String Exists Given Two Encoded Strings
// An original string, consisting of lowercase English letters, can be encoded by the following steps:
//     1. Arbitrarily split it into a sequence of some number of non-empty substrings.
//     2. Arbitrarily choose some elements (possibly none) of the sequence, 
//        and replace each with its length (as a numeric string).
//     3. Concatenate the sequence as the encoded string.

// For example, one way to encode an original string "abcdefghijklmnop" might be:
//     1. Split it as a sequence: ["ab", "cdefghijklmn", "o", "p"].
//     2. Choose the second and third elements to be replaced by their lengths, respectively. 
//        The sequence becomes ["ab", "12", "1", "p"].
//     3. Concatenate the elements of the sequence to get the encoded string: "ab121p".

// Given two encoded strings s1 and s2, consisting of lowercase English letters and digits 1-9 (inclusive), 
// return true if there exists an original string that could be encoded as both s1 and s2. 
// Otherwise, return false.

// Note: The test cases are generated such that the number of consecutive digits in s1 and s2 does not exceed 3.

// Example 1:
// Input: s1 = "internationalization", s2 = "i18n"
// Output: true
// Explanation: It is possible that "internationalization" was the original string.
// - "internationalization" 
//   -> Split:       ["internationalization"]
//   -> Do not replace any element
//   -> Concatenate:  "internationalization", which is s1.
// - "internationalization"
//   -> Split:       ["i", "nternationalizatio", "n"]
//   -> Replace:     ["i", "18",                 "n"]
//   -> Concatenate:  "i18n", which is s2

// Example 2:
// Input: s1 = "l123e", s2 = "44"
// Output: true
// Explanation: It is possible that "leetcode" was the original string.
// - "leetcode" 
//   -> Split:      ["l", "e", "et", "cod", "e"]
//   -> Replace:    ["l", "1", "2",  "3",   "e"]
//   -> Concatenate: "l123e", which is s1.
// - "leetcode" 
//   -> Split:      ["leet", "code"]
//   -> Replace:    ["4",    "4"]
//   -> Concatenate: "44", which is s2.

// Example 3:
// Input: s1 = "a5b", s2 = "c5b"
// Output: false
// Explanation: It is impossible.
// - The original string encoded as s1 must start with the letter 'a'.
// - The original string encoded as s2 must start with the letter 'c'.

// Constraints:
//     1 <= s1.length, s2.length <= 40
//     s1 and s2 consist of digits 1-9 (inclusive), and lowercase English letters only.
//     The number of consecutive digits in s1 and s2 does not exceed 3.

import "fmt"

// // 超出时间限制 97 / 215
// func possiblyEquals(s1 string, s2 string) bool {
//     cache := make(map[string]bool)
//     isDigit := func(b byte) bool { return b >= '0' && b <= '9' }
//     var checkPatterns func(s1, s2 string, diff int) bool
//     processDigsFor1 := func(s1, s2 string, diff int) bool {
//         p := fmt.Sprintf("%s-%s-%d", s1, s2, diff)
//         if cache[p] { return true }
//         for i := 0; i < len(s1) && isDigit(s1[i]); i++ {
//             var ln int
//             fmt.Sscanf(s1[:i+1], "%d", &ln)
//             if checkPatterns(s1[i+1:], s2, diff+ln) {
//                 cache[p] = true
//                 return true
//             }
//         }
//         cache[p] = false
//         return false
//     }
//     processDigsFor2 := func(s1, s2 string, diff int) bool {
//         p := fmt.Sprintf("%s-%s-%d", s1, s2, diff)
//         if cache[p] { return true }
//         for i := 0; i < len(s2) && isDigit(s2[i]); i++ {
//             var ln int
//             fmt.Sscanf(s2[:i+1], "%d", &ln)
//             if checkPatterns(s1, s2[i+1:], diff-ln) {
//                 cache[p] = true
//                 return true
//             }
//         }
//         cache[p] = false
//         return false
//     }
//     checkPatterns = func(s1, s2 string, diff int) bool {
//         if s1 == "" && s2 == "" { return diff == 0 }
//         if len(s1) > 0 && isDigit(s1[0]) { return processDigsFor1(s1, s2, diff) }
//         if len(s2) > 0 && isDigit(s2[0]) { return processDigsFor2(s1, s2, diff) }
//         if len(s1) > 0 && diff < 0 { return checkPatterns(s1[1:], s2, diff+1) }
//         if len(s2) > 0 && diff > 0 { return checkPatterns(s1, s2[1:], diff-1) }
//         if len(s1) > 0 && len(s2) > 0 && s1[0] == s2[0] { return checkPatterns(s1[1:], s2[1:], diff) }
//         return false
//     }
//     return checkPatterns(s1, s2, 0)
// }

func possiblyEquals(s1, s2 string) bool {
    n, m, bias := len(s1), len(s2), 1000
    visited := make([][][2000]bool, n + 1)
    for i := range visited {
        visited[i] = make([][2000]bool, m + 1)
    }
    isDigit := func(b byte) bool { return b >= '0' && b <= '9' }
    var dfs func(i, j, d int) bool
    dfs = func(i, j, d int) bool {
        if i == n && j == m { return d == 0 } // 匹配成功
        if visited[i][j][d + bias] { return false }
        visited[i][j][d + bias] = true
        // 原始字符串长度相同时，若 s1[i] == s2[j]，则 s1[:i] 和 s2[:j] 均可以向后扩展一个字母
        if d == 0 && i < n && j < m && s1[i] == s2[j] && dfs(i+1, j+1, 0) { return true }
        if d <= 0 && i < n { // s1[:i] 的原始字符串长度不超过 s2[:j] 的原始字符串长度时，扩展 s1[:i]
            if isDigit(s1[i]) { // 数字
                for p, v := i, 0; p < n && isDigit(s1[p]); p++ {
                    v = v*10 + int(s1[p]&15)
                    if dfs(p+1, j, d+v) { return true }
                }
            } else if d < 0 && dfs(i+1, j, d+1) { // 字符，扩展一位，注意这里 d 不能为 0
                return true
            }
        }
        if d >= 0 && j < m { // s2[:j] 的原始字符串长度不超过 s1[:i] 的原始字符串长度时，扩展 s2[:j]
            if isDigit(s2[j]) { // 数字
                for q, v := j, 0; q < m && isDigit(s2[q]); q++ {
                    v = v*10 + int(s2[q]&15)
                    if dfs(i, q+1, d-v) {
                        return true
                    }
                }
            } else if d > 0 && dfs(i, j+1, d-1) { // 字符，扩展一位，注意这里 d 不能为 0
                return true
            }
        }
        return false
    }
    return dfs(0, 0, 0)
}

func main() {
    // Example 1:
    // Input: s1 = "internationalization", s2 = "i18n"
    // Output: true
    // Explanation: It is possible that "internationalization" was the original string.
    // - "internationalization" 
    //   -> Split:       ["internationalization"]
    //   -> Do not replace any element
    //   -> Concatenate:  "internationalization", which is s1.
    // - "internationalization"
    //   -> Split:       ["i", "nternationalizatio", "n"]
    //   -> Replace:     ["i", "18",                 "n"]
    //   -> Concatenate:  "i18n", which is s2
    fmt.Println(possiblyEquals("internationalization", "i18n")) // true
    // Example 2:
    // Input: s1 = "l123e", s2 = "44"
    // Output: true
    // Explanation: It is possible that "leetcode" was the original string.
    // - "leetcode" 
    //   -> Split:      ["l", "e", "et", "cod", "e"]
    //   -> Replace:    ["l", "1", "2",  "3",   "e"]
    //   -> Concatenate: "l123e", which is s1.
    // - "leetcode" 
    //   -> Split:      ["leet", "code"]
    //   -> Replace:    ["4",    "4"]
    //   -> Concatenate: "44", which is s2.
    fmt.Println(possiblyEquals("l123e", "44")) // true
    // Example 3:
    // Input: s1 = "a5b", s2 = "c5b"
    // Output: false
    // Explanation: It is impossible.
    // - The original string encoded as s1 must start with the letter 'a'.
    // - The original string encoded as s2 must start with the letter 'c'.
    fmt.Println(possiblyEquals("a5b", "c5b")) // false

    fmt.Println(possiblyEquals("v223u2v84v95v219v748u6u5u18", "32v72u95u354v68v886u7u56v3v4")) // false
}
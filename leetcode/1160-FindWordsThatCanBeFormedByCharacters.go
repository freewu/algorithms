package main

// 1160. Find Words That Can Be Formed by Characters
// You are given an array of strings words and a string chars.

// A string is good if it can be formed by characters from chars (each character can only be used once).

// Return the sum of lengths of all good strings in words.

// Example 1:
// Input: words = ["cat","bt","hat","tree"], chars = "atach"
// Output: 6
// Explanation: The strings that can be formed are "cat" and "hat" so the answer is 3 + 3 = 6.

// Example 2:
// Input: words = ["hello","world","leetcode"], chars = "welldonehoneyr"
// Output: 10
// Explanation: The strings that can be formed are "hello" and "world" so the answer is 5 + 5 = 10.

// Constraints:
//     1 <= words.length <= 1000
//     1 <= words[i].length, chars.length <= 100
//     words[i] and chars consist of lowercase English letters.

import "fmt"

// func countCharacters(words []string, chars string) int {
//     res, mp := 0, make(map[byte]int)
//     for i := range chars {
//         mp[chars[i]]++
//     }
//     for _, word := range words {
//         flag, t := true, make(map[byte]int)
//         for i := range word { // 统计每个单词的字符出现次数
//             t[word[i]]++
//         }
//         for k, v := range t {
//             if c, ok := mp[k]; !ok && c < v { // 没有出现，或数量少于 直接跳出检测
//                 flag = false
//                 break
//             }
//         }
//         if flag {
//             res += len(word)
//         }
//     }
//     return res
// }

func countCharacters(words []string, chars string) int {
    res, n, mp := 0, len(words), map[uint8]int{}
    for _, v := range chars {
        mp[uint8(v)]++
    }
    for i := 0; i < n; i++ {
        flag, t := true, map[uint8]int{}
        for _, v := range words[i] { // 统计每个单词的字符出现次数
            t[uint8(v)]++
        }
        for k := range t {
            if t[k] > mp[k] {
                flag = false
                break
            }
        }
        if flag {
            res += len(words[i])
        }
    }
    return res
}

func countCharacters1(words []string, chars string) int {
    res, mp := 0, [26]int{}
    for i := 0; i < len(chars); i++ {
        mp[int(chars[i]-'a')]++
    }
    for _, word := range words {
        flag, t := true, [26]int{}
        for i := 0; i < len(word); i++ {
            t[int(word[i]-'a')]++
        }
        for i := 0; i < 26; i++ {
            if t[i] > mp[i] {
                flag = false
            }
        }
        if flag {
            res += len(word)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: words = ["cat","bt","hat","tree"], chars = "atach"
    // Output: 6
    // Explanation: The strings that can be formed are "cat" and "hat" so the answer is 3 + 3 = 6.
    fmt.Println(countCharacters([]string{"cat","bt","hat","tree"}, "atach")) // 6
    // Example 2:
    // Input: words = ["hello","world","leetcode"], chars = "welldonehoneyr"
    // Output: 10
    // Explanation: The strings that can be formed are "hello" and "world" so the answer is 5 + 5 = 10.
    fmt.Println(countCharacters([]string{"hello","world","leetcode"}, "welldonehoneyr")) // 10

    fmt.Println(countCharacters1([]string{"cat","bt","hat","tree"}, "atach")) // 6
    fmt.Println(countCharacters1([]string{"hello","world","leetcode"}, "welldonehoneyr")) // 10
}
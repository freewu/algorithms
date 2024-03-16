package main

// 389. Find the Difference
// You are given two strings s and t.
// String t is generated by random shuffling string s and then add one more letter at a random position.
// Return the letter that was added to t.

// Example 1:
// Input: s = "abcd", t = "abcde"
// Output: "e"
// Explanation: 'e' is the letter that was added.

// Example 2:
// Input: s = "", t = "y"
// Output: "y"

// Constraints:
//     0 <= s.length <= 1000
//     t.length == s.length + 1
//     s and t consist of lowercase English letters.

import "fmt"

// func findTheDifference(s string, t string) byte {
//     l := len(s)
//     for i := 0; i < l; i++ {
//         // 发现不同,直接返回
//         if s[i] != t[i] {
//             return t[i]
//         }
//     }
//     // 一定在队尾
//     return t[l]
// }

func findTheDifference(s string, t string) byte {
    // 把两个字符串 和出现次数保存到 map 中
    ms := make(map[byte]int)
    mt := make(map[byte]int)
    for i := 0; i < len(s); i++ {
        ms[s[i]]++
    }
    for i := 0; i < len(t); i++ {
        mt[t[i]]++
    }
    for k,v := range mt {
        // 出现不同即为答案 返回 
        if v != ms[k] {
            return k
        }
    }
    return ' '
}

// best solution
// 用异或
func findTheDifference1(s string, t string) byte {
    var res byte
    for _, c := range s {
        res ^= byte(c)
    }
    for _, c := range t {
        res ^= byte(c)
    } 
    return res
}

func main() {
    fmt.Println(string(findTheDifference("abcd","abcde"))) // e
    fmt.Println(string(findTheDifference("","y"))) // y
    fmt.Println(string(findTheDifference("abcd","abtcd"))) // t

    fmt.Println(string(findTheDifference1("abcd","abcde"))) // e
    fmt.Println(string(findTheDifference1("","y"))) // y
    fmt.Println(string(findTheDifference1("abcd","abtcd"))) // t
}
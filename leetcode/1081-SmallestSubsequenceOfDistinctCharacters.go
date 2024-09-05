package main

// 1081. Smallest Subsequence of Distinct Characters
// Given a string s, return the lexicographically smallest subsequence of s that contains all the distinct characters of s exactly once.

// Example 1:
// Input: s = "bcabc"
// Output: "abc"

// Example 2:
// Input: s = "cbacdcbc"
// Output: "acdb"

// Constraints:
//     1 <= s.length <= 1000
//     s consists of lowercase English letters.

import "fmt"
import "strings"

func smallestSubsequence(s string) string {
    exist := map[byte]struct{}{} // 用来记录字符是否出现过
    res := make([]byte, 0) // 用来记录字符的出现位置 保证最小位
    for i := range s {
        if _, ok := exist[s[i]]; ok { // 字符已出现不需要判断了
            continue
        }
        // repeat to test if the tail of ret is bigger then byte visiting now and the tail appears again after 
        // if yes, remove the tail
        for len(res) != 0 {
            tail := len(res) - 1
            // strings.LastIndexByte 最后出现的位置
            if res[tail] > s[i] && strings.LastIndexByte(s, res[tail]) > i {
                delete(exist, res[tail])
                res = res[0:tail]
                continue
            }
            break
        }
        exist[s[i]] = struct{}{}
        res = append(res, s[i])
    }
    return string(res)
}

func main() {
    fmt.Println(smallestSubsequence("bcabc")) // "abc"
    fmt.Println(smallestSubsequence("cbacdcbc")) // "acdb"
}
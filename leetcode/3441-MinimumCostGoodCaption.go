package main

// 3441. Minimum Cost Good Caption
// You are given a string caption of length n. 
// A good caption is a string where every character appears in groups of at least 3 consecutive occurrences.

// For example:
//     "aaabbb" and "aaaaccc" are good captions.
//     "aabbb" and "ccccd" are not good captions.

// You can perform the following operation any number of times:

// Choose an index i (where 0 <= i < n) and change the character at that index to either:
//     The character immediately before it in the alphabet (if caption[i] != 'a').
//     The character immediately after it in the alphabet (if caption[i] != 'z').

// Your task is to convert the given caption into a good caption using the minimum number of operations, 
// and return it. If there are multiple possible good captions, return the lexicographically smallest one among them. 
// If it is impossible to create a good caption, return an empty string "".

// Example 1:
// Input: caption = "cdcd"
// Output: "cccc"
// Explanation:
// It can be shown that the given caption cannot be transformed into a good caption with fewer than 2 operations. The possible good captions that can be created using exactly 2 operations are:
// "dddd": Change caption[0] and caption[2] to their next character 'd'.
// "cccc": Change caption[1] and caption[3] to their previous character 'c'.
// Since "cccc" is lexicographically smaller than "dddd", return "cccc".

// Example 2:
// Input: caption = "aca"
// Output: "aaa"
// Explanation:
// It can be proven that the given caption requires at least 2 operations to be transformed into a good caption. The only good caption that can be obtained with exactly 2 operations is as follows:
// Operation 1: Change caption[1] to 'b'. caption = "aba".
// Operation 2: Change caption[1] to 'a'. caption = "aaa".
// Thus, return "aaa".

// Example 3:
// Input: caption = "bc"
// Output: ""
// Explanation:
// It can be shown that the given caption cannot be converted to a good caption by using any number of operations.

// Constraints:
//     1 <= caption.length <= 5 * 10^4
//     caption consists only of lowercase English letters.

import "fmt"
import "slices"
import "bytes"

func minCostGoodCaption(s string) string {
    n := len(s)
    if n < 3 { return "" }
    f := make([]int, n+1)
    f[n-1], f[n-2] = 1 << 31, 1 << 31
    t := make([]byte, n + 1)
    size := make([]uint8, n)
    for i := n - 3; i >= 0; i-- {
        sub := []byte(s[i : i+3])
        slices.Sort(sub)
        a, b, c := sub[0], sub[1], sub[2]
        s3 := int(t[i+3])
        res := f[i+3] + int(c-a)
        mask := int(b)<<24 | s3<<16 | s3<<8 | s3 // 4 个 byte 压缩成一个 int，方便比较字典序
        size[i] = 3
        if i + 4 <= n {
            sub := []byte(s[i : i+4])
            slices.Sort(sub)
            a, b, c, d := sub[0], sub[1], sub[2], sub[3]
            s4 := int(t[i+4])
            res4 := f[i+4] + int(c-a+d-b)
            mask4 := int(b)<<24 | int(b)<<16 | s4<<8 | s4
            if res4 < res || res4 == res && mask4 < mask {
                res, mask = res4, mask4
                size[i] = 4
            }
        }
        if i + 5 <= n {
            sub := []byte(s[i : i+5])
            slices.Sort(sub)
            a, b, c, d, e := sub[0], sub[1], sub[2], sub[3], sub[4]
            res5 := f[i+5] + int(d-a+e-b)
            mask5 := int(c)<<24 | int(c)<<16 | int(c)<<8 | int(t[i+5])
            if res5 < res || res5 == res && mask5 < mask {
                res, mask = res5, mask5
                size[i] = 5
            }
        }
        f[i] = res
        t[i] = byte(mask >> 24)
    }
    res := make([]byte, 0, n)
    for i := 0; i < n; i += int(size[i]) {
        res = append(res, bytes.Repeat([]byte{t[i]}, int(size[i]))...)
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: caption = "cdcd"
    // Output: "cccc"
    // Explanation:
    // It can be shown that the given caption cannot be transformed into a good caption with fewer than 2 operations. The possible good captions that can be created using exactly 2 operations are:
    // "dddd": Change caption[0] and caption[2] to their next character 'd'.
    // "cccc": Change caption[1] and caption[3] to their previous character 'c'.
    // Since "cccc" is lexicographically smaller than "dddd", return "cccc".
    fmt.Println(minCostGoodCaption("cdcd")) // "cccc"
    // Example 2:
    // Input: caption = "aca"
    // Output: "aaa"
    // Explanation:
    // It can be proven that the given caption requires at least 2 operations to be transformed into a good caption. The only good caption that can be obtained with exactly 2 operations is as follows:
    // Operation 1: Change caption[1] to 'b'. caption = "aba".
    // Operation 2: Change caption[1] to 'a'. caption = "aaa".
    // Thus, return "aaa".
    fmt.Println(minCostGoodCaption("aca")) // "aaa"
    // Example 3:
    // Input: caption = "bc"
    // Output: ""
    // Explanation:
    // It can be shown that the given caption cannot be converted to a good caption by using any number of operations.
    fmt.Println(minCostGoodCaption("bc")) // ""

    fmt.Println(minCostGoodCaption("bluefrog")) // "fffffooo"
    fmt.Println(minCostGoodCaption("leetcode")) // "eeeedddd"
}
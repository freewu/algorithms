package main

// 541. Reverse String II
// Given a string s and an integer k, 
// reverse the first k characters for every 2k characters counting from the start of the string.

// If there are fewer than k characters left, reverse all of them. 
// If there are less than 2k but greater than or equal to k characters, 
// then reverse the first k characters and leave the other as original.

// Example 1:
// Input: s = "abcdefg", k = 2
// Output: "bacdfeg"

// Example 2:
// Input: s = "abcd", k = 2
// Output: "bacd"
 
// Constraints:
//     1 <= s.length <= 10^4
//     s consists of only lowercase English letters.
//     1 <= k <= 10^4

import "fmt"

// func reverseStr(s string, k int) string {
//     res := []byte(s)
//     left, right := 0, k - 1
//     for left < right {
//         res[left], res[right] = res[right], res[left]
//         left++
//         right--
//     }
//     return string(res)
// }

func reverseStr(s string, k int) string {
    if len(s) < k { // 如果剩余字符少于 k 个，则将剩余字符全部反转
        r := ""
        for i := len(s) - 1; i >= 0; i-- {
            r = r + string(s[i])
        }
        return r
    }
    r := []byte(s)
    for i := 0; i < len(s); i = i + 2 * k { // 如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。
        l, m := i, i + k-1
        if m >= len(s) {
            m = len(s)-1
        }
        for l < m {
            r[l], r[m] = r[m], r[l]
            l++
            m--
        }
    }
    return string(r)
}

func reverseStr1(s string, k int) string {
    chars, n := []byte(s), len(s)  // Convert string to slice of runes to handle Unicode characters
    reverse := func (chars []byte, left int, right int) {
        for left < right {
            chars[left], chars[right] = chars[right], chars[left]
            left++
            right--
        }
    }
    for i := 0; i < n; i += 2 * k {
        end := i + k
        if end > n {
            end = n
        }
        reverse(chars, i, end-1)
    }
    return string(chars)
}

func main() {
    // Example 1:
    // Input: s = "abcdefg", k = 2
    // Output: "bacdfeg"
    fmt.Println(reverseStr("abcdefg",2)) // bacdfeg
    // Example 2:
    // Input: s = "abcd", k = 2
    // Output: "bacd"
    fmt.Println(reverseStr("abcd",2)) // bacd

    fmt.Println(reverseStr1("abcdefg",2)) // bacdfeg
    fmt.Println(reverseStr1("abcd",2)) // bacd
}
package main

// 2663. Lexicographically Smallest Beautiful String
// A string is beautiful if:
//     It consists of the first k letters of the English lowercase alphabet.
//     It does not contain any substring of length 2 or more which is a palindrome.

// You are given a beautiful string s of length n and a positive integer k.
// Return the lexicographically smallest string of length n, which is larger than s and is beautiful. 
// If there is no such string, return an empty string.

// A string a is lexicographically larger than a string b (of the same length) if in the first position where a and b differ, 
// a has a character strictly larger than the corresponding character in b.
//     For example, "abcd" is lexicographically larger than "abcc" 
//     because the first position they differ is at the fourth character, 
//     and d is greater than c.

// Example 1:
// Input: s = "abcz", k = 26
// Output: "abda"
// Explanation: The string "abda" is beautiful and lexicographically larger than the string "abcz".
// It can be proven that there is no string that is lexicographically larger than the string "abcz", beautiful, and lexicographically smaller than the string "abda".

// Example 2:
// Input: s = "dc", k = 4
// Output: ""
// Explanation: It can be proven that there is no string that is lexicographically larger than the string "dc" and is beautiful.

// Constraints:
//     1 <= n == s.length <= 10^5
//     4 <= k <= 26
//     s is a beautiful string.

import "fmt"

func smallestBeautifulString(s string, k int) string {
    bs, n := []byte(s), len(s) - 1
    bs[n]++
    for i := n; i >= 0 && i <= n; {
        if bs[i] == byte('a' + k) {
            // 需要进位处理
            if i == 0 { 
                return "" // 前面也没法处理进位
            }
            bs[i] = 'a' // 当前位置0
            bs[i-1]++   // 上一位 +1
            i--         // 现在处理上一位
        } else if (i-1 >= 0 && bs[i] == bs[i-1]) || (i-2 >= 0 && bs[i] == bs[i-2]) {
            // 不需要进位处理，但是需要判断回文
            // 如果有回文，只能将当前的值 +1，这时候不能移动 i，因为 +1 之后可能还是回文，也可能需要进位
            bs[i]++ //
        } else {
            // 不需要进位处理,也没有回文，说明这一位暂时是稳住了，可以向后走了
            i++
        }
    }
    return string(bs)
}

func main() {
    // Example 1:
    // Input: s = "abcz", k = 26
    // Output: "abda"
    // Explanation: The string "abda" is beautiful and lexicographically larger than the string "abcz".
    // It can be proven that there is no string that is lexicographically larger than the string "abcz", beautiful, and lexicographically smaller than the string "abda".
    fmt.Println(smallestBeautifulString("abcz", 26)) // "abda"
    // Example 2:
    // Input: s = "dc", k = 4
    // Output: ""
    // Explanation: It can be proven that there is no string that is lexicographically larger than the string "dc" and is beautiful.
    fmt.Println(smallestBeautifulString("dc", 4)) // ""
}
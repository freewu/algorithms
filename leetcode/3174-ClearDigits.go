package main

// 3174. Clear Digits
// You are given a string s.

// Your task is to remove all digits by doing this operation repeatedly:
//     Delete the first digit and the closest non-digit character to its left.

// Return the resulting string after removing all digits.

// Example 1:
// Input: s = "abc"
// Output: "abc"
// Explanation:
// There is no digit in the string.

// Example 2:
// Input: s = "cb34"
// Output: ""
// Explanation:
// First, we apply the operation on s[2], and s becomes "c4".
// Then we apply the operation on s[1], and s becomes "".

// Constraints:
//     1 <= s.length <= 100
//     s consists only of lowercase English letters and digits.
//     The input is generated such that it is possible to delete all digits.

import "fmt"

func clearDigits(s string) string {
    stack := []rune{}
    for _, v := range s {
        if v >= '0' && v <= '9' {
            if len(stack) > 0 {
                stack = stack[:len(stack) - 1] // 删除 第一个数字字符 以及它左边 最近 的 非数字 字符
            }
        } else {
            stack = append(stack, v)
        }
    }
    return string(stack)
}

func main() {
    // Example 1:
    // Input: s = "abc"
    // Output: "abc"
    // Explanation:
    // There is no digit in the string.
    fmt.Println(clearDigits("abc")) // "abc"
    // Example 2:
    // Input: s = "cb34"
    // Output: ""
    // Explanation:
    // First, we apply the operation on s[2], and s becomes "c4".
    // Then we apply the operation on s[1], and s becomes "".
    fmt.Println(clearDigits("cb34")) // ""
    fmt.Println(clearDigits("34cb")) // "cb"
}
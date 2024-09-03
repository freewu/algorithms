package main

// 1003. Check If Word Is Valid After Substitutions
// Given a string s, determine if it is valid.

// A string s is valid if, starting with an empty string t = "", 
// you can transform t into s after performing the following operation any number of times:
//     Insert string "abc" into any position in t. 
//     More formally, t becomes tleft + "abc" + tright, where t == tleft + tright. 
//     Note that tleft and tright may be empty.

// Return true if s is a valid string, otherwise, return false.

// Example 1:
// Input: s = "aabcbc"
// Output: true
// Explanation:
// "" -> "abc" -> "aabcbc"
// Thus, "aabcbc" is valid.

// Example 2:
// Input: s = "abcabcababcc"
// Output: true
// Explanation:
// "" -> "abc" -> "abcabc" -> "abcabcabc" -> "abcabcababcc"
// Thus, "abcabcababcc" is valid.

// Example 3:
// Input: s = "abccba"
// Output: false
// Explanation: It is impossible to get "abccba" using the operation.

// Constraints:
//     1 <= s.length <= 2 * 10^4
//     s consists of letters 'a', 'b', and 'c'

import "fmt"

func isValid(s string) bool {
    isDeleted := true
    for isDeleted {
        isDeleted = false
        for i := 0; i < len(s) - 2; i++ {
            if s[i: i + 3] == "abc" { // 如果是 abc 则做截取操作
                s = s[:i] + s[i + 3:]
                isDeleted = true
                break
            }
        }
    }
    return len(s) == 0
}

// stack  abc 看成一组，遇到c 时需要从栈依次 pop 出b a 
func isValid1(s string) bool {
    stack := []byte{}
    for _, v := range s {
        if v == 'c' { // pop 两个值，且依次为 b a
            if len(stack) < 2 {
                return false
            }
            a, b := stack[len(stack)-2], stack[len(stack)-1] // pop twice
            stack = stack[:len(stack)-2]
            if a != 'a' || b != 'b' {
                return false
            } 
        } else {
            stack = append(stack, byte(v))
        }
    }
    return len(stack) == 0
}

func main() {
    // Example 1:
    // Input: s = "aabcbc"
    // Output: true
    // Explanation:
    // "" -> "abc" -> "aabcbc"
    // Thus, "aabcbc" is valid.
    fmt.Println(isValid("aabcbc")) // true
    // Example 2:
    // Input: s = "abcabcababcc"
    // Output: true
    // Explanation:
    // "" -> "abc" -> "abcabc" -> "abcabcabc" -> "abcabcababcc"
    // Thus, "abcabcababcc" is valid.
    fmt.Println(isValid("abcabcababcc")) // true
    // Example 3:
    // Input: s = "abccba"
    // Output: false
    // Explanation: It is impossible to get "abccba" using the operation.
    fmt.Println(isValid("abccba")) // false

    fmt.Println(isValid1("aabcbc")) // true
    fmt.Println(isValid1("abcabcababcc")) // true
    fmt.Println(isValid1("abccba")) // false
}
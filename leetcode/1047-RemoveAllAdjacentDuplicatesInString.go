package main

// 1047. Remove All Adjacent Duplicates In String
// You are given a string s consisting of lowercase English letters. 
// A duplicate removal consists of choosing two adjacent and equal letters and removing them.

// We repeatedly make duplicate removals on s until we no longer can.

// Return the final string after all such duplicate removals have been made. 
// It can be proven that the answer is unique.

// Example 1:
// Input: s = "abbaca"
// Output: "ca"
// Explanation: 
// For example, in "abbaca" we could remove "bb" since the letters are adjacent and equal, and this is the only possible move.  The result of this move is that the string is "aaca", of which only "aa" is possible, so the final string is "ca".

// Example 2:
// Input: s = "azxxzy"
// Output: "ay"
 
// Constraints:
//     1 <= s.length <= 10^5
//     s consists of lowercase English letters.

import "fmt"

// stack
func removeDuplicates(s string) string {
    stack := make([] byte, 0)
    for i := 0; i < len(s); i++ {
        if len(stack) == 0 { // 栈空了，直接进行栈
            stack = append(stack, s[i])
        } else {
            top := stack[len(stack) - 1]
            if top == s[i] { // 相邻字符相同则删除
                stack = stack[:len(stack) - 1]
            } else {
                stack = append(stack, s[i]) // 否则入栈
            }
        }
    }
    return string(stack)
}

func main() {
    // Example 1:
    // Input: s = "abbaca"
    // Output: "ca"
    // Explanation: 
    // For example, in "abbaca" we could remove "bb" since the letters are adjacent and equal, and this is the only possible move.  The result of this move is that the string is "aaca", of which only "aa" is possible, so the final string is "ca".
    fmt.Println(removeDuplicates("abbaca")) // "ca"
    // Example 2:
    // Input: s = "azxxzy"
    // Output: "ay"
    fmt.Println(removeDuplicates("azxxzy")) // "ay"
}
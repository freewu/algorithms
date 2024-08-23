package main

// 856. Score of Parentheses
// Given a balanced parentheses string s, return the score of the string.
// The score of a balanced parentheses string is based on the following rule:
//     "()" has score 1.
//     AB has score A + B, where A and B are balanced parentheses strings.
//     (A) has score 2 * A, where A is a balanced parentheses string.

// Example 1:
// Input: s = "()"
// Output: 1

// Example 2:
// Input: s = "(())"
// Output: 2

// Example 3:
// Input: s = "()()"
// Output: 2

// Constraints:
//     2 <= s.length <= 50
//     s consists of only '(' and ')'.
//     s is a balanced parentheses string.

import "fmt"

// stack
func scoreOfParentheses(s string) int {
    stack := make([]int, 0)
    stack = append(stack, 0)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, c := range s {
        if c == '(' { // push
            stack = append(stack, 0)
            continue
        }
        a, b := stack[len(stack)-1], stack[len(stack)-2]
        stack = stack[:len(stack)-2]
        stack = append(stack, b + max(2 * a, 1))
    }
    return stack[0]
}

func scoreOfParentheses1(s string) int {
    res, balance := 0, 0
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            balance++
            continue
        }
        balance--
        if s[i-1] == '(' {
            res += 1 << balance
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "()"
    // Output: 1
    fmt.Println(scoreOfParentheses("()")) // 1
    // Example 2:
    // Input: s = "(())"
    // Output: 2
    fmt.Println(scoreOfParentheses("(())")) // 2
    // Example 3:
    // Input: s = "()()"
    // Output: 2
    fmt.Println(scoreOfParentheses("()()")) // 2

    fmt.Println(scoreOfParentheses("()")) // 1     1
    fmt.Println(scoreOfParentheses("(())")) // 2   1 * 2
    fmt.Println(scoreOfParentheses("()()")) // 2   1 + 1
}
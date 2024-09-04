package main

// 1021. Remove Outermost Parentheses
// A valid parentheses string is either empty "", "(" + A + ")", or A + B, 
// where A and B are valid parentheses strings, and + represents string concatenation.
//     For example, "", "()", "(())()", and "(()(()))" are all valid parentheses strings.

// A valid parentheses string s is primitive if it is nonempty, 
// and there does not exist a way to split it into s = A + B, with A and B nonempty valid parentheses strings.

// Given a valid parentheses string s, consider its primitive decomposition: 
//     s = P1 + P2 + ... + Pk, where Pi are primitive valid parentheses strings.

// Return s after removing the outermost parentheses of every primitive string in the primitive decomposition of s.

// Example 1:
// Input: s = "(()())(())"
// Output: "()()()"
// Explanation: 
// The input string is "(()())(())", with primitive decomposition "(()())" + "(())".
// After removing outer parentheses of each part, this is "()()" + "()" = "()()()".

// Example 2:
// Input: s = "(()())(())(()(()))"
// Output: "()()()()(())"
// Explanation: 
// The input string is "(()())(())(()(()))", with primitive decomposition "(()())" + "(())" + "(()(()))".
// After removing outer parentheses of each part, this is "()()" + "()" + "()(())" = "()()()()(())".

// Example 3:
// Input: s = "()()"
// Output: ""
// Explanation: 
// The input string is "()()", with primitive decomposition "()" + "()".
// After removing outer parentheses of each part, this is "" + "" = "".

// Constraints:
//     1 <= s.length <= 10^5
//     s[i] is either '(' or ')'.
//     s is a valid parentheses string.

import "fmt"

func removeOuterParentheses(s string) string {
    res, count := "", 0
    for _, v := range s {
        if v == '(' { 
            count++
        } else {
            count--
        }
        if !(count == 0 && v == ')') && !(count == 1 && v == '(') {
            res = res + string(v)
        }
    }
    return res
}

func removeOuterParentheses1(s string) string {
    stack, left := make([]byte, 0), 0
    for _, v := range []byte(s) {
        if left == 0 {
            left++
        } else {
            if v == '(' {
                stack = append(stack, v)
                left++
            } else {
                if left == 1 {
                    left--
                    continue
                }
                stack = append(stack, v)
                left--
            }
        }
    }
    return string(stack)
}

func main() {
    // Example 1:
    // Input: s = "(()())(())"
    // Output: "()()()"
    // Explanation: 
    // The input string is "(()())(())", with primitive decomposition "(()())" + "(())".
    // After removing outer parentheses of each part, this is "()()" + "()" = "()()()".
    fmt.Println(removeOuterParentheses("(()())(())")) // "()()()"
    // Example 2:
    // Input: s = "(()())(())(()(()))"
    // Output: "()()()()(())"
    // Explanation: 
    // The input string is "(()())(())(()(()))", with primitive decomposition "(()())" + "(())" + "(()(()))".
    // After removing outer parentheses of each part, this is "()()" + "()" + "()(())" = "()()()()(())".
    fmt.Println(removeOuterParentheses("(()())(())(()(()))")) // "()()()()(())"
    // Example 3:
    // Input: s = "()()"
    // Output: ""
    // Explanation: 
    // The input string is "()()", with primitive decomposition "()" + "()".
    // After removing outer parentheses of each part, this is "" + "" = "".
    fmt.Println(removeOuterParentheses("()()")) // ""

    fmt.Println(removeOuterParentheses1("(()())(())")) // "()()()"
    fmt.Println(removeOuterParentheses1("(()())(())(()(()))")) // "()()()()(())"
    fmt.Println(removeOuterParentheses1("()()")) // ""
}
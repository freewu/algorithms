package main

// 32. Longest Valid Parentheses
// Given a string containing just the characters '(' and ')', 
// return the length of the longest valid (well-formed) parentheses substring.

// Example 1:
// Input: s = "(()"
// Output: 2
// Explanation: The longest valid parentheses substring is "()".

// Example 2:
// Input: s = ")()())"
// Output: 4
// Explanation: The longest valid parentheses substring is "()()".

// Example 3:
// Input: s = ""
// Output: 0
 
// Constraints:
//     0 <= s.length <= 3 * 104
//     s[i] is '(', or ')'.

import "fmt"

// stack
func longestValidParentheses(s string) int {
    stack, res := []int{}, 0 
    stack = append(stack, -1)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            stack = append(stack, i) // 发现 ( 入栈 push
        } else { // )
            stack = stack[:len(stack)-1] // 出栈操作 pop
            if len(stack) == 0 { // 如果没有
                stack = append(stack, i) // ?
            } else {
                res = max(res, i - stack[len(stack)-1])
            }
        }
    }
    return res
}

// 双指针
func longestValidParentheses1(s string) int {
    left, right, res := 0, 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            left++
        } else {
            right++
        }
        if left == right {
            res = max(res, 2 * right)
        } else if right > left {
            left, right = 0, 0
        }
    }
    left, right = 0, 0
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == '(' {
            left++
        } else {
            right++
        }
        if left == right {
            res = max(res, 2*left)
        } else if left > right {
            left, right = 0, 0
        }
    }
    return res
}

// dp
func longestValidParentheses2(s string) int {
    n, res := len(s), 0
    dp := make([]int, n) // dp[i]: 以第i个字符结尾的字符串的最长有效括号长度
    for i := 1; i < n; i++{
        if s[i] == ')' {
            if s[i-1] == '(' {
                dp[i] = 2
                if i - 2 >= 0{
                    dp[i] += dp[i-2]
                }
            } else if dp[i-1] > 0 && i - dp[i-1] - 1 >= 0 && s[i - dp[i-1] - 1] == '(' {
                dp[i] = dp[i-1] + 2
                if i - dp[i-1] - 2 >= 0{
                    dp[i] += dp[i - dp[i-1] - 2]
                }
            }
            if dp[i] > res {
                res = dp[i]
            }
        }
    }
    return res
}

func main() {
    fmt.Printf("longestValidParentheses(\"(()\") = %v\n",longestValidParentheses("(()")) // 2
    fmt.Printf("longestValidParentheses(\")()())\") = %v\n",longestValidParentheses(")()())")) // 4
    fmt.Printf("longestValidParentheses(\"\") = %v\n",longestValidParentheses("")) // 0
    fmt.Printf("longestValidParentheses(\")(\") = %v\n",longestValidParentheses(")(")) // 0

    fmt.Printf("longestValidParentheses1(\"(()\") = %v\n",longestValidParentheses1("(()")) // 2
    fmt.Printf("longestValidParentheses1(\")()())\") = %v\n",longestValidParentheses1(")()())")) // 4
    fmt.Printf("longestValidParentheses1(\"\") = %v\n",longestValidParentheses1("")) // 0
    fmt.Printf("longestValidParentheses1(\")(\") = %v\n",longestValidParentheses1(")(")) // 0

    fmt.Printf("longestValidParentheses2(\"(()\") = %v\n",longestValidParentheses2("(()")) // 2
    fmt.Printf("longestValidParentheses2(\")()())\") = %v\n",longestValidParentheses2(")()())")) // 4
    fmt.Printf("longestValidParentheses2(\"\") = %v\n",longestValidParentheses2("")) // 0
    fmt.Printf("longestValidParentheses2(\")(\") = %v\n",longestValidParentheses2(")(")) // 0
}

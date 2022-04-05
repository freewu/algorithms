package main

import "fmt"

/**
32. Longest Valid Parentheses
Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.

Constraints:

	0 <= s.length <= 3 * 10^4
	s[i] is '(', or ')'.

Example 1:

	Input: s = "(()"
	Output: 2
	Explanation: The longest valid parentheses substring is "()".

Example 2:

	Input: s = ")()())"
	Output: 4
	Explanation: The longest valid parentheses substring is "()()".

Example 3:

	Input: s = ""
	Output: 0

*/

// 解法一 栈
func longestValidParentheses(s string) int {
	var stack []int
	res := 0
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i) // 发现 ( 入栈
		} else { // )
			stack = stack[:len(stack)-1] // 出栈操作
			if len(stack) == 0 { // 如果没有
				stack = append(stack, i) // ?
			} else {
				res = max(res, i-stack[len(stack)-1])
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 解法二 双指针
func longestValidParentheses1(s string) int {
	left, right, maxLength := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2*right)
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
			maxLength = max(maxLength, 2*left)
		} else if left > right {
			left, right = 0, 0
		}
	}
	return maxLength
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
}

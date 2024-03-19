package main

// 20. Valid Parentheses
// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
// An input string is valid if:
//     Open brackets must be closed by the same type of brackets.
//     Open brackets must be closed in the correct order.

// Example 1:
//     Input: s = "()"
//     Output: true

// Example 2:
//     Input: s = "()[]{}"
//     Output: true

// Example 3:
//     Input: s = "(]"
//     Output: false

// Constraints:
//     1 <= s.length <= 10^4
//     s consists of parentheses only '()[]{}'.

import "fmt"

func isValid(s string) bool {
    // 判断长度是否小于2
    if len(s) < 2 {
        return false
    }
    // 用stack就可很好处理golang没有原生的stack 这里使用一个array 和 int来处理
    var a []string
    var l = 0

    for i := 0; i < len(s); i++ {
        // 遇到 ([{ 就入栈
        if '(' == s[i] || '[' == s[i] || '{' == s[i] {
            a = append(a, string(s[i]))
            l++
        }
        // 遇到)]} 就出栈 比对
        if ')' == s[i] || ']' == s[i] || '}' == s[i] {
            // 没有入栈过
            if 0 == l {
                return false
            }
            if ')' == s[i] && "(" != a[l-1] {
                return false
            }
            if ']' == s[i] && "[" != a[l-1] {
                return false
            }
            if '}' == s[i] && "{" != a[l-1] {
                return false
            }
            a = append(a[0 : l-1])
            l--
        }
    }
    // 数组里没有完全出栈
    if 0 != l {
        return false
    }
    return true
}

// best solution
func isValidBest(s string) bool {
    // 空字符串直接返回 true
    if len(s) == 0 {
        return true
    }
    stack := make([]rune, 0)
    for _, v := range s {
        if (v == '[') || (v == '(') || (v == '{') {
            stack = append(stack, v)
        } else if ((v == ']') && len(stack) > 0 && stack[len(stack)-1] == '[') ||
            ((v == ')') && len(stack) > 0 && stack[len(stack)-1] == '(') ||
            ((v == '}') && len(stack) > 0 && stack[len(stack)-1] == '{') {
            stack = stack[:len(stack)-1]
        } else {
            return false
        }
    }
    return len(stack) == 0
}

// stack + map
func isValid1(s string) bool {
    var pairs map[string]string = map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}
	var stack []string
	for _, c := range s {
		ch := string(c)
		size := len(stack)
		if left, ok := pairs[ch]; ok {
			if size == 0 || stack[size-1] != left {
				return false
			}
			stack = stack[:size-1]
		}else {
    		stack = append(stack, ch)
        }
	}
	return len(stack) == 0
}

func main() {
    fmt.Printf("isValid(\"((\") = %v\n",isValid("((")) // false
    fmt.Printf("isValid(\"(\") = %v\n",isValid("(")) // false
    fmt.Printf("isValid(\"(+\") = %v\n",isValid("()")) // true
    fmt.Printf("isValid(\"({[()]})\") = %v\n",isValid("({[()]})")) // true
    fmt.Printf("isValid(\"({[()}])\") = %v\n",isValid("({[()}])")) // false

    fmt.Printf("isValidBest(\"((\") = %v\n",isValidBest("((")) // false
    fmt.Printf("isValidBest(\"(\") = %v\n",isValidBest("(")) // false
    fmt.Printf("isValidBest(\"(+\") = %v\n",isValidBest("()")) // true
    fmt.Printf("isValidBest(\"({[()]})\") = %v\n",isValidBest("({[()]})")) // true
    fmt.Printf("isValidBest(\"({[()}])\") = %v\n",isValidBest("({[()}])")) // false

    fmt.Printf("isValid1(\"((\") = %v\n",isValid1("((")) // false
    fmt.Printf("isValid1(\"(\") = %v\n",isValid1("(")) // false
    fmt.Printf("isValid1(\"(+\") = %v\n",isValid1("()")) // true
    fmt.Printf("isValid1(\"({[()]})\") = %v\n",isValid1("({[()]})")) // true
    fmt.Printf("isValid1(\"({[()}])\") = %v\n",isValid1("({[()}])")) // false
}

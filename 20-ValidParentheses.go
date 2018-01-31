package main

/*
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

The brackets must close in the correct order, "()" and "()[]{}" are all valid but "(]" and "([)]" are not.
*/
import (
	"fmt"
)

func isValid(s string) bool {
	// 判断长度是否小于2
	if len(s) < 2 {
		return false
	}
	// 用stack就可很好处理golang没有原生的stack 这里使用一个array 和 int来处理
	var a = []string{}
	var l = 0

	for i := 0; i < len(s); i++ {
		// 遇到 ([{ 就入栈
		if '(' == s[i] || '[' == s[i] || '{' == s[i] {
			a = append(a,string(s[i]))
			l++
		}
		// 遇到)]} 就出栈 比对
		if ')' == s[i] || ']' == s[i] || '}' == s[i] {
			// 没有入栈过
			if 0 == l {
				return false
			}
			if ')' == s[i] && "(" != a[l - 1] {
				return false
			}
			if ']' == s[i] && "[" != a[l - 1] {
				return false
			}
			if '}' == s[i] && "{" != a[l - 1] {
				return false
			}
			a = append(a[0:l - 1])
			l--
		}
	}
	// 数组里没有完全出栈
	if 0 != l {
		return false
	}
    return true
}

func main() {
	fmt.Println(isValid("(("))
	fmt.Println(isValid("("))
	fmt.Println(isValid("()"))
	fmt.Println(isValid("({[()]})"))
	fmt.Println(isValid("({[()}])"))
}
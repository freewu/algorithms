package main

// 1249. Minimum Remove to Make Valid Parentheses
// Given a string s of '(' , ')' and lowercase English characters.
// Your task is to remove the minimum number of parentheses ( '(' or ')', in any positions ) so that the resulting parentheses string is valid and return any valid string.
// Formally, a parentheses string is valid if and only if:
//     It is the empty string, contains only lowercase characters, or
//     It can be written as AB (A concatenated with B), where A and B are valid strings, or
//     It can be written as (A), where A is a valid string.
 
// Example 1:
// Input: s = "lee(t(c)o)de)"
// Output: "lee(t(c)o)de"
// Explanation: "lee(t(co)de)" , "lee(t(c)ode)" would also be accepted.

// Example 2:
// Input: s = "a)b(c)d"
// Output: "ab(c)d"

// Example 3:
// Input: s = "))(("
// Output: ""
// Explanation: An empty string is also valid.
 
// Constraints:
//     1 <= s.length <= 10^5
//     s[i] is either'(' , ')', or lowercase English letter.

import "fmt"
import "strings"

// stack
func minRemoveToMakeValid(s string) string {
    stack, todel := make([]int,0), make([]int,0)
    for i, c := range s {
        if c == '('{
            // 入栈 push
            stack = append(stack, i)
            continue
        }
        if c == ')'{
            // 如果 没有 ( 就出现了 ) 加入到清除列表
            if len(stack) <= 0 {
                todel = append(todel, i)
            } else {
                // 出栈 pop
                stack = stack[:len(stack)-1]
            }
        }
    }
    todel = append(todel, stack...) // 只有 ( 没有 ）的都加入到待删除列表中
    //fmt.Println(todel)
    res := []byte{}
    for i, j := 0, 0; j < len(s); j++{
        if i < len(todel) && todel[i] == j {
            i++
            continue
        }
        res = append(res, s[j])
    }
    return string(res)
}

func minRemoveToMakeValid1(s string) string {
    stack1, stack2, res := []int{}, []int{}, []byte(s)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' { // 遇到 ( 入栈
			stack1 = append(stack1, i)
		} else if s[i] == ')' {
			if len(stack1) > 0 {
				stack1 = stack1[:len(stack1)-1] // 匹配到 ) 出栈
			} else {
				stack2 = append(stack2, i) // 没有配对的 ) 入栈到 stack2
			}
		}
	}
    // 将需要替换的 '('，')'  替换成 #
	for _, i := range stack1 { res[i] = '#'; }
	for _, i := range stack2 { res[i] = '#'; }
	return strings.ReplaceAll(string(res), "#", "")
}

func main() {
    // Explanation: "lee(t(co)de)" , "lee(t(c)ode)" would also be accepted.
    fmt.Println(minRemoveToMakeValid("lee(t(co)de)")) // lee(t(co)de)

    fmt.Println(minRemoveToMakeValid("a)b(c)d")) // ab(c)d

    // Explanation: An empty string is also valid.
    fmt.Println(minRemoveToMakeValid("))((")) // ""

    fmt.Println(minRemoveToMakeValid1("lee(t(co)de)")) // lee(t(co)de)
    fmt.Println(minRemoveToMakeValid1("a)b(c)d")) // ab(c)d
    fmt.Println(minRemoveToMakeValid1("))((")) // ""
}
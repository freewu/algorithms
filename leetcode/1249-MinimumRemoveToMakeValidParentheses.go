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

func main() {
    // Explanation: "lee(t(co)de)" , "lee(t(c)ode)" would also be accepted.
    fmt.Println(minRemoveToMakeValid("lee(t(co)de)")) // lee(t(co)de)

    fmt.Println(minRemoveToMakeValid("a)b(c)d")) // ab(c)d

    // Explanation: An empty string is also valid.
    fmt.Println(minRemoveToMakeValid("))((")) // ""
}
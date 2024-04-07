package main

// 678. Valid Parenthesis String
// Given a string s containing only three types of characters: '(', ')' and '*', return true if s is valid.
// The following rules define a valid string:
//     Any left parenthesis '(' must have a corresponding right parenthesis ')'.
//     Any right parenthesis ')' must have a corresponding left parenthesis '('.
//     Left parenthesis '(' must go before the corresponding right parenthesis ')'.
//     '*' could be treated as a single right parenthesis ')' or a single left parenthesis '(' or an empty string "".
    
// Example 1:
// Input: s = "()"
// Output: true
    
// Example 2:
// Input: s = "(*)"
// Output: true

// Example 3:
// Input: s = "(*))"
// Output: true
 
// Constraints:
//     1 <= s.length <= 100
//     s[i] is '(', ')' or '*'.

import "fmt"

// func checkValidString(s string) bool {
//     // 一个空字符串也被视为有效字符串
//     if len(s) == 0 {
//         return true
//     }

//     return true
// }

// stack
func checkValidString(s string) bool {
    // 找出 原始字符 的 ( 和 )
    traverse := func(s string) []byte {
        stack, rest, pairs := []byte{}, []byte{}, map[byte]byte { ')': '(' }
        for i := 0; i < len(s); i ++ {
            c := s[i]
            if c == '*' {
                rest = append(rest, c)
                continue
            }
            if pairs[c] > 0 { // 结束括号
                if len(stack) == 0 { // 如果当前是空栈 则 加入剩余字节数组
                    rest = append(rest, c)
                    continue
                }
                stack = stack[:len(stack)-1] // 出栈一个前括号 pop
                // 出栈 rest 的最后一个左括号出栈, 倒着遍历找最后一个左括号
                for i := len(rest) - 1; i >= 0; i-- {
                    if rest[i] == byte('(') {
                        rest = append(rest[:i], rest[i+1:]...)
                        break
                    }
                }
            } else { // 不匹配括号
                stack = append(stack, c) // 则为 左括号 ( 入栈
                rest = append(rest, c) // 维持前括号的顺序
            }
        }
        return rest
    }
    restToken := traverse(s)
    // 遍历找出能匹配剩余 * )( 的
    stack := []byte{}
    for _, c := range restToken {
        if c == '*' {
            // 看栈顶是不是 ( 不是的话入栈
            if len(stack) == 0 {// 啥都没有，匹配个毛，入栈
                stack = append(stack, c)
            } else {
                if stack[len(stack)- 1] == '(' {
                    stack = stack[:len(stack)-1] // ( 和 当前的 * 匹配 ，弹元素出栈 pop
                } else { // 栈顶就是 *, 无需匹配，把 * 入栈
                    stack = append(stack, c)
                }
            }
        } else if c == '(' { // 这个要等待 看后面有没有 * 所以要入栈
            stack = append(stack, c)
        } else { // c ==')'，这个要看栈顶有没有 * 不会有( 在等它，因为前面有一遍遍历已经把这个情况处理掉了
            if len(stack) == 0 {
                return false
            }
            stack = stack[:len(stack)-1] // 弹栈匹配 pop
        }
    }
    res := true
    for _, a := range stack {
        if a == '(' {
            res = false
            break
        }
    }
    return res
}

// 
func checkValidString1(s string) bool {
    min, max := 0, 0
    for i := range s {
        switch s[i] {
        case '(':
            min++
            max++
        case ')':
            min--
            max--
            if max < 0 { // 出现了未配对的 ) 直接 可以返回了
                return false
            }
        case '*':
            min--
            max++
        }
        if min < 0 {
            min = 0
        }
    }
    return min == 0
}

func main() {
    fmt.Println(checkValidString("()")) // true
    fmt.Println(checkValidString("(*)")) // true
    fmt.Println(checkValidString("(*))")) // true
    fmt.Println(checkValidString(")*(")) // false
    fmt.Println(checkValidString("(*))(")) // false

    fmt.Println(checkValidString1("()")) // true
    fmt.Println(checkValidString1("(*)")) // true
    fmt.Println(checkValidString1("(*))")) // true
    fmt.Println(checkValidString1(")*(")) // false
    fmt.Println(checkValidString1("(*))(")) // false
}
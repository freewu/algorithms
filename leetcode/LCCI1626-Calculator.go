package main

// 面试题 16.26. Calculator LCCI
// Given an arithmetic equation consisting of positive integers, +, -, * and / (no paren­theses), compute the result.

// The expression string contains only non-negative integers, +, -, *, / operators and empty spaces.
// The integer division should truncate toward zero.

// Example 1:
// Input: "3+2*2"
// Output: 7

// Example 2:
// Input: " 3/2 "
// Output: 1

// Example 3:
// Input: " 3+5 / 2 "
// Output: 5

// Note:
//     You may assume that the given expression is always valid.
//     Do not use the eval built-in library function.

import "fmt"

func calculate(s string) int {
    s += "+"
    res, stack, num, op := 0, []int{}, 0, byte('+')
    for i := range s {
        if s[i] >= '0' && s[i] <= '9' {
            num = num*10 + int(s[i] - '0')
        } else if s[i]== ' ' {
            continue
        } else {
            if op == '+' {
                stack = append(stack, num)
            } else if op == '-' {
                stack = append(stack, -num)
            } else if op == '*' {
                stack[len(stack)-1] = stack[len(stack)-1] * num
            } else if op == '/' {
                stack[len(stack)-1] = stack[len(stack)-1] / num
            }
            num, op = 0, s[i]
        }
    }
    for _, v := range stack {
        res += v
    }
    return res
}

func calculate1(s string) int {
    stack := []int{}
    res, num, op := 0, 0, '+'
    for i := 0; i < len(s);i++ {
        isNum:= ('0' <= s[i] && s[i] <= '9')
        if isNum {
            num = num * 10 + int(s[i] - '0')
        } 
        if (!isNum && s[i] != ' ') || i == len(s) - 1 { // 一定要空出来' '
            switch op {
            case '+':
                stack = append(stack, num)
            case '-':
                stack = append(stack, -num)
            case '*':
                stack[len(stack) - 1] *= num
            default:
                stack[len(stack) - 1] /= num
            }
            op, num = int32(s[i]), 0 // 操作完一定要把 num 重置为 0
        }
    }
    for _, v := range stack {
        res += v
    }
    return res
}

func main() {
    // Example 1:
    // Input: "3+2*2"
    // Output: 7
    fmt.Println(calculate("3+2*2")) // 7
    // Example 2:
    // Input: " 3/2 "
    // Output: 1
    fmt.Println(calculate("3/2")) // 1
    // Example 3:
    // Input: " 3+5 / 2 "
    // Output: 5
    fmt.Println(calculate(" 3+5 / 2 ")) // 5

    fmt.Println(calculate1("3+2*2")) // 7
    fmt.Println(calculate1("3/2")) // 1
    fmt.Println(calculate1(" 3+5 / 2 ")) // 5
}
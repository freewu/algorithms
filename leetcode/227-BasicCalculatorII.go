package main

// 227. Basic Calculator II
// Given a string s which represents an expression, evaluate this expression and return its value. 
// The integer division should truncate toward zero.
// You may assume that the given expression is always valid. All intermediate results will be in the range of [-2^31, 2^31 - 1].
// Note: You are not allowed to use any built-in function which evaluates strings as mathematical expressions, such as eval().

// Example 1:
// Input: s = "3+2*2"
// Output: 7

// Example 2:
// Input: s = " 3/2 "
// Output: 1

// Example 3:
// Input: s = " 3+5 / 2 "
// Output: 5

// Constraints:
//     1 <= s.length <= 3 * 10^5
//     s consists of integers and operators ('+', '-', '*', '/') separated by some number of spaces.
//     s represents a valid expression.
//     All the integers in the expression are non-negative integers in the range [0, 2^31 - 1].
//     The answer is guaranteed to fit in a 32-bit integer.

import "fmt"

// stack
type Stack []int  
  
func (stack *Stack) pop() int {  
    res := (*stack)[len(*stack)-1]  
    *stack = (*stack)[:len(*stack)-1]  
    return res  
}
func (stack *Stack) push(v int) {  
    *stack = append(*stack, v)  
}

func calculate(s string) int {
    res, stack, num, operator := 0, Stack{}, 0, byte('+')
    s += "+"  
    for i := range s {
        if s[i] == ' ' {
            continue
        }
        if s[i] >= '0' && s[i] <= '9' { // 处理数字
            num = (num * 10) + int(s[i]-'0')  
            continue  
        }
        switch operator {  
            case '+': stack.push(num)
            case '-': stack.push(-num)
            // 乘除法算完后入栈
            case '*': stack.push(stack.pop() * num) 
            case '/': stack.push(stack.pop() / num)  
        }
        operator = s[i]  
        num = 0  
    }
    for _, v := range stack {  
        res += v  
    }  
    return res  
}

func calculate1(s string) int {
    s += "+"
    res, stack, num, operator := 0, []int{}, 0, byte('+')
    for i := range s {
        if s[i] >= '0' && s[i] <= '9' {
            num = num*10 + int(s[i] - '0')
        } else if s[i]== ' ' {
            continue
        } else {
            if operator == '+' {
                stack = append(stack, num)
            } else if operator == '-' {
                stack = append(stack, -num)
            } else if operator == '*' {
                stack[len(stack)-1] = stack[len(stack)-1] * num
            } else if operator == '/' {
                stack[len(stack)-1] = stack[len(stack)-1] / num
            }
            num = 0
            operator = s[i]
        }
    }
    for _, v := range stack {
        res += v
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "3+2*2"
    // Output: 7
    fmt.Println(calculate("3+2*2")) // 7
    // Example 2:
    // Input: s = " 3/2 "
    // Output: 1
    fmt.Println(calculate(" 3/2 ")) // 1
    // Example 3:
    // Input: s = " 3+5 / 2 "
    // Output: 5
    fmt.Println(calculate(" 3+5 / 2 ")) // 5

    fmt.Println(calculate1("3+2*2")) // 7
    fmt.Println(calculate1(" 3/2 ")) // 1
    fmt.Println(calculate1(" 3+5 / 2 ")) // 5
}
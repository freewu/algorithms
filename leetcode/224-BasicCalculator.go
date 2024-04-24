package main

// 224. Basic Calculator
// Given a string s representing a valid expression, implement a basic calculator to evaluate it, and return the result of the evaluation.
// Note: You are not allowed to use any built-in function which evaluates strings as mathematical expressions, such as eval().

// Example 1:
// Input: s = "1 + 1"
// Output: 2

// Example 2:
// Input: s = " 2-1 + 2 "
// Output: 3

// Example 3:
// Input: s = "(1+(4+5+2)-3)+(6+8)"
// Output: 23

// Constraints:
//     1 <= s.length <= 3 * 10^5
//     s consists of digits, '+', '-', '(', ')', and ' '.
//     s represents a valid expression.
//     '+' is not used as a unary operation (i.e., "+1" and "+(2 + 3)" is invalid).
//     '-' could be used as a unary operation (i.e., "-1" and "-(2 + 3)" is valid).
//     There will be no two consecutive operators in the input.
//     Every number and running calculation will fit in a signed 32-bit integer.

import "fmt"

func calculate(s string) int {
    res, stack, sign, number := 0, []int{}, 1, 0
    for _, ch := range(s) {
        switch ch {
            case ' ': // 空格直接下个
                continue
            case '+':
                res += sign * number
                number = 0
                sign = 1
            case '-':
                res += sign * number
                number = 0
                sign = -1
            case '(':
                stack = append(stack, res, sign)
                res = 0
                sign = 1
            case ')':
                res += sign * number
                sign = stack[len(stack) - 1]
                res = stack[len(stack) - 2] + sign * res
                stack = stack[:len(stack) - 2]
                number = 0
            default:
                number = number * 10 + int(ch - '0')
        }
    }
    return res + sign * number
}

func main() {
    // Example 1:
    // Input: s = "1 + 1"
    // Output: 2
    fmt.Println(calculate("1 + 1")) // 2
    // Example 2:
    // Input: s = " 2-1 + 2 "
    // Output: 3
    fmt.Println(calculate(" 2-1 + 2")) // 3
    // Example 3:
    // Input: s = "(1+(4+5+2)-3)+(6+8)"
    // Output: 23
    fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)")) // 23
}
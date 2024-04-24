package main

// 150. Evaluate Reverse Polish Notation
// You are given an array of strings tokens that represents an arithmetic expression in a Reverse Polish Notation.
// Evaluate the expression. Return an integer that represents the value of the expression.
// Note that:
//     The valid operators are '+', '-', '*', and '/'.
//     Each operand may be an integer or another expression.
//     The division between two integers always truncates toward zero.
//     There will not be any division by zero.
//     The input represents a valid arithmetic expression in a reverse polish notation.
//     The answer and all the intermediate calculations can be represented in a 32-bit integer.
 
// Example 1:
// Input: tokens = ["2","1","+","3","*"]
// Output: 9
// Explanation: ((2 + 1) * 3) = 9

// Example 2:
// Input: tokens = ["4","13","5","/","+"]
// Output: 6
// Explanation: (4 + (13 / 5)) = 6

// Example 3:
// Input: tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
// Output: 22
// Explanation: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
// = ((10 * (6 / (12 * -11))) + 17) + 5
// = ((10 * (6 / -132)) + 17) + 5
// = ((10 * 0) + 17) + 5
// = (0 + 17) + 5
// = 17 + 5
// = 22

// Constraints:
//     1 <= tokens.length <= 10^4
//     tokens[i] is either an operator: "+", "-", "*", or "/", or an integer in the range [-200, 200].

import "fmt"
import "strconv"

func evalRPN(tokens []string) int {
    stack := make([]int, 0, len(tokens))
    for _, token := range tokens {
        v, err := strconv.Atoi(token)
        if err == nil { // 如果是数字 能被转换成功 压入栈中
            stack = append(stack, v)
        } else { // 不是数字说明是 符号
            // 取出最顶端的数字
            num1, num2 := stack[len(stack) - 2], stack[len(stack) - 1]
            stack = stack[:len(stack)-2]

            switch token { // 四则运算后,重新入栈
            case "+":
                stack = append(stack, num1 + num2)
            case "-":
                stack = append(stack, num1 - num2)
            case "*":
                stack = append(stack, num1 * num2)
            case "/":
                stack = append(stack, num1 / num2)
            }
        }
    }
    return stack[0]
}

// stack
func evalRPN1(tokens []string) int {
    stack := make([]int, 0)
    // 先判断符号 减少 strconv.Atoi 的调用次数
    for _, token := range tokens {
        if token == "-" {
            left := stack[len(stack) - 2]
            right := stack[len(stack) - 1]
            stack = stack[:len(stack) - 2]
            stack = append(stack, left - right)
        } else if token == "+" {
            left := stack[len(stack) - 2]
            right := stack[len(stack) - 1]
            stack = stack[:len(stack) - 2]
            stack = append(stack, left + right)
        } else if token == "/" {
            left := stack[len(stack) - 2]
            right := stack[len(stack) - 1]
            stack = stack[:len(stack) - 2]
            stack = append(stack, left / right)
        } else if token == "*" {
            left := stack[len(stack) - 2]
            right := stack[len(stack) - 1]
            stack = stack[:len(stack) - 2]
            stack = append(stack, left * right)
        } else {
            val, _ := strconv.Atoi(token)
            stack = append(stack, val)
        }
    }
    return stack[0]
}

func main() {
    // Example 1:
    // Input: tokens = ["2","1","+","3","*"]
    // Output: 9
    // Explanation: ((2 + 1) * 3) = 9
    fmt.Printf("evalRPN([]string{\"2\",\"1\",\"+\",\"3\",\"*\" }) = %v\n",evalRPN([]string{ "2","1","+","3","*" })) // 9  ((2 + 1) * 3)
    // Example 2:
    // Input: tokens = ["4","13","5","/","+"]
    // Output: 6
    // Explanation: (4 + (13 / 5)) = 6
    fmt.Printf("evalRPN([]string{\"4\",\"13\",\"5\",\"/\",\"+\"}) = %v\n",evalRPN([]string{ "4","13","5","/","+" })) // 6 (4 + (13 / 5))
    // Example 3:
    // Input: tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
    // Output: 22
    // Explanation: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
    // = ((10 * (6 / (12 * -11))) + 17) + 5
    // = ((10 * (6 / -132)) + 17) + 5
    // = ((10 * 0) + 17) + 5
    // = (0 + 17) + 5
    // = 17 + 5
    // = 22
    fmt.Printf("evalRPN([]string{\"10\",\"6\",\"9\",\"3\",\"+\",\"-11\",\"*\",\"/\",\"*\",\"17\",\"+\",\"5\",\"+\"}) = %v\n",evalRPN([]string{ "10","6","9","3","+","-11","*","/","*","17","+","5","+" })) // 22 ((10 * (6 / ((9 + 3) * -11))) + 17) + 5

    fmt.Printf("evalRPN1([]string{\"2\",\"1\",\"+\",\"3\",\"*\" }) = %v\n",evalRPN1([]string{ "2","1","+","3","*" })) // 9  ((2 + 1) * 3)
    fmt.Printf("evalRPN1([]string{\"4\",\"13\",\"5\",\"/\",\"+\"}) = %v\n",evalRPN1([]string{ "4","13","5","/","+" })) // 6 (4 + (13 / 5))
    fmt.Printf("evalRPN1([]string{\"10\",\"6\",\"9\",\"3\",\"+\",\"-11\",\"*\",\"/\",\"*\",\"17\",\"+\",\"5\",\"+\"}) = %v\n",evalRPN1([]string{ "10","6","9","3","+","-11","*","/","*","17","+","5","+" })) // 22 ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
}

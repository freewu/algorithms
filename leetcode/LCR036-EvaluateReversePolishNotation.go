package main

// LCR 036. 逆波兰表达式求值
// 根据 逆波兰表示法，求该后缀表达式的计算结果。
// 有效的算符包括 +、-、*、/ 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。

// 说明：
//     整数除法只保留整数部分。
//     给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。

// 示例 1：
// 输入：tokens = ["2","1","+","3","*"]
// 输出：9
// 解释：该算式转化为常见的中缀算术表达式为：((2 + 1) * 3) = 9

// 示例 2：
// 输入：tokens = ["4","13","5","/","+"]
// 输出：6
// 解释：该算式转化为常见的中缀算术表达式为：(4 + (13 / 5)) = 6

// 示例 3：
// 输入：tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
// 输出：22
// 解释：
// 该算式转化为常见的中缀算术表达式为：
//   ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
// = ((10 * (6 / (12 * -11))) + 17) + 5
// = ((10 * (6 / -132)) + 17) + 5
// = ((10 * 0) + 17) + 5
// = (0 + 17) + 5
// = 17 + 5
// = 22
 
// 提示：
//     1 <= tokens.length <= 10^4
//     tokens[i] 要么是一个算符（"+"、"-"、"*" 或 "/"），要么是一个在范围 [-200, 200] 内的整数

// 逆波兰表达式：
//     逆波兰表达式是一种后缀表达式，所谓后缀就是指算符写在后面。
//     平常使用的算式则是一种中缀表达式，如 ( 1 + 2 ) * ( 3 + 4 ) 。
//     该算式的逆波兰表达式写法为 ( ( 1 2 + ) ( 3 4 + ) * ) 。

// 逆波兰表达式主要有以下两个优点：
//     去掉括号后表达式无歧义，上式即便写成 1 2 + 3 4 + * 也可以依据次序计算出正确结果。
//     适合用栈操作运算：遇到数字则入栈；遇到算符则取出栈顶两个数字进行计算，并将结果压入栈中。

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

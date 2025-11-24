package main

// 3749. Evaluate Valid Expressions
// You are given a string expression that represents a nested mathematical expression in a simplified form.

// A valid expression is either an integer literal or follows the format op(a,b), where:
//     op is one of "add", "sub", "mul", or "div".
//     a and b are each valid expressions.

// The operations are defined as follows:
//     add(a,b) = a + b
//     sub(a,b) = a - b
//     mul(a,b) = a * b
//     div(a,b) = a / b
// Return an integer representing the result after fully evaluating the expression.

// Example 1:
// Input: expression = "add(2,3)"
// Output: 5
// Explanation:
// The operation add(2,3) means 2 + 3 = 5.

// Example 2:
// Input: expression = "-42"
// Output: -42
// Explanation:
// The expression is a single integer literal, so the result is -42.

// Example 3:
// Input: expression = "div(mul(4,sub(9,5)),add(1,1))"
// Output: 8
// Explanation:
// First, evaluate the inner expression: sub(9,5) = 9 - 5 = 4
// Next, multiply the results: mul(4,4) = 4 * 4 = 16
// Then, compute the addition on the right: add(1,1) = 1 + 1 = 2
// Finally, divide the two main results: div(16,2) = 16 / 2 = 8
// Therefore, the entire expression evaluates to 8.

// Constraints:
//     1 <= expression.length <= 10^5
//     expression is valid and consists of digits, commas, parentheses, the minus sign '-', and the lowercase strings "add", "sub", "mul", "div".
//     All intermediate results fit within the range of a long integer.
//     All divisions result in integer values.

import "fmt"
import "container/list"
import "strconv"
import "unicode"

func evaluateExpression(expression string) int64 {
    // 栈元素：操作符(string)和操作数列表([2]int64)
    stack := list.New()
    n, i := len(expression), 0
    for i < n {
        c := expression[i]
        if (c >= '0' && c <= '9') || c == '-' {
            // 解析整数（处理正负号）
            j := i
            if c == '-' {
                j++
            }
            // 收集所有连续数字字符
            for j < n && expression[j] >= '0' && expression[j] <= '9' {
                j++
            }
            // 解析数字
            num, _ := strconv.ParseInt(expression[i:j], 10, 64)
            i = j
            if stack.Len() == 0 {
                return num // 单独数字直接返回
            }
            // 填充栈顶操作的操作数
            top := stack.Back().Value.([]interface{})
            operands := top[1].([2]int64)
            if operands[0] == 0 && top[2].(bool) { // 标记第一个操作数是否未设置
                operands[0] = num
                top[1] = operands
                top[2] = false // 第一个操作数已设置
            } else {
                operands[1] = num
                top[1] = operands
            }
        } else if c == 'a' || c == 's' || c == 'm' || c == 'd' {
            // 解析操作符（add/sub/mul/div）
            var op string
            if i+3 <= n && expression[i:i+3] == "add" {
                op = "add"
                i += 3
            } else if i+3 <= n && expression[i:i+3] == "sub" {
                op = "sub"
                i += 3
            } else if i+3 <= n && expression[i:i+3] == "mul" {
                op = "mul"
                i += 3
            } else if i+3 <= n && expression[i:i+3] == "div" {
                op = "div"
                i += 3
            }
            i++ // 跳过 '('
            // 压入栈：操作符 + 操作数数组 + 第一个操作数是否未设置的标记
            stack.PushBack([]interface{}{op, [2]int64{0, 0}, true})
        } else if c == ',' {
            // 逗号分隔，准备解析第二个操作数
            i++
        } else if c == ')' {
            // 弹出栈顶操作计算结果
            topElem := stack.Remove(stack.Back())
            top := topElem.([]interface{})
            op := top[0].(string)
            operands := top[1].([2]int64)
            a, b := operands[0], operands[1]
            res := int64(0)
            // 执行对应运算
            switch op {
            case "add":
                res = a + b
            case "sub":
                res = a - b
            case "mul":
                res = a * b
            case "div":
                res = a / b // 题目保证除法结果为整数
            }
            if stack.Len() == 0 {
                return res // 栈空，当前结果为最终结果
            }
            // 将结果回写到上层操作
            upper := stack.Back().Value.([]interface{})
            upperOps := upper[1].([2]int64)
            if upper[2].(bool) { // 第一个操作数未设置
                upperOps[0] = res
                upper[1] = upperOps
                upper[2] = false
            } else {
                upperOps[1] = res
                upper[1] = upperOps
            }
            i++
        } else {
            // 跳过其他合法字符（如空格）
            i++
        }
    }
    return 0 // 输入合法，不会走到此分支
}

func evaluateExpression1(s string) int64 {
    stack, operations := []int{}, []byte{}
    n, sign, x := len(s), 1, 0
    for i := 0; i < n; i++ {
        c := s[i]
        if c == '-' {
            sign = -1
        } else if unicode.IsDigit(rune(c)) {
            x = x * 10 + int(c - '0')
            if i == n-1 || s[i+1] == ',' || s[i+1] == ')' {
                stack = append(stack, sign*x)
                sign, x = 1, 0
            }
        } else if c == ')' {
            v := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            op := operations[len(operations)-1]
            operations = operations[:len(operations)-1]
            switch op {
            case 'a': // add
                stack[len(stack)-1] += v
            case 's': // sub
                stack[len(stack)-1] -= v
            case 'm': // mul
                stack[len(stack)-1] *= v
            default: // div
                stack[len(stack)-1] /= v
            }
        } else if unicode.IsLower(rune(c)) {
            operations = append(operations, c)
            i += 3
        }
    }
    return int64(stack[len(stack) - 1])
}

func main() {
    // Example 1:
    // Input: expression = "add(2,3)"
    // Output: 5
    // Explanation:
    // The operation add(2,3) means 2 + 3 = 5.
    fmt.Println(evaluateExpression("add(2,3)")) // 5
    // Example 2:
    // Input: expression = "-42"
    // Output: -42
    // Explanation:
    // The expression is a single integer literal, so the result is -42.
    fmt.Println(evaluateExpression("-42")) // -42   
    // Example 3:
    // Input: expression = "div(mul(4,sub(9,5)),add(1,1))"
    // Output: 8
    // Explanation:
    // First, evaluate the inner expression: sub(9,5) = 9 - 5 = 4
    // Next, multiply the results: mul(4,4) = 4 * 4 = 16
    // Then, compute the addition on the right: add(1,1) = 1 + 1 = 2
    // Finally, divide the two main results: div(16,2) = 16 / 2 = 8
    // Therefore, the entire expression evaluates to 8.
    fmt.Println(evaluateExpression("div(mul(4,sub(9,5)),add(1,1))")) // 8 

    // 大数字测试（避免溢出）
    fmt.Println(evaluateExpression("mul(1603581729054586, 2)")) // 3207163458109172
    // 复杂嵌套测试
    fmt.Println(evaluateExpression("add(sub(1000000000000, 500000000000), div(800000000000, 2))")) // 900000000000

    fmt.Println(evaluateExpression1("add(2,3)")) // 5
    fmt.Println(evaluateExpression1("-42")) // -42   
    fmt.Println(evaluateExpression1("div(mul(4,sub(9,5)),add(1,1))")) // 8 
    fmt.Println(evaluateExpression1("mul(1603581729054586, 2)")) // 3207163458109172
    fmt.Println(evaluateExpression1("add(sub(1000000000000, 500000000000), div(800000000000, 2))")) // 900000000000
}
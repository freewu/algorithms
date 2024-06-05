package main

// 439. Ternary Expression Parser
// Given a string expression representing arbitrarily nested ternary expressions, 
// evaluate the expression, and return the result of it.

// You can always assume that the given expression is valid and only contains digits
// '?', ':', 'T', and 'F' where 'T' is true and 'F' is false. 
// All the numbers in the expression are one-digit numbers (i.e., in the range [0, 9]).

// The conditional expressions group right-to-left (as usual in most languages), 
// and the result of the expression will always evaluate to either a digit, 'T' or 'F'.

// Example 1:
// Input: expression = "T?2:3"
// Output: "2"
// Explanation: If true, then result is 2; otherwise result is 3.

// Example 2:
// Input: expression = "F?1:T?4:5"
// Output: "4"
// Explanation: The conditional expressions group right-to-left. Using parenthesis, it is read/evaluated as:
// "(F ? 1 : (T ? 4 : 5))" --> "(F ? 1 : 4)" --> "4"
// or "(F ? 1 : (T ? 4 : 5))" --> "(T ? 4 : 5)" --> "4"

// Example 3:
// Input: expression = "T?T?F:5:3"
// Output: "F"
// Explanation: The conditional expressions group right-to-left. Using parenthesis, it is read/evaluated as:
// "(T ? (T ? F : 5) : 3)" --> "(T ? F : 3)" --> "F"
// "(T ? (T ? F : 5) : 3)" --> "(T ? F : 5)" --> "F"
 
// Constraints:
//     5 <= expression.length <= 10^4
//     expression consists of digits, 'T', 'F', '?', and ':'.
//     It is guaranteed that expression is a valid ternary expression and that each number is a one-digit number.

import "fmt"

// 递归
func parseTernary(expression string) string {
    // 顺序解析，当前表达式必形为A？B：C三个部分或型如A的纯值，如果是前者顺序解析当前A，B，C并返回结果；
    // 怎么确定一个表达式部分（A，B，或者C）已经解析完成（即返回的是纯值）？通过它的右侧是否是“？”号来确定。
    n := len(expression)
    //i 为解析起始位，byte 为返回值，int 为该层表达式结尾下标
    var parse func(i int) (byte, int)
    parse = func(i int) (byte, int) {
        // 分别解析A，B，C
        A,iA := expression[i], i
        // 检查当前A表达式结果是否是更外一层表达式的判断条件，即 A 部分
        for iA+1<n && expression[iA+1] == '?' {
            B, iB := parse(iA+2)
            C, iC := parse(iB+2)
            if A == 'T' {
                A, iA = B, iC
            } else {
                A, iA = C, iC
            }
        }
        return A, iA
    }
    res, _ := parse(0)
    return string(res)
}

// stack
func parseTernary1(expression string) string {
    stack, n := []byte{}, len(expression)
    for i := n - 1; i >= 0; i-- {
        c := expression[i]
        if len(stack) > 0 && stack[len(stack)-1] == '?' {
            stack = stack[:len(stack)-1] // 移除 '?'
            trueExpr := stack[len(stack)-1] // 取出 true 处理部分
            stack = stack[:len(stack)-1] // 移除 trueExpr
            stack = stack[:len(stack)-1] // 移除 ':'
            falseExpr := stack[len(stack)-1] // 取出 false 处理部分
            stack = stack[:len(stack)-1] // 移除 falseExpr
            if c == 'T' {
                stack = append(stack, trueExpr)
            } else {
                stack = append(stack, falseExpr)
            }
        } else {
            stack = append(stack, c)
        }
    }
    return string(stack[0])
}

func main() {
    // Example 1:
    // Input: expression = "T?2:3"
    // Output: "2"
    // Explanation: If true, then result is 2; otherwise result is 3.
    fmt.Println(parseTernary("T?2:3")) // "2"
    // Example 2:
    // Input: expression = "F?1:T?4:5"
    // Output: "4"
    // Explanation: The conditional expressions group right-to-left. Using parenthesis, it is read/evaluated as:
    // "(F ? 1 : (T ? 4 : 5))" --> "(F ? 1 : 4)" --> "4"
    // or "(F ? 1 : (T ? 4 : 5))" --> "(T ? 4 : 5)" --> "4"
    fmt.Println(parseTernary("F?1:T?4:5")) // "4"
    // Example 3:
    // Input: expression = "T?T?F:5:3"
    // Output: "F"
    // Explanation: The conditional expressions group right-to-left. Using parenthesis, it is read/evaluated as:
    // "(T ? (T ? F : 5) : 3)" --> "(T ? F : 3)" --> "F"
    // "(T ? (T ? F : 5) : 3)" --> "(T ? F : 5)" --> "F"
    fmt.Println(parseTernary("T?T?F:5:3")) // "F"

    fmt.Println(parseTernary1("T?2:3")) // "2"
    fmt.Println(parseTernary1("F?1:T?4:5")) // "4"
    fmt.Println(parseTernary1("T?T?F:5:3")) // "F"
}
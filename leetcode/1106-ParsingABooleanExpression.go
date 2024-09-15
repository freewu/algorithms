package main

// 1106. Parsing A Boolean Expression
// A boolean expression is an expression that evaluates to either true or false. 
// It can be in one of the following shapes:
//     't' that evaluates to true.
//     'f' that evaluates to false.
//     '!(subExpr)' that evaluates to the logical NOT of the inner expression subExpr.
//     '&(subExpr1, subExpr2, ..., subExprn)' that evaluates to the logical AND of the inner expressions subExpr1, subExpr2, ..., subExprn where n >= 1.
//     '|(subExpr1, subExpr2, ..., subExprn)' that evaluates to the logical OR of the inner expressions subExpr1, subExpr2, ..., subExprn where n >= 1.

// Given a string expression that represents a boolean expression, return the evaluation of that expression.

// It is guaranteed that the given expression is valid and follows the given rules.

// Example 1:
// Input: expression = "&(|(f))"
// Output: false
// Explanation: 
// First, evaluate |(f) --> f. The expression is now "&(f)".
// Then, evaluate &(f) --> f. The expression is now "f".
// Finally, return false.

// Example 2:
// Input: expression = "|(f,f,f,t)"
// Output: true
// Explanation: The evaluation of (false OR false OR false OR true) is true.

// Example 3:
// Input: expression = "!(&(f,t))"
// Output: true
// Explanation: 
// First, evaluate &(f,t) --> (false AND true) --> false --> f. The expression is now "!(f)".
// Then, evaluate !(f) --> NOT false --> true. We return true.

// Constraints:
//     1 <= expression.length <= 2 * 10^4
//     expression[i] is one following characters: '(', ')', '&', '|', '!', 't', 'f', and ','.

import "fmt"
import "strings"

// S-> !(S) | |(L) | &(L) | E
// L-> S | S,L
// E-> t | f
func parseBoolExpr(expression string) bool {
    reader := strings.NewReader(expression)
    var parseExpr func(reader *strings.Reader) bool
    var parseOrExpList func(reader *strings.Reader) bool
    var parseAndExpList func(reader *strings.Reader) bool
    parseOrExpList = func(reader *strings.Reader) bool {
        res := parseExpr(reader)
        if reader.Len() > 0 {
            lookahead, _ := reader.ReadByte()
            if lookahead == ',' {
                if parseOrExpList(reader) {
                    res = true
                }
            }
        }
        return res
    }
    parseAndExpList = func(reader *strings.Reader) bool {
        res := parseExpr(reader)
        if reader.Len() > 0 {
            lookahead, _ := reader.ReadByte()
            if lookahead == ',' {
                if !parseAndExpList(reader) {
                    return false
                }
            }
            // or otherwise, lookahead is ')' and it was just consumed
        }
        return res
    }
    parseExpr = func(reader *strings.Reader) bool {
        lookahead, _ := reader.ReadByte()
        if lookahead == '!' {
            reader.ReadByte()
            flag := !parseExpr(reader)
            reader.ReadByte() // eliminate trailing ')'
            return flag
        } else if lookahead == '|' {
            reader.ReadByte()
            return parseOrExpList(reader)
        } else if lookahead == '&' {
            reader.ReadByte()
            return parseAndExpList(reader)
        } else if lookahead == 't' {
            return true
        }
        return false
    }
    return parseExpr(reader)
}

func parseBoolExpr1(expression string) bool {
    stack := []byte{}
    for i := range expression {
        c := expression[i]
        if c == ',' {
            continue
        } else if c != ')' {
            stack = append(stack, c) // push
            continue
        }
        t, f := 0, 0
        for stack[len(stack) - 1] != '(' {
            v := stack[len(stack) - 1] // pop
            stack = stack[:len(stack) - 1]
            if v == 't' {
                t++
            } else {
                f++
            }
        }
        stack = stack[:len(stack) - 1] // op + val
        op := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        b := byte('t')
        switch op {
            case '!':
                if t > 0 {
                    b = 'f'
                }
                stack = append(stack, b)
            case '&':
                if f > 0 {
                    b = 'f'
                }
                stack = append(stack, b)
            case '|':
                if t == 0 {
                    b = 'f'
                }
                stack = append(stack, b)
        }
    }
    return stack[len(stack) - 1] == 't'
}

func main() {
    // Example 1:
    // Input: expression = "&(|(f))"
    // Output: false
    // Explanation: 
    // First, evaluate |(f) --> f. The expression is now "&(f)".
    // Then, evaluate &(f) --> f. The expression is now "f".
    // Finally, return false.
    fmt.Println(parseBoolExpr("&(|(f))")) // false
    // Example 2:
    // Input: expression = "|(f,f,f,t)"
    // Output: true
    // Explanation: The evaluation of (false OR false OR false OR true) is true.
    fmt.Println(parseBoolExpr("|(f,f,f,t)")) // true
    // Example 3:
    // Input: expression = "!(&(f,t))"
    // Output: true
    // Explanation: 
    // First, evaluate &(f,t) --> (false AND true) --> false --> f. The expression is now "!(f)".
    // Then, evaluate !(f) --> NOT false --> true. We return true.
    fmt.Println(parseBoolExpr("!(&(f,t))")) // true

    fmt.Println(parseBoolExpr1("&(|(f))")) // false
    fmt.Println(parseBoolExpr1("|(f,f,f,t)")) // true
    fmt.Println(parseBoolExpr1("!(&(f,t))")) // true
}
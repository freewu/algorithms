package main

// 736. Parse Lisp Expression
// You are given a string expression representing a Lisp-like expression to return the integer value of.
// The syntax for these expressions is given as follows.
//     1. An expression is either an integer, let expression, add expression, mult expression, or an assigned variable. Expressions always evaluate to a single integer.
//     2. (An integer could be positive or negative.)
//     3. A let expression takes the form "(let v1 e1 v2 e2 ... vn en expr)", where let is always the string "let", then there are one or more pairs of alternating variables and expressions, meaning that the first variable v1 is assigned the value of the expression e1, the second variable v2 is assigned the value of the expression e2, and so on sequentially; and then the value of this let expression is the value of the expression expr.
//     4. An add expression takes the form "(add e1 e2)" where add is always the string "add", there are always two expressions e1, e2 and the result is the addition of the evaluation of e1 and the evaluation of e2.
//     5. A mult expression takes the form "(mult e1 e2)" where mult is always the string "mult", there are always two expressions e1, e2 and the result is the multiplication of the evaluation of e1 and the evaluation of e2.
//     6. For this question, we will use a smaller subset of variable names. A variable starts with a lowercase letter, then zero or more lowercase letters or digits. Additionally, for your convenience, the names "add", "let", and "mult" are protected and will never be used as variable names.
//     7. Finally, there is the concept of scope. When an expression of a variable name is evaluated, within the context of that evaluation, the innermost scope (in terms of parentheses) is checked first for the value of that variable, and then outer scopes are checked sequentially. It is guaranteed that every expression is legal. Please see the examples for more details on the scope.
 
// Example 1:
// Input: expression = "(let x 2 (mult x (let x 3 y 4 (add x y))))"
// Output: 14
// Explanation: In the expression (add x y), when checking for the value of the variable x,
// we check from the innermost scope to the outermost in the context of the variable we are trying to evaluate.
// Since x = 3 is found first, the value of x is 3.

// Example 2:
// Input: expression = "(let x 3 x 2 x)"
// Output: 2
// Explanation: Assignment in let statements is processed sequentially.

// Example 3:
// Input: expression = "(let x 1 y 2 x (add x y) (add x y))"
// Output: 5
// Explanation: The first (add x y) evaluates as 3, and is assigned to x.
// The second (add x y) evaluates as 3+2 = 5.

// Constraints:
//     1 <= expression.length <= 2000
//     There are no leading or trailing spaces in expression.
//     All tokens are separated by a single space in expression.
//     The answer and all intermediate calculations of that answer are guaranteed to fit in a 32-bit integer.
//     The expression is guaranteed to be legal and evaluate to an integer.

import "fmt"
import "strings"
import "strconv"

func evaluate(expression string) int {
    parse := func(str string) []string {
        parent, res, sb := 0, make([]string, 0), make([]byte, 0)
        for i := range str {
            if str[i] == '(' { parent++ }
            if str[i] == ')' { parent-- }
            if parent == 0 && str[i] == ' ' {
                res = append(res, string(sb))
                sb = make([]byte, 0)
            } else {
                sb = append(sb, str[i])
            }
        }
        if len(sb) > 0 {
            res = append(res, string(sb))
        }
        return res
    }
    var eval func(expression string, parent map[string]int) int
    eval = func(expression string, parent map[string]int) int {
        if expression[0] != '(' {
            if expression[0] == '-' || (expression[0] >= '0' && expression[0] <= '9') {
                res, _ := strconv.Atoi(expression)
                return res
            }
            return parent[expression]
        }
        cache, tokens := make(map[string]int), []string{}
        for k, v := range parent {
            cache[k] = v
        }
        if expression[1] == 'm' {
            tokens = parse(expression[6 : len(expression)-1])
        } else {
            tokens = parse(expression[5 : len(expression)-1])
        }
        if strings.HasPrefix(expression, "(a") { // + 
            return eval(tokens[0], cache) + eval(tokens[1], cache)
        } else if strings.HasPrefix(expression, "(m") { // *
            return eval(tokens[0], cache) * eval(tokens[1], cache)
        } else { // let
            for i := 0; i < len(tokens)-2; i += 2 {
                cache[tokens[i]] = eval(tokens[i+1], cache)
            }
            return eval(tokens[len(tokens)-1], cache)
        }
    }
    return eval(expression, make(map[string]int))
}

func main() {
    // Example 1:
    // Input: expression = "(let x 2 (mult x (let x 3 y 4 (add x y))))"
    // Output: 14
    // Explanation: In the expression (add x y), when checking for the value of the variable x,
    // we check from the innermost scope to the outermost in the context of the variable we are trying to evaluate.
    // Since x = 3 is found first, the value of x is 3.
    fmt.Println(evaluate("(let x 2 (mult x (let x 3 y 4 (add x y))))")) // 14
    // Example 2:
    // Input: expression = "(let x 3 x 2 x)"
    // Output: 2
    // Explanation: Assignment in let statements is processed sequentially.
    fmt.Println(evaluate("(let x 3 x 2 x)")) // 2
    // Example 3:
    // Input: expression = "(let x 1 y 2 x (add x y) (add x y))"
    // Output: 5
    // Explanation: The first (add x y) evaluates as 3, and is assigned to x.
    // The second (add x y) evaluates as 3+2 = 5.
    fmt.Println(evaluate("(let x 1 y 2 x (add x y) (add x y))")) // 5
}
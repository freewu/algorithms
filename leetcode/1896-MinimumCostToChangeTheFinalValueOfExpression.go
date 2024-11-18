package main

// 1896. Minimum Cost to Change the Final Value of Expression
// You are given a valid boolean expression as a string expression consisting of the characters '1','0','&' (bitwise AND operator),'|' (bitwise OR operator),'(', and ')'.
//     For example, "()1|1" and "(1)&()" are not valid while "1", "(((1))|(0))", and "1|(0&(1))" are valid expressions.

// Return the minimum cost to change the final value of the expression.
//     For example, if expression = "1|1|(0&0)&1", its value is 1|1|(0&0)&1 = 1|1|0&1 = 1|0&1 = 1&1 = 1. 
//     We want to apply operations so that the new expression evaluates to 0.

// The cost of changing the final value of an expression is the number of operations performed on the expression. 
// The types of operations are described as follows:
//     Turn a '1' into a '0'.
//     Turn a '0' into a '1'.
//     Turn a '&' into a '|'.
//     Turn a '|' into a '&'.

// Note: '&' does not take precedence over '|' in the order of calculation. 
// Evaluate parentheses first, then in left-to-right order.

// Example 1:
// Input: expression = "1&(0|1)"
// Output: 1
// Explanation: We can turn "1&(0|1)" into "1&(0&1)" by changing the '|' to a '&' using 1 operation.
// The new expression evaluates to 0. 

// Example 2:
// Input: expression = "(0&0)&(0&0&0)"
// Output: 3
// Explanation: We can turn "(0&0)&(0&0&0)" into "(0|1)|(0&0&0)" using 3 operations.
// The new expression evaluates to 1.

// Example 3:
// Input: expression = "(0|(1|0&1))"
// Output: 1
// Explanation: We can turn "(0|(1|0&1))" into "(0|(0|0&1))" using 1 operation.
// The new expression evaluates to 0.

// Constraints:
//     1 <= expression.length <= 10^5
//     expression only contains '1','0','&','|','(', and ')'
//     All parentheses are properly matched.
//     There will be no empty parentheses (i.e: "()" is not a substring of expression).

import "fmt"

func minOperationsToFlip(expression string) int {
    n := len(expression)
    pre, stack := make([]int, n), make([]int, n)
    for i := n - 1; i >= 0; i-- {
        if expression[i] == ')' {
            stack = append(stack, i)
        } else if expression[i] == '(' {
            pre[i] = stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var parse func(s []byte, pre []int, l, r int) [2]int
    parse = func(s []byte, pre []int, l, r int) [2]int {
        stack, op := [][2]int{}, []byte{}
        calc := func() {
            if len(op) == 0 { return }
            ch := op[0]
            op = op[:len(op) - 1]  // pop
            p1, p2 := stack[len(stack) - 2], stack[len(stack) - 1]
            p := [2]int{2e9, 2e9}
            if ch == '&' {
                p[0] = min(p1[0] + p2[0], p1[0] + p2[1])
                p[0] = min(p[0], p1[1] + p2[0])
                p[1] = min(p1[0] + p2[1] + 1, p1[1] + p2[0] + 1)
                p[1] = min(p[1], p1[1] + p2[1])
            } else {
                p[0] = min(p1[0] + p2[1] + 1, p1[1] + p2[0] + 1)
                p[0] = min(p[0], p1[0] + p2[0])
                p[1] = min(p1[0] + p2[1], p1[1] + p2[0])
                p[1] = min(p[1], p1[1] + p2[1])
            }
            stack[len(stack) - 2] = p
            stack = stack[:len(stack) - 1]
        }
        for i := l; i <= r; i++ {
            if s[i] == '(' {
                stack = append(stack, parse(s, pre, i + 1, pre[i] - 1))
                i = pre[i]
                calc()
            } else if s[i] == '&' || s[i] == '|' {
                op = append(op, s[i])
            } else if s[i] == '0' {
                stack = append(stack, [2]int{0, 1})
                calc()
            } else if s[i] == '1' {
                stack = append(stack, [2]int{1, 0})
                calc()
            }
        }
        return stack[0]
    }
    v := parse([]byte(expression), pre, 0, n - 1)
    return max(v[0], v[1])
}

func minOperationsToFlip1(expression string) int {
    exp := []byte(expression)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var parse func() (int, int)
    parse = func() (int, int) {
        c := exp[len(exp)-1]
        exp = exp[:len(exp)-1] // pop
        l0, l1 := 0, 0
        if c == ')' {
            l0, l1 = parse()
        } else {
            l0, l1 = int(c - '0'), 1 - int(c - '0')
        }
        if len(exp) == 0 || exp[len(exp)-1] == '(' {
            if len(exp) > 0 {
                exp = exp[:len(exp)-1]
            }
            return l0, l1
        }
        c = exp[len(exp)-1]
        exp = exp[:len(exp)-1]
        r0, r1 := parse()
        toZero, toOne := 0, 0 
        if c == '&' {
            toZero = min(l0, r0)
            toOne = min(l1+r1, min(l1, r1)+1)
        } else if c == '|' {
            toZero = min(l0+r0, min(l0, r0)+1)
            toOne = min(l1, r1)
        }
        return toZero, toOne
    }
    toZero, toOne := parse()
    return max(toZero, toOne)
}

func main() {
    // Example 1:
    // Input: expression = "1&(0|1)"
    // Output: 1
    // Explanation: We can turn "1&(0|1)" into "1&(0&1)" by changing the '|' to a '&' using 1 operation.
    // The new expression evaluates to 0. 
    fmt.Println(minOperationsToFlip("1&(0|1)")) // 1
    // Example 2:
    // Input: expression = "(0&0)&(0&0&0)"
    // Output: 3
    // Explanation: We can turn "(0&0)&(0&0&0)" into "(0|1)|(0&0&0)" using 3 operations.
    // The new expression evaluates to 1.
    fmt.Println(minOperationsToFlip("(0&0)&(0&0&0)")) // 3
    // Example 3:
    // Input: expression = "(0|(1|0&1))"
    // Output: 1
    // Explanation: We can turn "(0|(1|0&1))" into "(0|(0|0&1))" using 1 operation.
    // The new expression evaluates to 0.
    fmt.Println(minOperationsToFlip("(0|(1|0&1))")) // 1

    fmt.Println(minOperationsToFlip1("1&(0|1)")) // 1
    fmt.Println(minOperationsToFlip1("(0&0)&(0&0&0)")) // 3
    fmt.Println(minOperationsToFlip1("(0|(1|0&1))")) // 1
}
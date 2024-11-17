package main

// 2232. Minimize Result by Adding Parentheses to Expression
// You are given a 0-indexed string expression of the form "<num1>+<num2>" where <num1> and <num2> represent positive integers.

// Add a pair of parentheses to expression such that after the addition of parentheses, 
// expression is a valid mathematical expression and evaluates to the smallest possible value. 
// The left parenthesis must be added to the left of '+' and the right parenthesis must be added to the right of '+'.

// Return expression after adding a pair of parentheses such that expression evaluates to the smallest possible value. 
// If there are multiple answers that yield the same result, return any of them.

// The input has been generated such that the original value of expression, 
// and the value of expression after adding any pair of parentheses that meets the requirements fits within a signed 32-bit integer.

// Example 1:
// Input: expression = "247+38"
// Output: "2(47+38)"
// Explanation: The expression evaluates to 2 * (47 + 38) = 2 * 85 = 170.
// Note that "2(4)7+38" is invalid because the right parenthesis must be to the right of the '+'.
// It can be shown that 170 is the smallest possible value.

// Example 2:
// Input: expression = "12+34"
// Output: "1(2+3)4"
// Explanation: The expression evaluates to 1 * (2 + 3) * 4 = 1 * 5 * 4 = 20.

// Example 3:
// Input: expression = "999+999"
// Output: "(999+999)"
// Explanation: The expression evaluates to 999 + 999 = 1998.

// Constraints:
//     3 <= expression.length <= 10
//     expression consists of digits from '1' to '9' and '+'.
//     expression starts and ends with digits.
//     expression contains exactly one '+'.
//     The original value of expression, and the value of expression after adding any pair of parentheses that meets the requirements fits within a signed 32-bit integer.

import "fmt"
import "strings"
import "strconv"

func minimizeResult(expression string) string {
    plus := strings.Index(expression, "+") //  + 位置
    num1, num2, p1, p2 := expression[:plus], expression[plus + 1:], 0, 0
    lowest, res := 1 << 31, ""
    
    for i := 0; i < len(num1); i++ {
        for j := 0; j < len(num2); j++ {
            s1, s2 := num1[:i], num2[j + 1:]
            if s1 == "" {
                p1 = 1
            } else {
                p1, _ = strconv.Atoi(s1)
            }
            if s2 == "" {
                p2 = 1
            } else {
                p2, _ = strconv.Atoi(s2)
            }
            sum1, _ := strconv.Atoi(num1[i:])
            sum2, _ := strconv.Atoi(num2[:j + 1])
            sum := sum1 + sum2
            eval := p1 * sum * p2
            if lowest > eval {
                lowest = eval
                res = s1 + "(" + num1[i:] + "+" + num2[:j + 1] + ")" + s2
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: expression = "247+38"
    // Output: "2(47+38)"
    // Explanation: The expression evaluates to 2 * (47 + 38) = 2 * 85 = 170.
    // Note that "2(4)7+38" is invalid because the right parenthesis must be to the right of the '+'.
    // It can be shown that 170 is the smallest possible value.
    fmt.Println(minimizeResult("247+38")) // "2(47+38)"
    // Example 2:
    // Input: expression = "12+34"
    // Output: "1(2+3)4"
    // Explanation: The expression evaluates to 1 * (2 + 3) * 4 = 1 * 5 * 4 = 20.
    fmt.Println(minimizeResult("12+34")) // "1(2+3)4"
    // Example 3:
    // Input: expression = "999+999"
    // Output: "(999+999)"
    // Explanation: The expression evaluates to 999 + 999 = 1998.
    fmt.Println(minimizeResult("999+999")) // "(999+999)"
}
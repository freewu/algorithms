package main

// 964. Least Operators to Express Number
// Given a single positive integer x, we will write an expression of the form x (op1) x (op2) x (op3) x ... where each operator op1, op2, etc. 
// is either addition, subtraction, multiplication, or division (+, -, *, or /). 
// For example, with x = 3, we might write 3 * 3 / 3 + 3 - 3 which is a value of 3.

// When writing such an expression, we adhere to the following conventions:
//     The division operator (/) returns rational numbers.
//     There are no parentheses placed anywhere.
//     We use the usual order of operations: multiplication and division happen before addition and subtraction.
//     It is not allowed to use the unary negation operator (-). For example, "x - x" is a valid expression as it only uses subtraction, but "-x + x" is not because it uses negation.

// We would like to write an expression with the least number of operators such that the expression equals the given target. 
// Return the least number of operators used.

// Example 1:
// Input: x = 3, target = 19
// Output: 5
// Explanation: 3 * 3 + 3 * 3 + 3 / 3.
// The expression contains 5 operations.

// Example 2:
// Input: x = 5, target = 501
// Output: 8
// Explanation: 5 * 5 * 5 * 5 - 5 * 5 * 5 + 5 / 5.
// The expression contains 8 operations.

// Example 3
// Input: x = 100, target = 100000000
// Output: 3
// Explanation: 100 * 100 * 100 * 100.
// The expression contains 3 operations.

// Constraints:
//     2 <= x <= 100
//     1 <= target <= 2 * 10^8

import "fmt"

func leastOpsExpressTarget(x int, target int) int {
    memo := map[string]int{}
    min := func (x, y int) int { if x < y { return x; }; return y; }
    cost := func (x int) int{ if x > 0 { return x; };  return 2; }
    var dfs func(i, target int) int
    dfs = func(i, target int) int {
        code := fmt.Sprintf("%d#%d" , i, target)
        if res, ok := memo[code]; ok {
            return res
        }
        res := 0
        if target == 0 {
            res = 0
        } else if target == 1 {
            res = cost(i)
        } else if i >= 39 {
            res = target + 1
        } else {
            t, r := target / x, target % x
            res = min(r * cost(i) + dfs(i+1, t), (x-r) * cost(i) + dfs(i+1, t+1))
        }
        memo[code] = res
        return res
    }
    return dfs(0, target) - 1
}

func main() {
    // Example 1:
    // Input: x = 3, target = 19
    // Output: 5
    // Explanation: 3 * 3 + 3 * 3 + 3 / 3.
    // The expression contains 5 operations.
    fmt.Println(leastOpsExpressTarget(3, 19)) // 5  (3 * 3 + 3 * 3 + 3 / 3)
    // Example 2:
    // Input: x = 5, target = 501
    // Output: 8
    // Explanation: 5 * 5 * 5 * 5 - 5 * 5 * 5 + 5 / 5.
    // The expression contains 8 operations.
    fmt.Println(leastOpsExpressTarget(5, 501)) // 8 (5 * 5 * 5 * 5 - 5 * 5 * 5 + 5 / 5)
    // Example 3
    // Input: x = 100, target = 100000000
    // Output: 3
    // Explanation: 100 * 100 * 100 * 100.
    // The expression contains 3 operations.
    fmt.Println(leastOpsExpressTarget(100, 100000000)) // 3 (100 * 100 * 100 * 100)
}
package main

// 679. 24 Game
// You are given an integer array cards of length 4. 
// You have four cards, each containing a number in the range [1, 9]. 
// You should arrange the numbers on these cards in a mathematical expression using the operators ['+', '-', '*', '/'] and the parentheses '(' and ')' to get the value 24.

// You are restricted with the following rules:
//     1. The division operator '/' represents real division, not integer division.
//         For example, 4 / (1 - 2 / 3) = 4 / (1 / 3) = 12.
//     2. Every operation done is between two numbers. In particular, we cannot use '-' as a unary operator.
//         For example, if cards = [1, 1, 1, 1], the expression "-1 - 1 - 1 - 1" is not allowed.
//     3. You cannot concatenate numbers together
//         For example, if cards = [1, 2, 1, 2], the expression "12 + 12" is not valid.

// Return true if you can get such expression that evaluates to 24, and false otherwise.

// Example 1:
// Input: cards = [4,1,8,7]
// Output: true
// Explanation: (8-4) * (7-1) = 24

// Example 2:
// Input: cards = [1,2,1,2]
// Output: false

// Constraints:
//     cards.length == 4
//     1 <= cards[i] <= 9

import "fmt"

func judgePoint24(cards []int) bool {
    // 转换成 float64
    arr := []float64{}
    for _, c := range cards {
        arr = append(arr, float64(c))
    }
    abs := func (x float64) float64 { if x < 0 { return -x }; return x }
    operationsResult := func(a, b float64) []float64 { // 计算两个数的各种 + - * /  4则运算的各种可能
        res := []float64{ a - b, a + b, b - a, a * b}
        if a > 0 { res = append(res, b / a) }
        if b > 0 { res = append(res, a / b) }
        return res
    }
    var dfs func(cards []float64) bool 
    dfs = func (cards []float64) bool {
        n := len(cards)
        if n == 1 {
            return abs(cards[0] - 24) <= 0.1
        }
        for i := 0; i < n; i++ {
            for j := i + 1; j < n; j++ {
                var list []float64
                for k := 0; k < n; k++ {
                    if k != i && k != j {
                        list = append(list, cards[k])
                    }
                }
                ops := operationsResult(cards[i], cards[j])
                for _, op := range ops {
                    list = append(list, op)
                    if dfs(list) { // Recur
                        return true
                    }
                    list = list[:len(list)-1]
                }
            }
        }
        return false
    }
    return dfs(arr)
}

func main() {
    // Example 1:
    // Input: cards = [4,1,8,7]
    // Output: true
    // Explanation: (8-4) * (7-1) = 24
    fmt.Println(judgePoint24([]int{4,1,8,7})) // true
    // Example 2:
    // Input: cards = [1,2,1,2]
    // Output: false
    fmt.Println(judgePoint24([]int{1,2,1,2})) // false

    fmt.Println(judgePoint24([]int{1,2,3,4})) // true
    fmt.Println(judgePoint24([]int{4,3,2,1})) // false
}
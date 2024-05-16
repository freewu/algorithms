package main

// 241. Different Ways to Add Parentheses
// Given a string expression of numbers and operators, 
// return all possible results from computing all the different possible ways to group numbers and operators. 
// You may return the answer in any order.

// The test cases are generated such that the output values fit in a 32-bit integer and the number of different results does not exceed 10^4.

// Example 1:
// Input: expression = "2-1-1"
// Output: [0,2]
// Explanation:
// ((2-1)-1) = 0 
// (2-(1-1)) = 2

// Example 2:
// Input: expression = "2*3-4*5"
// Output: [-34,-14,-10,-10,10]
// Explanation:
// (2*(3-(4*5))) = -34 
// ((2*3)-(4*5)) = -14 
// ((2*(3-4))*5) = -10 
// (2*((3-4)*5)) = -10 
// (((2*3)-4)*5) = 10

// Constraints:
//     1 <= expression.length <= 20
//     expression consists of digits and the operator '+', '-', and '*'.
//     All the integer values in the input expression are in the range [0, 99].

import "fmt"
import "strconv"

func diffWaysToCompute(expression string) []int {
    res := make([]int, 0)
    for i := 0; i < len(expression); i++ {
        if expression[i] == '+' || expression[i] == '-' || expression[i] == '*' {
            leftExpAnswers := diffWaysToCompute(expression[:i])
            rightExpAnswers := diffWaysToCompute(expression[i+1:])
            for _, leftAnswer := range leftExpAnswers {
                for _, rightAnswer := range rightExpAnswers {
                    if expression[i] == '+' {
                        res = append(res, leftAnswer + rightAnswer)
                    } else if expression[i] == '-' {
                        res = append(res, leftAnswer - rightAnswer)
                    } else if expression[i] == '*' {
                        res = append(res, leftAnswer * rightAnswer)
                    }
                }
            }
        }
    }
    if len(res) == 0 {
        expressionInt, _ := strconv.Atoi(expression)
        res = append(res, expressionInt)
    }
    return res
}

func main() {
    // Example 1:
    // Input: expression = "2-1-1"
    // Output: [0,2]
    // Explanation:
    // ((2-1)-1) = 0 
    // (2-(1-1)) = 2
    fmt.Println(diffWaysToCompute("2-1-1")) // [0,2]
    // Example 2:
    // Input: expression = "2*3-4*5"
    // Output: [-34,-14,-10,-10,10]
    // Explanation:
    // (2*(3-(4*5))) = -34 
    // ((2*3)-(4*5)) = -14 
    // ((2*(3-4))*5) = -10 
    // (2*((3-4)*5)) = -10 
    // (((2*3)-4)*5) = 10
    fmt.Println(diffWaysToCompute("2*3-4*5")) // [-34,-14,-10,-10,10]
}
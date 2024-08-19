package main

// 991. Broken Calculator
// There is a broken calculator that has the integer startValue on its display initially. 
// In one operation, you can:
//     multiply the number on display by 2, or
//     subtract 1 from the number on display.

// Given two integers startValue and target, 
// return the minimum number of operations needed to display target on the calculator.

// Example 1:
// Input: startValue = 2, target = 3
// Output: 2
// Explanation: Use double operation and then decrement operation {2 -> 4 -> 3}.

// Example 2:
// Input: startValue = 5, target = 8
// Output: 2
// Explanation: Use decrement and then double {5 -> 4 -> 8}.

// Example 3:
// Input: startValue = 3, target = 10
// Output: 3
// Explanation: Use double, decrement and double {3 -> 6 -> 5 -> 10}.

// Constraints:
//     1 <= startValue, target <= 10^9

import "fmt"

func brokenCalc(startValue int, target int) int {
    res := 0 
    for target > startValue  {
        res++
        if target % 2 == 1 { // 如果为奇数 需要 + 1
            target++ 
            continue
        }
        target = target / 2 // multiply
    } 
    res += startValue - target // subtract
    return res
}

func main() {
    // Example 1:
    // Input: startValue = 2, target = 3
    // Output: 2
    // Explanation: Use double operation and then decrement operation {2 -> 4 -> 3}.
    fmt.Println(brokenCalc(2, 3)) // 2    2 -[双倍（Double）]-> 4 -[递减（Decrement）]-> 3
    // Example 2:
    // Input: startValue = 5, target = 8
    // Output: 2
    // Explanation: Use decrement and then double {5 -> 4 -> 8}.
    fmt.Println(brokenCalc(5, 8)) // 2   {5 -[递减（Decrement）]-> 4 -[双倍（Double）]-> 8}.
    // Example 3:
    // Input: startValue = 3, target = 10
    // Output: 3
    // Explanation: Use double, decrement and double {3 -> 6 -> 5 -> 10}.
    fmt.Println(brokenCalc(3, 10)) // 3   3 -[双倍（Double）]-> 6 -[递减（Decrement）]-> 5 -[双倍（Double）]-> 10
}
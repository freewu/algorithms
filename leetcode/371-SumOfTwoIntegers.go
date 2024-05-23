package main

// 371. Sum of Two Integers
// Given two integers a and b, return the sum of the two integers without using the operators + and -.

// Example 1:
// Input: a = 1, b = 2
// Output: 3

// Example 2:
// Input: a = 2, b = 3
// Output: 5
 
// Constraints:
//     -1000 <= a, b <= 1000

import "fmt"

func getSum(a int, b int) int {
    if a == 0 { return b }
    if b == 0 { return a }
    // (a & b) <<1 计算的是进位
    // a ^ b 计算的是不带进位的加法
    return getSum((a & b) << 1, a ^ b)
}

func main() {
    // Example 1:
    // Input: a = 1, b = 2
    // Output: 3
    fmt.Println(getSum(1, 2)) // 3
    // Example 2:
    // Input: a = 2, b = 3
    // Output: 5
    fmt.Println(getSum(2, 3)) // 5
}
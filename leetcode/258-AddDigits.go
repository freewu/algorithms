package main

// 258. Add Digits
// Given an integer num, repeatedly add all its digits until the result has only one digit, and return it.

// Example 1:
// Input: num = 38
// Output: 2
// Explanation: The process is
// 38 --> 3 + 8 --> 11
// 11 --> 1 + 1 --> 2 
// Since 2 has only one digit, return it.

// Example 2:
// Input: num = 0
// Output: 0
 
// Constraints:
//     0 <= num <= 2^31 - 1
 
// Follow up: Could you do it without any loop/recursion in O(1) runtime?

import "fmt"

// 递归处理
func addDigits(num int) int {
    if num < 10 {
        return num
    }
    return addDigits(num % 10 + num / 10)
}

// 迭代
func addDigits1(num int) int {
    res := num
    // 每次除10 到个位数
    for (res / 10) != 0 {
        t := res / 10
        t += res % 10
        res = t
    }
    return res
}

func main() {
    // 38 --> 3 + 8 --> 11
    // 11 --> 1 + 1 --> 2 
    fmt.Println(addDigits(38)) // 2
    fmt.Println(addDigits(0)) // 0

    fmt.Println(addDigits1(38)) // 2
    fmt.Println(addDigits1(0)) // 0
}
package main

// 面试题 08.05. Recursive Mulitply LCCI
// Write a recursive function to multiply two positive integers without using the * operator. 
// You can use addition, subtraction, and bit shifting, but you should minimize the number of those operations.

// Example 1:
// Input: A = 1, B = 10
// Output: 10

// Example 2:
// Input: A = 3, B = 4
// Output: 12

// Note:
//     The result will not overflow.

import "fmt"

func multiply(A int, B int) int {
    if B == 1 { return A }
    if B == 2 { return A << 1 }
    res := 0
    if B % 2 != 0 {
        res += A 
    }
    res += multiply(A, B  >> 1)  << 1
    return res 
}

func main() {
    // Example 1:
    // Input: A = 1, B = 10
    // Output: 10
    fmt.Println(multiply(1, 10)) // 10
    // Example 2:
    // Input: A = 3, B = 4
    // Output: 12
    fmt.Println(multiply(3, 4)) // 12
}
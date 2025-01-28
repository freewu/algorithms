package main

// 面试题 05.06. Convert Integer LCCI
// Write a function to determine the number of bits you would need to flip to convert integer A to integer B.

// Example1:
// Input: A = 29 (0b11101), B = 15 (0b01111)
// Output: 2

// Example2:
// Input: A = 1，B = 2
// Output: 2

// Note:
//     -2147483648 <= A, B <= 2147483647

import "fmt"

func convertInteger(A int, B int) int {
    res := 0
    for i := 0; i < 32; i++ {
        res += (A & 1) ^ (B & 1)
        A >>= 1
        B >>= 1
    }
    return res
}

func convertInteger1(A int, B int) int {
    res := 0
    for i := 0; i < 32; i++ {
        a, b := 0, 0
        if A != 0 { a = A & (1 << i) }
        if B != 0 { b = B & (1 << i) }
        if a != b { res++ }
    }
    return res
}

func main() {
    // Example1:
    // Input: A = 29 (0b11101), B = 15 (0b01111)
    // Output: 2
    fmt.Println(convertInteger(29, 15)) // 2
    // Example2:
    // Input: A = 1，B = 2
    // Output: 2
    fmt.Println(convertInteger(1, 2)) // 2

    fmt.Println(convertInteger(1024, 2046)) // 9
    fmt.Println(convertInteger(-1, 1)) // 31
    fmt.Println(convertInteger(-1, -1)) // 0

    fmt.Println(convertInteger1(29, 15)) // 2
    fmt.Println(convertInteger1(1, 2)) // 2
    fmt.Println(convertInteger1(1024, 2046)) // 9
    fmt.Println(convertInteger1(-1, 1)) // 31
    fmt.Println(convertInteger1(-1, -1)) // 0
}
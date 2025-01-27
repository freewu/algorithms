package main

// 面试题 17.01. Add Without Plus LCCI
// Write a function that adds two numbers. You should not use + or any arithmetic operators.

// Example:
// Input: a = 1, b = 1
// Output: 2

// Note:
//     a and b may be 0 or negative.
//     The result fits in 32-bit integer.

import "fmt"

// 考虑两个二进制位相加的四种情况如下：
//     0 + 0 = 0
//     0 + 1 = 1
//     1 + 0 = 1
//     1 + 1 = 0 (进位)
// 可以发现，对于整数 a 和 b：
//      1. 在不考虑进位的情况下，其无进位加法结果为 a⊕b。
//      2. 而所有需要进位的位为 a & b，进位后的进位结果为 (a & b) << 1。
// 于是，我们可以将整数 a 和 b 的和，拆分为 a 和 b 的无进位加法结果与进位结果的和。
// 因为每一次拆分都可以让需要进位的最低位至少左移一位，又因为 a 和 b 可以取到负数，所以我们最多需要 log(max_int) 次拆分即可完成运算。
func add(a int, b int) int {
    for b != 0 {
        carry := uint(a&b) << 1
        a ^= b
        b = int(carry)
    }
    return a
}

func main() {
    // Example:
    // Input: a = 1, b = 1
    // Output: 2
    fmt.Println(add(1, 1)) // 2

    // 5 + 17 = 22
    // 101 + 10001 = 10110
    // 101
    // 001
    fmt.Println(add(5, 17)) // 22

    fmt.Println(add(1024, 1)) // 1025
    fmt.Println(add(999_999_999, 1)) // 1_000_000_000
}
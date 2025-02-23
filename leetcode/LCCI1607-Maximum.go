package main

// 面试题 16.07. Maximum LCCI
// Write a method that finds the maximum of two numbers. 
// You should not use if-else or any other comparison operator.

// Example:
// Input:  a = 1, b = 2
// Output:  2

import "fmt"

func maximum(a int, b int) int {
    // 不许使用判断语句和比较运算符，可以用位运算
    // 负数的最高位为 1， 而正数的最高位为 0。
    // 右移63位（默认数据是64位整形）得到a-b的最高位 k （0或1）这样可以返回a*(^k)+b*k
    k := (a - b) >> 63     // 0 (a >= b) or -1 (not 1, a < b)
    return -(a*(^k) + b*k) // ^0 = -1, ^(-1) = 0
}

func maximum1(a int, b int) int {
    return a - ((a - b) & ((a - b) >> 63))
}

func main() {
    // Example:
    // Input:  a = 1, b = 2
    // Output:  2
    fmt.Println(maximum(1, 2)) // 2

    fmt.Println(maximum1(1, 2)) // 2
}
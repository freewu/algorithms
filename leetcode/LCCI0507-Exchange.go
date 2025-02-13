package main

// 面试题 05.07. Exchange LCCI
// Write a program to swap odd and even bits in an integer with as few instructions as possible 
// (e.g., bit 0 and bit 1 are swapped, bit 2 and bit 3 are swapped, and so on).

// Example1:
// Input: num = 2（0b10）
// Output 1 (0b01)

// Example2:
// Input: num = 3
// Output: 3

// Note:
//     0 <= num <= 2^30 - 1
//     The result integer fits into 32-bit integer.

import "fmt"

func exchangeBits(num int) int {
    even := num & 0xaaaaaaaa
    odd := num & 0x55555555
    return even >> 1 | odd << 1
}

func exchangeBits1(num int) int {
    return (num & 0xaaaaaaaa) >> 1 | (num & 0x55555555) << 1
}

func exchangeBits2(num int) int {
    w, a := 1, num 
    for i := 0;i < 32; i+=2 {
        p := a % 2 + ((a >> 1) % 2)*2
        if p == 1{
            num += w
        } else if p == 2 {
            num -= w
        }
        w *= 4
        a >>= 2
    }
    return num 
}

func main() {
    // Example1:
    // Input: num = 2（0b10）
    // Output 1 (0b01)
    fmt.Println(exchangeBits(2)) // 1
    // Example2:
    // Input: num = 3
    // Output: 3
    fmt.Println(exchangeBits(3)) // 3

    fmt.Println(exchangeBits(0)) // 0
    fmt.Println(exchangeBits(1)) // 2
    fmt.Println(exchangeBits(8)) // 4
    fmt.Println(exchangeBits(1024)) // 2048
    fmt.Println(exchangeBits(1 << 29)) // 268435456

    fmt.Println(exchangeBits1(2)) // 1
    fmt.Println(exchangeBits1(3)) // 3
    fmt.Println(exchangeBits1(0)) // 0
    fmt.Println(exchangeBits1(1)) // 2
    fmt.Println(exchangeBits1(8)) // 4
    fmt.Println(exchangeBits1(1024)) // 2048
    fmt.Println(exchangeBits1(1 << 29)) // 268435456

    fmt.Println(exchangeBits2(2)) // 1
    fmt.Println(exchangeBits2(3)) // 3
    fmt.Println(exchangeBits2(0)) // 0
    fmt.Println(exchangeBits2(1)) // 2
    fmt.Println(exchangeBits2(8)) // 4
    fmt.Println(exchangeBits2(1024)) // 2048
    fmt.Println(exchangeBits2(1 << 29)) // 268435456
}
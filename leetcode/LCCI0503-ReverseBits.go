package main

// 面试题 05.03. Reverse Bits LCCI
// You have an integer and you can flip exactly one bit from a 0 to a 1. 
// Write code to find the length of the longest sequence of 1s you could create.

// Example 1:
// Input: num = 1775(110111011112)
// Output: 8

// Example 2:
// Input: num = 7(01112)
// Output: 4

import "fmt"

func reverseBits(num int) int {
    res, prev, cur := 0, 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < 32; i++ {
        if num & 1 == 1 {
            cur++
        } else {
            prev = cur + 1
            cur = 0
        }
        res = max(res, prev + cur)
        num >>= 1
    }
    return res
}

func main() {
    // Example 1:
    // Input: num = 1775(110111011112)
    // Output: 8
    fmt.Println(reverseBits(1775)) // 8
    // Example 2:
    // Input: num = 7(01112)
    // Output: 4
    fmt.Println(reverseBits(7)) // 4

    fmt.Println(reverseBits(0)) // 1
    fmt.Println(reverseBits(1)) // 2
    fmt.Println(reverseBits(1024)) // 2
    fmt.Println(reverseBits(999_999_999)) // 10
    fmt.Println(reverseBits(1_000_000_000)) // 7
}
package main

// 461. Hamming Distance
// The Hamming distance between two integers is the number of positions at which the corresponding bits are different.
// Given two integers x and y, return the Hamming distance between them.

// Example 1:
// Input: x = 1, y = 4
// Output: 2
// Explanation:
// 1   (0 0 0 1)
// 4   (0 1 0 0)
//        ↑   ↑
// The above arrows point to positions where the corresponding bits are different.

// Example 2:
// Input: x = 3, y = 1
// Output: 1
 
// Constraints:
//     0 <= x, y <= 2^31 - 1

import "fmt"

func hammingDistance(x int, y int) int {
    res := 0
    if x == y {
        return res
    }
    // x 与 y 作异或运算，然后统计 1 的数量 x&(x-1) 将最后位置变成 0，直到 x 变成 0 的次数
    for xor := x ^ y; xor != 0; xor &= (xor - 1) {
        res++
    }
    return res
}

func main() {
    // Explanation:
    // 1   (0 0 0 1)
    // 4   (0 1 0 0)
    //        ↑   ↑
    // The above arrows point to positions where the corresponding bits are different.
    fmt.Println(hammingDistance(1, 4)) // 2
    // 1   (0 0 0 1)
    // 3   (0 0 1 1)
    //          ↑   
    fmt.Println(hammingDistance(3, 1)) // 1
}
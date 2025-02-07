package main

// 面试题 05.01. Insert Into Bits LCCI
// You are given two 32-bit numbers, N and M, and two bit positions, i and j. 
// Write a method to insert M into N such that M starts at bit j and ends at bit i. 
// You can assume that the bits j through i have enough space to fit all of M. 
// That is, if M = 10011, you can assume that there are at least 5 bits between j and i. 
// You would not, for example, have j = 3 and i = 2, because M could not fully fit between bit 3 and bit 2.

// Example1:
// Input: N = 10000000000, M = 10011, i = 2, j = 6
// Output: N = 10001001100

// Example2:
// Input:  N = 0, M = 11111, i = 0, j = 4
// Output: N = 11111

import "fmt"

func insertBits(N int, M int, i int, j int) int {
    nj := N >> (uint(j) + 1) << (uint(j) + 1)
    ni := (N >> uint(i) << uint(i)) ^ N
    return (ni | nj) | M << uint(i)
}

func insertBits1(N int, M int, i int, j int) int {
    count, data := j - i + 1, 0
    for k := 0; k < count; k++ {
        data <<= 1
        data |= 1
    }
    data <<= i
    data = ^data
    N &= data
    N |= M << i
    return N
}

func main() {
    // Example1:
    // Input: N = 10000000000, M = 10011, i = 2, j = 6
    // Output: 10000006252
    fmt.Println(insertBits(10000000000, 10011, 2, 6)) // 10000006252
    // Example2:
    // Input:  N = 0, M = 11111, i = 0, j = 4
    // Output: 11111
    fmt.Println(insertBits(0, 11111, 0, 4)) // 11111

    fmt.Println(insertBits1(10000000000, 10011, 2, 6)) // 10000006252
    fmt.Println(insertBits1(0, 11111, 0, 4)) // 11111
}
package main

// 762. Prime Number of Set Bits in Binary Representation
// Given two integers left and right,
// return the count of numbers in the inclusive range [left, right] having a prime number of set bits in their binary representation.

// Recall that the number of set bits an integer has is the number of 1's present when written in binary.
//     For example, 21 written in binary is 10101, which has 3 set bits.

// Example 1:
// Input: left = 6, right = 10
// Output: 4
// Explanation:
// 6  -> 110 (2 set bits, 2 is prime)
// 7  -> 111 (3 set bits, 3 is prime)
// 8  -> 1000 (1 set bit, 1 is not prime)
// 9  -> 1001 (2 set bits, 2 is prime)
// 10 -> 1010 (2 set bits, 2 is prime)
// 4 numbers have a prime number of set bits.

// Example 2:
// Input: left = 10, right = 15
// Output: 5
// Explanation:
// 10 -> 1010 (2 set bits, 2 is prime)
// 11 -> 1011 (3 set bits, 3 is prime)
// 12 -> 1100 (2 set bits, 2 is prime)
// 13 -> 1101 (3 set bits, 3 is prime)
// 14 -> 1110 (3 set bits, 3 is prime)
// 15 -> 1111 (4 set bits, 4 is not prime)
// 5 numbers have a prime number of set bits.
 
// Constraints:
//     1 <= left <= right <= 10^6
//     0 <= right - left <= 10^4

import "fmt"
import "math/bits"

func countPrimeSetBits(left int, right int) int {
    res, a := 0, make([]int, 32)
    a[2], a[3], a[5], a[7], a[11] = 1,1,1,1,1
    a[13], a[17], a[19], a[23], a[29], a[31] = 1,1,1,1,1,1
    for i := left; i < right + 1; i++ {
        count, temp := 0, i
        for temp != 0 {
            count += temp & 1
            temp >>= 1
        }
        res += a[count]
    }
    return res
}


func countPrimeSetBits1(left, right int) int {
    res := 0
    for i := left; i <= right; i++ {
        if 1 << bits.OnesCount(uint(i)) & 665772 != 0 {
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: left = 6, right = 10
    // Output: 4
    // Explanation:
    // 6  -> 110 (2 set bits, 2 is prime)
    // 7  -> 111 (3 set bits, 3 is prime)
    // 8  -> 1000 (1 set bit, 1 is not prime)
    // 9  -> 1001 (2 set bits, 2 is prime)
    // 10 -> 1010 (2 set bits, 2 is prime)
    // 4 numbers have a prime number of set bits.
    fmt.Println(countPrimeSetBits(6,10)) // 4
    // Example 2:
    // Input: left = 10, right = 15
    // Output: 5
    // Explanation:
    // 10 -> 1010 (2 set bits, 2 is prime)
    // 11 -> 1011 (3 set bits, 3 is prime)
    // 12 -> 1100 (2 set bits, 2 is prime)
    // 13 -> 1101 (3 set bits, 3 is prime)
    // 14 -> 1110 (3 set bits, 3 is prime)
    // 15 -> 1111 (4 set bits, 4 is not prime)
    // 5 numbers have a prime number of set bits.
    fmt.Println(countPrimeSetBits(10,15)) // 5

    fmt.Println(countPrimeSetBits(1,1)) // 0
    fmt.Println(countPrimeSetBits(1,10000)) // 4252
    fmt.Println(countPrimeSetBits(990_000,1_000_000)) // 3754

    fmt.Println(countPrimeSetBits1(6,10)) // 4
    fmt.Println(countPrimeSetBits1(10,15)) // 5
    fmt.Println(countPrimeSetBits1(1,1)) // 0
    fmt.Println(countPrimeSetBits1(1,10000)) // 4252
    fmt.Println(countPrimeSetBits1(990_000,1_000_000)) // 3754

    // Constraints:
//     1 <= left <= right <= 10^6
//     0 <= right - left <= 10^4
}
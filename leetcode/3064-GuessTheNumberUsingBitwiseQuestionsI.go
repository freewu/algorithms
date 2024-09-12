package main

// 3064. Guess the Number Using Bitwise Questions I
// There is a number n that you have to find.

// There is also a pre-defined API int commonSetBits(int num), 
// which returns the number of bits where both n and num are 1 in that position of their binary representation. 
// In other words, it returns the number of set bits in n & num, where & is the bitwise AND operator.

// Return the number n.

// Example 1:
// Input: n = 31
// Output: 31
// Explanation: It can be proven that it's possible to find 31 using the provided API.

// Example 2:
// Input: n = 33
// Output: 33
// Explanation: It can be proven that it's possible to find 33 using the provided API.

// Constraints:
//     1 <= n <= 2^30 - 1
//     0 <= num <= 2^30 - 1
//     If you ask for some num out of the given range, the output wouldn't be reliable.

import "fmt"

var val = 0

func commonSetBits(num int) int {
    return val
}

/** 
 * Definition of commonSetBits API.
 * func commonSetBits(num int) int;
 */
func findNumber() int {
    res := 0
    for i := 0; i < 30; i++ {
        if (commonSetBits(1 << i) > 0) {
            res += 1 << i
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 31
    // Output: 31
    // Explanation: It can be proven that it's possible to find 31 using the provided API.
    val = 31
    fmt.Println(findNumber()) // 31
    // Example 2:
    // Input: n = 33
    // Output: 33
    // Explanation: It can be proven that it's possible to find 33 using the provided API.
    val = 33
    fmt.Println(findNumber()) // 33
}
package main

// 1015. Smallest Integer Divisible by K
// Given a positive integer k, 
// you need to find the length of the smallest positive integer n such that n is divisible by k, 
// and n only contains the digit 1.

// Return the length of n. If there is no such n, return -1.
// Note: n may not fit in a 64-bit signed integer.

// Example 1:
// Input: k = 1
// Output: 1
// Explanation: The smallest answer is n = 1, which has length 1.

// Example 2:
// Input: k = 2
// Output: -1
// Explanation: There is no such positive integer n divisible by 2.

// Example 3:
// Input: k = 3
// Output: 3
// Explanation: The smallest answer is n = 111, which has length 3.
 
// Constraints:
//     1 <= k <= 10^5

import "fmt"

func smallestRepunitDivByK(k int) int {
    if k % 2 == 0 || k % 5 == 0 { 
        return -1 
    }
    cur, res := 0, 1
    for {
        cur = (cur * 10 + 1) % k
        if cur == 0 { 
            return res 
        }
        res++
    }
    return -1
}

func main() {
    // Example 1:
    // Input: k = 1
    // Output: 1
    // Explanation: The smallest answer is n = 1, which has length 1.
    fmt.Println(smallestRepunitDivByK(1)) // 1
    // Example 2:
    // Input: k = 2
    // Output: -1
    // Explanation: There is no such positive integer n divisible by 2.
    fmt.Println(smallestRepunitDivByK(2)) // -1
    // Example 3:
    // Input: k = 3
    // Output: 3
    // Explanation: The smallest answer is n = 111, which has length 3.
    fmt.Println(smallestRepunitDivByK(3)) // 3

    fmt.Println(smallestRepunitDivByK(10000)) // -1
    fmt.Println(smallestRepunitDivByK(9999)) // 36
    fmt.Println(smallestRepunitDivByK(99999)) // 45 
    fmt.Println(smallestRepunitDivByK(8)) // -1
    fmt.Println(smallestRepunitDivByK(64)) // -1
    fmt.Println(smallestRepunitDivByK(99)) // 18
    fmt.Println(smallestRepunitDivByK(100)) // -1
    fmt.Println(smallestRepunitDivByK(1024)) // -1
    fmt.Println(smallestRepunitDivByK(100_000)) // -1
}
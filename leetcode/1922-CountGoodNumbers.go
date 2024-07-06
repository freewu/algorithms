package main

// 1922. Count Good Numbers
// A digit string is good if the digits (0-indexed) at even indices are even and 
// the digits at odd indices are prime (2, 3, 5, or 7).

//     For example, "2582" is good because the digits (2 and 8) at even positions are even and the digits (5 and 2) at odd positions are prime. 
//     However, "3245" is not good because 3 is at an even index but is not even.
    
// Given an integer n, return the total number of good digit strings of length n. 
// Since the answer may be large, return it modulo 10^9 + 7.

// A digit string is a string consisting of digits 0 through 9 that may contain leading zeros.

// Example 1:
// Input: n = 1
// Output: 5
// Explanation: The good numbers of length 1 are "0", "2", "4", "6", "8".

// Example 2:
// Input: n = 4
// Output: 400

// Example 3:
// Input: n = 50
// Output: 564908303

// Constraints:
//     1 <= n <= 10^15

import "fmt"

func countGoodNumbers(n int64) int {
    mod, odd, even := int64(1e9 + 7), n / 2, (n / 2) + (n % 2)
    // power takes x, y intergars and returns the x^y with module 10^9+7
    var power func (x, y int64) int64
    power = func(x, y int64) int64 {
        if y == 0 { // x^0 = 1
            return 1
        }
        if y == 1 { // x^1 = x
            return x
        }
        res := power(x, y/2) % mod // Take the power of half
        res = (res * res) % mod // Multiply with same to make it x^y
        if (y % 2 == 1) { // If this is odd then one more multiplication is needed 奇数
            res = res * x
        }
        return res % mod // Take mod before returning
    }
    return int((power(5, even) * power(4, odd)) % mod)
}

func main() {
    // Example 1:
    // Input: n = 1
    // Output: 5
    // Explanation: The good numbers of length 1 are "0", "2", "4", "6", "8".
    fmt.Println(countGoodNumbers(1)) // 5
    // Example 2:
    // Input: n = 4
    // Output: 400
    fmt.Println(countGoodNumbers(4)) // 400
    // Example 3:
    // Input: n = 50
    // Output: 564908303
    fmt.Println(countGoodNumbers(50)) // 564908303
    fmt.Println(countGoodNumbers(1024)) // 267965558
}
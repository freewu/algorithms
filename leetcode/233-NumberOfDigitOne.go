package main

// 233. Number of Digit One
// Given an integer n, count the total number of digit 1 appearing in all non-negative integers less than or equal to n.

// Example 1:
// Input: n = 13
// Output: 6

// Example 2:
// Input: n = 0
// Output: 0

// Constraints:
//     0 <= n <= 10^9

import "fmt"

func countDigitOne(n int) int {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    count, base := 0, 1
    for base <= n {
        divider := base * 10
        count += (n / divider) * base + min(max(n % divider - base + 1, 0), base)
        base *= 10
    }
    return count
}

func main() {
    // Example 1:
    // Input: n = 13
    // Output: 6
    // 1 10 11 12 13
    fmt.Println(countDigitOne(13)) // 6
    // Example 2:
    // Input: n = 0
    // Output: 0
    fmt.Println(countDigitOne(0)) // 0

    fmt.Println(countDigitOne(1)) // 1
}
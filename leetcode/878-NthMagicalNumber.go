package main

// 878. Nth Magical Number
// A positive integer is magical if it is divisible by either a or b.

// Given the three integers n, a, and b, return the nth magical number. 
// Since the answer may be very large, return it modulo 10^9 + 7.

// Example 1:
// Input: n = 1, a = 2, b = 3
// Output: 2

// Example 2:
// Input: n = 4, a = 2, b = 3
// Output: 6

// Constraints:
//     1 <= n <= 10^9
//     2 <= a, b <= 4 * 10^4

import "fmt"

func nthMagicalNumber(n int, a int, b int) int {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    n1, n2, mod := a, b, 1_000_000_007
    l, r := min(a, b), n * min(a, b)
    for n2 > 0 {
        n2, n1 = n1 % n2, n2
    }
    lcm := (a * b) / n1
    for l < r {
        mid := l + (r - l) / 2
        if (mid / a) + (mid / b) - (mid / lcm) < n {
            l = mid + 1
        } else {
            r = mid
        }
    }
    return l % mod
}

func main() {
    // Example 1:
    // Input: n = 1, a = 2, b = 3
    // Output: 2
    fmt.Println(nthMagicalNumber(1,2,3)) // 2
    // Example 2:
    // Input: n = 4, a = 2, b = 3
    // Output: 6
    fmt.Println(nthMagicalNumber(4,2,3)) // 6
}
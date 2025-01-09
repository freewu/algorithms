package main

// 1201. Ugly Number III
// An ugly number is a positive integer that is divisible by a, b, or c.

// Given four integers n, a, b, and c, return the nth ugly number.

// Example 1:
// Input: n = 3, a = 2, b = 3, c = 5
// Output: 4
// Explanation: The ugly numbers are 2, 3, 4, 5, 6, 8, 9, 10... The 3rd is 4.

// Example 2:
// Input: n = 4, a = 2, b = 3, c = 4
// Output: 6
// Explanation: The ugly numbers are 2, 3, 4, 6, 8, 9, 10, 12... The 4th is 6.

// Example 3:
// Input: n = 5, a = 2, b = 11, c = 13
// Output: 10
// Explanation: The ugly numbers are 2, 4, 6, 8, 10, 11, 12, 13... The 5th is 10.

// Constraints:
//     1 <= n, a, b, c <= 10^9
//     1 <= a * b * c <= 10^18
//     It is guaranteed that the result will be in range [1, 2 * 10^9].

import "fmt"

func nthUglyNumber(n int, a int, b int, c int) int {
    gcd := func(x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    lcm := func(x int, y int) int { return (x * y) / gcd(x, y) }

    lcm_ab, lcm_ac, lcm_bc := lcm(a, b), lcm(a, c), lcm(b, c)
    lcm_abc := lcm(a, lcm_bc)
    uglies_less_equal := func(x int) int {
        return (x/a) + (x/b) + (x/c) - (x/lcm_ab) - (x/lcm_ac) - (x/lcm_bc) + (x/lcm_abc)
    }
    left, right, mid := 0, 1 << 31, 0
    for left < right {
        mid = left + (right - left) / 2
        if uglies_less_equal(mid) < n {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return left
}

func main() {
    // Example 1:
    // Input: n = 3, a = 2, b = 3, c = 5
    // Output: 4
    // Explanation: The ugly numbers are 2, 3, 4, 5, 6, 8, 9, 10... The 3rd is 4.
    fmt.Println(nthUglyNumber(3,2,3,5)) // 4
    // Example 2:
    // Input: n = 4, a = 2, b = 3, c = 4
    // Output: 6
    // Explanation: The ugly numbers are 2, 3, 4, 6, 8, 9, 10, 12... The 4th is 6.
    fmt.Println(nthUglyNumber(4,2,3,4)) // 6
    // Example 3:
    // Input: n = 5, a = 2, b = 11, c = 13
    // Output: 10
    // Explanation: The ugly numbers are 2, 4, 6, 8, 10, 11, 12, 13... The 5th is 10.
    fmt.Println(nthUglyNumber(5,2,11,13)) // 10
}
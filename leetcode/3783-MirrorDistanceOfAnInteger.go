package main

// 3783. Mirror Distance of an Integer
// You are given an integer n.

// Define its mirror distance as: abs(n - reverse(n))​​​​​​​ where reverse(n) is the integer formed by reversing the digits of n.

// Return an integer denoting the mirror distance of n​​​​​​​.

// abs(x) denotes the absolute value of x.

// Example 1:
// Input: n = 25
// Output: 27
// Explanation:
// reverse(25) = 52.
// Thus, the answer is abs(25 - 52) = 27

// Example 2:
// Input: n = 10
// Output: 9
// Explanation:
// reverse(10) = 01 which is 1.
// Thus, the answer is abs(10 - 1) = 9.

// Example 3:
// Input: n = 7
// Output: 0
// Explanation:
// reverse(7) = 7.
// Thus, the answer is abs(7 - 7) = 0.
 
// Constraints:
//     1 <= n <= 10^9

import "fmt"

// time: O(log n), space: O(1)
func mirrorDistance(n int) int {
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    reverse := func(n int) int {
        res := 0
        for ; n != 0; n /= 10 {
            res = res * 10 + n % 10
        }
        return res
    }
    return abs(n - reverse(n))
}

func mirrorDistance1(n int) int {
    rev, v := 0, n
    for v > 0 {
        rev = rev * 10 +  v% 10
        v /= 10 
    }
    if n >= rev { return n - rev }
    return rev - n
}

func main() {
    // Example 1:
    // Input: n = 25
    // Output: 27
    // Explanation:
    // reverse(25) = 52.
    // Thus, the answer is abs(25 - 52) = 27.
    fmt.Println(mirrorDistance(25)) // 27
    // Example 2:
    // Input: n = 10
    // Output: 9
    // Explanation:
    // reverse(10) = 01 which is 1.
    // Thus, the answer is abs(10 - 1) = 9.
    fmt.Println(mirrorDistance(10)) // 9
    // Example 3:
    // Input: n = 7
    // Output: 0
    // Explanation:
    // reverse(7) = 7.
    // Thus, the answer is abs(7 - 7) = 0.
    fmt.Println(mirrorDistance(7)) // 0

    fmt.Println(mirrorDistance(1)) // 0
    fmt.Println(mirrorDistance(99)) // 0
    fmt.Println(mirrorDistance(100)) // 99
    fmt.Println(mirrorDistance(1024)) // 3177
    fmt.Println(mirrorDistance(999_999_999)) // 0
    fmt.Println(mirrorDistance(1_000_000_000)) // 999_999_999

    fmt.Println(mirrorDistance1(25)) // 27
    fmt.Println(mirrorDistance1(10)) // 9
    fmt.Println(mirrorDistance1(7)) // 0
    fmt.Println(mirrorDistance1(1)) // 0
    fmt.Println(mirrorDistance1(99)) // 0
    fmt.Println(mirrorDistance1(100)) // 99
    fmt.Println(mirrorDistance1(1024)) // 3177
    fmt.Println(mirrorDistance1(999_999_999)) // 0
    fmt.Println(mirrorDistance1(1_000_000_000)) // 999_999_999
}
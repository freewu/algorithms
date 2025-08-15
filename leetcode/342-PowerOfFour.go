package main

// 342. Power of Four
// Given an integer n, return true if it is a power of four. Otherwise, return false.
// An integer n is a power of four, if there exists an integer x such that n == 4x.

// Example 1:
// Input: n = 16
// Output: true

// Example 2:
// Input: n = 5
// Output: false

// Example 3:
// Input: n = 1
// Output: true

// Constraints:
//     -2^31 <= n <= 2^31 - 1

// Follow up: Could you solve it without loops/recursion?

import "fmt"

// loop 
func isPowerOfFour(n int) bool {
    for n >= 4 {
        // 每次除4 判断是否能被 4 整除
        if n % 4 == 0 {
            n = n / 4
        } else {
            return false
        }
    }
    return n == 1
}

// 数论
// 证明 (4^n - 1) % 3 == 0，
//  (1) 4^n - 1 = (2^n + 1) * (2^n - 1)
//  (2) 在任何连续的 3 个数中 (2^n-1)，(2^n)，(2^n+1)，一定有一个数是 3 的倍数。
//      (2^n) 肯定不是 3 的倍数，
//      那么 (2^n-1) 或者 (2^n+1) 中一定有一个是 3 的倍数。
//      所以 4^n-1 一定是 3 的倍数
func isPowerOfFour1(num int) bool {
    return num > 0 && (num & (num-1)) == 0 && (num-1) % 3 == 0
}

func main() {
    // Example 1:
    // Input: n = 16
    // Output: true
    fmt.Println(isPowerOfFour(16)) // true
    // Example 2:
    // Input: n = 5
    // Output: false
    fmt.Println(isPowerOfFour(5)) // false
    // Example 3:
    // Input: n = 1
    // Output: true
    fmt.Println(isPowerOfFour(1)) // true

    fmt.Println(isPowerOfFour(-2147483648)) // false
    fmt.Println(isPowerOfFour(4)) // true
    fmt.Println(isPowerOfFour(2 << 31 - 1)) // false
    fmt.Println(isPowerOfFour(-2 << 31)) // false

    fmt.Println(isPowerOfFour1(16)) // true
    fmt.Println(isPowerOfFour1(5)) // false
    fmt.Println(isPowerOfFour1(1)) // true
    fmt.Println(isPowerOfFour1(-2147483648)) // false
    fmt.Println(isPowerOfFour1(4)) // true
    fmt.Println(isPowerOfFour1(2 << 31 - 1)) // false
    fmt.Println(isPowerOfFour1(-2 << 31)) // false
}
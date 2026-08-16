package main

// 4022. K-th Digit in Infinite String
// You are given an integer k.

// An infinite string is formed by concatenating the decimal representations of the positive integers, without separators.

// For every nonnegative integer b, block b contains the positive integers from 10 * b through 10 * b + 9. 
// The integers in each block are appended as follows:
//     1. If b is even, append the integers in increasing order.
//     2. If b is odd, append the integers in decreasing order.

// Therefore, the string starts with the integers 1 through 9, followed by 19 through 10, then 20 through 29, then 39 through 30, and so on.

// Return the kth digit (1-indexed) of this string.

// Example 1:
// Input: k = 4
// Output: 4
// Explanation:
// The string begins as "123456789..". The 4th digit is '4'.

// Example 2:
// Input: k = 15
// Output: 7
// Explanation:
// The string begins as "123456789191817..". The 15th digit is '7'.

// Example 3:
// Input: k = 11
// Output: 9
// Explanation:
// The string begins as "12345678919..". The 11th digit is '9'.

// Constraints:
//     1 <= k <= 10^15

import "fmt"

func kthDigit(k int64) int {
    k1 := int(k - 1) // k 改成从 0 开始，方便计算
    // 十进制长为 n 的正整数有 count = 9 * 10^(n-1) 个
    count, n := 9, 1
    for count * n <= k1 {
        k1 -= count * n // 这里减小了 k
        count *= 10
        n++
    }
    // k 在正整数 x 中
    x := count / 9 + k1 / n
    if x / 10 % 2 > 0 {
        // 改成递减顺序，例如 10 变成 19，11 变成 18 ……
        x += 9 - x % 10 * 2
    }
    // 计算 x 从高到低第 k%n（从 0 开始）个数字
    for range n - k1 % n - 1 {
        x /= 10
    }
    return x % 10
}

func main() {
    // Example 1:
    // Input: k = 4
    // Output: 4
    // Explanation:
    // The string begins as "123456789..". The 4th digit is '4'.
    fmt.Println(kthDigit(4)) // 4
    // Example 2:
    // Input: k = 15
    // Output: 7
    // Explanation:
    // The string begins as "123456789191817..". The 15th digit is '7'.
    fmt.Println(kthDigit(15)) // 7
    // Example 3:
    // Input: k = 11
    // Output: 9
    // Explanation:
    // The string begins as "12345678919..". The 11th digit is '9'.
    fmt.Println(kthDigit(11)) // 9

    fmt.Println(kthDigit(1)) // 1
    fmt.Println(kthDigit(2)) // 2
    fmt.Println(kthDigit(3)) // 3
    fmt.Println(kthDigit(8)) // 8
    fmt.Println(kthDigit(64)) // 3
    fmt.Println(kthDigit(99)) // 5
    fmt.Println(kthDigit(100)) // 5
    fmt.Println(kthDigit(101)) // 4
    fmt.Println(kthDigit(1024)) // 3
    fmt.Println(kthDigit(999_999_999_999_999)) // 2
    fmt.Println(kthDigit(1_000_000_000_000_000)) // 2
}
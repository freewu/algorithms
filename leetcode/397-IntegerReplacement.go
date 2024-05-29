package main

// 397. Integer Replacement
// Given a positive integer n, you can apply one of the following operations:
//     If n is even, replace n with n / 2.
//     If n is odd, replace n with either n + 1 or n - 1.

// Return the minimum number of operations needed for n to become 1.

// Example 1:
// Input: n = 8
// Output: 3
// Explanation: 8 -> 4 -> 2 -> 1

// Example 2:
// Input: n = 7
// Output: 4
// Explanation: 7 -> 8 -> 4 -> 2 -> 1
// or 7 -> 6 -> 3 -> 2 -> 1

// Example 3:
// Input: n = 4
// Output: 2
 
// Constraints:
//     1 <= n <= 2^31 - 1

import "fmt"

func integerReplacement1(n int) int {
    res := 0
    for ; n > 1; res++ {
        switch {
        case n % 2 == 0:
            n >>= 1
        case (n - 1) % 4 == 0 || n == 3:
            n--
        default:
            n++
        }
    }
    return res
}

// 当 n 为奇数的时候，什么时候需要加 1 ，什么时候需要减 1 ，
// 通过观察规律可以发现，除了 3 和 7 以外，所有加 1 就变成 4 的倍数的奇数，都适合先加 1 运算，
// 比如 15:
//     15 -> 16 -> 8 -> 4 -> 2 -> 1 (5)
//     15 -> 14 -> 7 -> 6 -> 3 -> 2 -> 1 (6)
func integerReplacement(n int) int {
    count := 0
    for {
        if n == 1 { break }
        if n % 2 == 1 { // If n is odd, replace n with either n + 1 or n - 1.
            if (n - 1) % 4 == 0 || n == 3 {  // 处理
                n--
            } else { // 除了 3 和 7 以外, 所有加 1 就变成 4 的倍数的奇数，都适合先加 1 运算
                n++
            }
        } else { // If n is even, replace n with n / 2.
            n = n >> 1
        }
        count++
    }
    return count
}

func main() {
    // Example 1:
    // Input: n = 8
    // Output: 3
    // Explanation: 8 -> 4 -> 2 -> 1
    fmt.Println(integerReplacement(8)) // 3  8 -> 4 -> 2 -> 1
    // Example 2:
    // Input: n = 7
    // Output: 4
    // Explanation: 7 -> 8 -> 4 -> 2 -> 1
    // or 7 -> 6 -> 3 -> 2 -> 1
    fmt.Println(integerReplacement(7)) // 4  7 -> 8 -> 4 -> 2 -> 1 || 7 -> 6 -> 3 -> 2 -> 1
    // Example 3:
    // Input: n = 4
    // Output: 2
    fmt.Println(integerReplacement(4)) // 2  4 -> 2 -> 1 

    fmt.Println(integerReplacement(1234)) // 14
    fmt.Println(integerReplacement(15)) // 5

    fmt.Println(integerReplacement1(8)) // 3  8 -> 4 -> 2 -> 1
    fmt.Println(integerReplacement1(7)) // 4  7 -> 8 -> 4 -> 2 -> 1 || 7 -> 6 -> 3 -> 2 -> 1
    fmt.Println(integerReplacement1(4)) // 2  4 -> 2 -> 1 
    fmt.Println(integerReplacement1(1234)) // 14
    fmt.Println(integerReplacement1(15)) // 5
}
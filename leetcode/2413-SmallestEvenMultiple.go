package main

// 2413. Smallest Even Multiple
// Given a positive integer n, return the smallest positive integer that is a multiple of both 2 and n.
 
// Example 1:
// Input: n = 5
// Output: 10
// Explanation: The smallest multiple of both 5 and 2 is 10.

// Example 2:
// Input: n = 6
// Output: 6
// Explanation: The smallest multiple of both 6 and 2 is 6. Note that a number is a multiple of itself.

// Constraints:
//     1 <= n <= 150

import "fmt"

func smallestEvenMultiple(n int) int {
    // 能被 2 整除返回本身
    if n % 2 == 0 {
        return n
    }
    // 否则 * 2 返回
    return n * 2
}

func main() {
    fmt.Println(smallestEvenMultiple(5)) // 10
    fmt.Println(smallestEvenMultiple(6)) // 6
    fmt.Println(smallestEvenMultiple(11)) // 22
    fmt.Println(smallestEvenMultiple(12)) // 12
}
package main

// 172. Factorial Trailing Zeroes
// Given an integer n, return the number of trailing zeroes in n!.
// Note that n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1.

// Example 1:
// Input: n = 3
// Output: 0
// Explanation: 3! = 6, no trailing zero.

// Example 2:
// Input: n = 5
// Output: 1
// Explanation: 5! = 120, one trailing zero.

// Example 3:
// Input: n = 0
// Output: 0
 
// Constraints:
//     0 <= n <= 10^4
 
// Follow up: Could you write a solution that works in logarithmic time complexity?

import "fmt"

func trailingZeroes(n int) int {
    // 这是一道数学题。计算 N 的阶乘有多少个后缀 0，
    // 即计算 N! 里有多少个 10，也是计算 N! 里有多少个 2 和 5（分解质因数），最后结果即 2 的个数和 5 的个数取较小值。
    // 每两个数字就会多一个质因数 2，而每五个数字才多一个质因数 5。每 5 个数字就会多一个质因数 5。
    //      0~4 的阶乘里没有质因数 5，
    //      5~9 的阶乘里有 1 个质因数 5，
    //      10~14 的阶乘里有 2 个质因数 5，
    //      依此类推。所以 0 的个数即为 min(阶乘中 5 的个数和 2 的个数)。
    // N! 有多少个后缀 0，即 N! 有多少个质因数 5。
    // N! 有多少个质因数 5，即 N 可以划分成多少组 5个数字一组，加上划分成多少组 25 个数字一组，加上划分多少组成 125 个数字一组，等等。
    // 即 res = N/5 + N/(5^2) + N/(5^3) + ... = ((N / 5) / 5) / 5 /... 。
    // 最终算法复杂度为 O(logN)。
    if n / 5 == 0 {
        return 0
    }
    return n / 5 + trailingZeroes(n / 5)
}

func main() {
    // Example 1:
    // Input: n = 3
    // Output: 0
    // Explanation: 3! = 6, no trailing zero.
    fmt.Println(trailingZeroes(3)) // 0
    // Example 2:
    // Input: n = 5
    // Output: 1
    // Explanation: 5! = 120, one trailing zero.
    fmt.Println(trailingZeroes(5)) // 1
    // Example 3:
    // Input: n = 0
    // Output: 0
    fmt.Println(trailingZeroes(0)) // 0
}
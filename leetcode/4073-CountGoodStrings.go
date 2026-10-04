package main 

// 4073. Count Good Strings
// You are given an integer n.

// A string is considered good if it consists only of the characters 'a' and 'b', and one of the following holds:
//     1. It contains exactly one distinct character, and its length is odd.
//     2. It can be written as s = s1 + s2, where s1 and s2 are non-empty good strings, 
//        and the last character of s1 is different from the first character of s2.

// Return the number of good strings of length n, modulo 109 + 7.

// Here, + denotes string concatenation.

// Example 1:
// Input: n = 4
// Output: 6
// Explanation:
// The good strings are "aaab", "abbb", "baaa", "bbba", "abab", and "baba".
// For example, "aaab" = "aaa" + "b". Both parts are good because each contains one distinct character and has odd length, and their characters at the boundary are different.
// Also, "ab" = "a" + "b" is good, so "abab" = "ab" + "ab" is good because the boundary characters are different.
// Thus, the answer is 6.

// Example 2:
// Input: n = 3
// Output: 4
// Explanation:
// The good strings are "aaa", "bbb", "aba", and "bab". Thus, the answer is 4.

// Example 3:
// Input: n = 2
// Output: 2
// Explanation:
// The good strings are "ab" and "ba". Thus, the answer is 2.

// Constraints:
//     1 <= n <= 10^15

import "fmt"

const MOD = 1_000_000_007

type Matrix [][]int

func newMatrix(n, m int) Matrix {
    matrix := make(Matrix, n)
    for i := range matrix {
        matrix[i] = make([]int, m)
    }
    return matrix
}

// 返回矩阵 a 和矩阵 b 相乘的结果
func (a Matrix) mul(b Matrix) Matrix {
    c := newMatrix(len(a), len(b[0]))
    for i, row := range a {
        for k, x := range row {
            if x == 0 {
                continue
            }
            for j, y := range b[k] {
                c[i][j] = (c[i][j] + x*y) % MOD
            }
        }
    }
    return c
}

// a^n * f
func (a Matrix) powMul(n int64, f Matrix) Matrix {
    res := f
    for ; n > 0; n /= 2 {
        if n%2 > 0 {
            res = a.mul(res)
        }
        a = a.mul(a)
    }
    return res
}

func countGoodStrings(n int64) int {
    m := Matrix{
        {1, 1},
        {1, 0},
    }
    f1 := Matrix{{2}, {0}} // 这里初始化成 2，就不用把答案乘以 2 了
    fn := m.powMul(n-1, f1)
    return fn[0][0]
}

func countGoodStrings1(n int64) int {
    const mod int64 = 1_000_000_007
    var fib func(int64) (int64, int64)
    fib = func(n int64) (int64, int64) {
        if n == 0 {
            return 0, 1
        }
        a, b := fib(n / 2)
        c := a * ((2*b - a + mod) % mod) % mod
        d := (a * a + b * b) % mod
        if n % 2 == 1 {
            return d, (c + d) % mod
        }
        return c, d
    }
    a, _ := fib(n)
    return int(2 * a % mod)
}

func main() {
    // Example 1:
    // Input: n = 4
    // Output: 6
    // Explanation:
    // The good strings are "aaab", "abbb", "baaa", "bbba", "abab", and "baba".
    // For example, "aaab" = "aaa" + "b". Both parts are good because each contains one distinct character and has odd length, and their characters at the boundary are different.
    // Also, "ab" = "a" + "b" is good, so "abab" = "ab" + "ab" is good because the boundary characters are different.
    // Thus, the answer is 6.
    fmt.Println(countGoodStrings(4)) // 6
    // Example 2:
    // Input: n = 3
    // Output: 4
    // Explanation:
    // The good strings are "aaa", "bbb", "aba", and "bab". Thus, the answer is 4.
    fmt.Println(countGoodStrings(3)) // 4
    // Example 3:
    // Input: n = 2
    // Output: 2
    // Explanation:
    // The good strings are "ab" and "ba". Thus, the answer is 2.
    fmt.Println(countGoodStrings(2)) // 2

    fmt.Println(countGoodStrings(1)) // 188417824
    fmt.Println(countGoodStrings(99)) // 375990357
    fmt.Println(countGoodStrings(100)) // 564408181
    fmt.Println(countGoodStrings(101)) // 564408181
    fmt.Println(countGoodStrings(1024)) // 509709173
    fmt.Println(countGoodStrings(999_999_999_999_999)) // 969940830
    fmt.Println(countGoodStrings(1_000_000_000_000_000)) // 2296650267

    fmt.Println(countGoodStrings1(4)) // 6
    fmt.Println(countGoodStrings1(3)) // 4
    fmt.Println(countGoodStrings1(2)) // 2
    fmt.Println(countGoodStrings1(1)) // 188417824
    fmt.Println(countGoodStrings1(99)) // 375990357
    fmt.Println(countGoodStrings1(100)) // 564408181
    fmt.Println(countGoodStrings1(101)) // 564408181
    fmt.Println(countGoodStrings1(1024)) // 509709173
    fmt.Println(countGoodStrings1(999_999_999_999_999)) // 969940830
    fmt.Println(countGoodStrings1(1_000_000_000_000_000)) // 2296650267
}
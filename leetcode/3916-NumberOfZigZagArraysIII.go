package main

// 3916. Number of ZigZag Arrays III
// You are given three integers n, l, and r.

// A ZigZag array of length n is defined as follows:
//     1. Each element lies in the range [l, r].
//     2. No two adjacent elements are equal.
//     3. No three consecutive elements form a strictly increasing or strictly decreasing sequence.

// Return the total number of valid ZigZag arrays.

// Since the answer may be large, return it modulo 10^9 + 7.

// Example 1:
// Input: n = 3, l = 4, r = 5
// Output: 2
// Explanation:
// There are only 2 valid ZigZag arrays of length n = 3 using values in the range [4, 5]:
// [4, 5, 4]
// [5, 4, 5]

// Example 2:
// Input: n = 3, l = 1, r = 3
// Output: 10
// Explanation:
// There are 10 valid ZigZag arrays of length n = 3 using values in the range [1, 3]:
// [1, 2, 1], [1, 3, 1], [1, 3, 2]
// [2, 1, 2], [2, 1, 3], [2, 3, 1], [2, 3, 2]
// [3, 1, 2], [3, 1, 3], [3, 2, 3]
// All arrays meet the ZigZag conditions.

// Constraints:
//     3 <= n <= 200
//     1 <= l < r <= 10​​​​​​^​9

import "fmt"

// 拉格朗日插值
func zigZagArrays(n, l, r int) int {
    k, mod := r - l + 1, 1_000_000_007
    dp := func(n, k int) int {
        res, f := 0, make([]int, k)
        for i := range f {
            f[i] = 1
        }
        for i := 1; i < n; i++ {
            if i % 2 > 0 { // 增
                pre := 0
                for j, v := range f {
                    f[j] = pre % mod
                    pre += v
                }
            } else { // 减
                suf := 0
                for j := k - 1; j >= 0; j-- {
                    v := f[j]
                    f[j] = suf % mod
                    suf += v
                }
            }
        }
        for _, v := range f {
            res += v
        }
        return res * 2 % mod
    }
    if k <= n + 1 {
        return dp(n, k)
    }
    pow := func(x, n int) int {
        res := 1
        for ; n > 0; n /= 2 {
            if n % 2 > 0 {
                res = res * x % mod
            }
            x = x * x % mod
        }
        return res
    }
    // 传入 n 次多项式 f(x) 上的 n+1 个点，返回 f(k)
    lagrange := func(xs, ys []int, k int) int {
        res := 0
        for i, xi := range xs {
            a, b := 1, 1
            for j, xj := range xs {
                if j != i {
                    a = a * (k - xj) % mod  // 分子
                    b = b * (xi - xj) % mod // 分母
                }
            }
            res += a * pow(b, mod-2) % mod * ys[i] % mod
        }
        res = (res % mod + mod) % mod
        return res
    }
    xs, ys := make([]int, n + 1), make([]int, n + 1)
    for i := range xs {
        xs[i] = i + 1
        ys[i] = dp(n, i+1)
    }
    return lagrange(xs, ys, k)
}

func main() {
    // Example 1:
    // Input: n = 3, l = 4, r = 5
    // Output: 2
    // Explanation:
    // There are only 2 valid ZigZag arrays of length n = 3 using values in the range [4, 5]:
    // [4, 5, 4]
    // [5, 4, 5]
    fmt.Println(zigZagArrays(3, 4, 5)) // 2
    // Example 2:
    // Input: n = 3, l = 1, r = 3
    // Output: 10
    // Explanation:
    // There are 10 valid ZigZag arrays of length n = 3 using values in the range [1, 3]:
    // [1, 2, 1], [1, 3, 1], [1, 3, 2]
    // [2, 1, 2], [2, 1, 3], [2, 3, 1], [2, 3, 2]
    // [3, 1, 2], [3, 1, 3], [3, 2, 3]
    // All arrays meet the ZigZag conditions.
    fmt.Println(zigZagArrays(3, 1, 3)) // 10

    fmt.Println(zigZagArrays(3, 1, 2)) // 2
    fmt.Println(zigZagArrays(3, 1, 1_000_000_000)) // 999999727
    fmt.Println(zigZagArrays(3, 999_999_999, 1_000_000_000)) // 2
    fmt.Println(zigZagArrays(200, 1, 2)) // 2
    fmt.Println(zigZagArrays(200, 1, 1_000_000_000)) // 122743172
    fmt.Println(zigZagArrays(200, 999_999_999, 1_000_000_000)) // 2
}
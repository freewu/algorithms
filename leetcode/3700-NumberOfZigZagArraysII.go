package main

// 3700. Number of ZigZag Arrays II
// You are given three integers n, l, and r.

// A ZigZag array of length n is defined as follows:
//     1. Each element lies in the range [l, r].
//     2. No two adjacent elements are equal.
//     3. No three consecutive elements form a strictly increasing or strictly decreasing sequence.

// Return the total number of valid ZigZag arrays.

// Since the answer may be large, return it modulo 10^9 + 7.

// A sequence is said to be strictly increasing if each element is strictly greater than its previous one (if exists).

// A sequence is said to be strictly decreasing if each element is strictly smaller than its previous one (if exists).

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
// ​​​​​​​There are 10 valid ZigZag arrays of length n = 3 using values in the range [1, 3]:
// [1, 2, 1], [1, 3, 1], [1, 3, 2]
// [2, 1, 2], [2, 1, 3], [2, 3, 1], [2, 3, 2]
// [3, 1, 2], [3, 1, 3], [3, 2, 3]
// All arrays meet the ZigZag conditions.

// Constraints:
//     3 <= n <= 10^9
//     1 <= l < r <= 75​​​​​​​

import "fmt"

const mod = 1_000_000_007

type Matrix [][]int

func newMatrix(n, m int) Matrix {
    arr := make(Matrix, n)
    for i := range arr {
        arr[i] = make([]int, m)
    }
    return arr
}

// 返回 a*b
func (a Matrix) mul(b Matrix) Matrix {
    c := newMatrix(len(a), len(b[0]))
    for i, row := range a {
        for k, x := range row {
            if x == 0 { continue }
            for j, y := range b[k] {
                c[i][j] = (c[i][j] + x*y) % mod
            }
        }
    }
    return c
}

// 返回 a^n * f1
func (a Matrix) powMul(n int, f1 Matrix) Matrix {
    res := f1
    for ; n > 0; n /= 2 {
        if n % 2 > 0 {
            res = a.mul(res)
        }
        a = a.mul(a)
    }
    return res
}

func zigZagArrays(n, l, r int) int {
    res, k := 0, r - l + 1
    matrix := newMatrix(k, k)
    for i := 0; i < k; i++ {
        for j := 0; j < k - 1 - i; j++ {
            matrix[i][j] = 1
        }
    }
    f1 := newMatrix(k, 1)
    for i := range f1 {
        f1[i][0] = 1
    }
    fn := matrix.powMul(n - 1, f1)
    for _, row := range fn {
        res += row[0]
    }
    return res * 2 % mod
}

func main() {
    // Example 1:
    // Input: n = 3, l = 4, r = 5
    // Output: 2
    // Explanation:
    // There are only 2 valid ZigZag arrays of length n = 3 using values in the range [4, 5]:
    // [4, 5, 4]
    // [5, 4, 5]​​​​​​​
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
    fmt.Println(zigZagArrays(1_000_000_000, 1, 2)) // 2
    fmt.Println(zigZagArrays(1_000_000_000, 1, 75)) // 16
    fmt.Println(zigZagArrays(3, 74, 75)) // 2
}
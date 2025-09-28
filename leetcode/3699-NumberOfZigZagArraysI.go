package main

// 3699. Number of ZigZag Arrays I
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
// [5, 4, 5]​​​​​​​

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
//     3 <= n <= 2000
//     1 <= l < r <= 2000

import "fmt"

func zigZagArrays(n int, l int, r int) int {
    res, mod, k := 0, 1_000_000_007, r - l + 1
    f := make([]int, k)
    for i := range f { // fill
        f[i] = 1
    }
    for i := 1; i < n; i++ {
        nf := make([]int, k)
        pre := 0
        for j, v := range f {
            nf[k - j - 1] = pre % mod
            pre += v
        }
        f = nf
    }
    for _, v := range f {
        res += v
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
    fmt.Println(zigZagArrays(2000, 1, 2)) // 2
    fmt.Println(zigZagArrays(2000, 1, 2000)) // 594850306
    fmt.Println(zigZagArrays(3, 1999, 2000)) // 2
}
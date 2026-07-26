package main

// 4002. Count Valid Sequences
// You are given two positive integers n and k.

// A valid sequence is a sequence of k positive integers such that:
//     1. The sum of all integers in the sequence is equal to n.
//     2. The product of all integers in the sequence is even.

// Return the number of valid sequences. Since the answer may be very large, return it modulo 10^9​​​​​​​ + 7.

// Two sequences are considered different if they differ at any index. 
// For example, [1, 1, 2] and [1, 2, 1] are considered different sequences.

// Example 1:
// Input: n = 5, k = 3
// Output: 3
// Explanation:
// The sequences of length k = 3 whose sum is 5 are:
// Sequence  | Product        | Parity
// [1, 1, 3] | 1 * 1 * 3 = 3  | Odd
// [1, 2, 2] | 1 * 2 * 2 = 4  | Even
// [2, 1, 2] | 2 * 1 * 2 = 4  | Even
// [2, 2, 1] | 2 * 2 * 1 = 4  | Even
// [1, 3, 1] | 1 * 3 * 1 = 3  | Odd
// [3, 1, 1] | 3 * 1 * 1 = 3  | Odd
// There are 3 sequences with an even product, thus the answer is 3.

// Example 2:
// Input: n = 3, k = 2
// Output: 2
// Explanation:
// The sequences of length k = 2 whose sum is 3 are:
// Sequence  | Product    | Parity
// [1, 2]    | 1 * 2 = 2  | Even
// [2, 1]    | 2 * 1 = 2  | Even   
// There are 2 sequences with an even product, thus the answer is 2.

// Example 3:
// Input: n = 5, k = 5
// Output: 0
// Explanation:
// The only possible sequence of length k = 5 whose sum is 5 is [1, 1, 1, 1, 1], which has an odd product. Thus, the answer is 0.

// Constraints:
//     1 <= n <= 5 * 10^5
//     1 <= k <= n

import "fmt"

const MOD = 1_000_000_007
const MX = 500_000

var fac [MX]int  // fac[i] = i!
var invF [MX]int // invF[i] = i!^-1 = pow(i!, mod-2)

func init() {
    fac[0] = 1
    for i := 1; i < MX; i++ {
        fac[i] = fac[i-1] * i % MOD
    }
    invF[MX-1] = pow(fac[MX-1], MOD-2)
    for i := MX - 1; i > 0; i-- {
        invF[i-1] = invF[i] * i % MOD
    }
}

func pow(x, n int) int {
    res := 1
    for ; n > 0; n /= 2 {
        if n % 2 > 0 {
            res = res * x % MOD
        }
        x = x * x % MOD
    }
    return res
}

// 从 n 个数中选 m 个数的方案数
func comb(n, m int) int {
    return fac[n] * invF[m] % MOD * invF[n-m] % MOD
}

func countValidSequences(n, k int) int {
    res := comb(n-1, k-1)
    if (n + k) % 2 == 0 {
        res = (res - comb((n + k) / 2-1, k - 1) + MOD) % MOD // +MOD 保证答案非负 
    }
    return res  
}

func main() {
    // Example 1:
    // Input: n = 5, k = 3
    // Output: 3
    // Explanation:
    // The sequences of length k = 3 whose sum is 5 are:
    // Sequence  | Product        | Parity
    // [1, 1, 3] | 1 * 1 * 3 = 3  | Odd
    // [1, 2, 2] | 1 * 2 * 2 = 4  | Even
    // [2, 1, 2] | 2 * 1 * 2 = 4  | Even
    // [2, 2, 1] | 2 * 2 * 1 = 4  | Even
    // [1, 3, 1] | 1 * 3 * 1 = 3  | Odd
    // [3, 1, 1] | 3 * 1 * 1 = 3  | Odd
    // There are 3 sequences with an even product, thus the answer is 3.
    fmt.Println(countValidSequences(5, 3)) // 3
    // Example 2:
    // Input: n = 3, k = 2
    // Output: 2
    // Explanation:
    // The sequences of length k = 2 whose sum is 3 are:
    // Sequence  | Product    | Parity
    // [1, 2]    | 1 * 2 = 2  | Even
    // [2, 1]    | 2 * 1 = 2  | Even   
    // There are 2 sequences with an even product, thus the answer is 2.
    fmt.Println(countValidSequences(3, 2)) // 2
    // Example 3:
    // Input: n = 5, k = 5
    // Output: 0
    // Explanation:
    // The only possible sequence of length k = 5 whose sum is 5 is [1, 1, 1, 1, 1], which has an odd product. Thus, the answer is 0.
    fmt.Println(countValidSequences(5, 5)) // 0

    fmt.Println(countValidSequences(1, 1)) // 0
    fmt.Println(countValidSequences(500_000, 1)) // 1
    fmt.Println(countValidSequences(500_000, 500_000)) // 0
}
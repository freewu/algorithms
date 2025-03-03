package main

// 3405. Count the Number of Arrays with K Matching Adjacent Elements
// You are given three integers n, m, k. A good array arr of size n is defined as follows:
//     1. Each element in arr is in the inclusive range [1, m].
//     2. Exactly k indices i (where 1 <= i < n) satisfy the condition arr[i - 1] == arr[i].

// Return the number of good arrays that can be formed.

// Since the answer may be very large, return it modulo 10^9 + 7.

// Example 1:
// Input: n = 3, m = 2, k = 1
// Output: 4
// Explanation:
// There are 4 good arrays. They are [1, 1, 2], [1, 2, 2], [2, 1, 1] and [2, 2, 1].
// Hence, the answer is 4.

// Example 2:
// Input: n = 4, m = 2, k = 2
// Output: 6
// Explanation:
// The good arrays are [1, 1, 1, 2], [1, 1, 2, 2], [1, 2, 2, 2], [2, 1, 1, 1], [2, 2, 1, 1] and [2, 2, 2, 1].
// Hence, the answer is 6.

// Example 3:
// Input: n = 5, m = 2, k = 0
// Output: 2
// Explanation:
// The good arrays are [1, 2, 1, 2, 1] and [2, 1, 2, 1, 2]. Hence, the answer is 2.

// Constraints:
//     1 <= n <= 10^5
//     1 <= m <= 10^5
//     0 <= k <= n - 1

import "fmt"

const mod = 1_000_000_007
const mx = 100_000

var F, invF [mx]int

func init() {
    F[0] = 1
    for i := 1; i < mx; i++ {
        F[i] = F[i-1] * i % mod
    }
    invF[mx-1] = pow(F[mx-1], mod-2)
    for i := mx - 1; i > 0; i-- {
        invF[i-1] = invF[i] * i % mod
    }
}

func pow(x, n int) int {
    res := 1
    for ; n > 0; n /= 2 {
        if n % 2 > 0 {
            res = res * x % mod
        }
        x = x * x % mod
    }
    return res
}

func countGoodArrays(n, m, k int) int {
    comb := func (n, m int) int { return F[n] * invF[m] % mod * invF[n-m] % mod }
    return comb(n-1, k) * m % mod * pow(m-1, n-k-1) % mod
}

func countGoodArrays1(n, m, k int) int {
    mod := 1_000_000_007
    pow := func(base, exp, mod int) int {
        res := 1
        for exp > 0 {
            if exp % 2 == 1 {
                res = res * base % mod
            }
            base = base * base % mod
            exp /= 2
        }
        return res
    }
    comb := func(n, r, mod int) int {
        if r > n { return 0 }
        if r == 0 || r == n { return 1 }
        // Use smaller r to optimize calculation
        if r > n / 2 {
            r = n - r
        }
        num, den := 1, 1
        for i := 0; i < r; i++ {
            num = (num * (n - i)) % mod
            den = (den * (i + 1)) % mod
        }
        denInv := pow(den, mod - 2, mod)
        return (num * denInv) % mod
    }
    res := m * pow(m - 1, n - k - 1, mod) % mod
    res = res * comb(n - 1, k, mod) % mod
    return res
}

func main() {
    // Example 1:
    // Input: n = 3, m = 2, k = 1
    // Output: 4
    // Explanation:
    // There are 4 good arrays. They are [1, 1, 2], [1, 2, 2], [2, 1, 1] and [2, 2, 1].
    // Hence, the answer is 4.
    fmt.Println(countGoodArrays(3, 2, 1)) // 4
    // Example 2:
    // Input: n = 4, m = 2, k = 2
    // Output: 6
    // Explanation:
    // The good arrays are [1, 1, 1, 2], [1, 1, 2, 2], [1, 2, 2, 2], [2, 1, 1, 1], [2, 2, 1, 1] and [2, 2, 2, 1].
    // Hence, the answer is 6.
    fmt.Println(countGoodArrays(4, 2, 2)) // 6
    // Example 3:
    // Input: n = 5, m = 2, k = 0
    // Output: 2
    // Explanation:
    // The good arrays are [1, 2, 1, 2, 1] and [2, 1, 2, 1, 2]. Hence, the answer is 2.
    fmt.Println(countGoodArrays(5, 2, 0)) // 2

    fmt.Println(countGoodArrays(1, 1, 0)) // 1
    fmt.Println(countGoodArrays(100000, 100000, 99999)) // 100000
    fmt.Println(countGoodArrays(100000, 1, 99999)) // 1
    fmt.Println(countGoodArrays(1, 100000, 0)) // 100000
    fmt.Println(countGoodArrays(100000, 100000, 0)) // 406242317

    fmt.Println(countGoodArrays1(3, 2, 1)) // 4
    fmt.Println(countGoodArrays1(4, 2, 2)) // 6
    fmt.Println(countGoodArrays1(5, 2, 0)) // 2
    fmt.Println(countGoodArrays1(1, 1, 0)) // 1
    fmt.Println(countGoodArrays1(100000, 100000, 99999)) // 100000
    fmt.Println(countGoodArrays1(100000, 1, 99999)) // 1
    fmt.Println(countGoodArrays1(1, 100000, 0)) // 100000
    fmt.Println(countGoodArrays1(100000, 100000, 0)) // 406242317
}
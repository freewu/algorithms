package main

// 3343. Count Number of Balanced Permutations
// You are given a string num. A string of digits is called balanced if the sum of the digits at even indices is equal to the sum of the digits at odd indices.

// Return the number of distinct permutations of num that are balanced.

// Since the answer may be very large, return it modulo 10^9 + 7.

// A permutation is a rearrangement of all the characters of a string.

// Example 1:
// Input: num = "123"
// Output: 2
// Explanation:
// The distinct permutations of num are "123", "132", "213", "231", "312" and "321".
// Among them, "132" and "231" are balanced. Thus, the answer is 2.

// Example 2:
// Input: num = "112"
// Output: 1
// Explanation:
// The distinct permutations of num are "112", "121", and "211".
// Only "121" is balanced. Thus, the answer is 1.

// Example 3:
// Input: num = "12345"
// Output: 0
// Explanation:
// None of the permutations of num are balanced, so the answer is 0.

// Constraints:
//     2 <= num.length <= 80
//     num consists of digits '0' to '9' only.

import "fmt"

// const MX  = 80
// const MOD = 1_000_000_007

// var c [MX][MX]int

// func init() {
//     c[0][0] = 1
//     for i := 1; i < MX; i++ {
//         c[i][0] = 1
//         for j := 1; j <= i; j++ {
//             c[i][j] = (c[i-1][j] + c[i-1][j-1]) % MOD
//         }
//     }
// }

// func countBalancedPermutations(num string) int {
//     count, sum := [10]int{}, 0
//     for _, v := range num {
//         count[v - '0']++
//         sum += int(v - '0')
//     }
//     if sum % 2 != 0 { return 0 }
//     n := len(num)
//     m := n / 2 + 1
//     dp := make([][][][]int, 10)
//     for i := range dp {
//         dp[i] = make([][][]int, sum / 2 + 1)
//         for j := range dp[i] {
//             dp[i][j] = make([][]int, m)
//             for k := range dp[i][j] {
//                 dp[i][j][k] = make([]int, m + 1)
//                 for l := range dp[i][j][k] {
//                     dp[i][j][k][l] = -1
//                 }
//             }
//         }
//     }
//     var dfs func(i, j, a, b int) int
//     dfs = func(i, j, a, b int) int {
//         if i > 9 {
//             if j == 0 && a == 0 && b == 0 {
//                 return 1
//             }
//             return 0
//         }
//         if a == 0 && j > 0 {
//             return 0
//         }
//         if dp[i][j][a][b] != -1 {
//             return dp[i][j][a][b]
//         }
//         res := 0
//         for l := 0; l <= min(count[i], a); l++ {
//             r := count[i] - l
//             if r >= 0 && r <= b && l*i <= j {
//                 t := c[a][l] * c[b][r] % MOD * dfs(i + 1, j - l * i, a - l, b - r) % MOD
//                 res = (res + t) % MOD
//             }
//         }
//         dp[i][j][a][b] = res
//         return res
//     }
//     return dfs(0, sum / 2, n / 2, (n + 1) / 2)
// }

const mod = 1_000_000_007
const mx = 40

var fac, invF [mx + 1]int

func init() {
    fac[0] = 1
    for i := 1; i <= mx; i++ {
        fac[i] = fac[i-1] * i % mod
    }
    invF[mx] = pow(fac[mx], mod-2)
    for i := mx; i > 0; i-- {
        invF[i-1] = invF[i] * i % mod
    }
}

func pow(x, n int) int {
    res := 1
    for ; n > 0; n /= 2 {
        if n%2 > 0 {
            res = res * x % mod
        }
        x = x * x % mod
    }
    return res
}

func countBalancedPermutations(num string) int {
    count := [10]int{}
    sum, n := 0, len(num)
    for _, c := range num {
        count[c-'0']++
        sum += int(c - '0')
    }
    if sum % 2 > 0 { return 0 }
    m := n / 2
    f := make([][]int, m + 1)
    for i := range f {
        f[i] = make([]int, sum / 2 + 1)
    }
    f[0][0] = 1
    sc, s := 0, 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, c := range count {
        sc += c
        s += c * i
        // 保证 left <= n - m，即 j >= sc - (n - m)
        for j := min(sc, m); j >= max(sc - (n - m), 0); j-- {
            left := sc - j
            // 保证分给第二个集合的元素和 <= sum / 2，即 p >= s - sum / 2
            for p := min(s, sum / 2); p >= max(s - sum / 2, 0); p-- {
                v := 0
                for k := max(c - left, 0); k <= min(c, j) && k*i <= p; k++ {
                    v = (v + f[j-k][p - k * i] * invF[k] % mod * invF[c - k]) % mod
                }
                f[j][p] = v
            }
        }
    }
    return fac[m] * fac[n - m] % mod * f[m][sum / 2] % mod
}

func main() {
    // Example 1:
    // Input: num = "123"
    // Output: 2
    // Explanation:
    // The distinct permutations of num are "123", "132", "213", "231", "312" and "321".
    // Among them, "132" and "231" are balanced. Thus, the answer is 2.
    fmt.Println(countBalancedPermutations("123")) // 2
    // Example 2:
    // Input: num = "112"
    // Output: 1
    // Explanation:
    // The distinct permutations of num are "112", "121", and "211".
    // Only "121" is balanced. Thus, the answer is 1.
    fmt.Println(countBalancedPermutations("112")) // 1
    // Example 3:
    // Input: num = "12345"
    // Output: 0
    // Explanation:
    // None of the permutations of num are balanced, so the answer is 0.
    fmt.Println(countBalancedPermutations("12345")) // 0

    fmt.Println(countBalancedPermutations("12345678")) // 4608
    fmt.Println(countBalancedPermutations("98765432")) // 4608
}
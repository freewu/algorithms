package main

// 3539. Find Sum of Array Product of Magical Sequences
// You are given two integers, m and k, and an integer array nums.
//     1. A sequence of integers seq is called magical if:
//     2. seq has a size of m.
//     3. 0 <= seq[i] < nums.length

// The binary representation of 2seq[0] + 2seq[1] + ... + 2seq[m - 1] has k set bits.

// The array product of this sequence is defined as prod(seq) = (nums[seq[0]] * nums[seq[1]] * ... * nums[seq[m - 1]]).

// Return the sum of the array products for all valid magical sequences.

// Since the answer may be large, return it modulo 10^9 + 7.

// A set bit refers to a bit in the binary representation of a number that has a value of 1.

// Example 1:
// Input: m = 5, k = 5, nums = [1,10,100,10000,1000000]
// Output: 991600007
// Explanation:
// All permutations of [0, 1, 2, 3, 4] are magical sequences, each with an array product of 1013.

// Example 2:
// Input: m = 2, k = 2, nums = [5,4,3,2,1]
// Output: 170
// Explanation:
// The magical sequences are [0, 1], [0, 2], [0, 3], [0, 4], [1, 0], [1, 2], [1, 3], [1, 4], [2, 0], [2, 1], [2, 3], [2, 4], [3, 0], [3, 1], [3, 2], [3, 4], [4, 0], [4, 1], [4, 2], and [4, 3].

// Example 3:
// Input: m = 1, k = 1, nums = [28]
// Output: 28
// Explanation:
// The only magical sequence is [0].

// Constraints:
//     1 <= k <= m <= 30
//     1 <= nums.length <= 50
//     1 <= nums[i] <= 10^8

import "fmt"
import "math/bits"

// func magicalSum(m int, k int, nums []int) int {
//     n, mod := len(nums), 1_000_000_007
//     f, inverse_f := make([]int, m + 1), make([]int, m + 1)
//     f[0] = 1
//     for i := 1; i <= m; i++ {
//         f[i] = f[i - 1] * i % mod
//     }
//     modPow := func(base, exp, mod int) int {
//         res := 1
//         base %= mod
//         if base == 0 { return 0 }
//         for ; exp > 0; exp >>= 1 {
//             if (exp & 1) == 1 { res = (res * base) % mod }
//             base = (base * base) % mod
//         }
//         return res
//     }
//     inverse_f[m] = modPow(f[m], mod - 2, mod)
//     for i := m; i >= 1; i-- {
//         inverse_f[i - 1] = inverse_f[i] * i % mod
//     }
//     pow_nums := make([][]int, n)
//     for i := 0; i < n; i++ {
//         pow_nums[i] = make([]int, m + 1)
//         pow_nums[i][0] = 1
//         for j := 1; j <= m; j++ {
//             pow_nums[i][j] = pow_nums[i][j - 1] * nums[i] % mod
//         }
//     }
//     dp := make([][][][]int, n + 1)
//     for i := range dp {
//         dp[i] = make([][][]int, m + 1)
//         for j := range dp[i] {
//             dp[i][j] = make([][]int, k + 1)
//             for g := range dp[i][j] {
//                 dp[i][j][g] = make([]int, n + 1)
//             }
//         }
//     }
//     dp[0][0][0][0] = 1
//     for i := 0; i < n; i++ {
//         for m1 := 0; m1 <= m; m1++ {
//             for k1 := 0; k1 <= k; k1++ {
//                 for m2 := 0; m2 <= m; m2++ {
//                     val := dp[i][m1][k1][m2]
//                     if val == 0 { continue }
//                     for c := 0; c <= m - m1; c++ {
//                         m12 := m1 + c
//                         s := c + m2
//                         bit := s & 1
//                         k2 := k1 + bit
//                         if k2 > k { continue} 
//                         m22 := s >> 1
//                         dp[i + 1][m12][k2][m22] = (dp[i + 1][m12][k2][m22] + val * inverse_f[c] % mod * pow_nums[i][c] % mod) % mod
//                     }
//                 }
//             }
//         }
//     }
//     res := 0
//     for k1 := 0; k1 <= k; k1++ {
//         for m2 := 0; m2 <= m; m2++ {
//             val := dp[n][m][k1][m2]
//             if val == 0 { continue }
//             bits := bits.OnesCount32(uint32(m2))
//             if (k1 + bits) == k {
//                 res = (res + val) % mod
//             }
//         }
//     }
//     res = res * f[m] % mod
//     return res
// }

func magicalSum(m int, k int, nums []int) int {
    n, mod := len(nums), 1_000_000_007
    powMod := func(base, exp, mod int) int {
        res := 1
        for exp > 0 {
            if exp&1 == 1 {
                res = res * base % mod
            }
            base = base * base % mod
            exp >>= 1
        }
        return res
    }
    bitCount := func(x int) int {
        count := 0
        for x != 0 {
            x &= x - 1
            count++
        }
        return count
    }
    f := make([]int, m + 1)
    inverseF := make([]int, m + 1)
    f[0] = 1
    inverseF[0] = 1
    for i := 1; i <= m; i++ {
        f[i] = f[i-1] * i % mod
    }
    inverseF[m] = powMod(f[m], mod - 2, mod)
    for i := m; i > 0; i-- {
        inverseF[i-1] = inverseF[i] * i % mod
    }
    powNums := make([][]int, n)
    for i := range powNums {
        powNums[i] = make([]int, m + 1)
        powNums[i][0] = 1
        for c := 1; c <= m; c++ {
            powNums[i][c] = powNums[i][c-1] * nums[i] % mod
        }
    }
    dp := make([][][][]int, n + 1)
    for i := range dp {
        dp[i] = make([][][]int, m + 1)
        for m1 := range dp[i] {
            dp[i][m1] = make([][]int, k + 1)
            for k1 := range dp[i][m1] {
                dp[i][m1][k1] = make([]int, m + 1)
            }
        }
    }
    dp[0][0][0][0] = 1
    for i := 0; i < n; i++ {
        for m1 := 0; m1 <= m; m1++ {
            for k1 := 0; k1 <= k; k1++ {
                for m2 := 0; m2 <= m; m2++ {
                    val := dp[i][m1][k1][m2]
                    if val == 0 { continue }
                    for c := 0; c <= m-m1; c++ {
                        m12 := m1 + c
                        s := c + m2
                        bit := s & 1
                        k2 := k1 + bit
                        if k2 > k { continue }
                        m22 := s >> 1
                        dp[i+1][m12][k2][m22] += val * inverseF[c] % mod * powNums[i][c] % mod
                        dp[i+1][m12][k2][m22] %= mod
                    }
                }
            }
        }
    }
    res := 0
    for k1 := 0; k1 <= k; k1++ {
        for m2 := 0; m2 <= m; m2++ {
            val := dp[n][m][k1][m2]
            if val == 0 { continue }
            bits := bitCount(m2)
            if k1 + bits == k {
                res += val
                res %= mod
            }
        }
    }
    res = res * f[m] % mod
    return res
}


const mod = 1_000_000_007
const mx = 31

var fac [mx]int  // fac[i] = i!
var invF [mx]int // invF[i] = i!^-1

func init() {
    fac[0] = 1
    for i := 1; i < mx; i++ {
        fac[i] = fac[i-1] * i % mod
    }

    invF[mx-1] = pow(fac[mx-1], mod-2)
    for i := mx - 1; i > 0; i-- {
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

func magicalSum1(m, k int, nums []int) int {
    n := len(nums)
    powV := make([][]int, n)
    for i, v := range nums {
        powV[i] = make([]int, m+1)
        powV[i][0] = 1
        for j := 1; j <= m; j++ {
            powV[i][j] = powV[i][j-1] * v % mod
        }
    }

    memo := make([][][][]int, n)
    for i := range memo {
        memo[i] = make([][][]int, m+1)
        for j := range memo[i] {
            memo[i][j] = make([][]int, m/2+1)
            for p := range memo[i][j] {
                memo[i][j][p] = make([]int, k+1)
                for q := range memo[i][j][p] {
                    memo[i][j][p][q] = -1
                }
            }
        }
    }
    var dfs func(int, int, int, int) int
    dfs = func(i, leftM, x, leftK int) (res int) {
        c1 := bits.OnesCount(uint(x))
        if c1+leftM < leftK { // 可行性剪枝
            return
        }
        if i == n || leftM == 0 || leftK == 0 { // 无法继续选数字
            if leftM == 0 && c1 == leftK {
                return 1
            }
            return
        }
        p := &memo[i][leftM][x][leftK]
        if *p != -1 {
            return *p
        }
        for j := range leftM + 1 { // 枚举 I 中有 j 个下标 i
            // 这 j 个下标 i 对 S 的贡献是 j * pow(2, i)
            // 由于 x = S >> i，转化成对 x 的贡献是 j
            bit := (x + j) & 1 // 取最低位，提前从 leftK 中减去，其余进位到 x 中
            r := dfs(i+1, leftM-j, (x+j)>>1, leftK-bit)
            res = (res + r*powV[i][j]%mod*invF[j]) % mod
        }
        *p = res
        return
    }
    return dfs(0, m, 0, k) * fac[m] % mod
}

func main() {
    // Example 1:
    // Input: m = 5, k = 5, nums = [1,10,100,10000,1000000]
    // Output: 991600007
    // Explanation:
    // All permutations of [0, 1, 2, 3, 4] are magical sequences, each with an array product of 1013.
    fmt.Println(magicalSum(5, 5, []int{1,10,100,10000,1000000})) // 991600007
    // Example 2:
    // Input: m = 2, k = 2, nums = [5,4,3,2,1]
    // Output: 170
    // Explanation:
    // The magical sequences are [0, 1], [0, 2], [0, 3], [0, 4], [1, 0], [1, 2], [1, 3], [1, 4], [2, 0], [2, 1], [2, 3], [2, 4], [3, 0], [3, 1], [3, 2], [3, 4], [4, 0], [4, 1], [4, 2], and [4, 3].
    fmt.Println(magicalSum(2, 2, []int{5,4,3,2,1})) // 170
    // Example 3:
    // Input: m = 1, k = 1, nums = [28]
    // Output: 28
    // Explanation:
    // The only magical sequence is [0].
    fmt.Println(magicalSum(1, 1, []int{28})) // 28

    fmt.Println(magicalSum(2, 1, []int{63})) // 3969
    fmt.Println(magicalSum(2, 2, []int{9,8,7,6,5,4,3,2,1})) // 1740
    fmt.Println(magicalSum(2, 2, []int{1,2,3,4,5,6,7,8,9})) // 1740

    fmt.Println(magicalSum1(5, 5, []int{1,10,100,10000,1000000})) // 991600007
    fmt.Println(magicalSum1(2, 2, []int{5,4,3,2,1})) // 170
    fmt.Println(magicalSum1(1, 1, []int{28})) // 28
    fmt.Println(magicalSum1(2, 1, []int{63})) // 3969
    fmt.Println(magicalSum1(2, 2, []int{9,8,7,6,5,4,3,2,1})) // 1740
    fmt.Println(magicalSum1(2, 2, []int{1,2,3,4,5,6,7,8,9})) // 1740
}


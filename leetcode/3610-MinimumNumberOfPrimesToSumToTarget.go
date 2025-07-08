package main

// 3610. Minimum Number of Primes to Sum to Target
// You are given two integers n and m.

// You have to select a multiset of prime numbers from the first m prime numbers such that the sum of the selected primes is exactly n. 
// You may use each prime number multiple times.

// Return the minimum number of prime numbers needed to sum up to n, or -1 if it is not possible.

// Example 1:
// Input: n = 10, m = 2
// Output: 4
// Explanation:
// The first 2 primes are [2, 3]. The sum 10 can be formed as 2 + 2 + 3 + 3, requiring 4 primes.

// Example 2:
// Input: n = 15, m = 5
// Output: 3
// Explanation:
// The first 5 primes are [2, 3, 5, 7, 11]. The sum 15 can be formed as 5 + 5 + 5, requiring 3 primes.

// Example 3:
// Input: n = 7, m = 6
// Output: 1
// Explanation:
// The first 6 primes are [2, 3, 5, 7, 11, 13]. The sum 7 can be formed directly by prime 7, requiring only 1 prime.

// Constraints:
//     1 <= n <= 1000
//     1 <= m <= 1000

import "fmt"

func minNumberOfPrimes(n int, m int) int {
    // 生成前 m 个质数
    getFirstMPrimes := func(m int) []int {
        primes := make([]int, 0, m)
        candidate := 2
        for len(primes) < m {
            isPrime := true
            // 检查是否能被已找到的质数整除（只需检查到 sqrt(candidate)）
            for _, p := range primes {
                if p*p > candidate {
                    break
                }
                if candidate%p == 0 {
                    isPrime = false
                    break
                }
            }
            if isPrime {
                primes = append(primes, candidate)
            }
            candidate++
        }
        return primes
    }
    // 生成前 m 个质数
    primes := getFirstMPrimes(m)
    // 初始化 DP 数组，dp[i] 表示和为 i 所需的最少质数数量
    inf := 1 << 61
    dp := make([]int, n + 1)
    for i := range dp {
        dp[i] = inf
    }
    dp[0] = 0 // 基准情况：和为 0 需要 0 个质数
    min := func (x, y int) int { if x < y { return x; }; return y; }
    // 动态规划更新
    for _, p := range primes {
        for i := p; i <= n; i++ {
            if dp[i-p] != inf {
                dp[i] = min(dp[i], dp[i-p]+1)
            }
        }
    }
    if dp[n] == inf { return -1 }
    return dp[n]
}

var primes []int

func init() {
    x, m := 2, 1000
    for len(primes) < m {
        is_prime := true
        for _, p := range primes {
            if p*p > x {
                break
            }
            if x % p == 0 {
                is_prime = false
                break
            }
        }
        if is_prime {
            primes = append(primes, x)
        }
        x++
    }
}

func minNumberOfPrimes1(n int, m int) int {
    inf := 1 << 61
    dp := make([]int, n + 1)
    for i := 1; i <= n; i++ {
        dp[i] = inf
    }
    dp[0] = 0
    for _, v := range primes[:m] {
        for i := v; i <= n; i++ {
            if dp[i - v] < inf && dp[i - v] + 1 < dp[i] {
                dp[i] = dp[i - v] + 1
            }
        }
    }
    if dp[n] < inf { return dp[n] }
    return -1
}

func main() {
    // Example 1:
    // Input: n = 10, m = 2
    // Output: 4
    // Explanation:
    // The first 2 primes are [2, 3]. The sum 10 can be formed as 2 + 2 + 3 + 3, requiring 4 primes.
    fmt.Println(minNumberOfPrimes(10, 2)) // 4
    // Example 2:
    // Input: n = 15, m = 5
    // Output: 3
    // Explanation:
    // The first 5 primes are [2, 3, 5, 7, 11]. The sum 15 can be formed as 5 + 5 + 5, requiring 3 primes.
    fmt.Println(minNumberOfPrimes(15, 5)) // 3
    // Example 3:
    // Input: n = 7, m = 6
    // Output: 1
    // Explanation:
    // The first 6 primes are [2, 3, 5, 7, 11, 13]. The sum 7 can be formed directly by prime 7, requiring only 1 prime.
    fmt.Println(minNumberOfPrimes(7, 6)) // 1

    fmt.Println(minNumberOfPrimes(1, 1)) // -1
    fmt.Println(minNumberOfPrimes(1000, 1000)) // 2
    fmt.Println(minNumberOfPrimes(1, 1000)) // -1
    fmt.Println(minNumberOfPrimes(1000, 1)) // 500

    fmt.Println(minNumberOfPrimes1(10, 2)) // 4
    fmt.Println(minNumberOfPrimes1(15, 5)) // 3
    fmt.Println(minNumberOfPrimes1(7, 6)) // 1
    fmt.Println(minNumberOfPrimes1(1, 1)) // -1
    fmt.Println(minNumberOfPrimes1(1000, 1000)) // 2
    fmt.Println(minNumberOfPrimes1(1, 1000)) // -1
    fmt.Println(minNumberOfPrimes1(1000, 1)) // 500
}
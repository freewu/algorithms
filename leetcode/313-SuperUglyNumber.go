package main

// 313. Super Ugly Number
// A super ugly number is a positive integer whose prime factors are in the array primes.
// Given an integer n and an array of integers primes, return the nth super ugly number.
// The nth super ugly number is guaranteed to fit in a 32-bit signed integer.

// Example 1:
// Input: n = 12, primes = [2,7,13,19]
// Output: 32
// Explanation: [1,2,4,7,8,13,14,16,19,26,28,32] is the sequence of the first 12 super ugly numbers given primes = [2,7,13,19].

// Example 2:
// Input: n = 1, primes = [2,3,5]
// Output: 1
// Explanation: 1 has no prime factors, therefore all of its prime factors are in the array primes = [2,3,5].

// Constraints:
//     1 <= n <= 10^5
//     1 <= primes.length <= 100
//     2 <= primes[i] <= 1000
//     primes[i] is guaranteed to be a prime number.
//     All the values of primes are unique and sorted in ascending order.

import "fmt"

func nthSuperUglyNumber(n int, primes []int) int {
    if n == 1 {
        return 1
    }
    dp, l, arr, inf := make([]int, n + 1), len(primes), make([]int, len(primes)), 1 << 32 - 1
    dp[1] = 1
    for i := 0; i < l; i++ { arr[i] = 1; } // fill 1
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 2; i <= n; i++ {
        mn := inf
        for j := 0; j < l; j++ {
            mn = min(mn, dp[arr[j]] * primes[j])
        }
        dp[i] = mn
        for j := 0; j < l; j++ {
            if dp[i] == dp[arr[j]] * primes[j] {
                arr[j]++
            }
        }
    }
    return dp[n]
}

func nthSuperUglyNumber1(n int, primes []int) int {
    l, inf := len(primes), 1 << 32 - 1
    dp, pointers, nums := make([]int, n + 1), make([]int, l), make([]int, l)
    for i := range nums { nums[i] = 1; } // fill 1
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i <= n; i++ {
        mn := inf
        for j := range pointers {
            mn = min(mn, nums[j])
        }
        dp[i] = mn
        for j := range nums {
            if nums[j] == mn {
                pointers[j]++
                nums[j] = dp[pointers[j]] * primes[j]
            }
        }
    }
    return dp[n]
}

func main() {
    // Example 1:
    // Input: n = 12, primes = [2,7,13,19]
    // Output: 32
    // Explanation: [1,2,4,7,8,13,14,16,19,26,28,32] is the sequence of the first 12 super ugly numbers given primes = [2,7,13,19].
    fmt.Println(nthSuperUglyNumber(12,[]int{2,7,13,19})) // 32
    // Example 2:
    // Input: n = 1, primes = [2,3,5]
    // Output: 1
    // Explanation: 1 has no prime factors, therefore all of its prime factors are in the array primes = [2,3,5].
    fmt.Println(nthSuperUglyNumber(1,[]int{2,3,5})) // 1

    fmt.Println(nthSuperUglyNumber1(12,[]int{2,7,13,19})) // 32
    fmt.Println(nthSuperUglyNumber1(1,[]int{2,3,5})) // 1
}
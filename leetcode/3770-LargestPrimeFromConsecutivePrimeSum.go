package main

// 3770. Largest Prime from Consecutive Prime Sum
// You are given an integer n.

// Return the largest prime number less than or equal to n that can be expressed as the sum of one or more consecutive prime numbers starting from 2. 
// If no such number exists, return 0.

// Example 1:
// Input: n = 20
// Output: 17
// Explanation:
// The prime numbers less than or equal to n = 20 which are consecutive prime sums are:
// 2 = 2
// 5 = 2 + 3
// 17 = 2 + 3 + 5 + 7
// The largest is 17, so it is the answer.

// Example 2:
// Input: n = 2
// Output: 2
// Explanation:
// The only consecutive prime sum less than or equal to 2 is 2 itself.

// Constraints:
//     1 <= n <= 5 * 10^5

import "fmt"
import "sort"

const MX = 500_000

var primes []int
var np [MX + 1]bool
var specialPrimes = []int{0} // 哨兵

func init() {
    for i := 2; i <= MX; i++ {
        if !np[i] {
            primes = append(primes, i)
            for j := i * i; j <= MX; j += i {
                np[j] = true
            }
        }
    }
    sum := 0
    for _, p := range primes {
        sum += p
        if sum > MX { break }
        if !np[sum] {
            specialPrimes = append(specialPrimes, sum)
        }
    }
}

func largestPrime(n int) int {
    // 二分找 <= n 的最大特殊质数
    i := sort.SearchInts(specialPrimes, n+1) - 1
    return specialPrimes[i]
}

func main() {
    // Example 1:
    // Input: n = 20
    // Output: 17
    // Explanation:
    // The prime numbers less than or equal to n = 20 which are consecutive prime sums are:
    // 2 = 2
    // 5 = 2 + 3
    // 17 = 2 + 3 + 5 + 7
    // The largest is 17, so it is the answer.
    fmt.Println(largestPrime(20)) // 17
    // Example 2:
    // Input: n = 2
    // Output: 2
    // Explanation:
    // The only consecutive prime sum less than or equal to 2 is 2 itself.
    fmt.Println(largestPrime(20)) // 17

    fmt.Println(largestPrime(1)) // 0
    fmt.Println(largestPrime(8)) // 5
    fmt.Println(largestPrime(64)) // 41
    fmt.Println(largestPrime(999)) // 281
    fmt.Println(largestPrime(1000)) // 281
    fmt.Println(largestPrime(1024)) // 281
    fmt.Println(largestPrime(500_000)) // 398771
}
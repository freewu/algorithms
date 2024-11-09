package main

// 1808. Maximize Number of Nice Divisors
// You are given a positive integer primeFactors. 
// You are asked to construct a positive integer n that satisfies the following conditions:
//     1. The number of prime factors of n (not necessarily distinct) is at most primeFactors.
//     2. The number of nice divisors of n is maximized. 
//        Note that a divisor of n is nice if it is divisible by every prime factor of n. 
//        For example, if n = 12, then its prime factors are [2,2,3], then 6 and 12 are nice divisors, while 3 and 4 are not.

// Return the number of nice divisors of n. 
// Since that number can be too large, return it modulo 10^9 + 7.

// Note that a prime number is a natural number greater than 1 that is not a product of two smaller natural numbers. 
// The prime factors of a number n is a list of prime numbers such that their product equals n.

// Example 1:
// Input: primeFactors = 5
// Output: 6
// Explanation: 200 is a valid value of n.
// It has 5 prime factors: [2,2,2,5,5], and it has 6 nice divisors: [10,20,40,50,100,200].
// There is not other value of n that has at most 5 prime factors and more nice divisors.

// Example 2:
// Input: primeFactors = 8
// Output: 18

// Constraints:
//     1 <= primeFactors <= 10^9

import "fmt"

func maxNiceDivisors(primeFactors int) int {
    res, mod := 1, 1_000_000_007
    power := func (x, y int) int {
        res := 1
        for ; y > 0; y >>= 1 {
            if y&1 > 0 {
                res = (res * x) % mod
            }
            x = x * x % mod
        }
        return res
    }
    if primeFactors == 1 { return res }   
    if primeFactors % 3 == 0 {
        res = power(3, primeFactors / 3)
    } else if primeFactors % 3 == 1 {
        res = (4 * power(3, (primeFactors - 4) / 3)) % mod
    } else{
        res = (2 * power(3, (primeFactors - 2) / 3)) % mod
    }
    return res
}

func main() {
    // Example 1:
    // Input: primeFactors = 5
    // Output: 6
    // Explanation: 200 is a valid value of n.
    // It has 5 prime factors: [2,2,2,5,5], and it has 6 nice divisors: [10,20,40,50,100,200].
    // There is not other value of n that has at most 5 prime factors and more nice divisors.
    fmt.Println(maxNiceDivisors(5)) // 6
    // Example 2:
    // Input: primeFactors = 8
    // Output: 18
    fmt.Println(maxNiceDivisors(8)) // 18

    fmt.Println(maxNiceDivisors(1)) // 1
    fmt.Println(maxNiceDivisors(999_999_999)) // 172851386
    fmt.Println(maxNiceDivisors(1_000_000_000)) // 897135186
}
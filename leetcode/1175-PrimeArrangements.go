package main

// 1175. Prime Arrangements
// Return the number of permutations of 1 to n so that prime numbers are at prime indices (1-indexed.)

// (Recall that an integer is prime if and only if it is greater than 1, 
// and cannot be written as a product of two positive integers both smaller than it.)

// Since the answer may be large, return the answer modulo 10^9 + 7.

// Example 1:
// Input: n = 5
// Output: 12
// Explanation: For example [1,2,5,4,3] is a valid permutation, but [5,2,3,4,1] is not because the prime number 5 is at index 1.

// Example 2:
// Input: n = 100
// Output: 682289015

// Constraints:
//     1 <= n <= 100

import "fmt"

func numPrimeArrangements(n int) int {
    res, prime, noPrime := 1, 0, 0
    isPrime := func (n int) bool {
        if n == 1 { return false }
        for i := 2; i < n; i++ {
            if n % i == 0 { return false } // 能被其数整除
        }
        return true
    }
    for i := 1; i <= n; i++ {
        if isPrime(i) {
            prime++ 
        } else {
            noPrime++
        }
    }
    for i := 2; i <= prime; i++ {
        res *= i
        res %= 1_000_000_007
    }
    for i := 2; i <= noPrime; i++ {
        res *= i
        res %= 1_000_000_007
    }
    return res
}

func numPrimeArrangements1(n int) int {
    makePrimeTable := func (limit int) []bool { // 得到一个素数表 1 -> limit
        primes := make([]bool, limit+1)
        primes[0], primes[1] = true, true
        for i := 2; i < len(primes); i++ {
            if !primes[i] {
                for j := i * i; j < len(primes); j += i {
                    primes[j] = true
                }
            }
        }
        for i := range primes { // 翻转一下，true 表示是质数，false 表示非质数
            primes[i] = !primes[i]
        }
        return primes
    }
    isPrime := makePrimeTable(n)
    res, p, np := 1, 1, 1
    for i := 1; i <= n; i++ {
        if isPrime[i] {
            res *= p
            p++
        } else {
            res *= np
            np++
        }
        res %= 1_000_000_007
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 5
    // Output: 12
    // Explanation: For example [1,2,5,4,3] is a valid permutation, but [5,2,3,4,1] is not because the prime number 5 is at index 1.
    fmt.Println(numPrimeArrangements(5)) // 12
    // Example 2:
    // Input: n = 100
    // Output: 682289015
    fmt.Println(numPrimeArrangements(100)) // 682289015

    fmt.Println(numPrimeArrangements(1)) // 1
    fmt.Println(numPrimeArrangements(2)) // 1
    fmt.Println(numPrimeArrangements(3)) // 2
    fmt.Println(numPrimeArrangements(10)) // 17280
    fmt.Println(numPrimeArrangements(99)) // 75763854

    fmt.Println(numPrimeArrangements1(5)) // 12
    fmt.Println(numPrimeArrangements1(100)) // 682289015
    fmt.Println(numPrimeArrangements1(1)) // 1
    fmt.Println(numPrimeArrangements1(2)) // 1
    fmt.Println(numPrimeArrangements1(3)) // 2
    fmt.Println(numPrimeArrangements1(10)) // 17280
    fmt.Println(numPrimeArrangements1(99)) // 75763854
}
package main

// 2507. Smallest Value After Replacing With Sum of Prime Factors
// You are given a positive integer n.

// Continuously replace n with the sum of its prime factors.
//     Note that if a prime factor divides n multiple times, it should be included in the sum as many times as it divides n.

// Return the smallest value n will take on.

// Example 1:
// Input: n = 15
// Output: 5
// Explanation: Initially, n = 15.
// 15 = 3 * 5, so replace n with 3 + 5 = 8.
// 8 = 2 * 2 * 2, so replace n with 2 + 2 + 2 = 6.
// 6 = 2 * 3, so replace n with 2 + 3 = 5.
// 5 is the smallest value n will take on.

// Example 2:
// Input: n = 3
// Output: 3
// Explanation: Initially, n = 3.
// 3 is the smallest value n will take on.

// Constraints:
//     2 <= n <= 10^5

import "fmt"

func smallestValue1(n int) int {
    isPrime := func() []bool {
        isPrime := make([]bool, n + 1)
        for i := range isPrime {
            isPrime[i] = true
        }
        for p := 2; p*p <= n; p++ {
            if isPrime[p] {
                for i := p * p; i <= n; i += p {
                    isPrime[i] = false
                }
            }
        }
        return isPrime
    }()
    sop := func(n int) int {
        sum, i := 0, 2
        for i <= n {
            if isPrime[i] && n%i == 0 {
                sum += i
                n /= i
            } else {
                i += 1
            }
        }
        return sum
    }
    prev := n
    for !isPrime[n] {
        n = sop(n)
        if prev == n { break }
        prev = n
    }
    return n
}

func smallestValue(n int) int {
    s, t := 0, n
    for i := 2; i * i <= t; i++ {
        for t % i == 0 {
            s += i
            t/=i
        }
    }
    if t > 1 {
        s += t
    }
    if s == n {
        return n
    }
    return smallestValue(s)
}

func main() {
    // Example 1:
    // Input: n = 15
    // Output: 5
    // Explanation: Initially, n = 15.
    // 15 = 3 * 5, so replace n with 3 + 5 = 8.
    // 8 = 2 * 2 * 2, so replace n with 2 + 2 + 2 = 6.
    // 6 = 2 * 3, so replace n with 2 + 3 = 5.
    // 5 is the smallest value n will take on.
    fmt.Println(smallestValue(15)) // 5
    // Example 2:
    // Input: n = 3
    // Output: 3
    // Explanation: Initially, n = 3.
    // 3 is the smallest value n will take on.
    fmt.Println(smallestValue(3)) // 3

    fmt.Println(smallestValue(2)) // 2
    fmt.Println(smallestValue(999)) // 7
    fmt.Println(smallestValue(1024)) // 5
    fmt.Println(smallestValue(9999)) // 61
    fmt.Println(smallestValue(100000)) // 7

    fmt.Println(smallestValue1(15)) // 5
    fmt.Println(smallestValue1(3)) // 3
    fmt.Println(smallestValue1(2)) // 2
    fmt.Println(smallestValue1(999)) // 7
    fmt.Println(smallestValue1(1024)) // 5
    fmt.Println(smallestValue1(9999)) // 61
    fmt.Println(smallestValue1(100000)) // 7
}
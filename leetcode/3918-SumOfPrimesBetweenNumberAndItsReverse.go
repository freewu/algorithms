package main

// 3918. Sum of Primes Between Number and Its Reverse
// You are given an integer n.

// Let r be the integer formed by reversing the digits of n.

// Return the sum of all prime numbers between min(n, r) and max(n, r), inclusive.

// A prime number is a natural number greater than 1 with only two factors, 1 and itself.

// Example 1:
// Input: n = 13
// Output: 132
// Explanation:
// The reverse of 13 is 31. Thus, the range is [13, 31].
// The prime numbers in this range are 13, 17, 19, 23, 29, and 31.
// The sum of these prime numbers is 13 + 17 + 19 + 23 + 29 + 31 = 132.

// Example 2:
// Input: n = 10
// Output: 17
// Explanation:
// The reverse of 10 is 1. Thus, the range is [1, 10].
// The prime numbers in this range are 2, 3, 5, and 7.
// The sum of these prime numbers is 2 + 3 + 5 + 7 = 17.

// Example 3:
// Input: n = 8
// Output: 0
// Explanation:
// The reverse of 8 is 8. Thus, the range is [8, 8].
// There are no prime numbers in this range, so the sum is 0.

// Constraints:
//     1 <= n <= 1000

import "fmt"

const MX = 1001
var isPrime [MX]int

func init() {
    for i := 2; i < MX; i++ {
        isPrime[i] = 1
    }
    for i := 2; i*i < MX; i++ {
        if isPrime[i] > 0 {
            for j := i * i; j < MX; j += i {
                isPrime[j] = 0
            }
        }
    }
    // 原地计算 isPrime 的质数前缀和
    for i := 1; i < MX; i++ {
        if isPrime[i] > 0 {
            isPrime[i] = isPrime[i-1] + i
        } else {
            isPrime[i] = isPrime[i-1]
        }
    }
}

func sumOfPrimesInRange(n int) int {
    r := 0
    for x := n; x > 0; x /= 10 {
        r = r * 10 + x % 10
    }
    return isPrime[max(n, r)] - isPrime[min(n, r) - 1]
}

func sumOfPrimesInRange1(n int) int {
    res, rem, tmp := 0, 0, n
    for tmp > 0 {
        rem = rem*10 + tmp%10
        tmp /= 10
    }
    start, end := n, rem
    if start > end {
        start, end = end, start
    }
    for i := start; i <= end; i++ {
        isPrime := true
        if i <= 1 {
            isPrime = false
        } else if i == 2 {
            isPrime = true
        } else if i%2 == 0 {
            isPrime = false
        } else {
            for j := 3; j*j <= i; j += 2 {
                if i%j == 0 {
                    isPrime = false
                    break
                }
            }
        }
        if isPrime {
            res += i
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 13
    // Output: 132
    // Explanation:
    // The reverse of 13 is 31. Thus, the range is [13, 31].
    // The prime numbers in this range are 13, 17, 19, 23, 29, and 31.
    // The sum of these prime numbers is 13 + 17 + 19 + 23 + 29 + 31 = 132.
    fmt.Println(sumOfPrimesInRange(13)) // 132
    // Example 2:
    // Input: n = 10
    // Output: 17
    // Explanation:
    // The reverse of 10 is 1. Thus, the range is [1, 10].
    // The prime numbers in this range are 2, 3, 5, and 7.
    // The sum of these prime numbers is 2 + 3 + 5 + 7 = 17.
    fmt.Println(sumOfPrimesInRange(10)) // 17
    // Example 3:
    // Input: n = 8
    // Output: 0
    // Explanation:
    // The reverse of 8 is 8. Thus, the range is [8, 8].
    // There are no prime numbers in this range, so the sum is 0.
    fmt.Println(sumOfPrimesInRange(8)) // 0

    fmt.Println(sumOfPrimesInRange(1)) // 0
    fmt.Println(sumOfPrimesInRange(2)) // 2
    fmt.Println(sumOfPrimesInRange(3)) // 3
    fmt.Println(sumOfPrimesInRange(64)) // 220
    fmt.Println(sumOfPrimesInRange(99)) // 0
    fmt.Println(sumOfPrimesInRange(100)) // 1060
    fmt.Println(sumOfPrimesInRange(101)) // 101
    fmt.Println(sumOfPrimesInRange(512)) // 18110
    fmt.Println(sumOfPrimesInRange(999)) // 0
    fmt.Println(sumOfPrimesInRange(1000)) // 76127

    fmt.Println(sumOfPrimesInRange1(13)) // 132
    fmt.Println(sumOfPrimesInRange1(10)) // 17
    fmt.Println(sumOfPrimesInRange1(8)) // 0
    fmt.Println(sumOfPrimesInRange1(1)) // 0
    fmt.Println(sumOfPrimesInRange1(2)) // 2
    fmt.Println(sumOfPrimesInRange1(3)) // 3
    fmt.Println(sumOfPrimesInRange1(64)) // 220
    fmt.Println(sumOfPrimesInRange1(99)) // 0
    fmt.Println(sumOfPrimesInRange1(100)) // 1060
    fmt.Println(sumOfPrimesInRange1(101)) // 101
    fmt.Println(sumOfPrimesInRange1(512)) // 18110
    fmt.Println(sumOfPrimesInRange1(999)) // 0
    fmt.Println(sumOfPrimesInRange1(1000)) // 76127
}
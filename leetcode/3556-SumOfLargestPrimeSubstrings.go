package main

// 3556. Sum of Largest Prime Substrings
// Given a string s, find the sum of the 3 largest unique prime numbers that can be formed using any of its substrings.

// Return the sum of the three largest unique prime numbers that can be formed. 
// If fewer than three exist, return the sum of all available primes. If no prime numbers can be formed, return 0.

// Note: Each prime number should be counted only once, even if it appears in multiple substrings. 
// Additionally, when converting a substring to an integer, any leading zeros are ignored.

// Example 1:
// Input: s = "12234"
// Output: 1469
// Explanation:
// The unique prime numbers formed from the substrings of "12234" are 2, 3, 23, 223, and 1223.
// The 3 largest primes are 1223, 223, and 23. Their sum is 1469.

// Example 2:
// Input: s = "111"
// Output: 11
// Explanation:
// The unique prime number formed from the substrings of "111" is 11.
// Since there is only one prime number, the sum is 11.
 
// Constraints:
//     1 <= s.length <= 10
//     s consists of only digits.

import "fmt"
import "slices"
import "math"
import "strconv"

func sumOfLargestPrimes(s string) int64 {
    res, n, primes := 0, len(s), []int{}
    isPrime := func(n int) bool {
        for i := 2; i*i <= n; i++ {
            if n % i == 0 { return false }
        }
        return n >= 2
    }
    for i := 0; i < n; i++ {
        x := 0
        for _, b := range s[i:] {
            x = x * 10 + int(b-'0')
            if isPrime(x) {
                primes = append(primes, x)
            }
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    slices.Sort(primes)
    primes = slices.Compact(primes) // 去重
    for _, v := range primes[max(len(primes) - 3, 0):] {
        res += v
    }
    return int64(res)
}

func sumOfLargestPrimes1(s string) int64 {
    res, n := int64(0), len(s)
    arr := make([]int64, 0, 100)
    isPrime := func(n int64) bool {
        if n <= 1 { return false }
        if n <= 3 { return true  }
        if n % 2 == 0 || n % 3 == 0 { return false }
        sqrtN := int64(math.Sqrt(float64(n)))
        for i := int64(5); i <= sqrtN; i += 6 {
            if n % i == 0 || n % (i + 2) == 0 {
                return false
            }
        }
        return true
    }
    for i := 0; i < n; i++ {
        for j := i + 1; j <= n; j++ {
            n, _ := strconv.ParseInt(s[i:j], 10, 64)
            if isPrime(n) {
                if !slices.Contains(arr, n) {
                    arr = append(arr, n)
                }
            }
        }
    }
    slices.Sort(arr)
    if len(arr) >= 3 {
        for i := len(arr) - 1; i >= len(arr) - 3; i-- {
            res += arr[i]
        }
    } else {
        for i := len(arr) - 1; i >= 0; i-- {
            res += arr[i]
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "12234"
    // Output: 1469
    // Explanation:
    // The unique prime numbers formed from the substrings of "12234" are 2, 3, 23, 223, and 1223.
    // The 3 largest primes are 1223, 223, and 23. Their sum is 1469.
    fmt.Println(sumOfLargestPrimes("12234")) // 1469
    // Example 2:
    // Input: s = "111"
    // Output: 11
    // Explanation:
    // The unique prime number formed from the substrings of "111" is 11.
    // Since there is only one prime number, the sum is 11.
    fmt.Println(sumOfLargestPrimes("111")) // 11

    fmt.Println(sumOfLargestPrimes("123456789")) // 23461445
    fmt.Println(sumOfLargestPrimes("987654321")) // 76593

    fmt.Println(sumOfLargestPrimes1("12234")) // 1469
    fmt.Println(sumOfLargestPrimes1("111")) // 11
    fmt.Println(sumOfLargestPrimes1("123456789")) // 23461445
    fmt.Println(sumOfLargestPrimes1("987654321")) // 76593
}
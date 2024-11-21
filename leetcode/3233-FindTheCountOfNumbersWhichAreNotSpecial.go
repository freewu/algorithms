package main

// 3233. Find the Count of Numbers Which Are Not Special
// You are given 2 positive integers l and r. 
// For any number x, all positive divisors of x except x are called the proper divisors of x.

// A number is called special if it has exactly 2 proper divisors. For example:
//     The number 4 is special because it has proper divisors 1 and 2.
//     The number 6 is not special because it has proper divisors 1, 2, and 3.

// Return the count of numbers in the range [l, r] that are not special.

// Example 1:
// Input: l = 5, r = 7
// Output: 3
// Explanation:
// There are no special numbers in the range [5, 7].

// Example 2:
// Input: l = 4, r = 16
// Output: 11
// Explanation:
// The special numbers in the range [4, 16] are 4 and 9.

// Constraints:
//     1 <= l <= r <= 10^9

import "fmt"
import "math"

func nonSpecialCount(l int, r int) int {
    count := 0
    isPrime := func(v int) bool {
        if v == 1 { return false }
        if v <= 3 { return true } // 2, 3
        if v % 2 == 0 { return false } // 能被 整除的 都不是
        for i := 3; i * i <= v; i += 2 {
            if v % i == 0 { return false }
        }
        return true
    }
    for i := 2; i * i <= r; i++ {
        if i*i < l { continue }
        if isPrime(i) {
            count++
        }
    }
    return r - l + 1 - count
}

func nonSpecialCount1(l int, r int) int {
    res, n := r - l + 1, int(math.Sqrt(float64(r)))
    mp := make([]bool, n + 1)
    for i := 2; i <= n; i++ {
        if !mp[i] {
            if i * i >= l && i * i <= r {
                res--
            }
            for j := i * 2; j <= n; j += i {
                mp[j] = true
            }
        }
    }
    return res
}

func nonSpecialCount2(l int, r int) int {
    const mx = 31622
    pi := make([]int, mx + 1)
    init := func() { // 放到 init 中
        for i := 2; i <= mx; i++ {
            if pi[i] == 0 { // i 是质数
                pi[i] = pi[i-1] + 1
                for j := i * i; j <= mx; j += i {
                    pi[j] = -1 // 标记 i 的倍数为合数
                }
            } else {
                pi[i] = pi[i-1]
            }
        }
    }
    init()
    calc := func(v int) int { return int(math.Sqrt(float64(v))) }
    return r - l + 1 - (pi[calc(r)] - pi[calc(l - 1)])
}

func main() {
    // Example 1:
    // Input: l = 5, r = 7
    // Output: 3
    // Explanation:
    // There are no special numbers in the range [5, 7].
    fmt.Println(nonSpecialCount(5, 7)) // 3
    // Example 2:
    // Input: l = 4, r = 16
    // Output: 11
    // Explanation:
    // The special numbers in the range [4, 16] are 4 and 9.
    fmt.Println(nonSpecialCount(4, 16)) // 11

    fmt.Println(nonSpecialCount(1, 1)) // 1
    fmt.Println(nonSpecialCount(1, 2)) // 2
    fmt.Println(nonSpecialCount(1_000_000_000, 1_000_000_000)) // 1
    fmt.Println(nonSpecialCount(1, 1_000_000_000)) // 999996599
    fmt.Println(nonSpecialCount(999_999_999, 1_000_000_000)) // 2

    fmt.Println(nonSpecialCount1(5, 7)) // 3
    fmt.Println(nonSpecialCount1(4, 16)) // 11
    fmt.Println(nonSpecialCount1(1, 1)) // 1
    fmt.Println(nonSpecialCount1(1, 2)) // 2
    fmt.Println(nonSpecialCount1(1_000_000_000, 1_000_000_000)) // 1
    fmt.Println(nonSpecialCount1(1, 1_000_000_000)) // 999996599
    fmt.Println(nonSpecialCount1(999_999_999, 1_000_000_000)) // 2

    fmt.Println(nonSpecialCount2(5, 7)) // 3
    fmt.Println(nonSpecialCount2(4, 16)) // 11
    fmt.Println(nonSpecialCount2(1, 1)) // 1
    fmt.Println(nonSpecialCount2(1, 2)) // 2
    fmt.Println(nonSpecialCount2(1_000_000_000, 1_000_000_000)) // 1
    fmt.Println(nonSpecialCount2(1, 1_000_000_000)) // 999996599
    fmt.Println(nonSpecialCount2(999_999_999, 1_000_000_000)) // 2
}
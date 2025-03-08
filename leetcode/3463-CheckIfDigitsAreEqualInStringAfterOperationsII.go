package main

// 3463. Check If Digits Are Equal in String After Operations II
// You are given a string s consisting of digits. 
// Perform the following operation repeatedly until the string has exactly two digits:
//     1. For each pair of consecutive digits in s, starting from the first digit, calculate a new digit as the sum of the two digits modulo 10.
//     2. Replace s with the sequence of newly calculated digits, maintaining the order in which they are computed.

// Return true if the final two digits in s are the same; otherwise, return false.

// Example 1:
// Input: s = "3902"
// Output: true
// Explanation:
// Initially, s = "3902"
// First operation:
// (s[0] + s[1]) % 10 = (3 + 9) % 10 = 2
// (s[1] + s[2]) % 10 = (9 + 0) % 10 = 9
// (s[2] + s[3]) % 10 = (0 + 2) % 10 = 2
// s becomes "292"
// Second operation:
// (s[0] + s[1]) % 10 = (2 + 9) % 10 = 1
// (s[1] + s[2]) % 10 = (9 + 2) % 10 = 1
// s becomes "11"
// Since the digits in "11" are the same, the output is true.

// Example 2:
// Input: s = "34789"
// Output: false
// Explanation:
// Initially, s = "34789".
// After the first operation, s = "7157".
// After the second operation, s = "862".
// After the third operation, s = "48".
// Since '4' != '8', the output is false.

// Constraints:
//     3 <= s.length <= 10^5
//     s consists of only digits.

import "fmt"
import "math/bits"

func hasSameDigits(s string) bool {
    n := len(s)
    calc := func(v, mod int) (int, int) {
        count := 0
        for v > 0 && v % mod == 0 {
            count++
            v /= mod
        }
        return v % mod, count
    }
    pow := func(base, exp, mod int) int {
        res := 1
        base %= mod
        if base == 0 { return 0 }
        for ; exp > 0; exp >>= 1 {
            if (exp & 1) == 1 { res = (res * base) % mod }
            base = (base * base) % mod
        }
        return res
    }
    helper := func(mod int) bool {
        res, r, c := 0, 1, 0
        for i := 0; i < n - 1; i++ {
            if c == 0 {
                res += r * (int(s[i]- '0') - int(s[i + 1]- '0'))
            }
            rr, cc := calc(n - 2 - i, mod)
            r = r * rr % mod
            c += cc
            rr, cc = calc(i + 1, mod)
            r = r * pow(rr, mod - 2, mod) % mod
            c -= cc
        }
        return res % mod == 0
    }
    return helper(2) && helper(5)
}

const mod = 10
const mx = 100_000

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

var pow2 = [4]int{2, 4, 8, 6}
var f, invF, p2, p5 [mx + 1]int

func init() {
    f[0] = 1
    for i := 1; i <= mx; i++ {
        x := i
        // 2 的幂次
        e2 := bits.TrailingZeros(uint(x))
        x >>= e2
        // 5 的幂次
        e5 := 0
        for x % 5 == 0 {
            e5++
            x /= 5
        }
        f[i] = f[i-1] * x % mod
        p2[i] = p2[i-1] + e2
        p5[i] = p5[i-1] + e5
    }
    invF[mx] = pow(f[mx], 3) // 欧拉定理
    for i := mx; i > 0; i-- {
        x := i
        x >>= bits.TrailingZeros(uint(x))
        for x % 5 == 0 {
            x /= 5
        }
        invF[i - 1] = invF[i] * x % mod
    }
}

func comb(n, k int) int {
    res := f[n] * invF[k] * invF[n - k]
    e2 := p2[n] - p2[k] - p2[n - k]
    if e2 > 0 {
        res *= pow2[(e2 - 1) % 4]
    }
    if p5[n] - p5[k] - p5[n - k] > 0 {
        res *= 5
    }
    return res
}

func hasSameDigits1(s string) bool {
    diff := 0
    for i := 0 ; i < len(s) - 1; i++ {
        diff += comb(len(s) - 2, i) * (int(s[i]) - int(s[i + 1]))
    }
    return diff % mod == 0
}

func main() {
    // Example 1:
    // Input: s = "3902"
    // Output: true
    // Explanation:
    // Initially, s = "3902"
    // First operation:
    // (s[0] + s[1]) % 10 = (3 + 9) % 10 = 2
    // (s[1] + s[2]) % 10 = (9 + 0) % 10 = 9
    // (s[2] + s[3]) % 10 = (0 + 2) % 10 = 2
    // s becomes "292"
    // Second operation:
    // (s[0] + s[1]) % 10 = (2 + 9) % 10 = 1
    // (s[1] + s[2]) % 10 = (9 + 2) % 10 = 1
    // s becomes "11"
    // Since the digits in "11" are the same, the output is true.
    fmt.Println(hasSameDigits("3902")) // true
    // Example 2:
    // Input: s = "34789"
    // Output: false
    // Explanation:
    // Initially, s = "34789".
    // After the first operation, s = "7157".
    // After the second operation, s = "862".
    // After the third operation, s = "48".
    // Since '4' != '8', the output is false.
    fmt.Println(hasSameDigits("34789")) // false

    fmt.Println(hasSameDigits("123456789")) // false
    fmt.Println(hasSameDigits("987654321")) // false

    fmt.Println(hasSameDigits1("3902")) // true
    fmt.Println(hasSameDigits1("34789")) // false
    fmt.Println(hasSameDigits1("123456789")) // false
    fmt.Println(hasSameDigits1("987654321")) // false
}
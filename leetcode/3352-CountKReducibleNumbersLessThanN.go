package main

// 3352. Count K-Reducible Numbers Less Than N
// You are given a binary string s representing a number n in its binary form.

// You are also given an integer k.

// An integer x is called k-reducible if performing the following operation at most k times reduces it to 1:
//     Replace x with the count of set bits in its binary representation.

// For example, the binary representation of 6 is "110". 
// Applying the operation once reduces it to 2 (since "110" has two set bits). 
// Applying the operation again to 2 (binary "10") reduces it to 1 (since "10" has one set bit).

// Return an integer denoting the number of positive integers less than n that are k-reducible.

// Since the answer may be too large, return it modulo 10^9 + 7.

// Example 1:
// Input: s = "111", k = 1
// Output: 3
// Explanation:
// n = 7. The 1-reducible integers less than 7 are 1, 2, and 4.

// Example 2:
// Input: s = "1000", k = 2
// Output: 6
// Explanation:
// n = 8. The 2-reducible integers less than 8 are 1, 2, 3, 4, 5, and 6.

// Example 3:
// Input: s = "1", k = 3
// Output: 0
// Explanation:
// There are no positive integers less than n = 1, so the answer is 0.

// Constraints:
//     1 <= s.length <= 800
//     s has no leading zeros.
//     s consists only of the characters '0' and '1'.
//     1 <= k <= 5

import "fmt"
import "math/bits"

func countKReducibleNumbers(s string, k int) int {
    res, n, mod := 0, len(s), 1_000_000_007
    pow := func(x, n, mod int) int {
        x %= mod
        res := 1 % mod
        for ; n > 0; n /= 2 {
            if n%2 > 0 {
                res = res * x % mod
            }
            x = x * x % mod
        }
        return res
    }
    // 打表
    inv := make([]int,n + 1)
    for i := 1; i <= n; i++ {
        inv[i] = pow(i, mod - 2, mod)
    }
    mem := make([][]int, n + 1)
    for i := 1; i <= n; i++ {
        mem[i] = make([]int, i + 1)
        mem[i][0], mem[i][i] = 1, 1
        for j := 1; j < i; j++ {
            mem[i][j] = (mem[i][j - 1] * ((i - j + 1) * inv[j] % mod)) % mod
        }
    }
    var check func(x, c int) bool
    check = func(x, c int) bool {
        if x == 1 { return true }
        c--
        if c == 0 { return false }
        return check(bits.OnesCount(uint(x)), c)
    }
    one := 0
    for i := 0; i < n; i++ {
        if s[i] == '1' {
            one++
        }
    }
    l, f := 0, one
    for i := n - 1; i >= 0; i-- {
        l++
        if s[i] == '1' {
            if f != one && check(one, k) {
                res = (res + 1) % mod
            }
            one--
            for j := 1; j < l; j++ {
                if check(one + j, k) {
                    res = (res + mem[l-1][j]) % mod
                }
            }
        }
    }
    return res
}

func countKReducibleNumbers1(s string, k int) int {
    res, n, mod := 0, len(s), 1_000_000_007
    if n == 1 { return 0 }
    if k == 1 {
        c := 0
        for i := 0; i < n; i++ {
            if s[i] == '1' {
                c++
            }
        }
        if c == 1 {
            return n - 1
        }
        return n
    }
    memo := make([]int, n)
    for i := 2; i <= n; i++ {
        c := 0
        for j := i; j > 0; {
            j ^= j & (-j) // unset one bit
            c++
        }
        memo[i - 1] = 1 + memo[c - 1]
    }
    count := make([]int, 1, n)
    for l, c := 1, 0; l < n; l++ {
        count = append(count, 0)
        for i := l; i > 0; i-- {
            count[i] = (count[i] + count[i - 1]) % mod
        }
        count[0]++
        if s[l] == '1' {
            count[c] = (count[c] + 1) % mod
            c++
        }
    }
    for b, c := range count {
        if memo[b] < k { // b + 1 is (k-1)-reducible
            res = (res + c) % mod
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "111", k = 1
    // Output: 3
    // Explanation:
    // n = 7. The 1-reducible integers less than 7 are 1, 2, and 4.
    fmt.Println(countKReducibleNumbers("111", 1)) // 3
    // Example 2:
    // Input: s = "1000", k = 2
    // Output: 6
    // Explanation:
    // n = 8. The 2-reducible integers less than 8 are 1, 2, 3, 4, 5, and 6.
    fmt.Println(countKReducibleNumbers("1000", 2)) // 6
    // Example 3:
    // Input: s = "1", k = 3
    // Output: 0
    // Explanation:
    // There are no positive integers less than n = 1, so the answer is 0.
    fmt.Println(countKReducibleNumbers("1", 3)) // 0

    fmt.Println(countKReducibleNumbers("1111111111", 3)) // 902
    fmt.Println(countKReducibleNumbers("0000000000", 3)) // 0
    fmt.Println(countKReducibleNumbers("1010101010", 3)) // 637
    fmt.Println(countKReducibleNumbers("0101010101", 3)) // 331
    fmt.Println(countKReducibleNumbers("1111100000", 3)) // 881
    fmt.Println(countKReducibleNumbers("0000011111", 3)) // 30

    fmt.Println(countKReducibleNumbers1("111", 1)) // 3
    fmt.Println(countKReducibleNumbers1("1000", 2)) // 6
    fmt.Println(countKReducibleNumbers1("1", 3)) // 0
    fmt.Println(countKReducibleNumbers1("1111111111", 3)) // 902
    fmt.Println(countKReducibleNumbers1("0000000000", 3)) // 0
    fmt.Println(countKReducibleNumbers1("1010101010", 3)) // 637
    fmt.Println(countKReducibleNumbers1("0101010101", 3)) // 331
    fmt.Println(countKReducibleNumbers1("1111100000", 3)) // 881
    fmt.Println(countKReducibleNumbers1("0000011111", 3)) // 30
}
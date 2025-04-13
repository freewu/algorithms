package main

// 3519. Count Numbers with Non-Decreasing Digits 
// You are given two integers, l and r, represented as strings, and an integer b. 
// Return the count of integers in the inclusive range [l, r] whose digits are in non-decreasing order when represented in base b.

// An integer is considered to have non-decreasing digits if, 
// when read from left to right (from the most significant digit to the least significant digit), each digit is greater than or equal to the previous one.

// Since the answer may be too large, return it modulo 10^9 + 7.

// Example 1:
// Input: l = "23", r = "28", b = 8
// Output: 3
// Explanation:
// The numbers from 23 to 28 in base 8 are: 27, 30, 31, 32, 33, and 34.
// Out of these, 27, 33, and 34 have non-decreasing digits. Hence, the output is 3.

// Example 2:
// Input: l = "2", r = "7", b = 2
// Output: 2
// Explanation:
// The numbers from 2 to 7 in base 2 are: 10, 11, 100, 101, 110, and 111.
// Out of these, 11 and 111 have non-decreasing digits. Hence, the output is 2.

// Constraints:
//     1 <= l.length <= r.length <= 100
//     2 <= b <= 10
//     l and r consist only of digits.
//     The value represented by l is less than or equal to the value represented by r.
//     l and r do not contain leading zeros.

import "fmt"
import "math/big"
import "strings"

func countNumbers(l, r string, b int) int {
    mod := 1_000_000_007
    trans := func(s string, b int) string {
        x := big.Int{}
        fmt.Fscan(strings.NewReader(s), &x)
        return x.Text(b) // 转成 b 进制
    }
    lowS, highS := trans(l, b), trans(r, b)
    n := len(highS)
    diff := n - len(lowS)
    memo := make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, b)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(i, pre int, limitLow, limitHigh bool) int 
    dfs = func(i, pre int, limitLow, limitHigh bool) int {
        if i == n { return 1 }
        res, low, high := 0, 0, b - 1
        if !limitLow && !limitHigh {
            p := &memo[i][pre]
            if *p >= 0 { return *p }
            defer func() { *p = res }()
        }
        if limitLow && i >= diff {
            low = int(lowS[i - diff] - '0')
        }
        if limitHigh {
            high = int(highS[i] - '0')
        }
        for d := max(low, pre); d <= high; d++ {
            res += dfs(i + 1, d, limitLow && d == low, limitHigh && d == high)
        }
        return res % mod
    }
    return dfs(0, 0, true, true)
}

func main() {
    // Example 1:
    // Input: l = "23", r = "28", b = 8
    // Output: 3
    // Explanation:
    // The numbers from 23 to 28 in base 8 are: 27, 30, 31, 32, 33, and 34.
    // Out of these, 27, 33, and 34 have non-decreasing digits. Hence, the output is 3.
    fmt.Println(countNumbers("23", "28", 8)) // 3
    // Example 2:
    // Input: l = "2", r = "7", b = 2
    // Output: 2
    // Explanation:
    // The numbers from 2 to 7 in base 2 are: 10, 11, 100, 101, 110, and 111.
    // Out of these, 11 and 111 have non-decreasing digits. Hence, the output is 2.
    fmt.Println(countNumbers("2", "7", 2)) // 2

    fmt.Println(countNumbers("1", "1", 2)) // 1
    fmt.Println(countNumbers("123456789", "123456789", 2)) // 0
    fmt.Println(countNumbers("1", "123456789", 2)) // 26
    fmt.Println(countNumbers("123456789", "1", 2)) // 1
    fmt.Println(countNumbers("987654321", "987654321", 2)) // 0
    fmt.Println(countNumbers("123456789", "987654321", 2)) // 3
    fmt.Println(countNumbers("987654321", "123456789", 2)) // 1
}
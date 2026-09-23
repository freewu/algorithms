package main

// 4060. Count Evenly Good Integers
// You are given two integers l and r.

// An integer is called evenly good if it contains an even number of even digits.

// Return the number of evenly good integers in the inclusive range [l, r].

// Example 1:
// Input: l = 18, r = 22
// Output: 3
// Explanation:
// The evenly good integers in the range [18, 22] are:
// 19, because it contains 0 even digits.
// 20, because it contains 2 even digits.
// 22, because it contains 2 even digits.
// Thus, the answer is 3.

// Example 2:
// Input: l = 98, r = 101
// Output: 2
// Explanation:
// The evenly good integers in the range [98, 101] are:
// 99, because it contains 0 even digits.
// 100, because it contains 2 even digits.
// Thus, the answer is 2.

// Example 3:
// Input: l = 1, r = 10
// Output: 5
// Explanation:
// The evenly good integers in the range [1, 10] are 1, 3, 5, 7, and 9, because each of them contains 0 even digits. 
// Thus, the answer is 5.

// Constraints:
//     1 <= l <= r <= 10^15

import "fmt"
import "strconv"

func countEvenlyGoodIntegers(l int64, r int64) int64 {
    var digits []int
    var dp [][]int64
    var dfs func(pos int, parity int, tight bool, leadingZero bool) int64
    dfs = func(pos int, parity int, tight bool, leadingZero bool) int64 {
        if pos == len(digits) {
            // 不是前导零，并且偶数数字个数是偶数，才算
            if leadingZero {
                return 0
            }
            if parity == 0 {
                return 1
            }
            return 0
        }
        // 状态编码 parity(1bit), tight(1bit), leadingZero(1bit) → 0~7
        state := 0
        if parity == 1 {
            state |= 1
        }
        if tight {
            state |= 2
        }
        if leadingZero {
            state |= 4
        }
        if dp[pos][state] != -1 {
            return dp[pos][state]
        }
        limit := 9
        if tight {
            limit = digits[pos]
        }
        res := int64(0)
        for d := 0; d <= limit; d++ {
            newTight := tight && (d == limit)
            newLZ := leadingZero && (d == 0)
            newP := parity
            if !newLZ {
                // d是偶数数字(0,2,4,6,8)，翻转奇偶计数
                if d%2 == 0 {
                    newP = parity ^ 1
                }
            }
            res += dfs(pos+1, newP, newTight, newLZ)
        }
        dp[pos][state] = res
        return res
    }
    calc := func(x int64) int64 {
        if x < 0 {
            return 0
        }
        s := strconv.FormatInt(x, 10)
        digits = make([]int, 0, len(s))
        for _, c := range s {
            digits = append(digits, int(c-'0'))
        }
        n := len(digits)
        // pos最多16位，state 0~7
        dp = make([][]int64, n)
        for i := range dp {
            dp[i] = make([]int64, 8)
            for j := range dp[i] {
                dp[i][j] = -1
            }
        }
        return dfs(0, 0, true, true)
    }
    return calc(r) - calc(l-1)
}

func main() {
    // Example 1:
    // Input: l = 18, r = 22
    // Output: 3
    // Explanation:
    // The evenly good integers in the range [18, 22] are:
    // 19, because it contains 0 even digits.
    // 20, because it contains 2 even digits.
    // 22, because it contains 2 even digits.
    // Thus, the answer is 3.
    fmt.Println(countEvenlyGoodIntegers(18, 22)) // 3
    // Example 2:
    // Input: l = 98, r = 101
    // Output: 2
    // Explanation:
    // The evenly good integers in the range [98, 101] are:
    // 99, because it contains 0 even digits.
    // 100, because it contains 2 even digits.
    // Thus, the answer is 2.
    fmt.Println(countEvenlyGoodIntegers(98, 101)) // 2
    // Example 3:
    // Input: l = 1, r = 10
    // Output: 5
    // Explanation:
    // The evenly good integers in the range [1, 10] are 1, 3, 5, 7, and 9, because each of them contains 0 even digits. 
    // Thus, the answer is 5.
    fmt.Println(countEvenlyGoodIntegers(1, 10)) // 5

    fmt.Println(countEvenlyGoodIntegers(1, 1)) // 1
    fmt.Println(countEvenlyGoodIntegers(1, 1024)) // 512
    fmt.Println(countEvenlyGoodIntegers(1024, 1024)) // 0
    fmt.Println(countEvenlyGoodIntegers(1, 1_000_000_000_000_000)) // 500000000000000
    fmt.Println(countEvenlyGoodIntegers(1_000_000_000_000_000, 1_000_000_000_000_000)) // 0
}
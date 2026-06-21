package main

// 3966. Count Good Integers in a Range
// You are given three integers l, r and k.

// A number is considered good if the absolute difference between every pair of adjacent digits is at most k.

// Return the number of good integers in the range [l, r] (inclusive).

// The absolute difference between values x and y is defined as abs(x - y).

// Example 1:
// Input: l = 10, r = 15, k = 1
// Output: 3
// Explanation:
// The good integers in the range are 10, 11, and 12.
// For 10, abs(1 - 0) = 1.
// For 11, abs(1 - 1) = 0.
// For 12, abs(1 - 2) = 1.
// All these differences are at most k = 1. Thus, the answer is 3.

// Example 2:
// Input: l = 201, r = 204, k = 2
// Output: 2
// Explanation:
// The good integers in the range are 201 and 202.
// For 201, abs(2 - 0) = 2 and abs(0 - 1) = 1.
// For 202, abs(2 - 0) = 2 and abs(0 - 2) = 2.
// Thus, the answer is 2.

// Constraints:
//     10 <= l <= r <= 10^15
//     0 <= k <= 9

import "fmt"
import "strconv"

func goodIntegers(l, r int64, k int) int64 {
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    lowS, highS := strconv.FormatInt(l, 10), strconv.FormatInt(r, 10)
    n := len(highS)
    diff := n - len(lowS)
    memo := make([][10]int64, n)
    for i := range memo {
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
    var dfs func(i, pre int, limitLow, limitHigh bool) int64
    dfs = func(i, pre int, limitLow, limitHigh bool) (res int64) {
        if i == n {
            return 1 // 找到一个好数
        }
        if !limitLow && !limitHigh {
            p := &memo[i][pre]
            if *p >= 0 {
                return *p
            }
            defer func() { *p = res }()
        }
        low, high := 0, 9
        if limitLow && i >= diff {
            low = int(lowS[i-diff] - '0')
        }
        if limitHigh {
            high = int(highS[i] - '0')
        }
        d := low
        if limitLow && i < diff {
            // 不填数字，上界不受约束
            res = dfs(i+1, 0, true, false)
            d = 1 // 下面填数字，从 1 开始填
        }
        // 如果在 diff 之前填过数字，那么 limitLow 一定是 false
        isFirst := limitLow && i <= diff
        for ; d <= high; d++ {
            if isFirst || abs(d-pre) <= k {
                res += dfs(i+1, d, limitLow && d == low, limitHigh && d == high)
            }
        }
        return
    }
    // pre 的初始值随意
    return dfs(0, 0, true, true)
}

func main() {
    // Example 1:
    // Input: l = 10, r = 15, k = 1
    // Output: 3
    // Explanation:
    // The good integers in the range are 10, 11, and 12.
    // For 10, abs(1 - 0) = 1.
    // For 11, abs(1 - 1) = 0.
    // For 12, abs(1 - 2) = 1.
    // All these differences are at most k = 1. Thus, the answer is 3.
    fmt.Println(goodIntegers(10, 15, 1)) // 3
    // Example 2:
    // Input: l = 201, r = 204, k = 2
    // Output: 2
    // Explanation:
    // The good integers in the range are 201 and 202.
    // For 201, abs(2 - 0) = 2 and abs(0 - 1) = 1.
    // For 202, abs(2 - 0) = 2 and abs(0 - 2) = 2.
    // Thus, the answer is 2.
    fmt.Println(goodIntegers(201, 204, 2)) // 2

    fmt.Println(goodIntegers(10, 10, 0)) // 0
    fmt.Println(goodIntegers(10, 10, 9)) // 1
    fmt.Println(goodIntegers(10, 1_000_000_000_000_000, 0)) // 126
    fmt.Println(goodIntegers(10, 1_000_000_000_000_000, 9)) // 999999999999991
    fmt.Println(goodIntegers(1_000_000_000_000_000, 1_000_000_000_000_000, 0)) // 0
    fmt.Println(goodIntegers(1_000_000_000_000_000, 1_000_000_000_000_000, 9)) // 1
}
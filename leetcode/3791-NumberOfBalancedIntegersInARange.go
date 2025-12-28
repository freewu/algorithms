package main

// 3791. Number of Balanced Integers in a Range
// You are given two integers low and high.

// An integer is called balanced if it satisfies both of the following conditions:
//     1. It contains at least two digits.
//     2. The sum of digits at even positions is equal to the sum of digits at odd positions (the leftmost digit has position 1).

// Return an integer representing the number of balanced integers in the range [low, high] (both inclusive).

// Example 1:
// Input: low = 1, high = 100
// Output: 9
// Explanation:
// The 9 balanced numbers between 1 and 100 are 11, 22, 33, 44, 55, 66, 77, 88, and 99.

// Example 2:
// Input: low = 120, high = 129
// Output: 1
// Explanation:
// Only 121 is balanced because the sum of digits at even and odd positions are both 2.

// Example 3:
// Input: low = 1234, high = 1234
// Output: 0
// Explanation:
// 1234 is not balanced because the sum of digits at odd positions (1 + 3 = 4) does not equal the sum at even positions (2 + 4 = 6).

// Constraints:
//     1 <= low <= high <= 10^15

import "fmt"
import "strconv"

func countBalanced(low, high int64) int64 {
    if high < 11 { return 0 } // 最小的满足要求的数是 11
    low = max(low, 11)
    lowS := strconv.FormatInt(low, 10)
    highS := strconv.FormatInt(high, 10)
    n := len(highS)
    diffLH := n - len(lowS)
    memo := make([][]int64, n)
    for i := range memo {
        // diff 至少 floor(n/2) * 9，至多 ceil(n/2) * 9，值域大小 n * 9
        memo[i] = make([]int64, n*9+1)
    }
    var dfs func(int, int, bool, bool) int64
    dfs = func(i, diff int, limitLow, limitHigh bool) (res int64) {
        if i == n {
            if diff != 0 { return 0 } // 不合法
            return 1
        }
        if !limitLow && !limitHigh {
            p := &memo[i][diff+n/2*9] // 保证下标非负
            if *p > 0 {
                return *p - 1
            }
            defer func() { *p = res + 1 }() // 记忆化的时候加一，这样 memo 可以初始化成 0
        }
        lo := 0
        if limitLow && i >= diffLH {
            lo = int(lowS[i-diffLH] - '0')
        }
        hi := 9
        if limitHigh {
            hi = int(highS[i] - '0')
        }
        d := lo
        // 通过 limit_low 和 i 可以判断能否不填数字，无需 isNum 参数
        if limitLow && i < diffLH { // 可以不填任何数
            res = dfs(i+1, diff, true, false) // 上界无约束
            d = 1 // 下面填数字，至少从 1 开始填
        }
        for ; d <= hi; d++ {
            // 下一个位置奇偶性翻转
            res += dfs(i+1, diff+(1-i%2*2)*d,
                limitLow && d == lo, limitHigh && d == hi)
        }
        return
    }
    return dfs(0, 0, true, true)
}

func main() {
    // Example 1:
    // Input: low = 1, high = 100
    // Output: 9
    // Explanation:
    // The 9 balanced numbers between 1 and 100 are 11, 22, 33, 44, 55, 66, 77, 88, and 99.
    fmt.Println(countBalanced(1, 100)) // 9
    // Example 2:
    // Input: low = 120, high = 129
    // Output: 1
    // Explanation:
    // Only 121 is balanced because the sum of digits at even and odd positions are both 2.
    fmt.Println(countBalanced(120, 129)) // 1
    // Example 3:
    // Input: low = 1234, high = 1234
    // Output: 0
    // Explanation:
    // 1234 is not balanced because the sum of digits at odd positions (1 + 3 = 4) does not equal the sum at even positions (2 + 4 = 6).
    fmt.Println(countBalanced(1234, 1234)) // 0

    fmt.Println(countBalanced(1, 1)) // 0
    fmt.Println(countBalanced(1, 1_000_000_000_000_000)) // 32811494188974
    fmt.Println(countBalanced(1024, 1_000_000_000_000_000)) // 32811494188917
    fmt.Println(countBalanced(1_000_000_000_000_000, 1_000_000_000_000_000)) // 0
}
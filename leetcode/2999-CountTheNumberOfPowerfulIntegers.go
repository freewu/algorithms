package main

// 2999. Count the Number of Powerful Integers
// You are given three integers start, finish, and limit. 
// You are also given a 0-indexed string s representing a positive integer.

// A positive integer x is called powerful if it ends with s (in other words, s is a suffix of x) 
// and each digit in x is at most limit.

// Return the total number of powerful integers in the range [start..finish].

// A string x is a suffix of a string y if and only if x is a substring of y 
// that starts from some index (including 0) in y and extends to the index y.length - 1. 
// For example, 25 is a suffix of 5125 whereas 512 is not.

// Example 1:
// Input: start = 1, finish = 6000, limit = 4, s = "124"
// Output: 5
// Explanation: The powerful integers in the range [1..6000] are 124, 1124, 2124, 3124, and, 4124. All these integers have each digit <= 4, and "124" as a suffix. Note that 5124 is not a powerful integer because the first digit is 5 which is greater than 4.
// It can be shown that there are only 5 powerful integers in this range.

// Example 2:
// Input: start = 15, finish = 215, limit = 6, s = "10"
// Output: 2
// Explanation: The powerful integers in the range [15..215] are 110 and 210. All these integers have each digit <= 6, and "10" as a suffix.
// It can be shown that there are only 2 powerful integers in this range.

// Example 3:
// Input: start = 1000, finish = 2000, limit = 4, s = "3000"
// Output: 0
// Explanation: All integers in the range [1000..2000] are smaller than 3000, hence "3000" cannot be a suffix of any integer in this range.

// Constraints:
//     1 <= start <= finish <= 10^15
//     1 <= limit <= 9
//     1 <= s.length <= floor(log10(finish)) + 1
//     s only consists of numeric digits which are at most limit.
//     s does not have leading zeros.

import "fmt"
import "strconv"
import "math"
import "strings"

func numberOfPowerfulInt(start int64, finish int64, limit int, s string) int64 {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    countSatisfyingNumbers := func(target int) int { // 计算小于等于 target 的数中符合条件的数的数量
        sInt, _ := strconv.Atoi(s) // 将字符串转换为整数
        if sInt > target { return 0 } // 如果字符串表示的整数大于 target，返回 0
        targetStr := strconv.Itoa(target) // 将 target 转换为字符串
        tn, sn := len(targetStr), len(s)
        if sn == tn { return 1 } // 如果 target 和 s 的长度相同，返回 1
        dp := make([]int, tn + 1) // 动态规划数组，用于存储中间结果
        if targetStr[tn - sn:] >= s {
            dp[sn] = 1
        }
        for i := sn + 1; i <= tn; i++ { // 填充动态规划数组
            t, _ := strconv.Atoi(string(targetStr[tn - i]))
            // 这里计算第i位的符合条件的数，是i-1位的长度的limit+1次方(因为取值是[0, limit])，乘以当前位的可选值
            dp[i] += int(math.Pow(float64(limit + 1), float64(i - sn - 1))) * min(t, limit + 1)
            // 当 t > limit 时，当前位的值本身就是无效的。这意味着，即使 min(t, limit+1) 限制了它的影响，当前位上的实际数字 t 已经违反了条件（即超过了 limit）。
            // 因此，我们不能简单地将 dp[i-1] 加到 dp[i] 上，因为这会计算那些实际上不满足条件的数字。 
            // 这就是为什么我们只在 t <= limit 的情况下才考虑累加 dp[i-1] 的原因。
            if t <= limit {
                dp[i] += dp[i-1]
            }
        }
        return dp[tn]
    }
    // start finish 区间改为求两次上限
    return int64(countSatisfyingNumbers(int(finish)) - countSatisfyingNumbers(int(start)-1))
}

func numberOfPowerfulInt1(start int64, finish int64, limit int, s string) int64 {
    num1, num2 := strconv.Itoa(int(start)), strconv.Itoa(int(finish))
    num1 = strings.Repeat("0", len(num2) - len(num1)) + num1
    preLen := len(num2) - len(s)
    memo := make([]int, len(num2)+1)
    for i := range memo {
        memo[i] = -1
    }
    var dp func(idx int, limit_high bool, limit_low bool) int
    dp = func(idx int, limit_low bool, limit_high bool) int {
        if idx == len(num2) { return 1 }
        if !limit_high && !limit_low {
            if memo[idx] != -1 {
                return memo[idx]
            }
        }
        res, low, high := 0, 0, 9
        if limit_low  { low  = int(num1[idx] - '0') }
        if limit_high { high = int(num2[idx] - '0') }
        if idx < preLen {
            // 不限制limit
            for i := low; i <= min(limit, high); i++ {
                res += dp(idx + 1, limit_low && i == low, limit_high && i == high )
            }
        } else {
            digit := int(s[idx - preLen] - '0')
            if low <= digit && digit <= high {
                res += dp(idx + 1, limit_low && digit == low, limit_high && digit == high)
            }
        }
        if !limit_high && !limit_low {
            memo[idx] = res
        }
        return res
    }
    return int64(dp(0, true, true))
}

func main() {
    // Example 1:
    // Input: start = 1, finish = 6000, limit = 4, s = "124"
    // Output: 5
    // Explanation: The powerful integers in the range [1..6000] are 124, 1124, 2124, 3124, and, 4124. All these integers have each digit <= 4, and "124" as a suffix. Note that 5124 is not a powerful integer because the first digit is 5 which is greater than 4.
    // It can be shown that there are only 5 powerful integers in this range.
    fmt.Println(numberOfPowerfulInt(1, 6000, 4, "124")) // 5
    // Example 2:
    // Input: start = 15, finish = 215, limit = 6, s = "10"
    // Output: 2
    // Explanation: The powerful integers in the range [15..215] are 110 and 210. All these integers have each digit <= 6, and "10" as a suffix.
    // It can be shown that there are only 2 powerful integers in this range.
    fmt.Println(numberOfPowerfulInt(15, 215, 6, "10")) // 2
    // Example 3:
    // Input: start = 1000, finish = 2000, limit = 4, s = "3000"
    // Output: 0
    // Explanation: All integers in the range [1000..2000] are smaller than 3000, hence "3000" cannot be a suffix of any integer in this range.
    fmt.Println(numberOfPowerfulInt(1000, 2000, 4, "3000")) // 0

    fmt.Println(numberOfPowerfulInt1(1, 6000, 4, "124")) // 5
    fmt.Println(numberOfPowerfulInt1(15, 215, 6, "10")) // 2
    fmt.Println(numberOfPowerfulInt1(1000, 2000, 4, "3000")) // 0
}
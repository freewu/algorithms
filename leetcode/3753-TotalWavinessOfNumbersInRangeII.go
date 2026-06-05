package main

// 3753. Total Waviness of Numbers in Range II
// You are given two integers num1 and num2 representing an inclusive range [num1, num2].

// The waviness of a number is defined as the total count of its peaks and valleys:
//     1. A digit is a peak if it is strictly greater than both of its immediate neighbors.
//     2. A digit is a valley if it is strictly less than both of its immediate neighbors.
//     3. The first and last digits of a number cannot be peaks or valleys.
//     4. Any number with fewer than 3 digits has a waviness of 0.

// Return the total sum of waviness for all numbers in the range [num1, num2].

// Example 1:
// Input: num1 = 120, num2 = 130
// Output: 3
// Explanation:
// In the range [120, 130]:
// 120: middle digit 2 is a peak, waviness = 1.
// 121: middle digit 2 is a peak, waviness = 1.
// 130: middle digit 3 is a peak, waviness = 1.
// All other numbers in the range have a waviness of 0.
// Thus, total waviness is 1 + 1 + 1 = 3.

// Example 2:
// Input: num1 = 198, num2 = 202
// Output: 3
// Explanation:
// In the range [198, 202]:
// 198: middle digit 9 is a peak, waviness = 1.
// 201: middle digit 0 is a valley, waviness = 1.
// 202: middle digit 0 is a valley, waviness = 1.
// All other numbers in the range have a waviness of 0.
// Thus, total waviness is 1 + 1 + 1 = 3.

// Example 3:
// Input: num1 = 4848, num2 = 4848
// Output: 2
// Explanation:
// Number 4848: the second digit 8 is a peak, and the third digit 4 is a valley, giving a waviness of 2.

// Constraints:
//     1 <= num1 <= num2 <= 10^15

import "fmt"
import "strconv"
import "cmp"

func totalWaviness(num1, num2 int64) int64 {
    lowS,highS := strconv.FormatInt(num1, 10),  strconv.FormatInt(num2, 10)
    n := len(highS)
    diff := n - len(lowS)
    memo := make([][][3][10]int, n)
    for i := range memo {
        memo[i] = make([][3][10]int, n-1) // 一个数至多包含 n-2 个峰或谷
    }

    var dfs func(int, int, int, int, bool, bool) int
    dfs = func(i, waviness, lastCmp, lastDigit int, limitLow, limitHigh bool) (res int) {
        if i == n {
            return waviness
        }
        if !limitLow && !limitHigh {
            p := &memo[i][waviness][lastCmp+1][lastDigit]
            if *p > 0 {
                return *p - 1
            }
            defer func() { *p = res + 1 }()
        }
        lo := 0
        if limitLow && i >= diff {
            lo = int(lowS[i-diff] - '0')
        }
        hi := 9
        if limitHigh {
            hi = int(highS[i] - '0')
        }
        isNum := !limitLow || i > diff // 前面是否填过数字
        for d := lo; d <= hi; d++ {
            c := 0
            if isNum { // 当前填的数不是最高位
                c = cmp.Compare(d, lastDigit)
            }
            w := waviness
            if c*lastCmp < 0 { // 形成了一个峰或谷
                w++
            }
            res += dfs(i+1, w, c, d, limitLow && d == lo, limitHigh && d == hi)
        }
        return
    }
    return int64(dfs(0, 0, 0, 0, true, true))
}

func totalWaviness1(num1, num2 int64) int64 {
    calc := func(n int64) int64 { // 计算 [1, n] 中的整数的波动值之和
        res := int64(0)
        // 把整数拆分成五段：prefix | l | m | r | suffix
        // 从低到高枚举 (l, m, r) 的位置，计算 (l, m, r) 对答案的贡献
        for pow10 := int64(1); n >= pow10 * 100; pow10 *= 10 {
            maxPrefix := n / (pow10 * 1000)
            // 1. prefix < maxPrefix 时，低位不受约束
            // 但 prefix=0 且 l=0 的情况是不合法的，需要减掉
            count := maxPrefix * 570 - 45 // 先不与 pow10 相乘
            n2 := n / pow10
            L, M, R := n2/100%10, n2/10%10, n2%10
            // 2. prefix = maxPrefix 且 l < L
            count += (242 + L*30 - L*L*2) * L / 6
            // 3. prefix = maxPrefix 且 l = L 且 m < M
            count += (L + M) * max(M-L-1, 0) / 2      // 峰
            count += (19 - min(L, M)) * min(L, M) / 2 // 谷
            // 4. prefix = maxPrefix 且 l = L 且 m = M 且 r < R
            if L < M { // 只能是峰
                count += min(M, R)
            } else if L > M { // 只能是谷
                count += max(R-M-1, 0)
            }
            // 到此为止，suffix 可以随便填，有 pow10 种填法
            res += count * pow10
            // 5. prefix = maxPrefix 且 l = L 且 m = M 且 r = R
            if (L - M) * (M - R) < 0 { // 峰或谷
                maxSuffix := n % pow10
                res += maxSuffix + 1 // suffix 可以填 [0, maxSuffix] 中的任意整数
            }
        }
        return res
    }
    return calc(num2) - calc(num1 - 1)
}

func main() {
    // Example 1:
    // Input: num1 = 120, num2 = 130
    // Output: 3
    // Explanation:
    // In the range [120, 130]:
    // 120: middle digit 2 is a peak, waviness = 1.
    // 121: middle digit 2 is a peak, waviness = 1.
    // 130: middle digit 3 is a peak, waviness = 1.
    // All other numbers in the range have a waviness of 0.
    // Thus, total waviness is 1 + 1 + 1 = 3.
    fmt.Println(totalWaviness(120, 130)) // 3   
    // Example 2:
    // Input: num1 = 198, num2 = 202
    // Output: 3
    // Explanation:
    // In the range [198, 202]:
    // 198: middle digit 9 is a peak, waviness = 1.
    // 201: middle digit 0 is a valley, waviness = 1.
    // 202: middle digit 0 is a valley, waviness = 1.
    // All other numbers in the range have a waviness of 0.
    // Thus, total waviness is 1 + 1 + 1 = 3.
    fmt.Println(totalWaviness(120, 130)) // 3   
    // Example 3:
    // Input: num1 = 4848, num2 = 4848
    // Output: 2
    // Explanation:
    // Number 4848: the second digit 8 is a peak, and the third digit 4 is a valley, giving a waviness of 2.
    fmt.Println(totalWaviness(4848, 4848)) // 2

    fmt.Println(totalWaviness(198, 202)) // 3
    fmt.Println(totalWaviness(1, 1)) // 0
    fmt.Println(totalWaviness(1, 1024)) // 543
    fmt.Println(totalWaviness(1, 1_000_000)) // 2230005
    fmt.Println(totalWaviness(1024, 1024)) // 1
    fmt.Println(totalWaviness(1024, 1_000_000)) // 2233939
    fmt.Println(totalWaviness(1_000_000, 1_000_000)) // 0
    fmt.Println(totalWaviness(1, 1_000_000_000_000_000)) // 7360000000000005
    fmt.Println(totalWaviness(1_000_000_000_000_000, 1_000_000_000_000_000)) // 0

    fmt.Println(totalWaviness1(120, 130)) // 3   
    fmt.Println(totalWaviness1(120, 130)) // 3   
    fmt.Println(totalWaviness1(4848, 4848)) // 2
    fmt.Println(totalWaviness1(198, 202)) // 3
    fmt.Println(totalWaviness1(1, 1)) // 0
    fmt.Println(totalWaviness1(1, 1024)) // 543
    fmt.Println(totalWaviness1(1, 1_000_000)) // 2230005
    fmt.Println(totalWaviness1(1024, 1024)) // 1
    fmt.Println(totalWaviness1(1024, 1_000_000)) // 2233939
    fmt.Println(totalWaviness1(1_000_000, 1_000_000)) // 0
    fmt.Println(totalWaviness1(1, 1_000_000_000_000_000)) // 7360000000000005
    fmt.Println(totalWaviness1(1_000_000_000_000_000, 1_000_000_000_000_000)) // 0
}
package main

// 4115. Sum of Compatible Numbers in Range II
// You are given two integers n and k.

// A positive integer x is called compatible if it satisfies both of the following conditions:
//     1. abs(n - x) <= k
//     2. (n & x) == 0

// Return the sum of all compatible integers x.

// Since the answer may be very large, return it modulo 109 + 7.

// Note:
//     1. Here, & denotes the bitwise AND operator.
//     2. The absolute difference between integers i and j is defined as abs(i - j).

// Example 1:
// Input: n = 2, k = 3
// Output: 10
// Explanation:
// The compatible integers are:
// x = 1, since abs(2 - 1) = 1 and 2 & 1 = 0.
// x = 4, since abs(2 - 4) = 2 and 2 & 4 = 0.
// x = 5, since abs(2 - 5) = 3 and 2 & 5 = 0.
// Thus, the answer is 1 + 4 + 5 = 10.

// Example 2:
// Input: n = 5, k = 1
// Output: 0
// Explanation:
// There are no compatible integers in the range [4, 6]. Thus, the answer is 0.

// Constraints:
//     1 <= n <= 10^15
//     0 <= k <= 10^15

import "fmt"

func sumOfGoodIntegers(n int64, k int64) int {
    const MOD = 1_000_000_007
    type State struct {
        count int64
        sum int64
    }
    low := n - k
    if low < 1 {
        low = 1
    }
    high := n + k
    // f(x): [1, x] 中满足 (n & x) == 0 的所有数字之和 mod MOD
    calc := func(x int64) int64 {
        if x < 1 {
            return 0
        }
        // 提取二进制，高位在前
        var b []int
        tmp := x
        for tmp > 0 {
            b = append(b, int(tmp&1))
            tmp >>= 1
        }
        for l, r := 0, len(b)-1; l < r; l, r = l+1, r-1 {
            b[l], b[r] = b[r], b[l]
        }
        m := len(b)
        // memo[pos][tight][leadingZero]
        var memo [64][2][2]*State
        var dfs func(pos int, tight int, leadingZero int) State
        dfs = func(pos int, tight int, leadingZero int) State {
            if pos == m {
                if leadingZero == 1 {
                    return State{count: 0, sum: 0}
                }
                return State{count: 1, sum: 0}
            }
            if p := memo[pos][tight][leadingZero]; p != nil {
                return *p
            }
            limit := 1
            if tight == 1 {
                limit = b[pos]
            }
            var res State
            for d := 0; d <= limit; d++ {
                newTight := 0
                if tight == 1 && d == limit {
                    newTight = 1
                }
                newLead := 0
                if leadingZero == 1 && d == 0 {
                    newLead = 1
                }
                // n 的这一位如果是1，x不能取1
                bitN := (n >> (m - 1 - pos)) & 1
                if bitN == 1 && int64(d) == 1 {
                    continue
                }
                sub := dfs(pos+1, newTight, newLead)
                // 当前位权值： 1 << (m-1-pos)
                pow := int64(1) << (m - 1 - pos)
                add := (int64(d) * pow % MOD) * sub.count % MOD
                res.count = (res.count + sub.count) % MOD
                res.sum = (res.sum + add + sub.sum) % MOD
            }
            memo[pos][tight][leadingZero] = &res
            return res
        }
        ret := dfs(0, 1, 1)
        return ret.sum % MOD
    }
    res := (calc(high) - calc(low - 1) + MOD) % MOD
    return int(res)
}

func main() {
    // Example 1:
    // Input: n = 2, k = 3
    // Output: 10
    // Explanation:
    // The compatible integers are:
    // x = 1, since abs(2 - 1) = 1 and 2 & 1 = 0.
    // x = 4, since abs(2 - 4) = 2 and 2 & 4 = 0.
    // x = 5, since abs(2 - 5) = 3 and 2 & 5 = 0.
    // Thus, the answer is 1 + 4 + 5 = 10.
    fmt.Println(sumOfGoodIntegers(2, 3)) // 10
    // Example 2:
    // Input: n = 5, k = 1
    // Output: 0
    // Explanation:
    // There are no compatible integers in the range [4, 6]. Thus, the answer is 0.
    fmt.Println(sumOfGoodIntegers(5, 1)) // 0

    fmt.Println(sumOfGoodIntegers(756756, 319514)) // 901966969
    fmt.Println(sumOfGoodIntegers(431156, 375863)) // 250345465
    fmt.Println(sumOfGoodIntegers(142, 140)) // 1961
    fmt.Println(sumOfGoodIntegers(4, 1)) // 3
    fmt.Println(sumOfGoodIntegers(1, 0)) // 0
    fmt.Println(sumOfGoodIntegers(1, 1_000_000_000_000_000)) // 1269736067576315904
    fmt.Println(sumOfGoodIntegers(1_000_000_000_000_000, 0)) // 0
    fmt.Println(sumOfGoodIntegers(1_000_000_000_000_000, 1_000_000_000_000_000)) // 6255042484506722304
}
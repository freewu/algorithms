package main

// 3869. Count Fancy Numbers in a Range
// You are given two integers l and r.

// An integer is called good if its digits form a strictly monotone sequence, meaning the digits are strictly increasing or strictly decreasing. 
// All single-digit integers are considered good.

// An integer is called fancy if it is good, or if the sum of its digits is good.

// Return an integer representing the number of fancy integers in the range [l, r] (inclusive).

// A sequence is said to be strictly increasing if each element is strictly greater than its previous one (if exists).

// A sequence is said to be strictly decreasing if each element is strictly less than its previous one (if exists).

// Example 1:
// Input: l = 8, r = 10
// Output: 3
// Explanation:
//     8 and 9 are single-digit integers, so they are good and therefore fancy.
//     10 has digits [1, 0], which form a strictly decreasing sequence, so 10 is good and thus fancy.
// Therefore, the answer is 3.

// Example 2:
// Input: l = 12340, r = 12341
// Output: 1
// Explanation:
// 12340
//     12340 is not good because [1, 2, 3, 4, 0] is not strictly monotone.
//     The digit sum is 1 + 2 + 3 + 4 + 0 = 10.
//     10 is good as it has digits [1, 0], which is strictly decreasing. Therefore, 12340 is fancy.
// 12341
//     12341 is not good because [1, 2, 3, 4, 1] is not strictly monotone.
//     The digit sum is 1 + 2 + 3 + 4 + 1 = 11.
//     11 is not good as it has digits [1, 1], which is not strictly monotone. Therefore, 12341 is not fancy.
// Therefore, the answer is 1.

// Example 3:
// Input: l = 123456788, r = 123456788
// Output: 0
// Explanation:
//     123456788 is not good because its digits are not strictly monotone.
//     The digit sum is 1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 8 = 44.
//     44 is not good as it has digits [4, 4], which is not strictly monotone. Therefore, 123456788 is not fancy.
// Therefore, the answer is 0.

// Constraints:
//     1 <= l <= r <= 10^15

import "fmt"
import "sort"
import "strconv"
import "slices"

var goodNums []int

func init() {
    for mask := 1; mask < 1<<10; mask++ {
        // 构造严格递减好数
        x, sum := 0, 0
        for i := 9; i >= 0; i-- {
            if mask >> i & 1 > 0 {
                x = x * 10 + i
                sum += i
            }
        }
        if !isGood(sum) {
            goodNums = append(goodNums, x)
        }
        // 构造严格递增好数
        if mask & 1 > 0 { continue }// 不能包含 0
        x, sum = 0, 0
        for i := 1; i < 10; i++ {
            if mask >> i & 1 > 0 {
                x = x * 10 + i
                sum += i
            }
        }
        if !isGood(sum) {
            goodNums = append(goodNums, x)
        }
    }
    slices.Sort(goodNums) // 方便二分求个数
}

// 判断数位和 s 是否为好数
func isGood(s int) bool {
    if s < 100 { // s 是个位数或者两位数
        return s / 10 != s % 10 // 十位和个位不相等即为好数
    }
    // s 是三位数，其百位一定是 1
    return 1 < s / 10 % 10 && s / 10 % 10 < s % 10 // 只能严格递增
}

func countFancy(l, r int64) int64 {
    // 计算 [l, r] 内的数位和不是好数的好数的个数
    res := int64(sort.SearchInts(goodNums, int(r+1)) - sort.SearchInts(goodNums, int(l)))
    // 计算 [l, r] 内的数位和是好数的数的个数（这个数是不是好数都可以）
    lowS, highS := strconv.FormatInt(l, 10), strconv.FormatInt(r, 10)
    n := len(highS)
    diffLH := n - len(lowS)
    memo := make([][]int64, n)
    for i := range memo {
        memo[i] = make([]int64, n*9+1) // 数位和最大 9n
    }
    var dfs func(int, int, bool, bool) int64
    dfs = func(i, digitSum int, limitLow, limitHigh bool) (res int64) {
        if i == n {
            if isGood(digitSum) { return 1 } // 合法
            return 0 // 不合法
        }
        if !limitLow && !limitHigh {
            dv := &memo[i][digitSum]
            if *dv > 0 {
                return *dv - 1
            }
            defer func() { *dv = res + 1 }() // 这样写无需初始化 memo 为 -1
        }
        low := 0
        if limitLow && i >= diffLH {
            low = int(lowS[i-diffLH] - '0')
        }
        high := 9
        if limitHigh {
            high = int(highS[i] - '0')
        }
        d := low
        // 通过 limitLow 和 i 可以判断能否不填数字，无需 isNum 参数
        if limitLow && i < diffLH {
            // 不填数字，上界不受约束
            res = dfs(i+1, 0, true, false)
            d = 1 // 下面填数字，从 1 开始填
        }
        for ; d <= high; d++ {
            res += dfs(i+1, digitSum + d, limitLow && d == low, limitHigh && d == high)
        }
        return
    }
    return res + dfs(0, 0, true, true)
}

func main() {
    // Example 1:
    // Input: l = 8, r = 10
    // Output: 3
    // Explanation:
    //     8 and 9 are single-digit integers, so they are good and therefore fancy.
    //     10 has digits [1, 0], which form a strictly decreasing sequence, so 10 is good and thus fancy.
    // Therefore, the answer is 3.
    fmt.Println(countFancy(8, 10)) // 3
    // Example 2:
    // Input: l = 12340, r = 12341
    // Output: 1
    // Explanation:
    // 12340
    //     12340 is not good because [1, 2, 3, 4, 0] is not strictly monotone.
    //     The digit sum is 1 + 2 + 3 + 4 + 0 = 10.
    //     10 is good as it has digits [1, 0], which is strictly decreasing. Therefore, 12340 is fancy.
    // 12341
    //     12341 is not good because [1, 2, 3, 4, 1] is not strictly monotone.
    //     The digit sum is 1 + 2 + 3 + 4 + 1 = 11.
    //     11 is not good as it has digits [1, 1], which is not strictly monotone. Therefore, 12341 is not fancy.
    // Therefore, the answer is 1.
    fmt.Println(countFancy(12340, 12341)) // 1
    // Example 3:
    // Input: l = 123456788, r = 123456788
    // Output: 0
    // Explanation:
    //     123456788 is not good because its digits are not strictly monotone.
    //     The digit sum is 1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 8 = 44.
    //     44 is not good as it has digits [4, 4], which is not strictly monotone. Therefore, 123456788 is not fancy.
    // Therefore, the answer is 0.
    fmt.Println(countFancy(123456788, 123456788)) // 0

    fmt.Println(countFancy(1, 1)) // 1
    fmt.Println(countFancy(1, 1_000_000_000_000_000)) // 907441159188136
    fmt.Println(countFancy(1_000_000_000_000_000, 1)) // -137
    fmt.Println(countFancy(1_000_000_000_000_000, 1_000_000_000_000_000)) // 1
}
package main

// 3490. Count Beautiful Numbers
// You are given two positive integers, l and r. 
// A positive integer is called beautiful if the product of its digits is divisible by the sum of its digits.

// Return the count of beautiful numbers between l and r, inclusive.

// Example 1:
// Input: l = 10, r = 20
// Output: 2
// Explanation:
// The beautiful numbers in the range are 10 and 20.

// Example 2:
// Input: l = 1, r = 15
// Output: 10
// Explanation:
// The beautiful numbers in the range are 1, 2, 3, 4, 5, 6, 7, 8, 9, and 10.

// Constraints:
//      <= l <= r < 109

import "fmt"
import "strconv"
import "slices"

func beautifulNumbers(l int, r int) int {
    res, start, end := 0, strconv.FormatInt(int64(l), 10), strconv.FormatInt(int64(r), 10)
    dp := make(map[[3]int]int)
    var getNumberOfBeauty func(sum, product, index int, bt []byte) int
    getNumberOfBeauty = func(sum, product, index int, bt []byte) int {
        res := 0
        validCache := (len(bt) > len(start) || string(bt[:index]) > start[:index]) && (len(bt) < len(end) || string(bt[:index]) < end[:index])
        if validCache {
            if v, ok := dp[[3]int{sum, product, len(bt) - index}]; ok {
                return v
            }
        }
        if product == 0 && validCache {
            res = 1
            for i := 1; i <= len(bt) - index; i++ {
                res *= 10
            }
        } else {
            for t := 0; t <= 9; t++ {
                bt[index] = byte(t + 48)
                if len(bt) == len(start) && string(bt[:index+1]) < start[:index + 1] { continue }
                if len(bt) == len(end) && string(bt[:index+1]) > end[:index+1] { break }
                if index == len(bt) - 1 {
                    if (product * t) % (sum + t) == 0 {
                        res++
                    }
                } else {
                    res += getNumberOfBeauty(sum + t, product * t, index + 1, bt)
                }
            }
        }
        if validCache {
            dp[[3]int{sum, product, len(bt) - index}] = res
        }
        return res
    }
    for i := len(start); i <= len(end); i++ {
        bt := make([]byte, i)
        for t := 1; t <= 9; t++ {
            bt[0] = byte(t + 48)
            if len(bt) == len(start) && bt[0] < start[0] { continue }
            if len(bt) == len(end) && bt[0] > end[0] { break }
            if i == 1 {
                res++
            } else {
                res += getNumberOfBeauty(t, t, 1, bt)
            }
        }
    }
    return res
}

func beautifulNumbers1(l, r int) int {
    type Tuple struct{ i, m, s int }
    memo := make(map[Tuple]int)
    var dfs func(i, m, s int, isLimit, isNum bool, high []byte) int
    dfs = func(i, m, s int, isLimit, isNum bool, high []byte) int {
        if i < 0 {
            if s == 0 || m % s > 0 {
                return 0
            }
            return 1
        }
        res, hi := 0, 9
        if !isLimit && isNum {
            t := Tuple{i, m, s}
            if v, ok := memo[t]; ok {
                return v
            }
            defer func() { memo[t] = res }()
        }
        if isLimit {
            hi = int(high[i] - '0')
        }
        d := 0
        if !isNum {
            res, d = dfs(i - 1, m, s, false, false, high), 1
        }
        // 枚举填数位 d
        for ; d <= hi; d++ {
            res += dfs(i-1, m*d, s+d, isLimit && d == hi, true, high)
        }
        return res
    }
    calc := func(v int) int {
        high := []byte(strconv.Itoa(v))
        slices.Reverse(high)
        return dfs(len(high) - 1, 1, 0, true, false, high)
    }
    return calc(r) - calc(l - 1)
}

func main() {
    // Example 1:
    // Input: l = 10, r = 20
    // Output: 2
    // Explanation:
    // The beautiful numbers in the range are 10 and 20.
    fmt.Println(beautifulNumbers(10, 20)) // 2
    // Example 2:
    // Input: l = 1, r = 15
    // Output: 10
    // Explanation:
    // The beautiful numbers in the range are 1, 2, 3, 4, 5, 6, 7, 8, 9, and 10.
    fmt.Println(beautifulNumbers(1, 15)) // 10

    fmt.Println(beautifulNumbers(1, 1)) // 1
    fmt.Println(beautifulNumbers(1_000_000_000, 1_000_000_000)) // 1
    fmt.Println(beautifulNumbers(1, 1_000_000_000)) // 670349659

    fmt.Println(beautifulNumbers1(10, 20)) // 2
    fmt.Println(beautifulNumbers1(1, 15)) // 10
    fmt.Println(beautifulNumbers1(1, 1)) // 1
    fmt.Println(beautifulNumbers1(1_000_000_000, 1_000_000_000)) // 1
    fmt.Println(beautifulNumbers1(1, 1_000_000_000)) // 670349659
}
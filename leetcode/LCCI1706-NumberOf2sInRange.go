package main

// 面试题 17.06. Number Of 2s In Range LCCI
// Write a method to count the number of 2s that appear in all the numbers between 0 and n (inclusive).

// Example:
// Input: 25
// Output: 9
// Explanation: (2, 12, 20, 21, 22, 23, 24, 25)(Note that 22 counts for two 2s.)

// Note:
//     n <= 10^9

import "fmt"
import "strconv"

func numberOf2sInRange(n int) int {
    s := strconv.Itoa(n)
    m := len(s)
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, m)
        for j := range dp[i] {
            dp[i][j] = -1
        }
    }
    var dfs func(i, count int, isLimit bool) int
    dfs = func(i, count int, isLimit bool) int {
        if i == m { return count }
        res, up := 0, 9
        if !isLimit {
            dv := &dp[i][count]
            if *dv >= 0 {  return *dv }
            defer func() { *dv = res }()
        }
        if isLimit {
            up = int(s[i] - '0')
        }
        for d := 0; d <= up; d++ { // 枚举要填入的数字 d
            c := count
            if d == 2 {
                c++
            }
            res += dfs(i + 1, c, isLimit && d == up)
        }
        return res
    }
    return dfs(0, 0, true)
}

func main() {
    // Example:
    // Input: 25
    // Output: 9
    // Explanation: (2, 12, 20, 21, 22, 23, 24, 25)(Note that 22 counts for two 2s.)
    fmt.Println(numberOf2sInRange(25)) // 9

    fmt.Println(numberOf2sInRange(1)) // 0
    fmt.Println(numberOf2sInRange(2)) // 1
    fmt.Println(numberOf2sInRange(3)) // 1
    fmt.Println(numberOf2sInRange(8)) // 1
    fmt.Println(numberOf2sInRange(999)) // 300
    fmt.Println(numberOf2sInRange(1024)) // 308
    fmt.Println(numberOf2sInRange(999_999_999)) // 900000000
    fmt.Println(numberOf2sInRange(1_000_000_000)) // 900000000
}
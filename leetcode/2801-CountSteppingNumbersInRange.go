package main

// 2801. Count Stepping Numbers in Range
// Given two positive integers low and high represented as strings, find the count of stepping numbers in the inclusive range [low, high].

// A stepping number is an integer such that all of its adjacent digits have an absolute difference of exactly 1.

// Return an integer denoting the count of stepping numbers in the inclusive range [low, high].

// Since the answer may be very large, return it modulo 10^9 + 7.

// Note: A stepping number should not have a leading zero.

// Example 1:
// Input: low = "1", high = "11"
// Output: 10
// Explanation: The stepping numbers in the range [1,11] are 1, 2, 3, 4, 5, 6, 7, 8, 9 and 10. There are a total of 10 stepping numbers in the range. Hence, the output is 10.

// Example 2:
// Input: low = "90", high = "101"
// Output: 2
// Explanation: The stepping numbers in the range [90,101] are 98 and 101. There are a total of 2 stepping numbers in the range. Hence, the output is 2. 

// Constraints:
//     1 <= int(low) <= int(high) < 10100
//     1 <= low.length, high.length <= 100
//     low and high consist of only digits.
//     low and high don't have any leading zeros.

import "fmt"

func countSteppingNumbers(low string, high string) int {
    const mod = 1_000_000_007
    facts := [110][10]int{}
    for i := range facts {
        for j := range facts[i] {
            facts[i][j] = -1
        }
    }
    num := high
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    var dfs func(pos, pre int, lead bool, limit bool) int
    dfs = func(pos, pre int, lead bool, limit bool) int {
        if pos >= len(num) {
            if lead { return 0 }
            return 1
        }
        if !lead && !limit && facts[pos][pre] != -1 { return facts[pos][pre] }
        res, up := 0, 9
        if limit {
            up = int(num[pos] - '0')
        }
        for i := 0; i <= up; i++ {
            if i == 0 && lead {
                res += dfs(pos + 1, pre, true, limit && i == up)
            } else if pre == -1 || abs(pre - i) == 1 {
                res += dfs(pos + 1, i, false, limit && i == up)
            }
            res %= mod
        }
        if !lead && !limit {
            facts[pos][pre] = res
        }
        return res
    }
    a := dfs(0, -1, true, true)
    t := []byte(low)
    for i := len(t) - 1; i >= 0; i-- {
        if t[i] != '0' {
            t[i]--
            break
        }
        t[i] = '9'
    }
    num = string(t)
    facts = [110][10]int{}
    for i := range facts {
        for j := range facts[i] {
            facts[i][j] = -1
        }
    }
    b := dfs(0, -1, true, true)
    return (a - b + mod) % mod
}

func main() {
    // Example 1:
    // Input: low = "1", high = "11"
    // Output: 10
    // Explanation: The stepping numbers in the range [1,11] are 1, 2, 3, 4, 5, 6, 7, 8, 9 and 10. There are a total of 10 stepping numbers in the range. Hence, the output is 10.
    fmt.Println(countSteppingNumbers("1", "11")) // 10
    // Example 2:
    // Input: low = "90", high = "101"
    // Output: 2
    // Explanation: The stepping numbers in the range [90,101] are 98 and 101. There are a total of 2 stepping numbers in the range. Hence, the output is 2. 
    fmt.Println(countSteppingNumbers("90", "101")) // 2

    fmt.Println(countSteppingNumbers("1", "10100")) // 119
    fmt.Println(countSteppingNumbers("1", "1")) // 1
    fmt.Println(countSteppingNumbers("10100", "10100")) // 0
    fmt.Println(countSteppingNumbers("10100", "1")) // 999999889
}
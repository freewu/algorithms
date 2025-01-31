package main

// 2767. Partition String Into Minimum Beautiful Substrings
// Given a binary string s, partition the string into one or more substrings such that each substring is beautiful.

// A string is beautiful if:
//     It doesn't contain leading zeros.
//     It's the binary representation of a number that is a power of 5.

// Return the minimum number of substrings in such partition. 
// If it is impossible to partition the string s into beautiful substrings, return -1.

// A substring is a contiguous sequence of characters in a string.

// Example 1:
// Input: s = "1011"
// Output: 2
// Explanation: We can paritition the given string into ["101", "1"].
// - The string "101" does not contain leading zeros and is the binary representation of integer 51 = 5.
// - The string "1" does not contain leading zeros and is the binary representation of integer 50 = 1.
// It can be shown that 2 is the minimum number of beautiful substrings that s can be partitioned into.

// Example 2:
// Input: s = "111"
// Output: 3
// Explanation: We can paritition the given string into ["1", "1", "1"].
// - The string "1" does not contain leading zeros and is the binary representation of integer 50 = 1.
// It can be shown that 3 is the minimum number of beautiful substrings that s can be partitioned into.

// Example 3:
// Input: s = "0"
// Output: -1
// Explanation: We can not partition the given string into beautiful substrings.

// Constraints:
//     1 <= s.length <= 15
//     s[i] is either '0' or '1'.

import "fmt"
import "strconv"

func minimumBeautifulSubstrings(s string) int {
    power, n := 1, len(s)
    mp, dp := make(map[int]bool), make([]int, n + 1)
    for i := 0; i <= n; i++ {
        mp[power] = true
        power *= 5
        dp[i] = -1
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    var dfs func(int) int
    dfs = func(i int) int {
        if i >= n { return 0 }
        if s[i] == '0' { return n + 1 }
        if dp[i] != -1 { return dp[i] }
        dp[i] = n + 1
        x := 0
        for j := i; j < n; j++ {
            x = x << 1 | int(s[j] - '0')
            if mp[x] {
                dp[i] = min(dp[i], 1 + dfs(j + 1))
            }
        }
        return dp[i]
    }
    res := dfs(0)
    if res > n { return -1 }
    return res
}

func minimumBeautifulSubstrings1(s string) int {
    n := len(s)
    dp := make([]int, n + 1)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    beautiful := func (s string) bool { // 判断一个字符串是否是美丽的
        if len(s) == 0 || s[0] == '0' { return false }
        t, _ := strconv.ParseInt(s, 2, 64)
        for t != 1 {
            if t % 5 == 0 {
                t /= 5
            } else {
                return false
            }
        }
        return true
    }
    for i := 1; i <= n; i++ {
        dp[i] = 1 << 31
        for j := i - 1; j >= 0; j-- {
            if beautiful(s[j:i]) {
                dp[i] = min(dp[i], dp[j] + 1)
            }
        }
    }
    if dp[n] >= 1 << 31 { return -1 }
    return dp[n]
}

func main() {
    // Example 1:
    // Input: s = "1011"
    // Output: 2
    // Explanation: We can paritition the given string into ["101", "1"].
    // - The string "101" does not contain leading zeros and is the binary representation of integer 51 = 5.
    // - The string "1" does not contain leading zeros and is the binary representation of integer 50 = 1.
    // It can be shown that 2 is the minimum number of beautiful substrings that s can be partitioned into.
    fmt.Println(minimumBeautifulSubstrings("1011")) // 2
    // Example 2:
    // Input: s = "111"
    // Output: 3
    // Explanation: We can paritition the given string into ["1", "1", "1"].
    // - The string "1" does not contain leading zeros and is the binary representation of integer 50 = 1.
    // It can be shown that 3 is the minimum number of beautiful substrings that s can be partitioned into.
    fmt.Println(minimumBeautifulSubstrings("111")) // 3
    // Example 3:
    // Input: s = "0"
    // Output: -1
    // Explanation: We can not partition the given string into beautiful substrings.
    fmt.Println(minimumBeautifulSubstrings("0")) // -1

    fmt.Println(minimumBeautifulSubstrings1("1011")) // 2
    fmt.Println(minimumBeautifulSubstrings1("111")) // 3
    fmt.Println(minimumBeautifulSubstrings1("0")) // -1
}
package main

// 2539. Count the Number of Good Subsequences
// A subsequence of a string is good if it is not empty and the frequency of each one of its characters is the same.

// Given a string s, return the number of good subsequences of s. 
// Since the answer may be too large, return it modulo 10^9 + 7.

// A subsequence is a string that can be derived from another string by deleting some 
// or no characters without changing the order of the remaining characters.

// Example 1:
// Input: s = "aabb"
// Output: 11
// Explanation: The total number of subsequences is 24. There are five subsequences which are not good: "aabb", "aabb", "aabb", "aabb", and the empty subsequence. Hence, the number of good subsequences is 24-5 = 11.

// Example 2:
// Input: s = "leet"
// Output: 12
// Explanation: There are four subsequences which are not good: "leet", "leet", "leet", and the empty subsequence. Hence, the number of good subsequences is 24-4 = 12.

// Example 3:
// Input: s = "abcd"
// Output: 15
// Explanation: All of the non-empty subsequences are good subsequences. Hence, the number of good subsequences is 24-1 = 15.

// Constraints:
//     1 <= s.length <= 10^4
//     s consists of only lowercase English letters.

import "fmt"

func countGoodSubsequences(s string) int {
    count := [26]int{}
    res, mx, n, mod := 0, 1, 10001, 1_000_000_007
    f, g := make([]int, n), make([]int, n)
    qmi := func (a, k, p int) int {
        res := 1
        for k != 0 {
            if k & 1 == 1 { res = res * a % p }
            k >>= 1
            a = a * a % p
        }
        return res
    }
    init := func() {
        f[0], g[0] = 1, 1
        for i := 1; i < n; i++ {
            f[i] = f[i-1] * i % mod
            g[i] = qmi(f[i], mod - 2, mod)
        }
    }
    init() // 可以放到 init()方法里 
    max := func (x, y int) int { if x > y { return x; }; return y; }
    comb := func(n, k int) int { return (f[n] * g[k] % mod) * g[n-k] % mod }
    for _, v := range s {
        count[v - 'a']++
        mx = max(mx, count[v - 'a'])
    }
    for i := 1; i <= mx; i++ {
        x := 1
        for _, v := range count {
            if v >= i {
                x = (x * (comb(v, i) + 1)) % mod
            }
        }
        res = (res + x - 1) % mod
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "aabb"
    // Output: 11
    // Explanation: The total number of subsequences is 24. There are five subsequences which are not good: "aabb", "aabb", "aabb", "aabb", and the empty subsequence. Hence, the number of good subsequences is 24-5 = 11.
    fmt.Println(countGoodSubsequences("aabb")) // 11
    // Example 2:
    // Input: s = "leet"
    // Output: 12
    // Explanation: There are four subsequences which are not good: "leet", "leet", "leet", and the empty subsequence. Hence, the number of good subsequences is 24-4 = 12.
    fmt.Println(countGoodSubsequences("leet")) // 12
    // Example 3:
    // Input: s = "abcd"
    // Output: 15
    // Explanation: All of the non-empty subsequences are good subsequences. Hence, the number of good subsequences is 24-1 = 15.
    fmt.Println(countGoodSubsequences("abcd")) // 15
}
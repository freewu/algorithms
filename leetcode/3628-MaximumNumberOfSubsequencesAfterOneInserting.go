package main

// 3628. Maximum Number of Subsequences After One Inserting
// You are given a string s consisting of uppercase English letters.

// You are allowed to insert at most one uppercase English letter at any position (including the beginning or end) of the string.

// Return the maximum number of "LCT" subsequences that can be formed in the resulting string after at most one insertion.

// A subsequence is a non-empty string that can be derived from another string by deleting some 
// or no characters without changing the order of the remaining characters.

// Example 1:
// Input: s = "LMCT"
// Output: 2
// Explanation:
// We can insert a "L" at the beginning of the string s to make "LLMCT", which has 2 subsequences, at indices [0, 3, 4] and [1, 3, 4].

// Example 2:
// Input: s = "LCCT"
// Output: 4
// Explanation:
// We can insert a "L" at the beginning of the string s to make "LLCCT", which has 4 subsequences, at indices [0, 2, 4], [0, 3, 4], [1, 2, 4] and [1, 3, 4].

// Example 3:
// Input: s = "L"
// Output: 0
// Explanation:
// Since it is not possible to obtain the subsequence "LCT" by inserting a single letter, the result is 0.

// Constraints:
//     1 <= s.length <= 10^5
//     s consists of uppercase English letters.

import "fmt"
import "strings"

func numOfSubsequences(s string) int64 {
    numDistinct := func(s, t string) int { // 115. 不同的子序列
        n, m := len(s), len(t)
        if n < m { return 0 }
        f := make([]int, m + 1)
        f[0] = 1
        for i, x := range s {
            for j := min(i, m-1); j >= max(m-n+i, 0); j-- {
                if byte(x) == t[j] {
                    f[j+1] += f[j]
                }
            }
        }
        return f[m]
    }
    calcInsertT := func(s string) int { // 计算插入 T 产生的额外 LCT 子序列个数的最大值
        cntT := strings.Count(s, "T") // s[i+1] 到 s[n-1] 的 'T' 的个数
        res, cntL := 0, 0 // s[0] 到 s[i] 的 'L' 的个数
        for _, c := range s {
            if c == 'T' {
                cntT--
            }
            if c == 'L' {
                cntL++
            }
            res = max(res, cntL*cntT)
        }
        return res
    }
    extra := max(numDistinct(s, "CT"), numDistinct(s, "LC"), calcInsertT(s))
    return int64(numDistinct(s, "LCT") + extra)
}

func numOfSubsequences1(s string) int64 {
    l, t, c := 0, 0, 0
    for i := 0; i < len(s); i++ {
        if s[i] == 'T' {
            t++
        } else if s[i] == 'C' {
            c++
        }
    }
    seq, lextra, textra, cextra := 0, 0, 0, 0
    for i := 0; i < len(s); i++ {
        if s[i] == 'L' {
            l++
        } 
        if s[i] == 'C' {
            lextra += (l+1)*t
            textra += l*(t+1)
            seq += l*t
        }
        cextra = max(cextra, l*t)
        if s[i] == 'T' {
            t--
        }
    }
    res := max(lextra, max(textra, seq + cextra))
    return int64(res)
}

func main() {
    // Example 1:
    // Input: s = "LMCT"
    // Output: 2
    // Explanation:
    // We can insert a "L" at the beginning of the string s to make "LLMCT", which has 2 subsequences, at indices [0, 3, 4] and [1, 3, 4].
    fmt.Println(numOfSubsequences("LMCT")) // 2
    // Example 2:
    // Input: s = "LCCT"
    // Output: 4
    // Explanation:
    // We can insert a "L" at the beginning of the string s to make "LLCCT", which has 4 subsequences, at indices [0, 2, 4], [0, 3, 4], [1, 2, 4] and [1, 3, 4].
    fmt.Println(numOfSubsequences("LCCT")) // 4
    // Example 3:
    // Input: s = "L"
    // Output: 0
    // Explanation:
    // Since it is not possible to obtain the subsequence "LCT" by inserting a single letter, the result is 0.
    fmt.Println(numOfSubsequences("L")) // 0

    fmt.Println(numOfSubsequences("LEETCODE")) // 1
    fmt.Println(numOfSubsequences("BLUEFROG")) // 0

    fmt.Println(numOfSubsequences1("LMCT")) // 2
    fmt.Println(numOfSubsequences1("LCCT")) // 4
    fmt.Println(numOfSubsequences1("L")) // 0
    fmt.Println(numOfSubsequences1("LEETCODE")) // 1
    fmt.Println(numOfSubsequences1("BLUEFROG")) // 0
}
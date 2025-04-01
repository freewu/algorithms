package main

// 3504. Longest Palindrome After Substring Concatenation II
// You are given two strings, s and t.

// You can create a new string by selecting a substring from s (possibly empty) and a substring from t (possibly empty), then concatenating them in order.

// Return the length of the longest palindrome that can be formed this way.

// Example 1:
// Input: s = "a", t = "a"
// Output: 2
// Explanation:
// Concatenating "a" from s and "a" from t results in "aa", which is a palindrome of length 2.

// Example 2:
// Input: s = "abc", t = "def"
// Output: 1
// Explanation:
// Since all characters are different, the longest palindrome is any single character, so the answer is 1.

// Example 3:
// Input: s = "b", t = "aaaa"
// Output: 4
// Explanation:
// Selecting "aaaa" from t is the longest palindrome, so the answer is 4.

// Example 4:
// Input: s = "abcde", t = "ecdba"
// Output: 5
// Explanation:
// Concatenating "abc" from s and "ba" from t results in "abcba", which is a palindrome of length 5.

// Constraints:
//     1 <= s.length, t.length <= 1000
//     s and t consist of lowercase English letters.

import "fmt"
import "slices"
import "index/suffixarray"
import "math"
import "unsafe"

func longestPalindrome1(s, t string) int {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    reverse := func(s string) string {
        t := []byte(s)
        slices.Reverse(t)
        return string(t)
    }
    calc := func(s, t string) int {
        // ts = t + "#" + s
        ts := append([]byte(t), '#')
        tmp := []byte(s)
        slices.Reverse(tmp)
        ts = append(ts, tmp...)
        sa := (*struct {
            _  []byte
            sa []int32
        })(unsafe.Pointer(suffixarray.New(ts))).sa

        // 后缀名次数组 rank
        // 后缀 ts[i:] 位于后缀字典序中的第 rank[i] 个
        // 特别地，rank[0] 即 s 在后缀字典序中的排名，rank[n-1] 即 ts[n-1:] 在字典序中的排名
        rank := make([]int, len(sa))
        for i, p := range sa {
            rank[p] = i
        }
        // 高度数组 height
        // sa 中相邻后缀的最长公共前缀 LCP
        // height[0] = 0
        // height[i] = LCP(ts[sa[i]:], ts[sa[i-1]:])
        height := make([]int, len(sa))
        h := 0
        for i, rk := range rank {
            if h > 0 {
                h--
            }
            if rk > 0 {
                for j := int(sa[rk-1]); i+h < len(ts) && j+h < len(ts) && ts[i+h] == ts[j+h]; h++ {
                }
            }
            height[rk] = h
        }
        mx := make([]int, len(s)+1)
        lcp := 0
        // sa[0] 对应 '#' 开头的后缀，不遍历
        for i := 1; i < len(sa); i++ {
            if int(sa[i]) < len(t) {
                lcp = math.MaxInt
            } else {
                lcp = min(lcp, height[i])
                mx[int(sa[i])-len(t)-1] = lcp
            }
        }
        lcp = 0
        if int(sa[len(sa)-1]) < len(t) {
            lcp = math.MaxInt
        }
        for i := len(sa) - 2; i > 0; i-- {
            if int(sa[i]) < len(t) {
                lcp = math.MaxInt
            } else {
                lcp = min(lcp, height[i+1])
                j := int(sa[i]) - len(t) - 1
                mx[j] = max(mx[j], lcp)
            }
        }
        res := slices.Max(mx) * 2 // |x| = |y| 的情况
        slices.Reverse(mx)
        // 计算 |x| > |y| 的情况
        s2 := append(make([]byte, 0, len(s)*2+3), '^')
        for _, c := range s {
            s2 = append(s2, '#', byte(c))
        }
        s2 = append(s2, '#', '$')
        halfLen := make([]int, len(s2)-2)
        halfLen[1] = 1
        boxM, boxR := 0, 0
        for i := 2; i < len(halfLen); i++ {
            hl := 1
            if i < boxR {
                hl = min(halfLen[boxM*2-i], boxR-i)
            }
            for s2[i-hl] == s2[i+hl] {
                hl++
                boxM, boxR = i, i+hl
            }
            halfLen[i] = hl

            if hl > 1 {
                l := (i - hl) / 2
                res = max(res, hl-1+mx[l]*2)
            }
        }
        return res
    }
    return max(calc(s, t), calc(reverse(t), reverse(s)))
}

func longestPalindrome(s, t string) int {
    m, n := len(s), len(t)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    expand := func(s string, g []int, l, r int) {
        for l >= 0 && r < len(s) && s[l] == s[r] {
            g[l] = max(g[l], r-l+1)
            l, r = l-1, r+1
        }
    }
    calc := func(s string) []int {
        n, g := len(s), make([]int, len(s))
        for i := 0; i < n; i++ {
            expand(s, g, i, i)
            expand(s, g, i, i+1)
        }
        return g
    }
    reverse := func(s string) string {
        r := []rune(s)
        slices.Reverse(r)
        return string(r)
    }
    t = reverse(t)
    g1, g2 := calc(s), calc(t)
    res := max(slices.Max(g1), slices.Max(g2))
    f := make([][]int, m + 1)
    for i := range f {
        f[i] = make([]int, n + 1)
    }
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if s[i-1] == t[j-1] {
                f[i][j] = f[i-1][j-1] + 1
                a, b := 0, 0
                if i < m {
                    a = g1[i]
                }
                if j < n {
                    b = g2[j]
                }
                res = max(res, max(f[i][j] * 2 + a, f[i][j] * 2 + b))
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "a", t = "a"
    // Output: 2
    // Explanation:
    // Concatenating "a" from s and "a" from t results in "aa", which is a palindrome of length 2.
    fmt.Println(longestPalindrome("a", "a")) // 2
    // Example 2:
    // Input: s = "abc", t = "def"
    // Output: 1
    // Explanation:
    // Since all characters are different, the longest palindrome is any single character, so the answer is 1.
    fmt.Println(longestPalindrome("abc", "def")) // 1
    // Example 3:
    // Input: s = "b", t = "aaaa"
    // Output: 4
    // Explanation:
    // Selecting "aaaa" from t is the longest palindrome, so the answer is 4.
    fmt.Println(longestPalindrome("b", "aaaa")) // 1
    // Example 4:
    // Input: s = "abcde", t = "ecdba"
    // Output: 5
    // Explanation:
    // Concatenating "abc" from s and "ba" from t results in "abcba", which is a palindrome of length 5.
    fmt.Println(longestPalindrome("abcde", "ecdba")) // 5

    fmt.Println(longestPalindrome("bluefrog", "leetcode")) // 3
    fmt.Println(longestPalindrome("leetcode", "bluefrog")) // 4

    fmt.Println(longestPalindrome1("a", "a")) // 2
    fmt.Println(longestPalindrome1("abc", "def")) // 1
    fmt.Println(longestPalindrome1("b", "aaaa")) // 1
    fmt.Println(longestPalindrome1("abcde", "ecdba")) // 5
    fmt.Println(longestPalindrome1("bluefrog", "leetcode")) // 3
    fmt.Println(longestPalindrome1("leetcode", "bluefrog")) // 4
}